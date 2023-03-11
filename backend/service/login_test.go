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
		Code: "063E6lFa1APAVE0wfBIa12tXBP1E6lFF",
	}

	token, err := WeChatLogin(req)
	assert.Nil(t, err)
	assert.NotEqual(t, len(token), 0)
}
