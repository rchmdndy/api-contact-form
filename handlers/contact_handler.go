package handlers

// Package handlers contains the HTTP handler implementations for managing contacts.
//
// It defines the ContactHandler struct, which provides methods to handle
// CRUD (Create, Read, Update, Delete) operations for contact entities.

import (
	"api-contact-form/requests"
	"api-contact-form/responses"
	"api-contact-form/services"
	"log"
	"net/http"
	"strconv"
	
	"github.com/gin-gonic/gin"
)

// ContactHandler handles HTTP requests related to contact operations.
type ContactHandler struct {
	service services.ContactService
}

// NewContactHandler creates a new instance of ContactHandler with the provided ContactService.
func NewContactHandler(service services.ContactService) *ContactHandler {
	return &ContactHandler{service}
}

// CreateContact handles the creation of a new contact.
//
// It expects a JSON payload matching the ContactRequest structure.
// Upon successful creation, it returns the created contact with a 201 status code.
// If there's an error in binding the request or creating the contact, it returns an appropriate error response.
//
// This function performs the following steps:
// 1. Binds the incoming JSON request to a ContactRequest struct.
// 2. Validates the request and returns a 400 Bad Request response if validation fails.
// 3. Calls the CreateContact method of the ContactService to create a new contact.
// 4. Returns a 500 Internal Server Error response if the service call fails.
// 5. Returns a 201 Created response with the created contact data if successful.
//
// Parameters:
// - c: *gin.Context - The Gin context for the request, which provides request and response handling.
//
// Example usage:
//  router.POST("/contacts", contactHandler.CreateContact)
func (h *ContactHandler) CreateContact(c *gin.Context) {
	var req requests.ContactRequest;
	
	if err := c.ShouldBindJSON(&req) ; err != nil{
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code : "BAD_REQUEST",
			Message : err.Error(),
			Data: nil,
		})
		return
	}
	log.Println(req)
	
	contact, err := h.service.CreateContact(&req);
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.APIResponse{
			Code :  "INTERNAL_SERVER_ERROR",
			Message : err.Error(),
			Data : nil,
		})
		return
	}
	
	c.JSON(http.StatusCreated, responses.APIResponse{
		Code: "CREATED",
		Message : "Contact created succesfully",
		Data : responses.ContactResponseFromModel(contact),
	})
}

// GetContacts retrieves all contacts.
//
// It interacts with the service layer to fetch all contact records.
// On success, it returns the list of contacts with a 200 status code.
// In case of an error, it responds with a 500 status code and an error message.
func (h *ContactHandler) GetContacts (c *gin.Context){
	contacts, err := h.service.GetAllContacts()
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.APIResponse{
			Code : "INTERNAL_SERVER_ERROR",
			Message : err.Error(),
			Data : nil,
		})
		
		return
	}
	
	var contactResponses []responses.ContactResponse
	
	for _, contact := range contacts{
		contactResponses = append(contactResponses, responses.ContactResponseFromModel(&contact))
	}
	
	c.JSON(http.StatusOK, responses.APIResponse{
		Code : "SUCCESS",
		Message : "Contacts retrieved successfully",
		Data : contactResponses,
	})
}

// GetContact retrieves a single contact by its ID.
//
// It expects the contact ID as a URL parameter.
// If the ID is invalid or the contact does not exist, it returns an appropriate error response.
// On success, it returns the contact details with a 200 status code.
func (h *ContactHandler) GetContact(c *gin.Context){
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil{
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code : "BAD_REQUEST",
			Message : "Invalid ID",
			Data : nil,
		})
		return
	}
	
	contact , err := h.service.GetContactByID(uint(id))
	if err != nil{
		c.JSON(http.StatusNotFound, responses.APIResponse{
			Code : "NOT_FOUND",
			Message : "Contact not found",
			Data : nil,
		})
		return
	}
	
	c.JSON(http.StatusOK, responses.APIResponse{
		Code : "SUCCESS",
		Message : "Contact retrieved successfully",
		Data : responses.ContactResponseFromModel(contact),
	})
}

// UpdateContact updates an existing contact by its ID.
//
// It expects the contact ID as a URL parameter and a JSON payload matching the ContactRequest structure.
// If the ID is invalid or the contact does not exist, it returns an appropriate error response.
// On successful update, it returns the updated contact with a 200 status code.
func (h *ContactHandler) UpdateContact (c *gin.Context){
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil{
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code  : "BAD_REQUEST",
			Message : "Invalid ID",
			Data : nil,
		})
		return
	}
	
	var req requests.ContactRequest
	
	if err := c.ShouldBindJSON(&req) ; err != nil{
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code : "BAD_REQUEST",
			Message : err.Error(),
			Data : nil,
		})
		return
	}
	
	contact, err := h.service.UpdateContact(uint(id), &req)
	
	if err != nil{
		c.JSON(http.StatusInternalServerError, responses.APIResponse{
			Code : "INTERNAL_SERVER_ERROR",
			Message : err.Error(),
			Data : nil,})
		return
	}
	
	c.JSON(http.StatusOK, responses.APIResponse{
		Code : "SUCCESS",
		Message : "Contact updated successfully",
		Data : responses.ContactResponseFromModel(contact),
	})
}

// DeleteContact removes a contact by its ID.
//
// It expects the contact ID as a URL parameter.
// If the ID is invalid or the contact does not exist, it returns an appropriate error response.
// On successful deletion, it returns a success message with a 200 status code.
func (h *ContactHandler) DeleteContact (c *gin.Context){
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil{
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code : "BAD_REQUEST",
			Message : "Invalid ID",
			Data : nil,
		})
	}
	
	if err = h.service.DeleteContact(uint(id)) ; err != nil{
		c.JSON(http.StatusInternalServerError, responses.APIResponse{
			Code : "INTERNAL_SERVER_ERROR",
			Message : err.Error(),
			Data : nil,
		})
		return
	}
	
	c.JSON(http.StatusOK, responses.APIResponse{
		Code : "SUCCESS",
		Message : "Contact deleted successfully",
		Data : nil,
	})
}