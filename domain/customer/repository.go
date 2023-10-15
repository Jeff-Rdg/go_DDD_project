package customer

import (
	"DDD_Project/aggregate"
	"errors"
	"github.com/google/uuid"
)

var (
	ErrCustomerNotFound    = errors.New("the customer was not found in the repository")
	ErrFailedToAddCustomer = errors.New("failed to add the customer")
	ErrUpdateCustomer      = errors.New("failed to update the customer")
)

type CustomerRepository interface {
	Get(id uuid.UUID) (aggregate.Customer, error)
	Add(customer aggregate.Customer) error
	Update(customer aggregate.Customer) error
}
