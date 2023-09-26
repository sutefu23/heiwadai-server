package admin

import (
	"server/core/entity"
	"server/core/errors"
	queryservice "server/core/infra/queryService"
	"server/core/infra/queryService/types"
	"server/core/infra/repository"
	"time"

	"github.com/google/uuid"
)

type UserDataUsecase struct {
	userRepository repository.IUserRepository
	userQuery      queryservice.IUserQueryService
}

func NewUserDataUsecase(userRepository repository.IUserRepository, userQuery queryservice.IUserQueryService) *UserDataUsecase {
	return &UserDataUsecase{
		userRepository: userRepository,
		userQuery:      userQuery,
	}
}

func (u *UserDataUsecase) GetAll(query *types.UserQuery, pager *types.PageQuery) ([]*entity.User, *errors.DomainError) {
	users, err := u.userQuery.GetAll(query, pager)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	return users, nil
}

func (u *UserDataUsecase) Update(
	ID uuid.UUID,
	FirstName string,
	LastName string,
	FirstNameKana string,
	LastNameKana string,
	CompanyName *string,
	BirthDate time.Time,
	ZipCode *string,
	Prefecture string,
	City *string,
	Address *string,
	Tel *string,
	Mail string,
	AcceptMail bool, // メルマガ配信可
) (*entity.User, *errors.DomainError) {

	existUser, err := u.userQuery.GetByID(ID)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	if existUser == nil {
		return nil, errors.NewDomainError(errors.QueryDataNotFoundError, "登録されているユーザーが存在しません")
	}

	existUser, err = u.userQuery.GetByMail(Mail)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	if existUser == nil {
		return nil, errors.NewDomainError(errors.QueryDataNotFoundError, "このアドレスで登録されているユーザーが存在しません")
	}

	prefecture, domainErr := entity.StringToPrefecture(Prefecture)
	if domainErr != nil {
		return nil, domainErr
	}

	updateData := entity.RegenUser(
		ID,
		FirstName,
		LastName,
		FirstNameKana,
		LastNameKana,
		CompanyName,
		BirthDate,
		ZipCode,
		prefecture,
		City,
		Address,
		Tel,
		Mail,
		AcceptMail,
	)
	err = u.userRepository.Save(updateData, nil)
	if err != nil {
		return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}

	return updateData, nil
}