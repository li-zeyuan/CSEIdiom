package handler

import (
	"fmt"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/li-zeyuan/common/httptransfer"
)

func UserProfile(c *gin.Context) {
	uid := httptransfer.GetUid(c)
	fmt.Println(requestid.Get(c))
	httptransfer.SuccJSONResp(c, struct {
		Uid int64
	}{
		uid,
	})
}
