package controller

import (
	"github.com/vanilla/go-jwt-crud/api/payload"
	"net/http"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	payload.MessageResponse(w, "Connected to user service Successfully", nil, http.StatusOK)
}