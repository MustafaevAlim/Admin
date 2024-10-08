package socials

import (
	"myapp/internal/model"
)

func ToSocialInfoFromRepo(s SocialsRepo) model.SocialsInfo {
	return model.SocialsInfo{
		ID:         int(s.ID),
		SocialName: s.SocialName,
		SocialLink: s.SocialLink,
		CountViews: int(s.CountViews),
		Updated:    s.UpdatedAt,
		Created:    s.CreatedAt,
	}

}
