package routes

import (
	"github.com/shuklarituparn/VK-Marketplace/api/controllers"
	"github.com/shuklarituparn/VK-Marketplace/config"
	"github.com/shuklarituparn/VK-Marketplace/pkg/common"
	"github.com/shuklarituparn/VK-Marketplace/pkg/middleware"
	"net/http"
)

func AdRouter(mux *http.ServeMux) {
	const prefix = "/api/v1/ads"
	const createAd = prefix + "/create"
	const getAd = prefix + "/get"
	const getAds = prefix + "/get/all"
	const searchAds = prefix + "/search"

	mux.Handle(createAd, middleware.IsAuthenticated(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			controllers.CreateAd(config.GetInstance())(w, r)
		} else {
			common.ErrorResponse(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		}
	})))
	mux.Handle(getAd, middleware.IsAuthenticated(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			controllers.GetAd(config.GetInstance())(w, r)
		} else {
			common.ErrorResponse(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		}
	})))
	mux.Handle(getAds, middleware.IsAuthenticated(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			controllers.GetAllAds(config.GetInstance())(w, r)
		} else {
			common.ErrorResponse(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		}
	})))
	mux.Handle(searchAds, middleware.IsAuthenticated(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			controllers.SearchAd(config.GetInstance())(w, r)
		} else {
			common.ErrorResponse(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		}
	})))
}
