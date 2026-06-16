package bookCity

import (
	"fmt"
	"strconv"
)

type Puzzle struct {
	Name  string
	Price Price
	Piece any // interface{}
}

func (p *Puzzle) Print() {
	fmt.Printf("%q is priced at %s.\n", p.Name, p.Price)
}

func (p *Puzzle) Discount(percent float64) {
	p.Price = Price(float64(p.Price) * (100 - percent) / 100)
	fmt.Printf("The price of %q after discount is %s ($%[2]f).\n", p.Name, p.Price)
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
