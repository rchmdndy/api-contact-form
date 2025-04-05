package repositories

// Package repositories provides implementations for data persistence and retrieval
// related to contact entities in the API Contact Form application.
//
// It defines the ContactRepository interface and its GORM-based implementation
// for performing CRUD operations on contact records in the database.

import (
	"api-contact-form/models"
	
	"gorm.io/gorm"
)

type ContactRepository interface {
	Create (contact *models.Contact) error
	FindAll () ([]models.Contact, error)
	FindByID (id uint) (*models.Contact, error)
	Update (contact *models.Contact) error
	Delete (contact *models.Contact) error
}

type contactRepository struct {
	db *gorm.DB
}

func NewContactRepository(db *gorm.DB) ContactRepository{
	return &contactRepository{db}
}

func (c contactRepository) Create(contact *models.Contact) error {
	return c.db.Create(contact).Error
}

func (c contactRepository) FindAll() ([]models.Contact, error) {
	var contacts []models.Contact
	err := c.db.Find(&contacts).Error
	return contacts, err
}

func (c contactRepository) FindByID(id uint) (*models.Contact, error) {
	var contact models.Contact
	err := c.db.First(&contact, id).Error
	return &contact, err
}

func (c contactRepository) Update(contact *models.Contact) error {
	return c.db.Save(contact).Error
}

func (c contactRepository) Delete(contact *models.Contact) error {
	return c.db.Delete(contact).Error
}