package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/shuklarituparn/VK-Marketplace/api/models"
	"github.com/shuklarituparn/VK-Marketplace/pkg/common"

	"github.com/charmbracelet/log"
	"gorm.io/gorm"
)

// SearchAd searches for ads based on the provided query string and optional price range.
// @Summary Search for ads
// @ID search-ads
// @Produce json
// @Tags Search Ads
// @Security BearerAuth
// @Param q query string true "Search query"
// @Param sort_by query string false "Field to sort by (default rating)"
// @Param sort_order query string false "Sort order (ASC or DESC, default DESC)"
// @Param min query string false "Minimum price"
// @Param max query string false "Maximum price"
// @Success 200 {object} SearchAdResponse "List of matching ads"
// @Failure 400 {string} string "Invalid search query"
// @Failure 500 {string} string "Error encoding response"
// @Router /api/v1/ads/search [get]
func SearchAd(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		query := r.URL.Query().Get("q")
		sortBy := r.URL.Query().Get("sort_by")
		sortOrder := strings.ToUpper(r.URL.Query().Get("sort_order"))
		if sortBy == "" {
			sortBy = "created_at"
		}
		if sortOrder != "ASC" && sortOrder != "DESC" {
			sortOrder = "DESC"
		}
		if query == "" {
			common.ErrorResponse(w, http.StatusBadRequest, "Invalid search query")
			return
		}
		searchQuery := "%" + strings.TrimSpace(query) + "%"

		var ads []models.Ad
		if sortBy == "price" {
			max := r.URL.Query().Get("max")
			min := r.URL.Query().Get("min")
			if min == "" || max == "" {
				common.ErrorResponse(w, http.StatusBadRequest, "Min and max price are required")
				return
			}
			db.Model(&models.Ad{}).
				Where("LOWER(ad_text) LIKE ? OR LOWER(title) LIKE ? AND price BETWEEN ? AND ?", "%"+strings.ToLower(searchQuery)+"%", "%"+strings.ToLower(searchQuery)+"%", "%"+strings.ToLower(searchQuery)+"%", min, max).
				Order(fmt.Sprintf("%s %s", sortBy, sortOrder)).
				Distinct().
				Find(&ads)
		} else {
			db.Model(&models.Ad{}).
				Where("LOWER(ad_text) LIKE ? OR LOWER(title) LIKE ?", "%"+strings.ToLower(searchQuery)+"%", "%"+strings.ToLower(searchQuery)+"%", "%"+strings.ToLower(searchQuery)+"%").
				Order(fmt.Sprintf("%s %s", sortBy, sortOrder)).
				Distinct().
				Find(&ads)
		}
		if err := json.NewEncoder(w).Encode(map[string]interface{}{"data": ads}); err != nil {
			log.Error("Error encoding response:", err.Error())
			fileLogger.Println("Error encoding response:", err.Error())
			common.ErrorResponse(w, http.StatusInternalServerError, "Error encoding response")
			return
		}
	}
}
