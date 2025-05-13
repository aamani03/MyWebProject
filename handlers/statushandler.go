package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"MyWebProject/dbInterface"
	"github.com/gorilla/mux"
)

var StatusObj dbInterface.StatusService

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	// Allow only GET method
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return
	}

	vars := mux.Vars(r)
	statusIDStr := vars["id"]

	statusID, err := strconv.Atoi(statusIDStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid status ID"))
		return
	}

	status, err := StatusObj.GetStatus(statusID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("status not found"))
		return
	}

	statusJSON, err := json.Marshal(status)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error encoding status data"))
		return
	}

	// Respond with JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(statusJSON)
}
