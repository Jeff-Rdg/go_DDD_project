package aggregate

import (
	"DDD_Project/entity"
	"DDD_Project/valueObject"
	"errors"
	"github.com/google/uuid"
)

var ErrInvalidPerson = errors.New("a customer has to have a valid name")

type Customer struct {
	person       *entity.Person
	products     []*entity.Item
	transactions []valueObject.Transaction
}

// NewCustomer is a factory to create a new customer aggregate
func NewCustomer(name string) (Customer, error) {
	var customer Customer
	if name == "" {
		return customer, ErrInvalidPerson
	}

	person := &entity.Person{
		ID:   uuid.New(),
		Name: name,
	}

	customer = Customer{
		person:       person,
		products:     make([]*entity.Item, 0),
		transactions: make([]valueObject.Transaction, 0),
	}

	return customer, nil
}

func (c *Customer) GetId() uuid.UUID {
	return c.person.ID
}

func (c *Customer) SetId(id uuid.UUID) {
	if c.person == nil {
		c.person = &entity.Person{}
	}
	c.person.ID = id
}

func (c *Customer) GetName() string {
	return c.person.Name
}

func (c *Customer) SetName(name string) {
	if c.person == nil {
		c.person = &entity.Person{}
	}
	c.person.Name = name
}
