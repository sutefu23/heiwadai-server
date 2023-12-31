package user

import (
	"server/core/entity"
	"server/core/errors"
	queryservice "server/core/infra/queryService"

	"github.com/google/uuid"
)

type StoreUsecase struct {
	storeQuery queryservice.IStoreQueryService
}

func NewStoreUsecase(storeQuery queryservice.IStoreQueryService) *StoreUsecase {
	return &StoreUsecase{
		storeQuery: storeQuery,
	}
}

func (u *StoreUsecase) GetAll() ([]*entity.Store, *errors.DomainError) {

	stores, err := u.storeQuery.GetActiveAll()
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	return stores, nil
}

func (u *StoreUsecase) GetStayables() ([]*entity.StayableStore, *errors.DomainError) {

	stores, err := u.storeQuery.GetStayables()
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}

	return stores, nil
}

func (u *StoreUsecase) GetStayableByID(id uuid.UUID) (*entity.StayableStore, *errors.DomainError) {

	stayableStore, err := u.storeQuery.GetStayableByID(id)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}

	return stayableStore, nil
}

func (u *StoreUsecase) GetByID(ID uuid.UUID) (*entity.Store, *errors.DomainError) {

	store, err := u.storeQuery.GetByID(ID)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	return store, nil
}
