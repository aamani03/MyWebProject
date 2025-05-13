package handlers

import (
	"MyWebProject/dbInterface"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

var TaskObj dbInterface.TaskService

func Taskhandler(resp http.ResponseWriter, req http.Request) {

	if req.Method != http.MethodGet {
		resp.WriteHeader(503)
		resp.Write([]byte("invalid method"))
		return
	}

	vars := mux.Vars(req)
	taskID := vars["id"]

	task, err := TaskObj.GetTask(taskID)
	if err != nil {
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte("task not found"))
		return
	}

	taskJSON, err := json.Marshal(task)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("error encoding task data"))
		return
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)
	resp.Write(taskJSON)
}
