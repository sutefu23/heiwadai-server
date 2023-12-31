package repository

import (
	"context"
	"database/sql"

	"server/core/entity"
	"server/core/infra/repository"
	"server/db/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

var _ repository.ICouponRepository = &CouponRepository{}

type CouponRepository struct {
	db *sql.DB
}

func NewCouponRepository() *CouponRepository {
	db := InitDB()

	return &CouponRepository{
		db: db,
	}
}

func (pr *CouponRepository) Save(updateCoupon *entity.Coupon) error {
	coupon := models.Coupon{
		ID:                updateCoupon.ID.String(),
		Name:              updateCoupon.Name,
		CouponType:        int(updateCoupon.CouponType),
		DiscountAmount:    int(updateCoupon.DiscountAmount),
		ExpireAt:          updateCoupon.ExpireAt,
		IsCombinationable: updateCoupon.IsCombinationable,
		CreateAt:          updateCoupon.CreateAt,
		CouponStatus:      int(updateCoupon.Status),
	}
	ctx := context.Background()
	tx := NewTransaction()
	err := tx.Begin(ctx)
	if err != nil {
		tx.Rollback()
		return err
	}

	deleteNotices, err := models.FindCoupon(ctx, pr.db, updateCoupon.ID.String())
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = deleteNotices.Delete(ctx, pr.db)
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, notice := range updateCoupon.Notices {
		modelNotice := models.CouponNotice{
			CouponID: updateCoupon.ID.String(),
			Notice:   notice,
		}
		err = modelNotice.Insert(ctx, pr.db, boil.Infer())
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	deleteStores, err := models.FindCoupon(ctx, pr.db, updateCoupon.ID.String())
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = deleteStores.Delete(ctx, pr.db)

	if err != nil {
		tx.Rollback()
		return err
	}

	for _, store := range updateCoupon.TargetStore {
		modelStore := models.CouponStore{
			CouponID: updateCoupon.ID.String(),
			StoreID:  store.ID.String(),
		}
		err = modelStore.Insert(ctx, pr.db, boil.Infer())
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = coupon.Upsert(ctx, pr.db, true, []string{"id"}, boil.Infer(), boil.Infer())
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
