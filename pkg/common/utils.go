package common

import (
	"encoding/json"
	"errors"
	"github.com/charmbracelet/log"
	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt"
	"github.com/shuklarituparn/VK-Marketplace/internal/logger"
	jwt2 "github.com/shuklarituparn/VK-Marketplace/pkg/jwt_token"
	"net/http"
	"os"
)

var fileLogger = logger.SetupLogger()

func ErrorResponse(w http.ResponseWriter, status int, message string) {
	log.Error(message)
	fileLogger.Println(message)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	errorMsg := map[string]string{"error": message}
	err := json.NewEncoder(w).Encode(errorMsg)
	if err != nil {
		log.Error("Error encoding JSON:", err.Error())
		fileLogger.Println("Error encoding JSON:", err.Error())
	}
}

func ValidateAndRespond(w http.ResponseWriter, v interface{}) bool {
	validate := validator.New()
	if err := validate.Struct(v); err != nil {
		errorsMap := make(map[string]interface{})
		for _, e := range err.(validator.ValidationErrors) {
			errorsMap[e.Field()] = e.Tag()
		}
		errJSON, _ := json.Marshal(errorsMap)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write(errJSON)
		if err != nil {
			ErrorResponse(w, http.StatusInternalServerError, "Server Error!")
		}
		return false
	}
	return true
}

func ExtractUserIDFromToken(tokenString string) (uint, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt2.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return 0, err
	}
	if !token.Valid {
		return 0, errors.New("invalid token")
	}
	claims, ok := token.Claims.(*jwt2.Claims)
	if !ok {
		return 0, errors.New("invalid claims format")
	}
	return claims.UserId, nil
}
