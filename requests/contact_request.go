package requests

// Package requests defines the request payload structures for the API Contact Form application.
//
// It includes the ContactRequest struct, which represents the data required to create or update
// a contact message through the API.
//

type ContactRequest struct {
	Name string `json:"name" binding:"required,max=100"`
	Email string `json:"email" binding:"required,email,max=100"`
	Phone string `json:"phone" binding:"required,max=20"`
	Message string `json:"message" binding:"required"`
}