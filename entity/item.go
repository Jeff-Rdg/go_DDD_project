package entity

import "github.com/google/uuid"

// Item is an entity that represents a Item in all domains
type Item struct {
	ID          uuid.UUID
	Name        string
	Description string
}
