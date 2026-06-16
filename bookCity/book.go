package bookCity

import "fmt"

type Book struct {
	Name  string
	Price Price
}

func (b Book) Print() {
	fmt.Printf("%q is priced at %s.\n", b.Name, b.Price)
}

func newBooks() []Book {
	holiday := Book{
		Name:  "Fundamentals of Physics",
		Price: 105.96,
	}

	aria := Book{
		Name:  "Introduction to Classical Mechanics",
		Price: 13.32,
	}

	pathria := Book{
		Name:  "Statistical Mechanics",
		Price: 70.39,
	}

	return []Book{
		holiday,
		aria,
		pathria,
	}
}
