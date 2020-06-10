package random

import (
	"math/rand"

	"github.com/leegeobuk/go-gRPC/proto/pc"
)

// KeyboardLayout returns one of keyboard layout
func KeyboardLayout() pc.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return pc.Keyboard_QWERTY
	case 2:
		return pc.Keyboard_QWERTZ
	default:
		return pc.Keyboard_AZERTY
	}
}

// Bool returns either true or false with random possibility
func Bool() bool {
	return rand.Intn(2) == 1
}
