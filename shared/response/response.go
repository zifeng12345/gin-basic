package response

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/fatih/structs"
)

type responseCont struct {
	Code    int         `json:"code"` //Code work and code not work. So it must be set as can access
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Time    string      `json:"time"`
}

func Response(w http.ResponseWriter, code int, message string, data interface{}) {
	res := responseCont{
		Code:    code,
		Message: message,
		Data:    data,
		Time:    time.Now().Format(time.RFC3339),
	}

	w.Header().Add("Content-Type", "application/json")
	resData := structs.Map(&res)

	_ = json.NewEncoder(w).Encode(resData)
}
