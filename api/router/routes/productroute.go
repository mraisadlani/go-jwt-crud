package routes

import (
	"github.com/vanilla/go-jwt-crud/api/controller"
	"net/http"
)

var productRoute = []Route{
	Route{
		URI: "/getproducts",
		Method: http.MethodGet,
		Handler: controller.GetAllProduct,
		AuthRequired: true,
	},
	Route{
		URI: "/get_product/{id}",
		Method: http.MethodGet,
		Handler: controller.GetProduct,
		AuthRequired: true,
	},
	Route{
		URI: "/create_product",
		Method: http.MethodPost,
		Handler: controller.CreateProduct,
		AuthRequired: true,
	},
	Route{
		URI: "/update_product/{id}",
		Method: http.MethodPut,
		Handler: controller.UpdateProduct,
		AuthRequired: true,
	},
	Route{
		URI: "/delete_product/{id}",
		Method: http.MethodDelete,
		Handler: controller.DeleteProduct,
		AuthRequired: true,
	},
}
