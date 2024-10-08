package user

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type UserRepo struct {
	ID         uuid.UUID `gorm:"type:varchar(200)"`
	TgId       string    `gorm:"type:varchar(100);unique_index"`
	Wallet     string    `gorm:"type:text"`
	CountViews int       `gorm:"type:bigint"`
	Balance    float32   `gorm:"type:float"`
	Channels   int       `gorm:"type:int"`
	Referrals  int       `gorm:"type:int"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
