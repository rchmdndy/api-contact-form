package handlers

// Package handlers contains the HTTP handler implementations for various endpoints.
//
// Specifically, the MainHandler provides the root endpoint to verify
// that the API Contact Form is running correctly.
import (
	"api-contact-form/responses"
	"net/http"
	
	"github.com/gin-gonic/gin"
)

type MainHandler struct {
	
}

func NewMainHandler() *MainHandler {
	return &MainHandler{}
}


func (h *MainHandler) MainHandler (c *gin.Context){
	c.JSON(http.StatusOK, responses.APIResponse{
		Code : "SUCCESS",
		Message : "API Contact Form is running.",
	})
}
