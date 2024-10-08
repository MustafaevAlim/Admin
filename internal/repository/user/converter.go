package user

import (
	"myapp/internal/model"
)

func ToUserInfoFromRepo(s UserRepo) model.UserInfo {
	return model.UserInfo{
		TgId:       s.TgId,
		Wallet:     s.Wallet,
		CountViews: s.CountViews,
		Channels:   s.Channels,
		Referrals:  s.Referrals,
		Balance:    int(s.Balance),
	}

}
