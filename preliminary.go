package main

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

func first() {
	dir, file := path.Split("parent/child/file")

	fmt.Println("dir: ", dir)
	fmt.Println("file: ", file)

	fmt.Println("os.Args: ", os.Args)

	var s rune = 'س'
	fmt.Println("rune type variable 'س': ", s)

	name := "سینا"
	println("len: ", len(name))
	println("utf8 rune: ", utf8.RuneCountInString(name))
}

func banger() {
	input := os.Args[1]
	l := utf8.RuneCountInString(input)
	s := strings.Repeat("!", l)
	msg := s + strings.ToUpper(input) + s

	fmt.Println("Banger", msg)
}

func constants() {
	const (
		monday = iota
		tuesday
		wednesday
		thursday
		friday
		saturday
		sunday
	)

	fmt.Println("monday number day: ", monday)
	fmt.Println("tuesday number day: ", tuesday)
	fmt.Println("wednesday number day: ", wednesday)
	fmt.Println("thursday number day: ", thursday)
	fmt.Println("friday number day: ", friday)
	fmt.Println("saturday number day: ", saturday)
	fmt.Println("sunday number day: ", sunday)

	const (
		EST = -5 - iota // -5
		_               // -6
		MST             // -7
		PST             // -8
	)

	fmt.Println("EST timezone: ", EST)
	fmt.Println("MST timezone: ", MST)
	fmt.Println("PST timezone: ", PST)

}

func print() {
	brand := "Google"
	fmt.Printf("Q: %q\n", brand)
	fmt.Printf("S: %s\n", brand)

	integer := 32
	fmt.Printf("D: %d\n", integer)

	float := 32.604
	fmt.Printf("F: %f\n", float)
	fmt.Printf("F: %.0f\n", float)
	fmt.Printf("F: %.1f\n", float)
	fmt.Printf("F: %.2f\n", float)
	fmt.Printf("F: %.3f\n", float)
	fmt.Printf("F: %.4f\n", float)
	fmt.Printf("F: %.5f\n", float)
	fmt.Printf("F: %.6f\n", float)

	boolean := true
	fmt.Printf("T: %t\n", boolean)

	fmt.Printf("Total: %d, Success: %d, Fail: %d\n", 3500, 3200, 300)

	var speed int
	fmt.Printf("T: %T\n", speed)

	fmt.Printf("V: %v\n", brand)
	fmt.Printf("V: %v\n", integer)

	var (
		planet   = "Venus"
		distance = 261
		orbital  = 224.701
		hasLife  = false
	)

	fmt.Printf(
		"Planet: %v\nDistance: %v\nOrbital: %v\nDose %v have life? %v\n",
		planet, distance, orbital, planet, hasLife,
	)

	fmt.Printf(
		"%[1]v is %d away. Think! %[2]d kms! %[1]s OMG!\n",
		planet, distance,
	)
}

func ifStatements() {
	score := 16.5

	if score >= 18 {
		fmt.Print("Excellent 🎉\n")
	} else if score > 16 {
		fmt.Print("Good 🎊\n")
	} else if score >= 12 {
		fmt.Printf("Pass ✅\n")
	} else {
		fmt.Print("Fail ❌\n")
	}

	if x := 3; x < 4 {
		fmt.Printf("%d (x) is lower than 4\n", x)
	} else {
		fmt.Printf("%d (x) is greater than 4\n", x)
	}
}

func authentication() {
	const (
		usage    = "Usage: [username] [password]\n"
		errUser  = "Access denied for %q\n"
		errPwd   = "Invalid password for %q\n"
		accessOK = "Access granted to %q\n"

		firstUser, secondUser = "sina", "saba"
		firstPwd, secondPwd   = "1372", "1374"
	)

	if input := os.Args; len(input) != 3 {
		fmt.Print(usage)
	} else if user, pwd := input[1], input[2]; (user != firstUser) &&
		(user != secondUser) {
		fmt.Printf(errUser, user)
	} else if ((user == firstUser) && (pwd != firstPwd)) ||
		((user == secondUser) && (pwd != secondPwd)) {
		fmt.Printf(errPwd, user)
	} else {
		fmt.Printf(accessOK, user)
	}
}

func errorHandle() {
	age := os.Args[1]

	n, err := strconv.Atoi(age)

	if err != nil {
		fmt.Println("Convert has Error: ", err)
		return
	}

	fmt.Printf("Yo 💥; Convert %q to %d was success\n", age, n)
}

func feetToMeter() {
	if len(os.Args) != 2 {
		fmt.Println("Fill input")
		return
	}

	feet := os.Args[1]

	n, err := strconv.ParseFloat(feet, 64)

	if err != nil {
		fmt.Printf("error: %q is not a number\n", feet)
		return
	}

	m := n * 0.3048
	fmt.Printf("%g feet is %g meters\n", n, m)
}

func shadow() {
	var (
		n int
		// err error
	)

	// if n, err = strconv.Atoi("42"); err == nil {
	if n, err := strconv.Atoi("42"); err == nil {
		fmt.Printf("n inside of if scope: %d\n", n)
	}
	fmt.Printf("n outside of if scope: %d\n", n)
}

func switchExpression() {
	city := "Mashhad"

	switch city {
	case "Tehran", "Mashhad":
		fmt.Println("Iran")
	case "Paris":
		fmt.Println("France")
	default:
		fmt.Println("Where?")
	}

	n := -2
	switch {
	case n > 0:
		fmt.Println("Positive")
	case n < 0:
		fmt.Println("Negative")
	default:
		fmt.Println("Zero")
	}

	n = 120
	switch {
	case n > 100:
		fmt.Print("Big ")
		fallthrough
	case n > 0:
		fmt.Print("Positive ")
		fallthrough
	default:
		fmt.Print("Number\n")
	}
}

func greeting() {
	switch h := time.Now().Hour(); {
	case h < 6:
		fmt.Println("Good night")
	case h < 12:
		fmt.Println("Good morning")
	case h < 17:
		fmt.Println("Good afternoon")
	case h < 21:
		fmt.Println("Good evening")
	default:
		fmt.Println("Good night")
	}
}

func forStatements() {
	factorial := 1
	for i := 1; i <= 5; i++ {
		factorial *= i
	}
	fmt.Printf("5 factorial equal to %d\n", factorial)

	factorial = 1
	for i := -3; i <= 5; i++ {
		if i <= 0 {
			continue
		}

		factorial *= i
	}
	fmt.Printf("5 factorial equal to %d\n", factorial)

	factorial = 1
	j := 1
	for {
		if j > 5 {
			break
		}
		factorial *= j
		j++
	}
	fmt.Printf("5 factorial equal to %d\n", factorial)
}

func multiplicationTable() {
	if len(os.Args) != 2 {
		fmt.Println("Please give me a number.")
		return
	}

	max, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		fmt.Println("Please give me a number.")
		return
	}

	fmt.Printf("%5s", "X")
	for i := 0; i <= int(max); i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for i := 0; i <= int(max); i++ {
		fmt.Printf("%5d", i)
		for j := 0; j <= int(max); j++ {
			fmt.Printf("%5d", i*j)
		}
		fmt.Println()
	}
}

func loopOnSlice() {
	words := strings.Fields("lazy cats jumps again and again and again.")

	for i := 0; i < len(words); i++ {
		fmt.Printf("#%-2d: %q\n", i+1, words[i])
	}
	fmt.Println()

	for i, v := range words {
		fmt.Printf("##%-2d: %q\n", i+1, v)
	}
	fmt.Println()

	// for i, v := range words {
	for i, v := range words[1:] {
		// if i==0 {
		// 	continue
		// }
		fmt.Printf("###%-2d: %q\n", i+1, v)
	}
	fmt.Println()

	for _, v := range words {
		fmt.Printf("####  : %q\n", v)
	}
	fmt.Println()

	// for i, _ := range words {
	for i := range words {
		fmt.Printf("#####%-2d: %q\n", i+1, words[i])
	}
}
