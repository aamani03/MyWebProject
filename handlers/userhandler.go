package handlers

import (
	"encoding/json"
	"net/http"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		resp.WriteHeader(503)
		resp.Write([]byte("invalid method"))
	}

	userId := req.URL.Query().Get("id")

	userData, err := UserObj.GetUser(userId)
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
