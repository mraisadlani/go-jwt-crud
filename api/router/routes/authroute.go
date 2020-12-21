package routes

import (
	"github.com/vanilla/go-jwt-crud/api/controller"
	"net/http"
)

var authRoute = []Route{
	Route{
		URI: "/signin",
		Method: http.MethodPost,
		Handler: controller.DoSign,
		AuthRequired: false,
	},
}