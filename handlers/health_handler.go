package handlers

// Package handlers contains the HTTP handler implementations for various endpoints.
//
// Specifically, the HealthHandler provides a health check endpoint to verify
// that the API is running correctly.

import (
	"api-contact-form/responses"
	"net/http"
	
	"github.com/gin-gonic/gin"
)

type HealthHandler struct {

}


func NewHealthHandler() *HealthHandler{
	return &HealthHandler{}
}

func (h *HealthHandler) HealthCheck (c *gin.Context){
	c.JSON(http.StatusOK, responses.APIResponse{
		Code : "SUCCESS",
		Message : "API is running",
	})
}