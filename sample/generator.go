package sample

import (
	"math/rand"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"github.com/leegeobuk/go-gRPC/proto/pc"

	"github.com/leegeobuk/go-gRPC/sample/random"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// NewLaptop returns a new sample laptop
func NewLaptop() *pc.Laptop {
	brand := random.LaptopBrand()
	return &pc.Laptop{
		ID:       uuid.New().String(),
		Brand:    brand,
		Name:     random.LaptopName(brand),
		Cpu:      NewCPU(),
		Ram:      NewRAM(),
		Gpus:     []*pc.GPU{NewGPU()},
		Storages: []*pc.Storage{NewSSD(), NewHDD()},
		Screen:   NewScreen(),
		Keyboard: NewKeyboard(),
		Weight: &pc.Laptop_WeightKg{
			WeightKg: random.Float(1.0, 3.0),
		},
		PriceUsd:    random.Float(1500, 3000),
		ReleaseYear: uint32(random.Integer(2015, 2019)),
		UpdatedAt:   ptypes.TimestampNow(),
	}
}

// NewCPU returns new sample cpu
func NewCPU() *pc.CPU {
	brand := random.CPUBrand()
	cores := random.Integer(2, 8)
	minGhz := random.Float(2.0, 3.5)

	return &pc.CPU{
		Brand:   brand,
		Name:    random.CPUName(brand),
		Cores:   uint32(cores),
		Threads: uint32(random.Integer(cores, 12)),
		MinGhz:  minGhz,
		MaxGhz:  random.Float(minGhz, 5.0),
	}
}

// NewGPU returns new sample gpu
func NewGPU() *pc.GPU {
	brand := random.GPUBrand()
	minGhz := random.Float(1.0, 1.5)
	gpu := &pc.GPU{
		Brand:  brand,
		Name:   random.GPUName(brand),
		MinGhz: minGhz,
		MaxGhz: random.Float(minGhz, 2.0),
		Memory: &pc.Memory{
			Value: uint64(random.Integer(2, 6)),
			Unit:  pc.Memory_GIGABYTE,
		},
	}
	return gpu
}

// NewRAM returns new sample ram
func NewRAM() *pc.Memory {
	return &pc.Memory{
		Value: uint64(random.Integer(4, 64)),
		Unit:  pc.Memory_GIGABYTE,
	}
}

// NewSSD returns new sample ssd
func NewSSD() *pc.Storage {
	return &pc.Storage{
		Driver: pc.Storage_SDD,
		Memory: &pc.Memory{
			Value: uint64(random.Integer(128, 1024)),
			Unit:  pc.Memory_GIGABYTE,
		},
	}
}

// NewHDD returns new sample hdd
func NewHDD() *pc.Storage {
	return &pc.Storage{
		Driver: pc.Storage_HDD,
		Memory: &pc.Memory{
			Value: uint64(random.Integer(1, 6)),
			Unit:  pc.Memory_TERABYTE,
		},
	}
}

// NewScreen returns new sample screen
func NewScreen() *pc.Screen {
	return &pc.Screen{
		SizeInch:   float32(random.Float(13, 17)),
		Resolution: random.Resolution(),
		Panel:      random.Panel(),
		Multitouch: random.Bool(),
	}
}

// NewKeyboard returns new sample keyboard
func NewKeyboard() *pc.Keyboard {
	return &pc.Keyboard{
		Layout:  random.KeyboardLayout(),
		Backlit: random.Bool(),
	}
}
