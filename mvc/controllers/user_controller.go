package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/stepup99/microservices-p1/mvc/services"
	"github.com/stepup99/microservices-p1/mvc/utils"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	userId, err := strconv.Atoi(req.URL.Query().Get("user_id"))
	resp.Header().Set("Content-Type", "application/json")

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id is must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write([]byte(jsonValue))
		return
	}

	user, apiErr := services.GetUser(int64(userId))

	if apiErr != nil {
		resp.WriteHeader(apiErr.StatusCode)
		jsonValue, _ := json.Marshal(apiErr)
		resp.Write([]byte(jsonValue))
		return
	}

	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}
