package random

import (
	pc "github.com/leegeobuk/go-gRPC/pb"
	"math/rand"
)

// RandomKeyboardLayout returns one of keyboard layout
func RandomKeyboardLayout() pc.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return pc.Keyboard_QWERTY
	case 2:
		return pc.Keyboard_QWERTZ
	default:
		return pc.Keyboard_AZERTY
	}
}

// RandomBool returns either true or false with random possibility
func RandomBool() bool {
	return rand.Intn(2) == 1
}
