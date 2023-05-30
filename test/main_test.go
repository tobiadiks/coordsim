package main

import(
	"testing"
	"github.com/tobiadiks/coordsim/cmd"
)


// test for creating coordinate
func TestCreate(t *testing.T) {
	pv := &cmd.Coordinate{
		X: 0,
		Y: 0,
	}
	got1,got2 := pv.Move(0,0)
	want1,want2 := 0, 0
	if got1!=want1 && got2!=want2{
		t.Errorf("Expected %d %d, got %d %d", want1,want2, got1, got2)
	}
    
}

// test for moving coordinate
func TestMove(t *testing.T) {
	pv := &cmd.Coordinate{
		X: 3,
		Y: 8,
	}
	got1,got2 := pv.Move(2,5)
	want1,want2 := 5, 13
	if got1!=want1 && got2!=want2{
		t.Errorf("Expected %d %d, got %d %d", want1,want2, got1, got2)
	}
    
}