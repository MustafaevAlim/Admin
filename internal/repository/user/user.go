package user

import (
	"log"
	"myapp/internal/model"
	"myapp/internal/repository/socials"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

func InRepo(user model.UserAdd, db *gorm.DB) {
	uuid, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("failed to generate UUID: %v", err)
	}

	u := UserRepo{ID: uuid, TgId: user.TgId, Wallet: user.Wallet, CountViews: 0,
		Balance: 0, Referrals: 0, Channels: 0}
	res := db.Create(&u)
	if res.Error != nil {
		panic(res.Error.Error())
	}

}

func UpdateBalance(tgId string, value float32, db *gorm.DB) error {
	result := db.Model(&UserRepo{}).Where("tg_id = ?", tgId).Update("balance", gorm.Expr("balance + ?", value))
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetAllViews(id uuid.UUID, db *gorm.DB) int {
	var totalAmount int
	db.Model(&socials.SocialsRepo{}).
		Where("user_id = ?", id).
		Select("SUM(count_views) as total").
		Scan(&totalAmount)
	db.Model(&UserRepo{}).Where("id = ?", id).Update("count_views", totalAmount)
	return totalAmount

}

func Get(tgId string, db *gorm.DB) model.UserInfo {
	var user UserRepo
	result := db.Where("tg_id = ?", tgId).First(&user)
	if result.Error != nil {
		panic(result.Error.Error())
	}
	return model.UserInfo{
		TgId:       user.TgId,
		Wallet:     user.Wallet,
		Channels:   user.Channels,
		CountViews: user.CountViews,
		Referrals:  user.Referrals,
		Balance:    int(user.Balance),
	}

}

func GetAll(db *gorm.DB) []model.UserInfo {
	var AllNote = make([]UserRepo, 0)
	result := db.Find(&AllNote)
	if result.Error != nil {
		panic(result.Error.Error())
	}
	return Map(AllNote, db, ToUserInfoFromRepo)

}

func IncChannels(tgId string, db *gorm.DB) error {
	result := db.Model(&UserRepo{}).Where("tg_id = ?", tgId).Update("channels", gorm.Expr("channels + ?", 1))
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetFromTgId(tgId string, db *gorm.DB) (*model.UserInfo, error) {
	var u UserRepo
	result := db.Where("tg_id = ?", tgId).First(&u)
	if result.Error != nil {
		return nil, result.Error
	}
	mUser := model.UserInfo{
		UserID:     u.ID,
		TgId:       u.TgId,
		Wallet:     u.Wallet,
		Channels:   u.Channels,
		CountViews: u.CountViews,
		Balance:    int(u.Balance),
		Referrals:  u.Referrals,
	}
	return &mUser, nil
}

func GetIdFromTgId(tgId string, db *gorm.DB) (uuid.UUID, error) {
	var u UserRepo
	result := db.Where("tg_id = ?", tgId).First(&u)
	if result.Error != nil {
		return uuid.Nil, result.Error
	}

	return u.ID, nil
}

func Map(vs []UserRepo, db *gorm.DB, f func(repo UserRepo) model.UserInfo) []model.UserInfo {
	vsm := make([]model.UserInfo, len(vs))
	for i, v := range vs {
		v.CountViews = GetAllViews(v.ID, db)
		vsm[i] = f(v)
	}
	return vsm
}
