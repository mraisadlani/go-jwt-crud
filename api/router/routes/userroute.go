package routes

import (
	"github.com/vanilla/go-jwt-crud/api/controller"
	"net/http"
)

var userRoute = []Route{
	Route{
		URI: "/getusers",
		Method: http.MethodGet,
		Handler: controller.GetAllUser,
		AuthRequired: true,
	},
	Route{
		URI: "/get_user/{id}",
		Method: http.MethodGet,
		Handler: controller.GetUser,
		AuthRequired: true,
	},
	Route{
		URI: "/create_user",
		Method: http.MethodPost,
		Handler: controller.CreateUser,
		AuthRequired: true,
	},
	Route{
		URI: "/update_user/{id}",
		Method: http.MethodPut,
		Handler: controller.UpdateUser,
		AuthRequired: true,
	},
	Route{
		URI: "/delete_user/{id}",
		Method: http.MethodDelete,
		Handler: controller.DeleteUser,
		AuthRequired: true,
	},
}