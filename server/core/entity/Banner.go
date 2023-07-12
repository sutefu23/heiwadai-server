package entity

import (
	"time"

	"github.com/google/uuid"
)

type Banner struct {
	ID        uuid.UUID
	ImageURL  string
	Url       string
	Status    BannerStatus
	CreateAt  time.Time
	UpdatedAt time.Time
}

type BannerStatus int

const (
	BannerDraft BannerStatus = iota
	BannerPublish
)

func (b BannerStatus) String() string {
	switch b {
	case BannerDraft:
		return "Draft"
	case BannerPublish:
		return "Publish"
	default:
		return "Unknown"
	}
}

func CreateBanner(
	ImageURL string,
	Url string,
) *Banner {
	return &Banner{
		ID:        uuid.New(),
		ImageURL:  ImageURL,
		Url:       Url,
		Status:    BannerDraft,
		CreateAt:  time.Now(),
		UpdatedAt: time.Now(),
	}
}

func RegenBanner(
	ID uuid.UUID,
	ImageURL string,
	Url string,
	Status BannerStatus,
	CreateAt time.Time,
	UpdatedAt time.Time,
) *Banner {
	return &Banner{
		ID:        uuid.New(),
		ImageURL:  ImageURL,
		Url:       Url,
		Status:    BannerDraft,
		CreateAt:  CreateAt,
		UpdatedAt: UpdatedAt,
	}
}