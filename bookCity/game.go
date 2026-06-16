package bookCity

import "fmt"

type Game struct {
	Name  string
	Price Price
}

func (g *Game) Print() {
	fmt.Printf("%q is priced at %s.\n", g.Name, g.Price)
}

func (g *Game) Discount(p float64) {
	g.Price = Price(float64(g.Price) * (100 - p) / 100)
	fmt.Printf("The price of %q after discount is %s ($%[2]f).\n", g.Name, g.Price)
}

func newGames() []Game {
	COD := Game{
		Name:  "Call of Duty",
		Price: 59.99,
	}

	rayman := Game{
		Name:  "Rayman® Legends",
		Price: 29.99,
	}

	return []Game{
		COD,
		rayman,
	}
}
