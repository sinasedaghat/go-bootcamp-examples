package main

import (
	"fmt"
	"strings"
)

type price float64

func (p price) String() string {
	return fmt.Sprintf("$%.2f (USD)", p)
}

type book struct {
	name  string
	price price
}

func (b book) print() {
	fmt.Printf("%q is priced at %s.\n", b.name, b.price)
}

func (b book) wrongMethodDiscount(p float64) {
	b.price = price(float64(b.price)*(100-p)) / 100

	fmt.Printf("[wrongMethodDiscount method] The price of %q after discount is %s ($%[2]f).\n", b.name, b.price)
}

func (b *book) discount(p float64) {
	b.price = price(float64(b.price)*(100-p)) / 100

	fmt.Printf("[discount method] The price of %q after discount is %s ($%[2]f).\n", b.name, b.price)
}

type library []*book

func (l library) print() {
	for i, b := range l {
		fmt.Printf("%d book > ", i)
		b.print()
	}
}

type game struct {
	name  string
	price float64
}

func (g game) print() {
	fmt.Printf("%q is priced at %s.\n", g.name, price(g.price))
}

func (g game) score(s int) {
	fmt.Printf("%s's score is  %d\n", g.name, s)
}

type archive struct {
	game [1000000]game
}

func (a archive) show(n int) {
	fmt.Printf("%-2d show >>> %p\n", n, &a)
}

func (a *archive) showPointer(n int) {
	fmt.Printf("%-2d showPointer >>> %p\n", n, a)
}

func workWithBook() {
	holiday := book{
		name:  "Fundamentals of Physics",
		price: 105.96,
	}

	holiday.print()
	holiday.wrongMethodDiscount(10)
	fmt.Println("After call holiday.wrongMethodDiscount(10) >>>")
	holiday.print()
	holiday.discount(10)
	fmt.Println("After call holiday.discount(10) >>>")
	holiday.print()
}

func workWithGame() {
	COD := game{
		name:  "Call of Duty",
		price: 59.99,
	}

	COD.print()
	COD.score(90)

	rayman := game{
		name:  "Rayman® Legends",
		price: 29.99,
	}

	game.print(rayman)
	game.score(rayman, 98)
}

func startMethods() {
	workWithBook()
	fmt.Println(strings.Repeat("-", 100))
	workWithGame()
}

func checkPerformance() {
	// HINT: Run this with "time gor run ."
	var hugeArchive archive
	for i := range 10 {
		// hugeArchive.show(i)
		hugeArchive.showPointer(i)
	}
}

func workWithLibrary() {
	holiday := book{
		name:  "Fundamentals of Physics",
		price: 105.96,
	}

	aria := book{
		name:  "Introduction to Classical Mechanics",
		price: 13.32,
	}

	pathria := book{
		name:  "Statistical Mechanics",
		price: 70.39,
	}

	var books []*book

	books = append(books, &holiday, &aria, &pathria)
	myBooks := library(books)

	myBooks.print()
}
