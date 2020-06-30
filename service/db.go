package service

import (
	"errors"
	"fmt"
	"sync"

	"github.com/jinzhu/copier"
	"github.com/leegeobuk/go-gRPC/proto/pc"
)

var (
	// ErrorExistingID is thrown when laptop ID already exists in the store
	ErrorExistingID error = errors.New("Id already exists")
)

// DB is an interface for storing laptops
type DB interface {
	Save(laptop *pc.Laptop) error
	Find(id string) (*pc.Laptop, error)
}

// LaptopStore stores laptops
type LaptopStore struct {
	mutex sync.RWMutex
	data  map[string]*pc.Laptop
}

// NewLaptopStore returns a new LaptopStore
func NewLaptopStore() *LaptopStore {
	return &LaptopStore{
		data: make(map[string]*pc.Laptop),
	}
}

// Save saves the laptop to the store
func (store *LaptopStore) Save(laptop *pc.Laptop) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.data[laptop.ID] != nil {
		return ErrorExistingID
	}

	store.data[laptop.ID] = laptop
	return nil
}

// Find finds the laptop with given id
func (store *LaptopStore) Find(id string) (*pc.Laptop, error) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	laptop := store.data[id]
	if laptop == nil {
		return nil, nil
	}

	// deep copy
	other := new(pc.Laptop)
	err := copier.Copy(other, laptop)
	if err != nil {
		return nil, fmt.Errorf("cannot copy laptop data: %w", err)
	}

	return other, nil
}
