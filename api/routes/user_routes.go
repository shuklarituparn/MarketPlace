package routes

import (
	"github.com/shuklarituparn/VK-Marketplace/api/controllers"
	"github.com/shuklarituparn/VK-Marketplace/config"
	"github.com/shuklarituparn/VK-Marketplace/pkg/common"
	"github.com/shuklarituparn/VK-Marketplace/pkg/middleware"
	"net/http"
)

func UserRouter(mux *http.ServeMux) {
	const prefix = "/api/v1/users"

	registerRoute := prefix + "/register"
	loginRoute := prefix + "/login"
	refreshTokenRoute := prefix + "/refresh"

	mux.Handle(registerRoute, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			controllers.RegisterUser(config.GetInstance())(w, r)
		} else {
			common.ErrorResponse(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		}
	}))

	mux.Handle(loginRoute, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			controllers.LoginUser(config.GetInstance())(w, r)
		} else {
			common.ErrorResponse(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		}
	}))

	mux.Handle(refreshTokenRoute, middleware.IsAuthenticated(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			controllers.RefreshToken(w, r)
		} else {
			common.ErrorResponse(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		}
	})))
}
