package matrix

import (
	"fmt"
	"math/rand"
)

type Figure struct {
	vector uint64
	mW     uint8
	mH     uint8
}

func (f *Figure) Print() {
	var shift = uint64(1) << 63
	var i, j uint8
	for ; i < f.mH; i++ {
		j = 0
		fmt.Print("[")
		for ; j < f.mW; j++ {
			if f.vector&shift != 0 {
				fmt.Print(" # ")
			} else {
				fmt.Print(" _ ")
			}
			shift = shift >> 1
		}
		fmt.Print("]\n")
	}
}

func GenerateFigures(w, h uint8) []Figure {
	square := 0
	var figures []Figure
	for {
		x, y := rand.Uint32()%uint32(w)+1, rand.Uint32()%uint32(h)+1
		if x > uint32(w) || y > uint32(h) {
			fmt.Printf("[DEBUG] skip due oversize [%d | %d]", x, y)
			continue
		}
		square += int(x * y)
		if square > int(w*h) {
			return figures
		}
		vector := uint64(0)
		for i := uint32(0); i < y; i++ {
			shift := ^uint64(0)
			for j := uint32(0); j < x; j++ {
				shift = shift >> 1
			}
			shift = ^shift >> (i * uint32(w))
			vector += shift
		}
		if debug {
			fmt.Printf("[DEBUG] generated [w: %d | h: %d] figure with total square: %d\n", x, y, square)
			fmt.Printf("[DEBUG] vector: %064b\n", vector)
		}
		figures = append(figures, Figure{vector: vector, mW: w, mH: h})
	}
}
