package helper

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func ResponseWithJson(w *gin.Context, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Writer.Header().Set("Content-Type", "application/json")
	w.Writer.WriteHeader(code)
	w.Writer.Write(response)
}



