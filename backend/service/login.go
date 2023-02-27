package service

import (
	"fmt"
	"time"

	"github.com/li-zeyuan/CSEIdiom/backend/config"
	"github.com/li-zeyuan/CSEIdiom/backend/dao"
	"github.com/li-zeyuan/common/external"
	"github.com/li-zeyuan/common/model"
	"github.com/li-zeyuan/common/sequence"
	"github.com/li-zeyuan/common/utils"
	"gorm.io/gorm"
)

func WeChatLogin(req *model.WeChatLoginReq) (string, error) {
	wx := external.NewWechat(config.AppCfg.AppId, config.AppCfg.Secret)
	wxSession, err := wx.QueryWxSession(req.Code)
	if err != nil {
		return "", err
	}

	userProfile, err := dao.D.User.GetByOpenid(wxSession.OpenId)
	if err != nil && err != gorm.ErrRecordNotFound {
		return "", err
	}
	if err == gorm.ErrRecordNotFound {
		userProfile = new(model.UserProfileTable)
		userProfile.Uid = sequence.NewID()
		userProfile.Name = fmt.Sprintf("husband_%d", userProfile.Uid%1000000)
		userProfile.Openid = wxSession.OpenId
		userProfile.SessionKey = wxSession.SessionKey
		err = dao.D.User.Save([]*model.UserProfileTable{userProfile})
		if err != nil {
			return "", err
		}
	}

	token, err := utils.GenerateToken(userProfile.Uid, utils.RoleUser, time.Hour*24*30)
	if err != nil {
		return "", err
	}

	return token, nil
}