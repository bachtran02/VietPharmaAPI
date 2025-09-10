package main

import (
	"fmt"
	"log"
	"vietpharma-api/internal"
	"vietpharma-api/internal/handlers"
	"vietpharma-api/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	/* Load config file */
	config, err := internal.ReadConfig("config.yml")
	if err != nil {
		log.Printf("Failed to read config: %v", err)
	}

	/* Initialize services */
	medicineService := services.NewMedicalProductService()

	/* Initialize handlers */
	medicineHandler := handlers.NewMedicalProductHandler(medicineService)

	/* Initialize Gin router */
	r := gin.Default()

	/* Define routes */
	r.GET("/search-product", medicineHandler.SearchProduct)

	/* Start server */
	log.Printf("Server starting on port %d", config.Server.Port)
	r.Run(fmt.Sprintf(":%d", config.Server.Port))
}
