package products

import "fmt"

type product struct {
	title    string
	price    price
	released timestamp
}

func (p *product) String() string {
	return fmt.Sprintf("The %q costs %f, and it was released in %v.", p.title, p.price, p.released)
}

func (p *product) discount(dp float64) {
	p.price = price(float64(p.price) * (100 - dp) / 100)
	fmt.Printf("The price of %q after %f percent discount is %s.\n", p.title, dp, p.price)
}
