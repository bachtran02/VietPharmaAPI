package handlers

import (
	"net/http"
	"vietpharma-api/internal/services"

	"github.com/gin-gonic/gin"
)

// MedicalProductHandler handles HTTP requests for medical products
type MedicalProductHandler struct {
	productService services.MedicalProductService
}

// NewMedicalProductHandler creates a new instance of MedicalProductHandler
func NewMedicalProductHandler(productService services.MedicalProductService) *MedicalProductHandler {
	return &MedicalProductHandler{
		productService: productService,
	}
}

// SearchProduct handles the search request for medical products
func (h *MedicalProductHandler) SearchProduct(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Search query is required",
		})
		return
	}

	results, err := h.productService.SearchProduct(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"query":   query,
		"results": results,
	})
}
