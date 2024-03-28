package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/shuklarituparn/VK-Marketplace/api/models"
	"github.com/shuklarituparn/VK-Marketplace/internal/logger"
	"github.com/shuklarituparn/VK-Marketplace/internal/prometheus"
	"github.com/shuklarituparn/VK-Marketplace/pkg/common"
	"gorm.io/gorm"
)

var fileLogger = logger.SetupLogger()

// CreateAd creates a new advertisement.
// @Summary Create a new advertisement
// @ID create-advertisement
// @Accept json
// @Produce json
// @Tags Advertisements
// @Security BearerAuth
// @Param ad body models.CreateAd true "Ad object to be created"
// @Success 201 {object} CreateAdResponse "Ad Added"
// @Failure 400 {string} string "Invalid request payload"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/ads/create [post]
func CreateAd(db *gorm.DB) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		prometheus.AdCreateCounter.Inc()
		writer.Header().Set("Content-Type", "application/json")
		var ad models.Ad
		if err := json.NewDecoder(request.Body).Decode(&ad); err != nil {
			common.ErrorResponse(writer, http.StatusBadRequest, "Invalid request payload")
			return
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				common.ErrorResponse(writer, http.StatusInternalServerError, "Internal Server Error")
			}
		}(request.Body)
		userId, err := common.ExtractUserIDFromToken(strings.Split(request.Header.Get("Authorization"), " ")[1])
		if err != nil {
			common.ErrorResponse(writer, http.StatusInternalServerError, "Internal Server Error")
		}
		ad.UserID = userId
		if !common.ValidateAndRespond(writer, ad) {
			return
		}
		txErr := db.Transaction(func(tx *gorm.DB) error {
			if err := tx.Create(&ad).Error; err != nil {
				if strings.Contains(err.Error(), "duplicate key") {
					common.ErrorResponse(writer, http.StatusConflict, "Advertisement with the same title already exists")
					return nil
				}
				common.ErrorResponse(writer, http.StatusInternalServerError, "Failed to create advertisement")
				return err
			}
			return nil
		})
		if txErr != nil {
			return
		}
		writer.WriteHeader(http.StatusCreated)
		resErr := json.NewEncoder(writer).Encode(map[string]interface{}{"ad": ad, "message": "Ad Added"})
		if resErr != nil {
			common.ErrorResponse(writer, http.StatusInternalServerError, "Internal Server Error")
		}
	}
}

// GetAd retrieves a single advertisement by its ID.
// @Summary Retrieve a single advertisement
// @ID get-advertisement
// @Accept json
// @Produce json
// @Tags Advertisements
// @Security BearerAuth
// @Param id query string true "Advertisement ID"
// @Success 200 {object} ReadAdResponse "Advertisement Data"
// @Failure 400 {string} string "Advertisement ID is required"
// @Failure 404 {string} string "Advertisement not found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/ads/get [get]
func GetAd(db *gorm.DB) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		prometheus.AdGetCounter.Inc()
		writer.Header().Set("Content-Type", "application/json")
		var adId = request.URL.Query().Get("id")
		var ad models.Ad
		if adId == "" {
			common.ErrorResponse(writer, http.StatusBadRequest, "Advertisement ID is required")
			return
		}
		txErr := db.Transaction(func(tx *gorm.DB) error {
			if err := tx.First(&ad, adId).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					common.ErrorResponse(writer, http.StatusNotFound, "Advertisement not found")
					return err
				}
				log.Error("Error fetching ad:", err.Error())
				fileLogger.Println("Error fetching ad:", err.Error())
				common.ErrorResponse(writer, http.StatusInternalServerError, "Failed to fetch Advertisement")
				return err
			}
			return nil
		})
		if txErr != nil {
			return
		}
		writer.WriteHeader(http.StatusOK)
		resErr := json.NewEncoder(writer).Encode(map[string]interface{}{"data": ad})
		if resErr != nil {
			log.Error("Error encoding JSON:", resErr.Error())
			fileLogger.Println("Error encoding JSON:", resErr.Error())
			common.ErrorResponse(writer, http.StatusInternalServerError, "Internal Server Error")
		}
	}
}

// GetAllAds retrieves a list of advertisements with pagination and optional price range.
// @Summary Retrieve a list of advertisements
// @ID get-all-advertisements
// @Accept json
// @Produce json
// @Tags Advertisements
// @Security BearerAuth
// @Param page query integer false "Page number" default(1)
// @Param page_size query integer false "Page size" default(10)
// @Param sort_by query string false "Field to sort by" default(price)
// @Param sort_order query string false "Sort order (ASC or DESC)" default(DESC)
// @Param min query string false "Minimum price"
// @Param max query string false "Maximum price"
// @Success 200 {object} ReadAllAdsResponse "Advertisements Data"
// @Failure 400 {string} string "Invalid page_size or page"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/ads/get/all [get]
func GetAllAds(db *gorm.DB) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		prometheus.AdGetAllCounter.Inc()
		writer.Header().Set("Content-Type", "application/json")
		writer.Header().Set("Content-Type", "application/json")
		var ads []models.Ad
		var totalAdsCount int64
		pageSize, err := strconv.Atoi(request.URL.Query().Get("page_size"))
		if err != nil {
			common.ErrorResponse(writer, http.StatusBadRequest, "Invalid page_size")
			return
		}
		pageNum, err := strconv.Atoi(request.URL.Query().Get("page"))
		if err != nil || pageNum < 1 {
			common.ErrorResponse(writer, http.StatusBadRequest, "Invalid page")
			return
		}
		offset := (pageNum - 1) * pageSize
		sortBy := request.URL.Query().Get("sort_by")
		sortOrder := strings.ToUpper(request.URL.Query().Get("sort_order"))
		if sortBy == "" {
			sortBy = "created_at"
		}
		if sortOrder != "ASC" && sortOrder != "DESC" {
			sortOrder = "DESC"
		}
		txErr := db.Transaction(func(tx *gorm.DB) error {
			if err := tx.Model(&models.Ad{}).Select("COUNT(*)").Count(&totalAdsCount).Error; err != nil {
				log.Error("Error counting ads:", err.Error())
				fileLogger.Println("Error counting ads:", err.Error())
				common.ErrorResponse(writer, http.StatusInternalServerError, "Something went wrong")
				return err
			}
			query := tx.Model(&models.Ad{}).Limit(pageSize).Offset(offset).Order(fmt.Sprintf("%s %s", sortBy, sortOrder))
			if sortBy == "price" {
				max := request.URL.Query().Get("max")
				min := request.URL.Query().Get("min")
				if min == "" || max == "" {
					tx.Rollback()
					common.ErrorResponse(writer, http.StatusBadRequest, "Min and max price are required")
					return errors.New("Min and max price are required")
				}
				if err := query.Where("price BETWEEN ? AND ?", min, max).Find(&ads).Error; err != nil {
					log.Error("Error fetching ads:", err.Error())
					fileLogger.Println("Error fetching ads:", err.Error())
					common.ErrorResponse(writer, http.StatusInternalServerError, "Something went wrong")
					return err
				}
			} else {
				if err := query.Find(&ads).Error; err != nil {
					log.Error("Error fetching ads:", err.Error())
					fileLogger.Println("Error fetching ads:", err.Error())
					common.ErrorResponse(writer, http.StatusInternalServerError, "Something went wrong")
					return err
				}
			}
			return nil
		})
		if txErr != nil {
			return
		}
		totalPages := int(math.Ceil(float64(totalAdsCount) / float64(pageSize)))
		writer.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(writer).Encode(map[string]interface{}{"data": ads, "total_pages": totalPages}); err != nil {
			log.Error("Error encoding JSON:", err.Error())
			fileLogger.Println("Error encoding JSON:", err.Error())
			common.ErrorResponse(writer, http.StatusInternalServerError, "Internal Server Error")
		}
	}
}
