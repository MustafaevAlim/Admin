package admin

import "gorm.io/gorm"

type AdminRepo struct {
	gorm.Model
	Password string `gorm:"varchar(200)"`
	Name     string `gorm:"varchar(200)"`
}
