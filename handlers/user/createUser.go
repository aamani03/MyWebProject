package user

import (
	"MyWebProject/dbInterface"
	"MyWebProject/handlers"
	"encoding/json"
	"net/http"
)

type CreateUserRequestWrapper struct {
	Req     *http.Request
	UserSvc dbInterface.UserService
}

type CreateUserResponse struct {
	UserID string `json:"user_id"`
	Email  string `json:"email_id"`
}

func CreateUser(resp http.ResponseWriter, req CreateUserRequestWrapper) {
	if req.Req.Method != http.MethodPost {
		handlers.NewApiErrorResponse(resp, http.StatusMethodNotAllowed, "Invalid method")
		return
	}

	userData := dbInterface.User{}
	err := json.NewDecoder(req.Req.Body).Decode(&userData)
	if err != nil {
		handlers.NewApiErrorResponse(resp, http.StatusBadRequest, "Invalid request data")
		return
	}

	userId, err := req.UserSvc.CreateUser(&userData)
	if err != nil {
		handlers.NewApiErrorResponse(resp, http.StatusInternalServerError, "Something went wrong")
		return
	}

	respObj := &CreateUserResponse{
		UserID: userId,
		Email:  userData.Email,
	}

	jsonData, err := json.Marshal(respObj)
	if err != nil {
		handlers.NewApiErrorResponse(resp, http.StatusInternalServerError, "Unable to send userId")
		return
	}

	resp.WriteHeader(200)
	resp.Write(jsonData)
}
