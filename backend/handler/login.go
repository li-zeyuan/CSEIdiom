package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/li-zeyuan/CSEIdiom/backend/service"
	"github.com/li-zeyuan/common/httptransfer"
	"github.com/li-zeyuan/common/model"
)

func WechatLogin(c *gin.Context) {
	apiReq := new(model.WeChatLoginReq)
	err := httptransfer.ParseBody(c, apiReq)
	if err != nil {
		httptransfer.ErrJSONResp(c, err)
		return
	}

	token, err := service.WeChatLogin(apiReq)
	if err != nil {
		httptransfer.ErrJSONResp(c, err)
		return
	}

	resp := new(model.WeChatLoginResp)
	resp.Token = token
	httptransfer.SuccJSONResp(c, resp)
}
