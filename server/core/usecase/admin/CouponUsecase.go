package admin

import (
	"time"

	"server/core/entity"
	"server/core/errors"
	queryservice "server/core/infra/queryService"
	"server/core/infra/repository"

	"github.com/google/uuid"
)

type AdminCouponUsecase struct {
	couponRepository     repository.ICouponRepository
	couponQuery          queryservice.ICouponQueryService
	userCouponQuery      queryservice.IUserCouponQueryService
	usercouponRepository repository.IUserCouponRepository
	storeQuery           queryservice.IStoreQueryService
}

func NewAdminCouponUsecase(couponRepository repository.ICouponRepository, couponQuery queryservice.ICouponQueryService,
	userCouponQuery queryservice.IUserCouponQueryService, usercouponRepository repository.IUserCouponRepository, storeQuery queryservice.IStoreQueryService,
) *AdminCouponUsecase {
	return &AdminCouponUsecase{
		couponRepository:     couponRepository,
		couponQuery:          couponQuery,
		usercouponRepository: usercouponRepository,
		storeQuery:           storeQuery,
	}
}

func (u *AdminCouponUsecase) CreateDefaultCoupon() *errors.DomainError {
	// Seederで叩く想定。デフォルトのクーポンをDB生成＆保持
	store, err := u.storeQuery.GetActiveAll()
	if err != nil {
		return errors.NewDomainError(errors.RepositoryError, err.Error())
	}

	standard, domainErr := entity.CreateStandardCoupon(store)
	if err != nil {
		return domainErr
	}
	err = u.couponRepository.Save(standard)
	if err != nil {
		return errors.NewDomainError(errors.RepositoryError, err.Error())
	}
	return nil
}

func (u *AdminCouponUsecase) GetUsersCouponList(UserID uuid.UUID) ([]*entity.UserAttachedCoupon, *errors.DomainError) {
	coupons, err := u.userCouponQuery.GetActiveAll(UserID)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	return coupons, nil
}

func (u *AdminCouponUsecase) CreateCustomCoupon(
	Name string,
	DiscountAmount uint,
	ExpireAt time.Time,
	IsCombinationable bool,
	Notices []string,
) (*entity.Coupon, *errors.DomainError) {
	stores, err := u.storeQuery.GetActiveAll()
	if err != nil {
		return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}

	customCoupon, domainErr := entity.CreateCustomCoupon(
		Name,
		DiscountAmount,
		ExpireAt,
		IsCombinationable,
		Notices,
		stores,
	)
	if domainErr != nil {
		return nil, domainErr
	}
	err = u.couponRepository.Save(customCoupon)
	if err != nil {
		return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}

	return customCoupon, nil
}

func (u *AdminCouponUsecase) SaveCustomCoupon(couponID uuid.UUID) *errors.DomainError {
	coupon, err := u.couponQuery.GetByID(couponID)
	if err != nil {
		return errors.NewDomainError(errors.QueryError, err.Error())
	}
	if coupon == nil {
		return errors.NewDomainError(errors.QueryDataNotFoundError, "該当のクーポンIDが見つかりません。")
	}
	if coupon.Status != entity.CouponDraft {
		return errors.NewDomainError(errors.UnPemitedOperation, "下書き状態のクーポンではありません。")
	}

	saveCoupon, domainErr := entity.SaveCustomCoupon(coupon)
	if domainErr != nil {
		return domainErr
	}

	err = u.couponRepository.Save(saveCoupon)

	if err != nil {
		return errors.NewDomainError(errors.RepositoryError, err.Error())
	}
	return nil
}

func (u *AdminCouponUsecase) AttachCustomCouponToAllUser(couponID uuid.UUID) (*int, *errors.DomainError) {
	coupon, err := u.couponQuery.GetByID(couponID)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	if coupon == nil {
		return nil, errors.NewDomainError(errors.QueryDataNotFoundError, "該当のクーポンIDが見つかりません。")
	}
	if coupon.Status != entity.CouponSaved {
		return nil, errors.NewDomainError(errors.UnPemitedOperation, "保存済ステータスのクーポンではありません。")
	}

	count, err := u.usercouponRepository.IssueAll(coupon)
	if err != nil {
		return nil, errors.NewDomainError(errors.ActionError, err.Error())
	}
	if count == 0 {
		return nil, errors.NewDomainError(errors.ActionError, "クーポンの発行に失敗しました。")
	}
	return &count, nil
}
