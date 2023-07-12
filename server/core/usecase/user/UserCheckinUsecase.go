package user

import (
	"server/core/entity"
	"server/core/errors"
	queryService "server/core/infra/queryService"
	"server/core/infra/repository"
	"time"

	"github.com/google/uuid"
)

type UserCheckinUsecase struct {
	storeRepository   repository.IStoreRepository
	checkInRepository repository.ICheckinRepository
	couponRepository  repository.ICouponRepository
	storeQuery        queryService.IStoreQueryService
	checkinQuery      queryService.ICheckinQueryService
	couponQuery       queryService.ICouponQueryService
	transaction       repository.ITransaction
}

func NewUserCheckinUsecase(
	storeRepository repository.IStoreRepository,
	checkInRepository repository.ICheckinRepository,
	couponRepository repository.ICouponRepository,
	storeQuery queryService.IStoreQueryService,
	checkinQuery queryService.ICheckinQueryService,
	couponQuery queryService.ICouponQueryService,
	transaction repository.ITransaction,

) *UserCheckinUsecase {
	return &UserCheckinUsecase{
		storeRepository:   storeRepository,
		checkInRepository: checkInRepository,
		couponRepository:  couponRepository,
		storeQuery:        storeQuery,
		checkinQuery:      checkinQuery,
		couponQuery:       couponQuery,
		transaction:       transaction,
	}
}

func (u *UserCheckinUsecase) GetStampCard(user *entity.User) (*entity.StampCard, *errors.DomainError) {
	userCheckins, err := u.checkinQuery.GetActiveCheckin(user)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	return entity.NewStampCard(userCheckins)
}

// チェックインによってクーポンが付与された場合クーポンを返す
func (u *UserCheckinUsecase) Checkin(AuthUser *entity.User, QrHash uuid.UUID) (*entity.Coupon, *errors.DomainError) {

	allStores, err := u.storeQuery.GetActiveAll()
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	var checkInStore *entity.Store
	var isUnlimitQr bool
	lastCheckin, err := u.checkinQuery.GetLastStoreCheckin(AuthUser, checkInStore)

	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}

	for _, store := range allStores {
		//通常のQRコードでは24時間以内にチェックインした店舗はチェックインできない
		if store.QRCode == QrHash {
			checkInStore = store
			isUnlimitQr = false
		}
		//無制限のQRコード
		if store.UnLimitedQRCode == QrHash {
			checkInStore = store
			isUnlimitQr = true
		}
	}

	if checkInStore == nil {
		return nil, errors.NewDomainError(errors.QueryDataNotFoundError, "該当のQRコードの店舗が見つかりません。")
	}

	isSameStore := lastCheckin != nil && lastCheckin.Store.ID == checkInStore.ID

	if !isUnlimitQr && isSameStore && lastCheckin.CheckInAt.Add(24*time.Hour).After(time.Now()) {
		return nil, errors.NewDomainError(errors.UnPemitedOperation, "24時間以内にチェックインした店舗はチェックインできません。")
	}

	u.transaction.Begin()
	newCheckin := entity.CreateCheckin(*checkInStore, *AuthUser)
	myCheckins, err := u.checkinQuery.GetActiveCheckin(AuthUser)
	if err != nil {
		u.transaction.Rollback()
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	var newCoupon *entity.Coupon
	if len(myCheckins) >= 5 {
		newCoupon, domainErr := entity.CreateStandardCoupon(AuthUser, allStores)
		if domainErr != nil {
			u.transaction.Rollback()
			return nil, domainErr
		}
		err = u.couponRepository.Save(newCoupon)
		if err != nil {
			u.transaction.Rollback()
			return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
		}
	}

	err = u.checkInRepository.Save(newCheckin)
	if err != nil {
		u.transaction.Rollback()
		return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}
	u.transaction.Commit()

	return newCoupon, nil
}
