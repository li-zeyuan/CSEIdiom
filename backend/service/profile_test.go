package service

import (
	"testing"

	"github.com/li-zeyuan/CSEIdiom/backend/config"
	"github.com/li-zeyuan/CSEIdiom/backend/dao"
	"github.com/li-zeyuan/CSEIdiom/backend/testdata"
	"github.com/stretchr/testify/assert"
)

func initDao() {
	dao.New(config.AppCfg)
}

func TestDetail(t *testing.T) {
	err := testdata.InitServer()
	assert.Nil(t, err)
	initDao()

	profile, err := Profile.Detail(422664666169176832)
	assert.Nil(t, err)
	t.Log(profile)
}

