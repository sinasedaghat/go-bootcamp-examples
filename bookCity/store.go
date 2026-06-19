package bookCity

import "fmt"

type Printer interface {
	Print()
}

type Discounter interface {
	Discount(float64)
}

func Products() {
	books := newBooks()
	games := newGames()
	bl := len(books)
	gl := len(games)
	// fmt.Printf("books >> %#v\n\n games >> %#v\n", books, games)
	var products []Printer
	products = make([]Printer, bl+gl)
	for i := range bl {
		products[i] = &books[i]
	}
	for i := range gl {
		products[i+bl] = &games[i]
	}

	shelf(products)
}

func shelf(p []Printer) {
	for _, v := range p {
		v.Print()
	}
}

func Rebate() {
	bob := Puzzle{
		Product: Product{
			Name:  "Spongebob",
			Price: 49.99,
		},
	}

	var pBob Printer = &bob
	pBob.Print()

	if dBob, ok := pBob.(Discounter); ok {
		dBob.Discount(30)
	}

	holiday := Book{
		Product{
			Name:  "Fundamentals of Physics",
			Price: 105.96,
		},
	}

	var a any = holiday

	if d, ok := a.(Discounter); ok {
		d.Discount(30)
	} else {
		fmt.Println("does not implement Discounter")
	}
}

func PuzzleShelf() {
	starryNight := Puzzle{
		Product: Product{
			Name:  "The Starry Night",
			Price: 59.99,
		},
		Piece: "1000",
	}

	monaLisa := Puzzle{
		Product{
			Name:  "Mona Lisa",
			Price: 39.99,
		},
		500,
	}

	puzzles := [...]Puzzle{starryNight, monaLisa}

	for _, p := range puzzles {
		p.Print()
		p.PieceCounter()
		fmt.Println()
	}
}
