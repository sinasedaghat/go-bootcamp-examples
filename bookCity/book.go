package bookCity

import "fmt"

type Book struct {
	Product
}

func (b *Book) Print() {
	fmt.Print("Book ")
	b.Product.Print()
}

func newBooks() []Book {
	holiday := Book{
		Product{
			Name:  "Fundamentals of Physics",
			Price: 105.96,
		},
	}

	aria := Book{
		Product{
			Name:  "Introduction to Classical Mechanics",
			Price: 13.32,
		},
	}

	pathria := Book{
		Product{
			Name:  "Statistical Mechanics",
			Price: 70.39,
		},
	}

	return []Book{
		holiday,
		aria,
		pathria,
	}
}
