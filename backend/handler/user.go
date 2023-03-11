package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/li-zeyuan/CSEIdiom/backend/service"
	"github.com/li-zeyuan/common/httptransfer"
)

func UserProfile(c *gin.Context) {
	uid := httptransfer.GetUid(c)
	profile, err := service.Profile.Detail(uid)
	if err != nil {
		httptransfer.ErrJSONResp(c, http.StatusInternalServerError, err)
		return
	}

	httptransfer.SuccJSONResp(c, profile)
}
