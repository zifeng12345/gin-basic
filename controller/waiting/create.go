package waiting

import (
	"encoding/json"
	"fmt"
	"net/http"

	"nwd/model/request"
	"nwd/shared/response"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var params request.Waiting
	err := json.NewDecoder(c.Request.Body).Decode(&params)
	if err != nil {
		fmt.Println(err)
		response.Response(c.Writer, http.StatusInternalServerError, fmt.Sprintf("Waiting created failed, error:%v", err), "")
		return
	}

	err = wctl.service.Create(params)
	if err != nil {
		fmt.Println(err)
		response.Response(c.Writer, http.StatusInternalServerError, fmt.Sprintf("Waiting created failed, error:%v", err), "")
		return
	}

	response.Response(c.Writer, http.StatusOK, "Waiting created success", "")
}
