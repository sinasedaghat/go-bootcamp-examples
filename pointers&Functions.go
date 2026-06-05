package main

import (
	"fmt"
	"strings"
)

type House struct {
	name string
	room int
}

func pointerFunction() {
	nums := [...]int{2, 3, 4}
	fmt.Printf("1 nums: %#v\n\n", nums)

	incr(nums)
	fmt.Printf("2 nums: %#v\n\n", nums)

	incrByPtr(&nums)
	fmt.Printf("3 nums: %#v\n", nums)
	fmt.Printf("%s\n", strings.Repeat("-", 20))

	dirs := []string{"up", "right", "down", "left"}
	fmt.Printf("1 dirs: %#v\n\n", dirs)

	upper(dirs)
	fmt.Printf("2 dirs: %#v\n\n", dirs)

	upperByPtr(&dirs)
	fmt.Printf("3 dirs: %#v\n\n", dirs)
	fmt.Printf("%s\n", strings.Repeat("-", 20))

	confused := map[string]int{"one": 2, "tow": 1}
	fmt.Printf("1 confused: %#v\n\n", confused)

	fix(confused)
	fmt.Printf("2 confused: %#v\n\n", confused)

	fixByPtr(&confused)
	fmt.Printf("3 confused: %#v\n\n", confused)
	fmt.Printf("%s\n", strings.Repeat("-", 20))

	home := House{name: "My House", room: 10}
	fmt.Printf("1 home: %#v (ptr = %p)\n\n", home, &home)

	addRoom(home)
	fmt.Printf("2 home: %#v\n\n", home)

	addRoomByPtr(&home)
	fmt.Printf("3 home: %#v\n\n", home)
}

func incr(num [3]int) {
	for i, n := range num {
		n++ // Only `n` changes; there is no change in the `num` argument or in the original `num` array.
		// fmt.Printf("\t🔷 num[%d] == %d 🔷 n of %[1]d == %d\n", i, num[i], n)

		num[i] += 10 // Only the `num` argument changes; the original `num` array remains unchanged.
		// fmt.Printf("\t🔶 num[%d] == %d\n", i, num[i])
	}
	fmt.Printf("func num: %#v\n", num)
}

func incrByPtr(num *[3]int) {
	for i, n := range *num {
		n++ // Only `n` changes; there is no change in the original `num` array.

		num[i] += 20 // Change original `num` array
	}
	fmt.Printf("func *num: %#v\n", *num)
}

func upper(str []string) {
	for i, n := range str {
		n += "↬" // Only `n` changes; ...
		// fmt.Printf("%s\n", n)

		str[i] = strings.ToUpper(str[i])
		fmt.Printf("\tstr[%d] == %-8s(%p)\n", i, str[i], &str[i])
	}
	fmt.Printf("1 func str: %#v\n", str)
	str = append(str, "newDir")
	fmt.Printf("2 func str: %#v\n", str)
}

func upperByPtr(str *[]string) {
	for i := range *str {
		(*str)[i] = strings.ToUpper((*str)[i])
		fmt.Printf("\tstr[%d] == %-8s(%p)\n", i, (*str)[i], &(*str)[i])
	}
	fmt.Printf("func *str: %#v\n", str)
	*str = append(*str, "newDir")
	fmt.Printf("2 func *str: %#v\n", str)
}

func fix(m map[string]int) {
	for k, v := range m {
		if k == "one" {
			m[k] = 1
		} else {
			m[k] = 2
		}

		fmt.Printf("\tm[%s] == %-8d(v: %d)(ptr of k: %p | ptr of v: %p)\n", k, m[k], v, &k, &v)
	}
}

func fixByPtr(m *map[string]int) {
	for k, v := range *m {
		if k == "one" {
			(*m)[k] = 1
		} else {
			(*m)[k] = 2
		}

		fmt.Printf("\t(*m)[%s] == %-8d(v: %d)(ptr of k: %p | ptr of v: %p)\n", k, (*m)[k], v, &k, &v)
	}
}

func addRoom(house House) {
	fmt.Printf("1 func house: %#v (ptr = %p)\n", house, &house)
	house.room++
	fmt.Printf("2 func house: %#v (ptr = %p)\n", house, &house)
}

func addRoomByPtr(house *House) {
	fmt.Printf("1 func *house: %#v (ptr = %[1]p)\n", house)
	house.room++
	fmt.Printf("2 func *house: %#v (ptr = %[1]p)\n", house)
}
