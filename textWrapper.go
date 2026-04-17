package main

import (
	"errors"
	"fmt"
	"os"
	"unicode"
)

const (
	maxWidth = 40
)

func getText() (string, error) {
	args := os.Args[1:]
	if len(args) != 1 {
		return "", errors.New("Give me your text!")
	}
	return args[0], nil
}

func wrapper() {
	lw := 0
	text, err := getText()

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, r := range text {
		fmt.Printf("%c", r)

		switch lw++; {
		case lw > maxWidth && r != '\n' && unicode.IsSpace(r):
			fmt.Println()
			fallthrough
		case r == '\n':
			lw = 0
		}
	}
}
