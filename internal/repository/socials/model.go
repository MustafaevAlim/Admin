package socials

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type SocialsRepo struct {
	gorm.Model
	UserID     uuid.UUID `gorm:"type:varchar(200)"`
	SocialName string    `gorm:"type:varchar(200)"`
	SocialLink string    `gorm:"type:text"`
	CountViews int       `gorm:"type:bigint"`
}
