package middleware

import (
	"github.com/shuklarituparn/VK-Marketplace/internal/logger"
	"github.com/shuklarituparn/VK-Marketplace/pkg/common"
	jwt "github.com/shuklarituparn/VK-Marketplace/pkg/jwt_token"
	"net/http"
	"strings"

	"github.com/charmbracelet/log"
)

var fileLogger = logger.SetupLogger()

func IsAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info("Request received:", r.Method, r.URL.Path)
		fileLogger.Println("Request received:", r.Method, r.URL.Path)
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			common.ErrorResponse(w, http.StatusUnauthorized, "Authorization header missing")
			return
		}
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 {
			common.ErrorResponse(w, http.StatusBadRequest, "Invalid authorization header format. It should be Bearer <token>")
			return
		}
		token := tokenParts[1]
		_, err := jwt.VerifyToken(token)
		if err != nil {
			common.ErrorResponse(w, http.StatusForbidden, "Token Not Valid")
			return
		}
		next.ServeHTTP(w, r)
	})
}
