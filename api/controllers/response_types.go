package controllers

import (
	"time"

	"github.com/shuklarituparn/VK-Marketplace/api/models"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type CreateAdResponse struct {
	Ad      models.Ad `json:"ad"`
	Message string    `json:"message"`
}

type ReadAllAdsResponse struct {
	Data       []models.Ad `json:"data"`
	TotalPages int         `json:"total_pages"`
}

type ReadAdResponse struct {
	Data models.Ad `json:"data"`
}

type HealthCheckResponse struct {
	Author      string    `json:"author"`
	CurrentTime time.Time `json:"current_time"`
	Status      string    `json:"status"`
}

type SearchAdResponse struct {
	Data []models.Ad `json:"data"`
}

type RefreshTokenResponse struct {
	Token   string `json:"token"`
	Refresh string `json:"refresh"`
}
type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	Email        string `json:"email"`
	ID           int    `json:"id"`
	Message      string `json:"message"`
	RefreshToken string `json:"refresh_token"`
}

type CreateUserResponse struct {
	Email   string `json:"email"`
	ID      int    `json:"id"`
	Message string `json:"message"`
}
