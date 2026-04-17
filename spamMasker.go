package main

import (
	"bytes"
	"fmt"
	"os"
)

// VERSION 0
var (
	patterns = []string{"http://", "https://", "www."}
)

func masker() {
	var msg string
	if args := os.Args[1:]; len(args) == 1 {
		msg = args[0]
	} else {
		fmt.Println("Give me a something")
		return
	}

	var r [][2]int

	for b := 0; b < len(msg); {
		for _, pattern := range patterns {
			for p := 0; p < len(pattern); p++ {
				if pattern[p] != msg[b+p] {
					break
				}
				if p == len(pattern)-1 {
					var s, e int
					b += len(pattern)
					s = b
					for b < len(msg) && isDomainRune(msg[b]) {
						b++
					}
					e = b
					r = append(r, [2]int{s, e})
				}
			}
		}
		b++
	}
	writeMasked(msg, r)
}

func isDomainRune(b byte) bool {
	return (b >= '0' && b <= '9') ||
		(b >= 'a' && b <= 'z') ||
		(b >= 'A' && b <= 'Z') ||
		b == '.' || b == '-'
}

func writeMasked(str string, mp [][2]int) {
	b := []byte(str)

	for _, p := range mp {
		copy(
			b[p[0]:p[1]],
			bytes.Repeat([]byte{'*'}, p[1]-p[0]),
		)
	}

	fmt.Printf("%s\n", b)
}

// VERSION 1
const (
	link = "http://"

	mask = '*'
)

func spamMasker() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Give me a text")
		return
	}

	var (
		text = args[0]
		n    = len(text)
		buf  = []byte(text)
		c    byte
		in   bool

		nl = len(link)
	)

	for b := 0; b < n; b++ {
		if b+nl <= n && text[b:b+nl] == link {
			copy(buf[b:b+nl], text[b:b+nl])
			b += nl
			in = true
		}

		c = text[b]

		switch c {
		case ' ', '\n', '\t':
			in = false
		}

		if in {
			c = mask
		}
		buf[b] = c
	}

	fmt.Printf("%s\n", buf)
}
