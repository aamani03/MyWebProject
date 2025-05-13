package user

import (
	"MyWebProject/dbInterface"
	"encoding/json"
	"net/http"
)

type GetUserRequestWrapper struct {
	Req     *http.Request
	UserSvc dbInterface.UserService
}

func GetUser(resp http.ResponseWriter, reqObj GetUserRequestWrapper) {
	if reqObj.Req.Method != http.MethodGet {
		resp.WriteHeader(503)
		resp.Write([]byte("invalid method"))
	}

	userId := reqObj.Req.URL.Query().Get("id")

	userData, err := reqObj.UserSvc.GetUser(userId)
	if err != nil {
		resp.Write([]byte("Error getting user"))
		return
	}

	userDataJson, err := json.Marshal(userData)
	if err != nil {
		resp.Write([]byte("Error getting user"))
		return
	}

	resp.WriteHeader(200)
	resp.Write(userDataJson)
}

// API Design
/*
URL - /object/
method -> GET, PUT, POST, DELETE
Request format
Response format
Error Codes/messages defined
Auth
*/
