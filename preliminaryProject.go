package main

import (
	"fmt"
	"math/rand"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

const (
	maxTurn = 5
	usage   = `Welcome to the Lucky Number Game! 🍀

The program will pick %d random numbers.
Your mission is to guess one of those numbers.

The greater your number is, harder it gets.

Wanna play?
`
	giveNumber  = "Give me a positive number"
	give2Number = "Give me 2 positive numbers"
	win1st      = "🎉 You Win at first Try! 🎊"
	win         = "🎉 You Win!"
	awesome     = "🎉 You're Awesome!"
	perfect     = "🎉 Perfect!"
	loos        = "☠️ You Loos!"
	badLuck     = "☠️ Just a Bad Luck..."
)

const (
	corpus = "lazy cats jumps again and again and again."
)

func luckyNumberV0() {
	args := os.Args
	if len(args) != 2 {
		fmt.Printf(usage, maxTurn)
		return
	}

	gestes, err := strconv.Atoi(args[1])
	if err != nil || gestes < 0 {
		fmt.Println(giveNumber)
		return
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for t := 0; t <= maxTurn; t++ {
		n := r.Intn(gestes + 1)

		if n == gestes {
			fmt.Println(win)
			return
		}
	}

	fmt.Println(loos)
}

func luckyNumberV1() {
	args := os.Args
	if len(args) != 2 {
		fmt.Printf(usage, maxTurn)
		return
	}

	gestes, err := strconv.Atoi(args[1])
	if err != nil || gestes < 0 {
		fmt.Println(giveNumber)
		return
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for t := 1; t <= maxTurn; t++ {
		n := r.Intn(gestes + 1)

		if n == gestes && t == 1 {
			fmt.Println(win1st)
			return
		} else if n == gestes {
			fmt.Println(win)
			return
		}
	}

	fmt.Println(loos)
}

func luckyNumberV2() {
	args := os.Args
	if len(args) != 2 {
		fmt.Printf(usage, maxTurn)
		return
	}

	gestes, err := strconv.Atoi(args[1])
	if err != nil || gestes < 0 {
		fmt.Println(giveNumber)
		return
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for t := 0; ; {
		if t > maxTurn {
			break
		}

		t++
		n := r.Intn(gestes + 1)

		if n == gestes {
			switch r.Intn(3) {
			case 0:
				fmt.Println(win)
			case 1:
				fmt.Println(awesome)
			case 2:
				fmt.Println(perfect)
			}
			return
		}
	}

	switch r.Intn(2) {
	case 0:
		fmt.Println(loos)
	case 1:
		fmt.Println(badLuck)
	}
}

func luckyNumberV3() {
	args := os.Args
	if len(args) != 3 {
		fmt.Printf(usage, maxTurn)
		return
	}

	gestes, err := strconv.Atoi(args[1])
	if err != nil || gestes < 0 {
		fmt.Println(give2Number)
		return
	}

	gestes2, err := strconv.Atoi(args[2])
	if err != nil || gestes2 < 0 {
		fmt.Println(give2Number)
		return
	}

	max := slices.Max([]int{gestes, gestes2})
	n := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(max + 1)
	fmt.Printf("Random number is %d\n", n)

	if n == gestes || n == gestes2 {
		fmt.Println(win)
	} else {
		fmt.Println(loos)
	}
}

func wordFinderV0() {
	words := strings.Fields(corpus)

	queries := os.Args[1:]

	for _, q := range queries {
		for i, w := range words {
			// if strings.ToLower(q) == strings.ToLower(w) {
			if strings.EqualFold(q, w) {
				fmt.Printf("#%-2d: %q\n", i+1, w)
				break
			}
		}
	}
}

func wordFinderV1() {
	words := strings.Fields(corpus)

	queries := os.Args[1:]

queries:
	for _, q := range queries {
		for i, w := range words {
			if strings.EqualFold(q, w) {
				fmt.Printf("#%-2d: %q\n", i+1, w)
				break queries
			}
		}
	}
}

func wordFinderV2() {
	words := strings.Fields(corpus)

	queries := os.Args[1:]

queries:
	for _, q := range queries {
		for i, w := range words {
			if strings.EqualFold(q, w) {
				fmt.Printf("#%-2d: %q\n", i+1, w)
				continue queries
			}
		}
	}
}

func wordFinderV3() {
	words := strings.Fields(corpus)

	queries := os.Args[1:]

queries:
	for _, q := range queries {
	search:
		for i, w := range words {
			if strings.EqualFold(q, w) {
				switch strings.ToLower(w) {
				case "and", "or":
					break search
				}
				fmt.Printf("#%-2d: %q\n", i+1, w)
				continue queries
			}
		}
	}
}

func gotoLoop() {
	var i int
loop:
	if i < maxTurn {
		i++
		fmt.Printf("Turn %d\n", i)
		goto loop
	}

	fmt.Println("End of loop.")
}
