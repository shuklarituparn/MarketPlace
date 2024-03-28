package main

import (
	"fmt"
	"github.com/shuklarituparn/VK-Marketplace/api/controllers"
	"github.com/shuklarituparn/VK-Marketplace/api/routes"
	"github.com/shuklarituparn/VK-Marketplace/config"
	"github.com/shuklarituparn/VK-Marketplace/internal/logger"
	"github.com/shuklarituparn/VK-Marketplace/internal/prometheus"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

func swaggerHandler(w http.ResponseWriter, r *http.Request) {
	filePath := filepath.Join(".", "docs", "swagger.json")
	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "Failed to read swagger.json", http.StatusInternalServerError)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Error(err.Error())
		}
	}(file)

	w.Header().Set("Content-Type", "application/json")
	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
}

// @title Marketplace API
// @version 1.0
// @description Marketplace Api Docs
// @contact.name API Support
// @contact.url https://github.com/shuklarituparn
// @contact.email rtprnshukla@gmail.com
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	prometheus.Init()

	err := godotenv.Load()

	var fileLogger = logger.SetupLogger()

	if err != nil {
		log.Error("Error loading .env file")
		fileLogger.Println("Error loading .env file")
	}

	config.ConnectDb()

	if os.Getenv("PORT") == "" {
		log.Error("PORT environment variable not set")
		fileLogger.Println("PORT environment variable not set")
		os.Exit(1)
	}

	ServerPort, _ := strconv.Atoi(os.Getenv("PORT"))

	port := ServerPort

	rootMux := http.NewServeMux()
	rootMux.Handle("/metrics", promhttp.Handler())
	rootMux.HandleFunc("/swagger.json", swaggerHandler)
	rootMux.HandleFunc("/", controllers.HealthCheck)

	routes.UserRouter(rootMux)
	routes.AdRouter(rootMux)

	rootMux.HandleFunc("/docs/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger.json"),
	))

	log.Info(fmt.Sprintf("Server started on port %d", port))
	fileLogger.Printf("Server started on port %d", port)
	c := cors.Default()

	serverError := http.ListenAndServe(fmt.Sprintf(":%d", port), c.Handler(rootMux))
	if serverError != nil {
		log.Error("Something went wrong while starting the server!")
		fileLogger.Println("Something went wrong while starting the server!")
		panic(serverError)
	} else {
		log.Info("Server started successfully!")
		fileLogger.Println("Server started successfully!")
	}
}
