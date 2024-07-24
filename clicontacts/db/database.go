package db

import "github.com/msanatan/golangtutorials/clicontacts/models"

type Database interface {
	InitDB() error
	CreateContact(contact models.Contact) error
	// ListContacts() ([]models.Contact, error)
	// FindContact(id int) (models.Contact, error)
	// DeleteContact(id int) error
}
