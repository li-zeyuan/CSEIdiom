package dao

import (
	"github.com/li-zeyuan/common/model"
	"github.com/li-zeyuan/common/mylogger"
	"github.com/li-zeyuan/common/mysqlstore"
	"go.uber.org/zap"
)

type User struct {}

func (u *User) GetByOpenid(openid string) (*model.UserProfileTable, error) {
	m := new(model.UserProfileTable)
	err := mysqlstore.Db.Table(model.TableNameUserProfile).
		Where("openid = ?", openid).
		First(m).Error
	if err != nil {
		mylogger.Error("get user by openid error: ", zap.Error(err))
		return nil, err
	}

	return m, nil
}

func (u *User) Save( models []*model.UserProfileTable) error {
	if len(models) == 0 {
		return nil
	}

	err := mysqlstore.Db.Table(model.TableNameUserProfile).
		Create(&models).Error
	if err != nil {
		mylogger.Error("create users error: ", zap.Error(err))
		return err
	}

	return nil
}

func (u *User) GetOne(uid int64) (*model.UserProfileTable, error) {
	m := new(model.UserProfileTable)
	err := mysqlstore.Db.Table(model.TableNameUserProfile).
		Where("id = ?", uid).
		Where(model.DefaultDelCondition).
		First(m).Error
	if err != nil {
		mylogger.Error("get user by uid error: ", zap.Error(err))
		return nil, err
	}

	return m, nil
}