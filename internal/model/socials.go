package model

import (
	"time"
)

type SocialsInfo struct {
	SocialName string
	SocialLink string
	CountViews int
	ID         int
	Created    time.Time
	Updated    time.Time
}
