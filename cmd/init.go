package cmd

import "fmt"

type Coordinate struct {
	X int
	Y int
}

func (p *Coordinate) Move(X int, Y int) (int, int) {
	p.X += X
	p.Y += Y
	return p.X, p.Y
}

func (p *Coordinate) PrintCoordinte() {
	fmt.Printf("X: %d, Y: %d\n", p.X, p.Y)
}

