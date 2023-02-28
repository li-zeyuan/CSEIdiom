package dao

import (
	"github.com/li-zeyuan/common/mylogger"
	"gorm.io/gorm"
	"testing"

	"github.com/li-zeyuan/common/mysqlstore"
	"github.com/stretchr/testify/assert"
)

func newMysql() error {
	mylogger.Init(nil)

	conf := mysqlstore.Config{
		DSN:     "root:root@tcp(localhost:3306)/government_exam?charset=utf8mb4&parseTime=True&loc=Local",
		MaxConn: 1,
		MaxOpen: 1,
		Timeout: 10,
	}
	return mysqlstore.New(&conf)
}

func TestGetByOpenid(t *testing.T) {
	err := newMysql()
	assert.Nil(t, err)

	u := User{}
	user, err := u.GetByOpenid("open_id")
	assert.Equal(t, err, gorm.ErrRecordNotFound)
	assert.Nil(t, user)
}
