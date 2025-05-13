package handlers

import (
	"encoding/json"
	"net/http"
)

type ApiErrorResponse struct {
	error
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

func NewApiErrorResponse(resp http.ResponseWriter, code int, msg string) {
	erData, er := json.Marshal(&ApiErrorResponse{
		Code:    code,
		Message: msg,
	})

	if er != nil {
		resp.WriteHeader(code)
		resp.Write([]byte(er.Error()))
	}
	resp.WriteHeader(code)
	resp.Write(erData)
}
