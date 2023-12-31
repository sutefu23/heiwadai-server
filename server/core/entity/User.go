package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID            uuid.UUID
	FirstName     string
	LastName      string
	FirstNameKana string
	LastNameKana  string
	CompanyName   *string
	BirthDate     *time.Time
	ZipCode       *string
	Prefecture    Prefecture
	City          *string
	Address       *string
	Tel           *string
	Mail          string
	AcceptMail    bool
}

type UserWichLastCheckin struct {
	User          *User
	LastCheckinAt *time.Time
}

func CreateUser(
	ID uuid.UUID,
	FirstName string,
	LastName string,
	FirstNameKana string,
	LastNameKana string,
	CompanyName *string,
	BirthDate *time.Time,
	ZipCode *string,
	Prefecture Prefecture,
	City *string,
	Address *string,
	Tel *string,
	Mail string,
	AcceptMail bool,
) *User {
	return &User{
		ID:            ID,
		FirstName:     FirstName,
		LastName:      LastName,
		FirstNameKana: FirstNameKana,
		LastNameKana:  LastNameKana,
		CompanyName:   CompanyName,
		BirthDate:     BirthDate,
		ZipCode:       ZipCode,
		Prefecture:    Prefecture,
		City:          City,
		Address:       Address,
		Tel:           Tel,
		Mail:          Mail,
		AcceptMail:    AcceptMail,
	}
}

func RegenUser(
	ID uuid.UUID,
	FirstName string,
	LastName string,
	FirstNameKana string,
	LastNameKana string,
	CompanyName *string,
	BirthDate *time.Time,
	ZipCode *string,
	Prefecture Prefecture,
	City *string,
	Address *string,
	Tel *string,
	Mail string,
	AcceptMail bool,
) *User {
	return &User{
		ID:            ID,
		FirstName:     FirstName,
		LastName:      LastName,
		FirstNameKana: FirstNameKana,
		LastNameKana:  LastNameKana,
		CompanyName:   CompanyName,
		BirthDate:     BirthDate,
		ZipCode:       ZipCode,
		Prefecture:    Prefecture,
		City:          City,
		Address:       Address,
		Tel:           Tel,
		Mail:          Mail,
		AcceptMail:    AcceptMail,
	}
}
