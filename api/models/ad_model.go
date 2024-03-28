package models

import "time"

type Ad struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	UserID       uint      `json:"user_id"`
	Title        string    `json:"title" validate:"required,max=100" gorm:"unique"`
	AdText       string    `json:"ad_text" validate:"required,max=500"`
	ImageAddress string    `json:"image_address" validate:"required,url"`
	Price        float64   `json:"price" validate:"required"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type CreateAd struct {
	Title        string  `json:"title" validate:"required,max=100" gorm:"unique"`
	AdText       string  `json:"ad_text" validate:"required,max=500"`
	ImageAddress string  `json:"image_address" validate:"required,url"`
	Price        float64 `json:"price" validate:"required"`
}
