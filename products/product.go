package products

import "fmt"

type product struct {
	Title    string
	Price    price
	Released *timestamp
}

func (p *product) String() string {
	return fmt.Sprintf("The %q costs %f, and it was released in %v.", p.Title, p.Price, p.Released)
}

func (p *product) discount(dp float64) {
	p.Price = price(float64(p.Price) * (100 - dp) / 100)
	fmt.Printf("The price of %q after %f percent discount is %s.\n", p.Title, dp, p.Price)
}
