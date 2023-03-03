package service

import (
	"fmt"
	"time"

	"github.com/li-zeyuan/CSEIdiom/backend/config"
	"github.com/li-zeyuan/CSEIdiom/backend/dao"
	"github.com/li-zeyuan/common/external"
	"github.com/li-zeyuan/common/httptransfer"
	"github.com/li-zeyuan/common/model"
	"github.com/li-zeyuan/common/sequence"
	"github.com/li-zeyuan/common/utils"
	"gorm.io/gorm"
)

func WeChatLogin(req *model.WeChatLoginReq) (string, error) {
	wx := external.NewWechat(config.AppCfg.WxAppId, config.AppCfg.WxSecret)
	wxSession, err := wx.QueryWxSession(req.Code)
	if err != nil {
		return "", err
	}

	userProfile, err := dao.D.User.GetByOpenid(wxSession.OpenId)
	if err != nil && err != gorm.ErrRecordNotFound {
		return "", err
	}
	if userProfile != nil && userProfile.DeletedAt != 0 {
		return "", httptransfer.ErrorLoginForbid
	}
	if err == gorm.ErrRecordNotFound {
		userProfile = new(model.UserProfileTable)
		userProfile.ID = sequence.NewID()
		userProfile.Name = fmt.Sprintf("husband_%d", userProfile.ID%1000000)
		userProfile.Openid = wxSession.OpenId
		userProfile.SessionKey = wxSession.SessionKey
		err = dao.D.User.Save([]*model.UserProfileTable{userProfile})
		if err != nil {
			return "", err
		}
	}

	// todo redis token
	token, err := utils.GenerateToken(userProfile.ID, time.Hour*24*30, config.AppCfg.JwtSecret)
	if err != nil {
		return "", err
	}

	return token, nil
}
