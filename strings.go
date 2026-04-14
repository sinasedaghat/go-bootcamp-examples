package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
	"unsafe"
)

func runeDetails() {
	var start, end int

	if args := os.Args[1:]; len(args) == 2 {
		start, _ = strconv.Atoi(args[0])
		end, _ = strconv.Atoi(args[1])
	}
	if start == 0 && end == 0 {
		start = 'A'
		end = 'Z'
	}

	fmt.Printf(
		"%-10s %-10s %-10s %-12s\n%s\n",
		"literal", "dec", "hex", "encode",
		strings.Repeat("-", 45),
	)
	for i := start; i <= end; i++ {

		fmt.Printf(
			"%-10c %-10[1]d %-10[1]x % -12x\n",
			i, string(i),
		)
	}
}

func runeLoop() {
	ch := '¶'

	fmt.Printf(
		"ch := '¶'\nch: %15v\nch decimal: %7[1]d\nch hexadecimal: %[1]x\nstring(ch): % 9x\n\n",
		ch, string(ch),
	)

	s := "¶"
	fmt.Printf(
		"s := \"¶\"\ns: %6s\nlen(s): %d\n\n",
		s, len(s),
	)

	fmt.Printf("for i, r := range s {\n")
	for i, r := range s {
		fmt.Printf("\tindex(i): %d, byte(s[i]): %x, rune(r): %d, character(%%c): %[3]c\n", i, s[i], r)
	}
	fmt.Printf("}\n\n")

	fmt.Printf("for i := 0; i < len(s); i++ {\n")
	for i := 0; i < len(s); i++ {
		fmt.Printf("\tindex(i): %d, byte(s[i]): %x, rune(s[i]): %[2]d, character(%%c s[i]): %[2]c\n", i, s[i])
	}
	fmt.Printf("}\n\n")
}

func convertStrings() {
	str := "Yügen ☠️"
	bytes := []byte(str)
	runes := []rune(str)

	fmt.Printf("===> %s\n", string(bytes[1:3]))

	fmt.Printf(
		"str := %s\n\tlen(str) = %d\n\trune(str) = %d\n\n",
		str, len(str), utf8.RuneCountInString(str),
	)

	fmt.Printf(
		"bytes := %#v\n\tlen(bytes) = %d\n\trune(bytes) = %d\n\n",
		bytes, len(bytes), utf8.RuneCount(bytes),
	)

	fmt.Printf(
		"runes := %#v\n\tlen(runes) = %d\n\tsizeof(bytes) = %d\n\n",
		runes, len(runes), unsafe.Sizeof(runes[0]),
	)

	fmt.Printf(
		"%-10s %-10s %-12s \n%s\n",
		"index", "rune", "encode",
		strings.Repeat("-", 30),
	)
	for i, r := range str {
		fmt.Printf(
			"%-10d %-10c %-12[2]x\n",
			i, r,
		)
	}

	fmt.Printf("%s\n", strings.Repeat("-", 30))
	for i, b := range bytes {
		fmt.Printf(
			"%-10d %-10c %-12[2]x\n",
			i, b,
		)
	}

	fmt.Printf("%s\n", strings.Repeat("-", 30))
	for i, r := range runes {
		fmt.Printf(
			"%-10d %-10c %-12[2]x\n",
			i, r,
		)
	}
}

func runeDecoding() {
	str := "öykü"
	bytes := []byte(str)

	fmt.Println("⚠️ ", "for i := 0; i < len(bytes); i++ {")
	for i := 0; i < len(bytes); i++ {
		fmt.Printf("\tbytes[%d]: %c (% [2]x)\n", i, bytes[i])
	}
	fmt.Printf("}\n\n")

	fmt.Println("✅", "for i := 0; i < len(bytes); {")
	for i := 0; i < len(bytes); {
		r, s := utf8.DecodeRune(bytes[i:])
		fmt.Printf("\tbytes[%d:%d]: %c (% x)\n", i, i+s, r, bytes[i:i+s])
		i += s
	}
	fmt.Printf("}\n\n")

	fmt.Println("⚠️ ", "for i := 0; i < len(str); i++ {")
	for i := 0; i < len(str); i++ {
		fmt.Printf("\tstr[%d]: %c (% [2]x)\n", i, str[i])
	}
	fmt.Printf("}\n\n")

	fmt.Println("✅", "for i := 0; i < len(str); {")
	for i := 0; i < len(str); {
		r, s := utf8.DecodeRuneInString(str[i:])
		fmt.Printf("\tstr[%d:%d]: %c (% x)\n", i, i+s, r, str[i:i+s])
		i += s
	}
	fmt.Printf("}\n\n")

	fmt.Println("✅ ✅", "for i, r:= range str {")
	for i, r := range str {
		fmt.Printf("\tstr[%d:?]: %c\n", i, r)
	}
	fmt.Printf("}\n\n")
}

func upperRune() {
	word := []byte("öykü")
	fmt.Printf("%s: % [1]x\n", word)

	var s int
	for i := range string(word) {
		if i > 0 {
			s = i
			break
		}
	}

	firstCharV0 := bytes.ToUpper(word[:s])
	fmt.Printf("\nUse for loop to find first rune\n%s: % [1]x\n", firstCharV0)

	_, s = utf8.DecodeRune(word)
	firstCharV1 := bytes.ToUpper(word[:s])
	fmt.Printf("\nUse DecodeRune() function to find first rune\n%s: % [1]x\n", firstCharV1)

	copy(word[:s], firstCharV1)
	fmt.Printf("\n%s: % [1]x\n", word)
}

func stringHeader() {
	str0 := "hello"
	str1 := "hello"
	str2 := "hell"
	str3 := str0[:4]

	stringDetails(&str0)
	stringDetails(&str1)
	stringDetails(&str2)
	stringDetails(&str3)

	fmt.Printf("str2 == str3 = %t\n", str2 == str3)
}

func stringDetails(s *string) {
	pointer := unsafe.StringData(*s)
	length := len(*s)
	fmt.Printf("\nDetails of:\n\t   %s\n", *s)
	fmt.Printf("StringHeader:{\n  pointer: %x\n  length: %d\n}\n", pointer, length)
}
