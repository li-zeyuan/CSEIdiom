package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/li-zeyuan/common/httptransfer"
)

func UserProfile(c *gin.Context) {

	httptransfer.SuccJSONResp(c, struct {
		
	}{})
}
