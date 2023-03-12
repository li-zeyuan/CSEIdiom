package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/li-zeyuan/CSEIdiom/backend/service"
	"github.com/li-zeyuan/common/httptransfer"
	"github.com/li-zeyuan/common/model"
)

var LoginHandler = new(login)

type login struct {}

func (l *login)WechatLogin(c *gin.Context) {
	apiReq := new(model.WeChatLoginReq)
	err := httptransfer.ParseBody(c, apiReq)
	if err != nil {
		httptransfer.ErrJSONResp(c,http.StatusInternalServerError, err)
		return
	}

	token, err := service.Login.WeChatLogin(apiReq)
	if err != nil {
		httptransfer.ErrJSONResp(c,http.StatusInternalServerError, err)
		return
	}

	resp := new(model.WeChatLoginResp)
	resp.Token = token
	httptransfer.SuccJSONResp(c, resp)
}
