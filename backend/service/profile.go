package service

import (
	"github.com/li-zeyuan/CSEIdiom/backend/dao"
	"github.com/li-zeyuan/CSEIdiom/backend/model"
	"github.com/li-zeyuan/common/utils"
)

var Profile = profileService{}

type profileService struct{}

func (l *profileService) Detail(uid int64) (*model.ProfileApiDetailResp, error) {
	userProfile, err := dao.D.User.GetOne(uid)
	if err != nil {
		return nil, err
	}

	resp := new(model.ProfileApiDetailResp)
	resp.UpdatedAt = utils.Time2TimeStamp(userProfile.UpdatedAt)
	resp.Uid = userProfile.ID
	resp.Name = userProfile.Name
	resp.Gender = userProfile.Gender
	resp.Portrait = userProfile.Portrait
	resp.CurrentSubjectId = userProfile.CurrentSubjectId
	resp.StudyTotal = userProfile.StudyTotal

	return resp, nil
}
