package repository

import (
	"context"
	"database/sql"
	"errors"

	"server/core/entity"
	queryservice "server/core/infra/queryService"
	"server/core/infra/queryService/types"
	"server/db/models"

	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var _ queryservice.IUserCouponQueryService = &UserCouponQueryService{}

type UserCouponQueryService struct {
	db *sql.DB
}

func NewUserCouponQueryService() *UserCouponQueryService {
	db := InitDB()

	return &UserCouponQueryService{
		db: db,
	}
}

func (pq *UserCouponQueryService) GetByID(userID uuid.UUID, couponID uuid.UUID) (*entity.UserAttachedCoupon, error) {
	// userCoupon, err := models.FindCouponAttachedUser(context.Background(), pq.db, couponID.String(), userID.String())
	userCoupon, err := models.CouponAttachedUsers(
		models.CouponAttachedUserWhere.UserID.EQ(userID.String()),
		qm.Load(models.CouponAttachedUserRels.Coupon),
		models.CouponAttachedUserWhere.CouponID.EQ(couponID.String())).One(context.Background(), pq.db)
	if err != nil {
		return nil, err
	}

	if userCoupon == nil {
		return nil, errors.New("該当のクーポンIDが見つかりません。")
	}
	coupon := userCoupon.R.Coupon
	if coupon == nil {
		return nil, errors.New("該当のクーポンが見つかりません。")
	}
	entityCoupon := CouponModelToEntity(coupon, nil, nil)
	return entity.RegenUserAttachedCoupon(
		userID,
		entityCoupon,
		userCoupon.UsedAt.Ptr(),
	), nil
}

func (pq *UserCouponQueryService) GetActiveAll(userID uuid.UUID) ([]*entity.UserAttachedCoupon, error) {
	userCoupons, err := models.CouponAttachedUsers(models.CouponAttachedUserWhere.UserID.EQ(userID.String()),
		qm.Load(models.CouponAttachedUserRels.Coupon),
		models.CouponAttachedUserWhere.UsedAt.IsNull()).All(context.Background(), pq.db)
	if err != nil {
		return nil, err
	}
	if userCoupons == nil {
		return nil, nil
	}

	var result []*entity.UserAttachedCoupon
	for _, userCoupon := range userCoupons {
		coupon := userCoupon.R.Coupon
		entityCoupon := CouponModelToEntity(coupon, nil, nil)
		entityUserCoupon := entity.RegenUserAttachedCoupon(
			userID,
			entityCoupon,
			userCoupon.UsedAt.Ptr(),
		)
		result = append(result, entityUserCoupon)
	}
	return result, nil
}

func (pq *UserCouponQueryService) GetAll(userID uuid.UUID, pager *types.PageQuery) ([]*entity.UserAttachedCoupon, error) {
	userCoupons, err := models.CouponAttachedUsers(models.CouponAttachedUserWhere.UserID.EQ(userID.String()),
		qm.Load(models.CouponAttachedUserRels.Coupon),
		qm.Limit(pager.Limit()), qm.Offset(pager.Offset())).All(context.Background(), pq.db)
	if err != nil {
		return nil, err
	}
	if userCoupons == nil {
		return nil, nil
	}
	var result []*entity.UserAttachedCoupon
	for _, userCoupon := range userCoupons {
		coupon := userCoupon.R.Coupon
		entityCoupon := CouponModelToEntity(coupon, nil, nil)
		entityUserCoupon := entity.RegenUserAttachedCoupon(
			userID,
			entityCoupon,
			userCoupon.UsedAt.Ptr(),
		)
		result = append(result, entityUserCoupon)
	}
	return result, nil
}
