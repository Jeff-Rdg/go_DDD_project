package memory

import (
	"DDD_Project/aggregate"
	"DDD_Project/domain/customer"
	"fmt"
	"github.com/google/uuid"
	"sync"
)

type MemoryRepository struct {
	customers map[uuid.UUID]aggregate.Customer
	sync.Mutex
}

func New() *MemoryRepository {
	return &MemoryRepository{
		customers: make(map[uuid.UUID]aggregate.Customer),
	}
}

func (mr *MemoryRepository) Get(id uuid.UUID) (aggregate.Customer, error) {
	cst, ok := mr.customers[id]
	if ok {
		return cst, nil
	}

	return cst, customer.ErrCustomerNotFound
}

func (mr *MemoryRepository) Add(cst aggregate.Customer) error {
	if mr.customers == nil {
		mr.Lock()
		mr.customers = map[uuid.UUID]aggregate.Customer{}
		mr.Unlock()
	}

	_, ok := mr.customers[cst.GetId()]
	if ok {
		return fmt.Errorf("customer already exists: %w", customer.ErrFailedToAddCustomer)
	}
	mr.Lock()
	mr.customers[cst.GetId()] = cst
	mr.Unlock()
	return nil
}

func (mr *MemoryRepository) Update(customer aggregate.Customer) error {
	return nil
}
