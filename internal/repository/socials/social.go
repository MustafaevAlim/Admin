package socials

import (
	"myapp/internal/model"
	"strconv"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

func InRepo(s model.UserAdd, userID uuid.UUID, db *gorm.DB) {
	u := SocialsRepo{SocialName: s.SocialName, SocialLink: s.SocialLink, CountViews: s.CountViews, UserID: userID}
	res := db.Create(&u)
	if res.Error != nil {
		panic(res.Error.Error())
	}
}

func GetViews(id uuid.UUID, social string, db *gorm.DB) string {
	var views int
	result := db.Model(SocialsRepo{}).
		Where("user_id = ? AND social_name = ?", id, social).
		Select("count_views").Scan(&views)
	if result.Error != nil {
		panic(result.Error.Error())
	}
	count_views := strconv.Itoa(views)
	return count_views
}

func GetLink(id uuid.UUID, social string, db *gorm.DB) string {
	var link string
	result := db.Model(SocialsRepo{}).
		Where("user_id = ? AND social_name = ?", id, social).
		Select("social_link").Scan(&link)
	if result.Error != nil {
		panic(result.Error.Error())
	}
	return link
}

func GetAllViews(social string, db *gorm.DB) int {
	var totalAmount int
	db.Model(&SocialsRepo{}).
		Where("social_name = ?", social).
		Select("SUM(count_views) as total").
		Scan(&totalAmount)
	return totalAmount
}

func Map(vs []SocialsRepo, f func(repo SocialsRepo) model.SocialsInfo) []model.SocialsInfo {
	vsm := make([]model.SocialsInfo, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}
