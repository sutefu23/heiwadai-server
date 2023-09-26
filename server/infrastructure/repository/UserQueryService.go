package repository

import (
	"context"
	"database/sql"
	"server/core/entity"
	queryservice "server/core/infra/queryService"
	"server/core/infra/queryService/types"
	"server/db/models"

	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var _ queryservice.IUserQueryService = &UserQueryService{}

type UserQueryService struct {
	db *sql.DB
}

func NewUserQueryService() *UserQueryService {
	db := InitDB()

	return &UserQueryService{
		db: db,
	}
}

func (pq *UserQueryService) GetByID(id uuid.UUID) (*entity.User, error) {
	user, err := models.FindUserDatum(context.Background(), pq.db, id.String())
	if err != nil {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return UserModelToEntity(user, user.R.User.Email), nil
}
func (pq *UserQueryService) GetByMail(mail string) (*entity.User, error) {
	usermanager, err := models.UserManagers(models.UserManagerWhere.Email.EQ(mail)).One(context.Background(), pq.db)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return UserModelToEntity(usermanager.R.UserUserDatum, usermanager.Email), nil
}

func (pq *UserQueryService) GetUserByPrefecture(prefectures []*entity.Prefecture) ([]*entity.User, error) {
	var preIds []int
	for _, prefecture := range prefectures {
		preIds = append(preIds, prefecture.ToInt())
	}

	users, err := models.UserData(models.UserDatumWhere.Prefecture.IN(preIds)).All(context.Background(), pq.db)
	if err != nil {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	var entities []*entity.User
	for _, user := range users {
		entity := UserModelToEntity(user, user.R.User.Email)
		entities = append(entities, entity)
	}
	return entities, nil

}

func (pq *UserQueryService) GetAll(query *types.UserQuery, pager *types.PageQuery) ([]*entity.User, error) {
	var firstNameQuery qm.QueryMod = nil
	if query.FirstName != nil {
		firstNameQuery = models.UserDatumWhere.FirstName.EQ("%" + *query.FirstName + "%")
	}
	var lastNameQuery qm.QueryMod = nil
	if query.FirstName != nil {
		lastNameQuery = models.UserDatumWhere.LastName.EQ("%" + *query.LastName + "%")
	}
	var firstNameKanaQuery qm.QueryMod = nil
	if query.FirstNameKana != nil {
		firstNameKanaQuery = models.UserDatumWhere.FirstNameKana.EQ("%" + *query.FirstNameKana + "%")
	}
	var lastNameKanaQuery qm.QueryMod = nil
	if query.LastNameKana != nil {
		lastNameKanaQuery = models.UserDatumWhere.LastNameKana.EQ("%" + *query.LastNameKana + "%")
	}
	var prefectureQuery qm.QueryMod = nil
	if query.LastNameKana != nil {
		prefectureQuery = models.UserDatumWhere.Prefecture.EQ(query.Prefecture.ToInt())
	}

	userdata, err := models.UserData(firstNameQuery, lastNameQuery, firstNameKanaQuery, lastNameKanaQuery, prefectureQuery, qm.Limit(pager.Offset()), qm.Offset(pager.Offset())).All(context.Background(), pq.db)

	if err != nil {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	var result []*entity.User
	for _, user := range userdata {
		result = append(result, UserModelToEntity(user, user.R.User.Email))
	}
	return result, nil
}

func UserModelToEntity(model *models.UserDatum, email string) *entity.User {
	return entity.RegenUser(
		uuid.MustParse(model.UserID),
		model.FirstName,
		model.LastName,
		model.FirstNameKana,
		model.LastNameKana,
		model.CompanyName.Ptr(),
		model.BirthDate,
		&model.ZipCode.String,
		entity.Prefecture(model.Prefecture),
		model.City.Ptr(),
		model.Address.Ptr(),
		model.Tel.Ptr(),
		email,
		model.AcceptMail,
	)
}
