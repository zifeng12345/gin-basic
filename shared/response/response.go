package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
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
	resd, err := json.Marshal(&res)
	if err != nil {
		fmt.Printf("Response err : %v", err)
	}

	w.Write(resd)
}
