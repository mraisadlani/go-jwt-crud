package routes

import (
	"github.com/gorilla/mux"
	"github.com/vanilla/go-jwt-crud/api/middleware"
	"net/http"
)

type Route struct {
	URI string
	Method string
	Handler func(w http.ResponseWriter, r *http.Request)
	AuthRequired bool
}

func Load() []Route {
	routes := welcomeRoute
	routes = append(routes, userRoute...)
	routes = append(routes, productRoute...)
	routes = append(routes, authRoute...)

	return routes
}

func SetupRouteWithMiddleware(r *mux.Router) *mux.Router {
	for _, route := range Load() {
		api := r.PathPrefix("/api/v1").Subrouter()
		if route.AuthRequired {
			api.HandleFunc(route.URI,
				middleware.SetMiddlewareCors(
					middleware.SetMiddlewareJSON(
						middleware.SetMiddlewareLogger(
							middleware.SetMiddlewareAuthentication(
								route.Handler,
							),
						),
					),
				),
			).Methods(route.Method)
		} else {
			api.HandleFunc(route.URI,
				middleware.SetMiddlewareCors(
					middleware.SetMiddlewareJSON(
						middleware.SetMiddlewareLogger(route.Handler),
					),
				),
			).Methods(route.Method)
		}
	}

	return r
}
