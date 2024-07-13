package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/stepup99/microservices-p1/mvc/services"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	userId, err := strconv.Atoi(req.URL.Query().Get("user_id"))
	resp.Header().Set("Content-Type", "application/json")
	if err != nil {
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte("user_id must be a number"))
		return
	}
	// resp.Write([]byte("GetUser ................. %v"))
	fmt.Printf("this is user %v", userId)
	user, err := services.GetUser(int64(userId))
	if err != nil {
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte(err.Error()))
		return
	}
	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}
