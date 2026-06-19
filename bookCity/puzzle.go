package bookCity

import (
	"fmt"
	"strconv"
)

type Puzzle struct {
	Product
	Piece any // interface{}
}

func (p *Puzzle) PieceCounter() {
	switch v := p.Piece.(type) {
	case int:
		fmt.Printf("%q has %d piece.\n", p.Name, v)
	case float64, float32:
		fmt.Printf("%q has %.0f piece.\n", p.Name, v)
	case string:
		if n, err := strconv.Atoi(v); err == nil {
			fmt.Printf("%q has %.d piece.\n", p.Name, n)
		} else {
			fmt.Println("The number of pieces is unknown.")
		}
	default:
		fmt.Println("The number of pieces is unknown.")
	}
}
