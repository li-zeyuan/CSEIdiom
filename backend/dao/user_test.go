package dao

import (
	"fmt"
	"testing"

	"github.com/li-zeyuan/common/model"
	"github.com/li-zeyuan/common/sequence"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestGetByOpenid(t *testing.T) {
	err := InitDao()
	assert.Nil(t, err)

	u := User{}
	user, err := u.GetByOpenid("open_id")
	assert.Equal(t, err, gorm.ErrRecordNotFound)
	assert.Nil(t, user)
}

func TestSave(t *testing.T) {
	err := InitDao()
	assert.Nil(t, err)

	userProfile := new(model.UserProfileTable)
	userProfile.ID = sequence.NewID()
	userProfile.Name = fmt.Sprintf("husband_%d", userProfile.ID%1000000)
	userProfile.Openid = "test"
	userProfile.SessionKey = "deef"

	u := User{}
	err = u.Save([]*model.UserProfileTable{userProfile})
	assert.Nil(t, err)
}
