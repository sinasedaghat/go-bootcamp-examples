package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

// VERSION 1
type placeholder [5]string

var (
	zero = placeholder{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}

	one = placeholder{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}

	two = placeholder{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}

	three = placeholder{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}

	four = placeholder{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}

	five = placeholder{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}

	six = placeholder{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}

	seven = placeholder{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}

	eight = placeholder{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}

	nine = placeholder{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}

	colon = placeholder{
		"   ",
		" ░ ",
		"   ",
		" ░ ",
		"   ",
	}

	digits = [...]placeholder{
		0: zero,
		1: one,
		2: two,
		3: three,
		4: four,
		5: five,
		6: six,
		7: seven,
		8: eight,
		9: nine,
	}
)

func now() ([8]placeholder, int) {
	now := time.Now()

	hour := now.Hour()
	min := now.Minute()
	sec := now.Second()

	return [8]placeholder{
		digits[hour/10], digits[hour%10],
		colon,
		digits[min/10], digits[min%10],
		colon,
		digits[sec/10], digits[sec%10],
	}, sec
}

func clock() {
	screen.Clear()
	for {
		screen.MoveTopLeft()
		now, sec := now()
		for r := range len(digits[0]) {
			for _, digit := range now {
				if digit == colon && sec%2 == 0 {
					fmt.Print("    ")
					continue
				}
				fmt.Printf("%s ", digit[r])
			}
			fmt.Println()
		}
		fmt.Println()
		time.Sleep(time.Second)
	}
}

// VERSION 0
var (
	retroNumbers = [10][5]string{
		0: {
			0: "███",
			1: "█ █",
			2: "█ █",
			3: "█ █",
			4: "███",
		},
		1: {
			0: "██ ",
			1: " █ ",
			2: " █ ",
			3: " █ ",
			4: "███",
		},
		2: {
			0: "███",
			1: "  █",
			2: "███",
			3: "█  ",
			4: "███",
		},
		3: {
			0: "███",
			1: "  █",
			2: "███",
			3: "  █",
			4: "███",
		},
		4: {
			0: "█ █",
			1: "█ █",
			2: "███",
			3: "  █",
			4: "  █",
		},
		5: {
			0: "███",
			1: "█  ",
			2: "███",
			3: "  █",
			4: "███",
		},
		6: {
			0: "███",
			1: "█  ",
			2: "███",
			3: "█ █",
			4: "███",
		},
		7: {
			0: "███",
			1: "  █",
			2: "  █",
			3: "  █",
			4: "  █",
		},
		8: {
			0: "███",
			1: "█ █",
			2: "███",
			3: "█ █",
			4: "███",
		},
		9: {
			0: "███",
			1: "█ █",
			2: "███",
			3: "  █",
			4: "███",
		},
	}
	separator = [5]string{
		0: "   ",
		1: " ░ ",
		2: "   ",
		3: " ░ ",
		4: "   ",
	}
)

func retroClock() {
	screen.Clear()
	for {
		screen.MoveTopLeft()
		_createClock()
		time.Sleep(time.Second)
	}
}

func _createClock() {
	fmt.Printf("\f")
	t := time.Now()
	clock := [...][2]int{
		_convertIntToArr(t.Hour()),
		_convertIntToArr(t.Minute()),
		_convertIntToArr(t.Second()),
	}

	for r := range 5 {
		for i, c := range clock {
			fmt.Printf("%s %s", retroNumbers[c[0]][r], retroNumbers[c[1]][r])
			if i != 2 {
				fmt.Printf(" %s ", separator[r])
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func _convertIntToArr(n int) [2]int {
	return [2]int{
		n / 10,
		n % 10,
	}
}
