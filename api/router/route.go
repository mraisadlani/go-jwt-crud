package router

import (
	"github.com/gorilla/mux"
	"github.com/vanilla/go-jwt-crud/api/router/routes"
)

func New() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	return routes.SetupRouteWithMiddleware(r)
}
