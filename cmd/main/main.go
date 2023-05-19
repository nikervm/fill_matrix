package main

import (
	"fmt"
	"math/rand"
	"matrix/matrix"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	x, y := rand.Uint32()%9+1, rand.Uint32()%9+1
	m := matrix.InitMatrix(uint8(x), uint8(y))
	m.Print()

	figures := matrix.GenerateFigures(m.GetSize())

	if figures == nil {
		return
	}
	fmt.Println("solving:")
	for _, fig := range figures {
		fmt.Println("put figure to matrix:")
		fig.Print()
		fmt.Println()
		m.Print()
		m.AddFigure(fig)
		fmt.Println()
	}
}
