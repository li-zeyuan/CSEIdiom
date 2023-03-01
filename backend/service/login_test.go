package service

import (
	"testing"

	"github.com/li-zeyuan/CSEIdiom/backend/testdata"
	"github.com/li-zeyuan/common/model"
	"github.com/stretchr/testify/assert"
)

func TestWeChatLogin(t *testing.T) {
	err := testdata.InitServer()
	assert.Nil(t, err)

	req := &model.WeChatLoginReq{
		Code: "12343",
	}

	token, err := WeChatLogin(req)
	assert.Nil(t, err)
	assert.NotEqual(t, len(token), 0)
}
