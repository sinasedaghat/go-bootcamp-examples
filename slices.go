package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	s "github.com/inancgumus/prettyslice"
)

var cities = []string{
	"Tehran", "Mashhad", "Isfahan", "Shiraz",
	"Tabriz", "Ahvaz", "Karaj", "Qom", "Kermanshah",
	"Rasht", "Urmia", "Yazd", "Kerman", "Bandar Abbas",
	"Zanjan", "Arak", "Hamedan", "Sanandaj", "Sari", "Qazvin",
}

const costs string = `23 78
100 365 98 100
73 65 87 98
89
99 100 250
`

func startSlice() {
	var sV0 []int
	fmt.Printf("var sV0 []int:\n	sV0: %[1]v; sV0(with type): %#[1]v; sV0 == nil: %t, len(sV0): %d\n", sV0, sV0 == nil, len(sV0))

	var sV1 []int = []int{}
	fmt.Printf("var sV1 []int = []int{}:\n	sV1: %[1]v; sV1(with type): %#[1]v; sV1 == nil: %t, len(sV1): %d\n", sV1, sV1 == nil, len(sV1))

	sV2 := []int{}
	fmt.Printf("sV2 := []int{}:\n	sV2: %[1]v; sV2(with type): %#[1]v; sV2 == nil: %t, len(sV2): %d\n", sV2, sV2 == nil, len(sV2))

	sV2[0] = 23 // Error Runtime: panic: runtime error: index out of range [0] with length 0
	fmt.Printf("sV2: %#[1]v", sV2)
}

func compareSlice() {
	// var sV0 []string // sV0 := []string{}
	// var sV1 []string // sV1 := []string{}
	// fmt.Printf("sV0 == sV1: %t\n", sV0 ==	sV1) // Error Compiler : sV0 == sV1 (slice can only be compared to nil)

	sV0 := []string{
		"zero", "one", "two",
	}
	fmt.Printf("sV0: %#v\n", sV0)

	sV1 := []string{}
	fmt.Printf("sV1: %#v\n", sV1)

	sV1 = sV0
	fmt.Printf("after sV1 = sV0 statement; sV0: %#v\n", sV0)
	fmt.Printf("after sV1 = sV0 statement; sV1: %#v\n", sV1)

	sV0[2] = "new0 two"
	fmt.Printf("after sV0[2] = \"new0\" two statement; sV0: %#v\n", sV0)
	fmt.Printf("after sV0[2] = \"new0\" two statement; sV1: %#v\n", sV1)

	sV1[1] = "new1 one"
	fmt.Printf("after sV1[1] = \"new1 one\" statement; sV0: %#v\n", sV0)
	fmt.Printf("after sV1[1] = \"new1 one\" statement; sV1: %#v\n", sV1)

	sV2 := []string{
		"zero", "one", "two", "three", "four",
	}
	fmt.Printf("sV0: %#v\n", sV0)
	fmt.Printf("sV1: %#v\n", sV1)
	fmt.Printf("sV2: %#v\n", sV2)

	sV0 = sV2
	fmt.Printf("after sV0 = sV2 statement; sV0: %#v\n", sV0)
	fmt.Printf("after sV0 = sV2 statement; sV1: %#v\n", sV1)
	fmt.Printf("after sV0 = sV2 statement; sV2: %#v\n", sV2)

	// sV0 = nil
	// sV0 = []string{}
	// sV0 = sV1
	// sV2 = sV1
	// sV2 = nil // Runtime Error: panic: runtime error: index out of range [0] with length 0
	var equal string
	if len(sV0) != len(sV2) {
		equal = "not "
	}
	for i := range sV0 {
		if sV0[i] != sV2[i] {
			equal = "not "
			break
		}
	}

	fmt.Printf("\nsV0: %#v\nsV2: %#v\nsV0 is %sequal to sV2\n", sV0, sV2, equal)
}

func generateUnixNumber() {
	r := rand.New(rand.NewSource(time.Now().UnixMilli()))

	max, _ := strconv.Atoi(os.Args[1])
	unixNumbers := []int{}

loop:
	for len(unixNumbers) < max {
		n := r.Intn(max)
		fmt.Printf("%d  ", n)

		for _, u := range unixNumbers {
			if u == n {
				continue loop
			}
		}

		unixNumbers = append(unixNumbers, n)
	}

	fmt.Printf("\nunix numbers: %#v\n", unixNumbers)

	sort.Ints(unixNumbers)
	fmt.Printf("\nsorted unix numbers: %#v\n", unixNumbers)

	nums := [5]int{4, 6, 9, 1, 3}
	sort.Ints(nums[:])
	fmt.Printf("\nsorted numbers: %#v\n", nums)
}

func appendSlice() {
	sV0 := []int{0, 1, 2, 3}
	sV1 := append(sV0, 5)

	fmt.Printf("sV0: %#v\n", sV0)
	fmt.Printf("sV1: %#v\n", sV1)

	sV0[3] = 33

	fmt.Printf("after sV0[3] = 33; sV0: %#v\n", sV0)
	fmt.Printf("after sV0[3] = 33; sV1: %#v\n", sV1)

	sV2 := append(sV0) // reference assignment
	sV3 := sV0
	sV0[3] = 3

	fmt.Printf("after sV0[3] == 3; sV0: %#v\n", sV0)
	fmt.Printf("after sV0[3] == 3; sV2: %#v\n", sV2)
	fmt.Printf("after sV0[3] == 3; sV3: %#v\n", sV3)

	s := []int{0, 1, 2, 3, 4}
	fmt.Printf("s: %#v\n", s)

	s = append(s, 5, 6, 7)
	fmt.Printf("after s = append(s, 5, 6, 7); s: %#v\n", s)

	s = append(s, sV1...)
	fmt.Printf("after s = append(s, sV1...); s: %#v\n", s)

	sV1[1] = 11
	fmt.Printf("after sV1[1] = 11; s: %#v\n", s)
}

func sliceExpressionsV0() {
	fmt.Printf("These are Iran's cities: %#v | (len: %d)\n", cities, len(cities))
	top4 := cities[:4]
	// top4 := cities[0:4]
	fmt.Printf("\ntop 4 of Iran's cities: %#v | (len: %d)\n", top4, len(top4))

	last4 := cities[len(cities)-4:]
	// last4 := cities[len(cities)-4 : len(cities)]
	fmt.Printf("\nlast 4 of Iran's cities: %#v | (len: %d)\n", last4, len(last4))

	// wrong := cities[4:25] // Runtime Error: slice bounds out of range [:25] with capacity 20

	penultimateCityV0 := cities[len(cities)-2 : len(cities)-1]
	penultimateCityV1 := last4[2:3]
	fmt.Printf("\npenultimate Iran's city (slice);\nV0: %#v\nV1: %#v\n", penultimateCityV0, penultimateCityV1)
	fmt.Printf("\npenultimate Iran's city (index);\nV0: %#v\nV1: %#v\n", cities[len(cities)-2], last4[2])
}

func sliceExpressionsV1() {
	var target []int = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	middle := target[4:7]
	fmt.Printf("\nafter middle := target[4:7];\ntarget slice(%[3]d length): %#[1]v\nmiddle slice(%[4]d length): %#[2]v\n", target, middle, len(target), len(middle))

	target[5] = 55
	fmt.Printf("\nafter target[5] = 55;\ntarget slice(%[3]d length): %#[1]v\nmiddle slice(%[4]d length): %#[2]v\n", target, middle, len(target), len(middle))

	middle[2] = 62
	fmt.Printf("\nafter middle[2] = 62;\ntarget slice(%[3]d length): %#[1]v\nmiddle slice(%[4]d length): %#[2]v\n", target, middle, len(target), len(middle))

	midMiddle := middle[1:2]
	fmt.Printf("\nafter midMiddle := middle[1: 3];\ntarget slice(%[3]d length): %#[1]v\nmiddle slice(%[4]d length): %#[2]v\nmiddle slice(%[6]d length): %#[5]v\n", target, middle, len(target), len(middle), midMiddle, len(midMiddle))

	midMiddle[0] = 554
	fmt.Printf("\nafter midMiddle[0] = 554;\ntarget slice(%[3]d length): %#[1]v\nmiddle slice(%[4]d length): %#[2]v\nmiddle slice(%[6]d length): %#[5]v\n", target, middle, len(target), len(middle), midMiddle, len(midMiddle))

	midMiddle = append(midMiddle, 623)
	fmt.Printf("\nafter midMiddle = append(midMiddle, 623);\ntarget slice(%[3]d length): %#[1]v\nmiddle slice(%[4]d length): %#[2]v\nmiddle slice(%[6]d length): %#[5]v\n", target, middle, len(target), len(middle), midMiddle, len(midMiddle))

	middle = append(middle, 72)
	fmt.Printf("\nafter middle = append(middle, 72);\ntarget slice(%[3]d length): %#[1]v\nmiddle slice(%[4]d length): %#[2]v\n", target, middle, len(target), len(middle))

	middle = middle[:len(middle)-1]
	fmt.Printf("\nafter middle = append(middle, 72);\ntarget slice(%[3]d length): %#[1]v\nmiddle slice(%[4]d length): %#[2]v\n", target, middle, len(target), len(middle))

	middle = append(middle, 722)
	fmt.Printf("\nafter middle = append(middle, 722);\ntarget slice(%[3]d length): %#[1]v\nmiddle slice(%[4]d length): %#[2]v\n", target, middle, len(target), len(middle))

	last := target[len(target)-3:]
	fmt.Printf("\nafter last := target[len(target)-3:];\ntarget slice(%[3]d length): %#[1]v\nlast slice(%[4]d length): %#[2]v\n", target, last, len(target), len(last))

	last = append(last, 113)
	fmt.Printf("\nafter last = append(last, 113);\ntarget slice(%[3]d length): %#[1]v\nlast slice(%[4]d length): %#[2]v\n", target, last, len(target), len(last))

	target = append(target, 11)
	fmt.Printf("\nafter target = append(target, 11);\ntarget slice(%[3]d length): %#[1]v\nlast slice(%[4]d length): %#[2]v\n", target, last, len(target), len(last))

	last[3] = 1133
	fmt.Printf("\nafter last[3] = 1133;\ntarget slice(%[3]d length): %#[1]v\nlast slice(%[4]d length): %#[2]v\n", target, last, len(target), len(last))

	last[0] = 83
	fmt.Printf("\nafter last[0] = 83;\ntarget slice(%[3]d length): %#[1]v\nlast slice(%[4]d length): %#[2]v\n", target, last, len(target), len(last))

	most := target[7:10]
	fmt.Printf("\nafter most := target[7: 10];\ntarget slice(%[3]d length): %#[1]v\nmost slice(%[4]d length): %#[2]v\n", target, most, len(target), len(most))

	most = append(most, 105, 115)
	fmt.Printf("\nafter most = append(most, 105, 115);\ntarget slice(%[3]d length): %#[1]v\nmost slice(%[4]d length): %#[2]v\n", target, most, len(target), len(most))

	most[1] = 85
	fmt.Printf("\nafter most[1] = 85;\ntarget slice(%[3]d length): %#[1]v\nmost slice(%[4]d length): %#[2]v\n", target, most, len(target), len(most))
}

func pagination() {
	const pageSize = 6
	max := len(cities)

	for from := 0; from < max; from += pageSize {
		page := fmt.Sprintf("page: #%d", (from/pageSize)+1)
		// fmt.Printf("from: %d\n%s\n", from, page)

		to := from + pageSize
		l := len(cities)
		if to > l {
			to = l
		}

		c := cities[from:to]
		fmt.Printf("%s ==> (len %d)%#v\n", page, len(c), c)
	}
}

func backingArray() {
	fmt.Println("⚠️ all pointer recommended pointer of slice in this function not valid for this duty we was use `unsafe.SliceData(slice)` to read pointer value in `sliceHeader`")
	ba := [...]int{22, 44, 77, 11, 88, 33, 66, 99}

	s1 := ba[:]

	fmt.Printf(
		"after s1 := ba[:]\nbacking array (len %d  ptr: %p): %#v\nfirst slice   (len %d  ptr: %p): %#v\n",
		len(ba), &ba, ba, len(s1), &s1, s1,
	)

	s1[0] = 223
	fmt.Printf(
		"\nafter s1[0] = 223\nbacking array (len %d  ptr: %p): %#v\nfirst slice   (len %d  ptr: %p): %#v\n",
		len(ba), &ba, ba, len(s1), &s1, s1,
	)

	s2 := s1[:3]
	fmt.Printf(
		"\nafter s2 := s1[:3]\nbacking array (len %d  ptr: %p): %#v\nfirst slice   (len %d  ptr: %p): %#v\nsecond slice  (len %d  ptr: %p): %#v\n",
		len(ba), &ba, ba, len(s1), &s1, s1, len(s2), &s2, s2,
	)

	s2 = append(s2, 226)
	fmt.Printf(
		"\nafter s2 = append(s2, 226)\nbacking array (len %d  ptr: %p): %#v\nfirst slice   (len %d  ptr: %p): %#v\nsecond slice  (len %d  ptr: %p): %#v\n",
		len(ba), &ba, ba, len(s1), &s1, s1, len(s2), &s2, s2,
	)

	s2[1] = 4462
	fmt.Printf(
		"\nafter s2[1] = 4462\nbacking array (len %d  ptr: %p): %#v\nfirst slice   (len %d  ptr: %p): %#v\nsecond slice  (len %d  ptr: %p): %#v\n",
		len(ba), &ba, ba, len(s1), &s1, s1, len(s2), &s2, s2,
	)

	s3 := append([]int{}, s1...)
	fmt.Printf(
		"\nafter s3 := append([]int{}, s1...)\nbacking array (len %d  ptr: %p): %#v\nfirst slice   (len %d  ptr: %p): %#v\nsecond slice  (len %d  ptr: %p): %#v\nthird slice   (len %d  ptr: %p): %#v\n",
		len(ba), &ba, ba, len(s1), &s1, s1, len(s2), &s2, s2, len(s3), &s3, s3,
	)

	s3[7] = 877
	fmt.Printf(
		"\nafter s3[7] = 877\nbacking array (len %d  ptr: %p): %#v\nfirst slice   (len %d  ptr: %p): %#v\nsecond slice  (len %d  ptr: %p): %#v\nthird slice   (len %d  ptr: %p): %#v\n",
		len(ba), &ba, ba, len(s1), &s1, s1, len(s2), &s2, s2, len(s3), &s3, s3,
	)

	s0 := []int{2, 4, 7, 3}
	s01 := s0

	fmt.Printf(
		"\nafter s0 := []int{2, 4, 7, 3}; s01 := s0\nzero slice     (len %d  ptr: %p): %#v\nzero one slice (len %d  ptr: %p): %#v\n",
		len(s0), &s0, s0, len(s01), &s01, s01,
	)

	s01[2] = 9
	fmt.Printf(
		"\nafter s01[2] = 9\nzero slice     (len %d  ptr: %p): %#v\nzero one slice (len %d  ptr: %p): %#v\n",
		len(s0), &s0, s0, len(s01), &s01, s01,
	)

	s02 := s0[:]
	s0 = append(s0, 8)
	fmt.Printf(
		"\nafter s02 := s0[:]; s0 = append(s0, 8)\nzero slice     (len %d  ptr: %p): %#v\nzero two slice (len %d  ptr: %p): %#v\n",
		len(s0), &s0, s0, len(s02), &s02, s02,
	)

	s02[3] = 7
	fmt.Printf(
		"\nafter s02[3] = 7\nzero slice     (len %d  ptr: %p): %#v\nzero two slice (len %d  ptr: %p): %#v\n",
		len(s0), &s0, s0, len(s02), &s02, s02,
	)
}

func capacitySlice() {
	src := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Printf(
		"src slice & backing array items: %#v\ncapacity of src slice and length od backing array: %d\nlength od src slice: %d\n",
		src, cap(src), len(src),
	)
	fmt.Print("---          ---          ---")

	part := src
	fmt.Printf(
		"\nStatement: part := src\n\nsrc   = %#v\npart  = %#v\n\ncap(src)   = %d; len(src)   = %d\ncap(part)  = %d; len(part)  = %d\n",
		src, part, cap(src), len(src), cap(part), len(part),
	)
	fmt.Print("---          ---          ---")

	part[4] = 44
	fmt.Printf(
		"\nStatement: part[4] = 44\n\nsrc   = %#v\npart  = %#v\n\ncap(src)   = %d; len(src)   = %d\ncap(part)  = %d; len(part)  = %d\n",
		src, part, cap(src), len(src), cap(part), len(part),
	)
	fmt.Print("---          ---          ---")

	part = part[1:4]
	fmt.Printf(
		"\nStatement: part = part[1:4]\n\nsrc   = %#v\npart  = %#v\n\ncap(src)   = %d; len(src)   = %d\ncap(part)  = %d; len(part)  = %d\n",
		src, part, cap(src), len(src), cap(part), len(part),
	)
	fmt.Print("---          ---          ---")

	part = part[:5]
	fmt.Printf(
		"\nStatement: part = part[:5]\n\nsrc   = %#v\npart  = %#v\n\ncap(src)   = %d; len(src)   = %d\ncap(part)  = %d; len(part)  = %d\n",
		src, part, cap(src), len(src), cap(part), len(part),
	)
	fmt.Print("---          ---          ---")

	part = append(part, 66)
	fmt.Printf(
		"\nStatement: part = append(part, 66)\n\nsrc   = %#v\npart  = %#v\n\ncap(src)   = %d; len(src)   = %d\ncap(part)  = %d; len(part)  = %d\n",
		src, part, cap(src), len(src), cap(part), len(part),
	)
	fmt.Print("---          ---          ---")

	part = part[:cap(part)]
	fmt.Printf(
		"\nStatement: part = part[:cap(part)]\n\nsrc   = %#v\npart  = %#v\n\ncap(src)   = %d; len(src)   = %d\ncap(part)  = %d; len(part)  = %d\n",
		src, part, cap(src), len(src), cap(part), len(part),
	)
	fmt.Print("---          ---          ---")

	part = append(part, 99)
	fmt.Printf(
		"\nStatement: part = append(part, 99)\n\nsrc   = %#v\npart  = %#v\n\ncap(src)   =  %d; len(src)   =  %d\ncap(part)  = %d; len(part)  = %d\n",
		src, part, cap(src), len(src), cap(part), len(part),
	)
	fmt.Println("Change backing array of these two slices (src and part)")
	fmt.Print("---          ---          ---")

	part[6] = 77
	fmt.Printf(
		"\nStatement: part[6] = 77\n\nsrc   = %#v\npart  = %#v\n\ncap(src)   =  %d; len(src)   =  %d\ncap(part)  = %d; len(part)  = %d\n",
		src, part, cap(src), len(src), cap(part), len(part),
	)
	fmt.Print("---          ---          ---")

	part = part[:cap(part)]
	fmt.Printf(
		"\nStatement: part = part[:cap(part)]\n\nsrc   = %#v\npart  = %#v\n\ncap(src)   =  %d; len(src)   =  %d\ncap(part)  = %d; len(part)  = %d\n",
		src, part, cap(src), len(src), cap(part), len(part),
	)
	fmt.Print("---          ---          ---")
}

func fullSliceExpressions() {
	s.PrintBacking = true
	nums := []int{1, 3, 2, 4}

	odds := nums[:2:3]
	odds = append(odds, 5)
	odds = append(odds, 7)
	evens := nums[2:4:len(nums)]

	s.Show("nums", nums)
	s.Show("odds", odds)
	s.Show("evens", evens)
}

func makeFunction() {
	s.PrintBacking = true
	s.MaxPerLine = 10

	tasks := []string{"read", "run", "fight"}
	s.Show("task", tasks)

	tasksV0 := tasks
	s.Show("tasksV0", tasksV0)

	var tasksV1 []string
	tasksV1 = tasks
	s.Show("tasksV1", tasksV1)

	tasksV2 := make([]string, 0, len(tasks))
	for _, task := range tasks {
		tasksV2 = append(tasksV2, task)
	}
	s.Show("tasksV2", tasksV2)

	tasksV2[0] = "read book"
	tasksV1[2] = "play"

	s.Show("task after tasksV2[0] = \"read book\";	tasksV1[2] = \"play\"", tasks)
	s.Show("tasksV0 after tasksV2[0] = \"read book\";	tasksV1[2] = \"play\"", tasksV0)
	s.Show("tasksV1 after tasksV2[0] = \"read book\";	tasksV1[2] = \"play\"", tasksV1)
	s.Show("tasksV2 after tasksV2[0] = \"read book\";	tasksV1[2] = \"play\"", tasksV2)
}

func copyFunction() {
	s.PrintBacking = true
	s.MaxPerLine = 10

	data := []int{0, 1, 2, 3}
	copyDataV1 := []int{}
	var copyDataV2 []int
	copyDataV3 := []int{4, 5}
	copyDataV4 := []int{4, 5, 6, 7, 8, 9}
	copyDataV5 := make([]int, 0, len(data))
	copyDataV6 := make([]int, len(data))
	copyDataV7 := append([]int(nil), data...)

	copy(copyDataV1, data)
	copy(copyDataV2, data)
	copy(copyDataV3, data)
	copy(copyDataV4, data)
	copy(copyDataV5, data)
	copy(copyDataV6, data)

	s.Show("data", data)
	s.Show("copyDataV1 after copy(copyDataV1, data)", copyDataV1)
	s.Show("copyDataV2 after copy(copyDataV2, data)", copyDataV2)
	s.Show("copyDataV3 after copy(copyDataV3, data)", copyDataV3)
	s.Show("copyDataV4 after copy(copyDataV4, data)", copyDataV4)
	s.Show("copyDataV5 after copy(copyDataV5, data)", copyDataV5)
	s.Show("copyDataV6 after copy(copyDataV6, data)", copyDataV6)
	s.Show("copyDataV7 after copyDataV7 := append([]int(nil), data...)", copyDataV7)

	data[1] = 11
	copyDataV3[0] = 66
	s.Show("data after data[1] = 11; copyDataV3[0] = 66", data)
	s.Show("copyDataV3 after data[1] = 11; copyDataV3[0] = 66", copyDataV3)
}

func multiDimensionSlice() {
	daySpending := strings.Split(costs, "\n")
	spending := make([][]int, len(daySpending))
	day := make([]int, len(daySpending))

	for i, d := range daySpending {
		daySpend := strings.Fields(d)
		spending[i] = make([]int, len(daySpend))
		for j, spend := range daySpend {
			cost, _ := strconv.Atoi(spend)
			spending[i][j] = cost
			day[i] += cost
		}
	}

	for i, d := range day {
		fmt.Printf("DAY %d\nTotal spend: $%d\n", i+1, d)
		if len(spending[i]) != 0 {
			fmt.Print("Detail:\n\n")
		}
		for j, s := range spending[i] {
			fmt.Printf("\t spend %d: $%d\n", j+1, s)
		}
		fmt.Print("\n-------------------------------------------\n\n")
	}
}
