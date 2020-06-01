package random

import (
	"math/rand"

	"github.com/leegeobuk/go-gRPC/pb/pc"
)

// Resolution returns random resolution
func Resolution() *pc.Screen_Resolution {
	height := Integer(1080, 4320)
	return &pc.Screen_Resolution{
		Width:  uint32(height * 16 / 9),
		Height: uint32(height),
	}
}

// Panel returns random panel
func Panel() pc.Screen_Panel {
	switch rand.Intn(2) {
	case 1:
		return pc.Screen_OLED
	default:
		return pc.Screen_IPS
	}
}
