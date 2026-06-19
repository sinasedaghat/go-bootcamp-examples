package bookCity

import "fmt"

type Product struct {
	Name  string
	Price Price
}

func (p *Product) Print() {
	fmt.Printf("%q is priced at %s.\n", p.Name, p.Price)
}

func (p *Product) Discount(f float64) {
	p.Price = Price(float64(p.Price) * (100 - f) / 100)
	fmt.Printf("The price of %q after discount is %s ($%[2]f).\n", p.Name, p.Price)
}
