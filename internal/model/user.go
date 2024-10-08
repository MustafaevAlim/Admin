package model

import "github.com/gofrs/uuid"

type UserAdd struct {
	TgId       string `json:"tgId" validate:"required"`
	Wallet     string `json:"wallet" validate:"required"`
	SocialName string `json:"social_name" validate:"required"`
	SocialLink string `json:"social_link" validate:"required,min=4"`
	CountViews int    `json:"count_views" validate:"required"`
}

type UserInfo struct {
	TgId       string
	Wallet     string
	CountViews int
	UserID     uuid.UUID
	Channels   int
	Referrals  int
	Balance    int
}
