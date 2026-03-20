package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	winter = 1
	summer = 3
)

func startArr() {
	var arr [2]string
	fmt.Println(arr)

	a := [4]string{}
	a[0] = "Zero"
	a[1] = "One"
	a[2] = "Two"
	a[3] = "Three"
	// i := 4
	// a[i] = "Four"

	fmt.Println(a)

	b := [2]bool{}

	fmt.Printf("%#v\n", b)

}

func bookstore() {
	// books := [winter + summer]string{}
	// books[0] = "The body never lies"
	// books[1] = "Atomic habits"
	// books[2] = "The compound effect"
	// books[3] = books[2] + "(2nd edition)"

	books := [...]string{
		"The body never lies",
		"Atomic habits",
		"The compound effect",
		"The compound effect (2nd edition)",
	}

	fmt.Printf("\nBooks: %#v\n", books)

	var (
		wBooks [winter]string
		sBooks [summer]string
	)

	wBooks[0] = books[0]
	fmt.Printf("\nWinter Books: %#v\n", wBooks)

	// for i := 0; i < len(sBooks); i++ {
	for i := range sBooks {
		sBooks[i] = books[i+1]
	}
	fmt.Printf("\nSummer Books: %#v\n", sBooks)

	var publishedBooks [len(books)]bool

	publishedBooks[0] = true
	publishedBooks[len(publishedBooks)-1] = true

	fmt.Println("\nPublished Books:")
	for i, ok := range publishedBooks {
		if ok {
			fmt.Printf("- %s\n", books[i])
		}
	}
}

func moodlyV0() {
	mood := [...]string{
		"%s walked with a spring in his step, humming a tune without realizing it. (😊 Happy)\n",
		"The shadows seemed to move when %s looked away. (😨 Fearful)\n",
		"%s’s voice trembled — even silence felt heavy around him. (😢 Sad)\n",
		"The door slammed behind %s before anyone could answer. (😡 Angry)\n",
		"%s kept replaying the scene, wishing his words could be taken back. (😔 Regretful)\n",
		"The world lost its colors for %s that morning. (😢 Sad)\n",
		"%s closed his eyes — not to escape, but to stay a little longer in that stillness. (😌 Calm)\n",
		"His eyes searched the ground — guilt written in every breath %s took. (😔 Regretful)\n",
		"Tears rolled down %s’s cheeks, but he made no sound. (😢 Sad)\n",
		"The wind brushed %s’s face, and everything felt right again. (😌 Calm)\n",
		"Even %s’s silence burned like fire. (😡 Angry)\n",
		"%s stared at the floor, trying to find answers between the tiles. (😢 Sad)\n",
		"Laughter spilled out of %s before he could hold it back. (😊 Happy)\n",
		"%s’s fists clenched so hard the knuckles turned white. (😡 Angry)\n",
		"%s’s heart thumped so loudly he thought others could hear it. (😨 Fearful)\n",
		"“I shouldn’t have said that,” %s whispered to no one. (😔 Regretful)\n",
		"The quiet around %s was not empty; it was comforting. (😌 Calm)\n",
		"%s’s eyes sparkled as if the sun itself lived behind them. (😊 Happy)\n",
		"%s couldn’t stop smiling — the day felt lighter than air. (😊 Happy)\n",
		"Every creak of the floor sounded to %s like footsteps approaching. (😨 Fearful)\n",
		"The apology sat heavy in %s’s chest, too late to fix what broke. (😔 Regretful)\n",
	}

	args := os.Args
	if len(args) < 2 {
		fmt.Println("What's your name?")
		return
	}

	n := rand.New(rand.NewSource(time.Now().UnixMicro())).Intn(len(mood))
	fmt.Printf("random number: %d\n", n)
	fmt.Printf(mood[n], args[1])
}

func compareArrV0() {
	var (
		blue = [...]int{3, 9, 6}
		red  = [3]int{3, 9, 6}
	)

	fmt.Printf("Blue Array: %#v\n", blue)
	fmt.Printf("Red Array : %#v\n", red)

	fmt.Printf("Blue array is equal to Red array? %t\n", blue == red)
}

func assignArr() {
	blue := [3]int{3, 9, 6}
	fmt.Printf("Blue Array: %#v\n", blue)

	var red1 [3]int
	red1 = blue
	fmt.Printf("Red1 Array: %#v\n\n", red1)

	blue[1] = 4
	fmt.Printf("Blue Array after seconde index value of Blue Array: %#v\n", blue)
	fmt.Printf("Red1 Array after seconde index value of Blue Array: %#v\n\n", red1)

	red1[2] = 1
	fmt.Printf("Blue Array after third index value of Red1 Array: %#v\n", blue)
	fmt.Printf("Red1 Array after third index value of Red1 Array: %#v\n\n", red1)

	fmt.Printf("//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\\n\n")

	fmt.Printf("Blue Array: %#v\n", blue)
	red2 := blue
	fmt.Printf("Red2 Array: %#v\n\n", red2)

	blue[1] = 6
	fmt.Printf("Blue Array after seconde index value of Blue Array: %#v\n", blue)
	fmt.Printf("Red2 Array after seconde index value of Blue Array: %#v\n\n", red2)

	red2[2] = 7
	fmt.Printf("Blue Array after third index value of Red2 Array: %#v\n", blue)
	fmt.Printf("Red2 Array after third index value of Red2 Array: %#v\n\n", red2)
}

func multiDimensionArr() {
	// students := [2][3]float64{
	// 	[3]float64{4, 8, 5},
	// 	[3]float64{8, 5, 9},
	// }
	// students := [2][3]float64{
	// 	{4, 8, 5},
	// 	{8, 5, 9},
	// }

	students := [...][3]float64{
		{4, 8, 5},
		{8, 5, 9},
	}

	var (
		sum float64
		n   int // = len(students) * len(students[0])
	)
	for _, grades := range students {
		for _, grade := range grades {
			sum += grade
			n++
		}
	}

	fmt.Printf("Average Grades: %g\n", sum/float64(n))
}

func moodlyV1() {
	moods := [...][4]string{
		{
			"%s walked with a spring in his step, humming a tune without realizing it. (😊 Happy)\n",
			"The wind brushed %s’s face, and everything felt right again. (😌 Calm)\n",
			"%s’s eyes sparkled as if the sun itself lived behind them. (😊 Happy)\n",
			"%s closed his eyes — not to escape, but to stay a little longer in that stillness. (😌 Calm)\n",
		},
		{
			"Every creak of the floor sounded to %s like footsteps approaching. (😨 Fearful)\n",
			"The apology sat heavy in %s’s chest, too late to fix what broke. (😔 Regretful)\n",
			"Even %s’s silence burned like fire. (😡 Angry)\n",
			"%s stared at the floor, trying to find answers between the tiles. (😢 Sad)\n",
		},
	}

	const (
		positive = 0
		negative = 1
	)

	args := os.Args
	if len(args) != 3 {
		fmt.Println("Enter [user name] [positive|negative]")
		return
	}

	var m int
	switch strings.ToLower(args[2]) {
	case "positive":
		m = positive
	case "negative":
		m = negative
	default:
		fmt.Println("Please Enter valid mood [positive|negative]")
		return
	}

	n := rand.New(rand.NewSource(time.Now().UnixMicro())).Intn(len(moods[0]))
	fmt.Printf(moods[m][n], args[1])
}

func keyedElementsV0() {
	arr1 := [...]string{
		1: "1",
		4: "4",
		"5",
		0: "0",
		7: "7",
		2: "2",
		"3",
	} // [8]string{"0", "1", "2", "3", "4", "5", "", "7"}
	fmt.Printf("Arr1 Array: %#v\n\n", arr1)

	arr2 := [...]string{
		1: "1",
		4: "4",
		0: "0",
		2: "2",
		7: "7",
		"8",
	} // Arr2 Array: [9]string{"0", "1", "2", "", "4", "", "", "7", "8"}
	fmt.Printf("Arr2 Array: %#v\n\n", arr2)
}

func cryptoExchangeRatios() {
	const (
		ETH = 9 - iota
		WAN
		ICX
	)

	rate := [...]float64{
		ETH: 25.5,
		WAN: 120.5,
		ICX: 20,
	}

	fmt.Printf("Rate array: %#v\n", rate)

	fmt.Printf("1 BTC is %g ETH\n", rate[ETH])
	fmt.Printf("1 BTC is %g WAN\n", rate[WAN])
	fmt.Printf("1 BTC is %g ICX\n", rate[ICX])
}

func compareArrV1() {
	type (
		integer int

		blue [3]int
		red  [3]int

		nBlue [3]integer
		nRed  [3]integer
	)
	var (
		blue0 = [3]int{3, 9, 6}
		red0  = [3]int{3, 9, 6}

		blue1 = blue{3, 9, 6}
		red1  = red{3, 9, 6}

		blue2 = nBlue{3, 9, 6}
		red2  = nRed{3, 9, 6}
	)

	fmt.Printf("blue0 Array: %#v\n", blue0)
	fmt.Printf("red0 Array : %#v\n\n", red0)

	fmt.Printf("blue1 Array: %#v\n", blue1)
	fmt.Printf("red1 Array : %#v\n\n", red1)

	fmt.Printf("red(blue1) Array: %#v\n", red(blue1))
	fmt.Printf("blue(red1) Array : %#v\n\n", blue(red1))

	fmt.Printf("blue2 Array: %#v\n", blue2)
	fmt.Printf("red2 Array : %#v\n\n", red2)

	fmt.Printf("blue0 == red0: %t\n\n", blue0 == red0)
	fmt.Printf("blue1 == red0: %t\n\n", blue1 == red0)
	// fmt.Printf("blue1 == red1: %t\n\n", blue1 == red1) // Compiler Error: mismatched types blue and red
	fmt.Printf("blue1 == blue(red1): %t\n\n", blue1 == blue(red1))
	// fmt.Printf("red(blue1) == blue(red1): %t\n\n", red(blue1) == blue(red1)) // Compiler Error: mismatched types red and blue
	// fmt.Printf("blue2 == red0: %t\n\n", blue2 == red0) // Compiler Error: mismatched types nBlue and [3]int
	// fmt.Printf("blue2 == blue1: %t\n\n", blue2 == blue1) // Compiler Error: mismatched types nBlue and blue
	fmt.Printf("blue2 == nBlue(red2): %t\n\n", blue2 == nBlue(red2))
	// fmt.Printf("blue2 == nBlue(blue0): %t\n\n", blue2 == nBlue(blue0)) // Compiler Error: cannot convert blue0 (variable of type [3]int) to type nBlue
	// fmt.Printf("[3]integer{} == [3]int{}: %t\n\n", [3]integer{} == [3]int{}) // Compiler Error: mismatched types [3]integer and [3]int
	fmt.Printf("[3]integer{} == nRed{}: %t\n\n", [3]integer{} == nRed{})
}
