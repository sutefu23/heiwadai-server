package repository

import (
	"context"
	"database/sql"
	"errors"

	"server/core/entity"
	queryservice "server/core/infra/queryService"
	"server/db/models"

	"github.com/google/uuid"
)

var _ queryservice.IStoreQueryService = &StoreQueryService{}

type StoreQueryService struct {
	db *sql.DB
}

func NewStoreQueryService() *StoreQueryService {
	db := InitDB()

	return &StoreQueryService{
		db: db,
	}
}

func (pq *StoreQueryService) GetByID(id uuid.UUID) (*entity.Store, error) {
	ctx := context.Background()
	store, err := models.FindStore(ctx, pq.db, id.String())
	if err != nil {
		return nil, err
	}
	if store == nil {
		return nil, nil
	}
	info, err := models.StayableStoreInfos(models.StayableStoreInfoWhere.StoreID.EQ(store.ID)).One(ctx, InitDB())
	if err != nil {
		return nil, err
	}
	if info == nil {
		info = nil
	}

	return StoreModelToEntity(store, info), nil
}

func (pq *StoreQueryService) GetActiveAll() ([]*entity.Store, error) {
	stores, err := models.Stores(models.StoreWhere.IsActive.EQ(true)).All(context.Background(), pq.db)
	if err != nil {
		return nil, err
	}
	if stores == nil {
		return nil, nil
	}
	var result []*entity.Store
	for _, store := range stores {
		result = append(result, StoreModelToEntity(store, nil))
	}
	return result, nil
}

func (pq *StoreQueryService) GetStayableByID(id uuid.UUID) (*entity.StayableStore, error) {
	store, err := models.FindStore(context.Background(), pq.db, id.String())
	if err != nil {
		return nil, err
	}
	if store == nil {
		return nil, nil
	}
	infoModel, err := store.StayableStoreInfo().One(context.Background(), InitDB())
	if err != nil {
		return nil, err
	}

	return StayableStoreToEntity(store, infoModel), nil
}

func (pq *StoreQueryService) GetStayables() ([]*entity.StayableStore, error) {
	stores, err := models.Stores(models.StoreWhere.IsActive.EQ(true), models.StoreWhere.Stayable.EQ(true)).All(context.Background(), pq.db)
	if err != nil {
		return nil, err
	}
	if stores == nil {
		return nil, nil
	}
	var result []*entity.StayableStore
	for _, store := range stores {
		infoModel, _ := store.StayableStoreInfo().One(context.Background(), InitDB())
		stayable := StayableStoreToEntity(store, infoModel)
		result = append(result, stayable)
	}
	return result, nil
}

func (pq *StoreQueryService) GetAll() ([]*entity.Store, error) {
	stores, err := models.Stores().All(context.Background(), pq.db)
	if err != nil {
		return nil, err
	}
	if stores == nil {
		return nil, nil
	}
	var result []*entity.Store
	for _, store := range stores {
		result = append(result, StoreModelToEntity(store, nil))
	}
	return result, nil
}

func (pq *StoreQueryService) GetStayableByBookingID(bookingID string) (*entity.StayableStore, error) {
	stayables, err := pq.GetStayables()
	if err != nil {
		return nil, err
	}
	for _, stayable := range stayables {
		if stayable.BookingSystemID == bookingID {
			return stayable, nil
		}
	}
	return nil, errors.New("該当のStayableStoreがBookingIDから見つけることが出来ません。")
}

func StoreModelToEntity(model *models.Store, info *models.StayableStoreInfo) *entity.Store {
	var stayableInfo *entity.StayableStoreInfo
	if info != nil {
		stayableInfo = StayableInfoToEntity(info)
	}

	return entity.RegenStore(
		uuid.MustParse(model.ID),
		model.Name,
		model.BranchName.Ptr(),
		model.ZipCode,
		model.Address,
		model.Tel,
		model.SiteURL,
		model.StampImageURL,
		model.Stayable,
		model.IsActive,
		uuid.MustParse(model.QRCode),
		uuid.MustParse(model.UnLimitedQRCode),
		stayableInfo,
	)
}

func StayableInfoToEntity(info *models.StayableStoreInfo) *entity.StayableStoreInfo {
	return entity.RegenStayableStoreInfo(
		info.Parking,
		info.Latitude,
		info.Longitude,
		info.AccessInfo,
		info.RestAPIURL,
		info.BookingSystemID,
	)
}

func StayableStoreToEntity(store *models.Store, info *models.StayableStoreInfo) *entity.StayableStore {
	var stayableInfo *entity.StayableStoreInfo
	if info != nil {
		stayableInfo = StayableInfoToEntity(info)
	}
	var storeEntity *entity.Store
	if store != nil {
		storeEntity = StoreModelToEntity(store, nil)
	}

	return entity.RegenStayableStore(
		storeEntity,
		stayableInfo,
	)
}
