package main

import (
	"github.com/tobiadiks/coordsim/cmd"
)
func main() {
	pv := &cmd.Coordinate{
		X: 0,
		Y: 0,
	}
	pv.PrintCoordinte()
	pv.Move(3, 1)
	pv.PrintCoordinte()
	pv.Move(1, 0)
	pv.PrintCoordinte()

}
