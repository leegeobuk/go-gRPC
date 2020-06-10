package service

import (
	"errors"
	"sync"

	"github.com/leegeobuk/go-gRPC/proto/pc"
)

var (
	// ErrorExistingID is thrown when laptop ID already exists in the store
	ErrorExistingID error = errors.New("Id already exists")
)

// DB is an interface for storing laptops
type DB interface {
	Save(laptop *pc.Laptop) error
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

	if store.data[laptop.Id] != nil {
		return ErrorExistingID
	}

	store.data[laptop.Id] = laptop
	return nil
}
