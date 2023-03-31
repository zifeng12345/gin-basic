package users

import (
	"encoding/json"
	"fmt"
	"net/http"
	"nwd/model/request"
	"nwd/shared/log"
	"nwd/shared/response"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var params request.Login
	err := json.NewDecoder(c.Request.Body).Decode(&params)
	if err != nil {
		log.GetLog().Error("", fmt.Sprintf("User Login failed, error:%v", err))
		response.Response(c.Writer, http.StatusInternalServerError, fmt.Sprintf("User Login failed, error:%v", err), "")
		return
	}

	jwt, err := uctl.service.Login(params)
	if err != nil {
		log.GetLog().Error("", fmt.Sprintf("User Login failed, error:%v", err))
		response.Response(c.Writer, http.StatusInternalServerError, fmt.Sprintf("User Login failed, error:%v", err), "")
		return
	}

	log.GetLog().Info("", "user  %v login success", params.UserName)
	response.Response(c.Writer, http.StatusOK, "User login success", map[string]string{"token": jwt})
}
