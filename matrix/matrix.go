package matrix

import (
	"fmt"
)

const (
	debug = true
)

type Matrix struct {
	w     uint8
	h     uint8
	array uint64
}

func InitMatrix(x, y uint8) *Matrix {
	if x == 0 || y == 0 || x*y > 64 {
		return nil
	}
	return &Matrix{w: x, h: y, array: 0}
}

func (m *Matrix) GetSize() (uint8, uint8) {
	return m.w, m.h
}

func (m *Matrix) AddFigure(a Figure) {
	end := uint64(1)
	shift := 0
	notFit := false
	vec := a.vector
	for {
		if debug {
			fmt.Printf("[DEBUG] array:\t%064b\n", m.array)
			fmt.Printf("[DEBUG] vector:\t%064b\n", vec)
		}
		if m.array&vec == 0 {
			break
		}
		if vec&end == 1 {
			notFit = true
			break
		}
		shift++
		vec = vec >> 1
	}
	if notFit {
		fmt.Printf("skip figure (doesn't fit): %064b\n", a.vector)
		return
	}
	m.array += vec
	fmt.Printf("added figure [%064b] with shift [%d]\n", a.vector, shift)
	m.Print()
}

func (m *Matrix) Print() {
	var shift = uint64(1) << 63
	var i, j uint8
	fmt.Printf("current matrix (w: %d, h: %d):\n", m.w, m.h)
	if debug {
		fmt.Printf("[DEBUG] %064b\n", m.array)
	}
	for ; i < m.h; i++ {
		j = 0
		fmt.Print("[")
		for ; j < m.w; j++ {
			if m.array&shift != 0 {
				fmt.Print(" # ")
			} else {
				fmt.Print(" _ ")
			}
			shift = shift >> 1
		}
		fmt.Print("]\n")
	}
}
