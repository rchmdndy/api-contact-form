package models
// Package models defines the data models for the API Contact Form application.
//
// It includes the Contact struct, which represents a contact message submitted
// through the API. The struct is configured for use with GORM, an ORM library
// for Go, to handle database interactions.
//

import (
	"gorm.io/gorm"
	"time"
)

type Contact struct {
	ID uint `gorm:"primaryKey;column:id;type:BIGINT UNSIGNED AUTO_INCREMENT"`
	FullName string `gorm:"column:full_name;type:VARCHAR(100);not null"`
	Email string `gorm:"column:email_address;type:VARCHAR(100);not null"`
	Phone string `gorm:"column:phone_number;type:VARCHAR(20);not null"`
	Message string `gorm:"column:message_text;type:TEXT;not null"`
	CreatedAt time.Time `gorm:"column:created_at;type:DATETIME;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:DATETIME;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (Contact) TableName() string{
	return "contact_messages"
}