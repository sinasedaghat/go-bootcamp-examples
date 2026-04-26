package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"unicode/utf8"
)

func inputScanning() {
	fmt.Println("Yo Scanner warm up...")
}

func lineCounter() {
	// os.Stdin.Close()
	in := bufio.NewScanner(os.Stdin)

	lines := 1
	for in.Scan() {
		fmt.Printf("content of line %d (text):\n\t%v\n", lines, in.Text())
		fmt.Printf("content of line %d (bytes):\n\t%v\n", lines, in.Bytes())
		lines++
	}

	err := in.Err()
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	fmt.Printf("\nYour target context has %d lines\n", lines)
}

func wordFinder() {
	rx := regexp.MustCompile(`[^a-z]+`)

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("You want to find which word?")
		return
	}

	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	words := make(map[string]bool) // set data structure
	wc := 0
	for in.Scan() {

		word := rx.ReplaceAllString(strings.ToLower(in.Text()), "")
		if utf8.RuneCountInString(word) > 2 {
			words[word] = true
			wc++
		}
	}

	// for k := range words {
	// 	fmt.Println(k)
	// }

	fmt.Printf("\n\t\t\t✅Your reference text has %d word, and has %d specific word\n\n", wc, len(words))
	if words[strings.ToLower(args[0])] {
		fmt.Printf("✅ The reference text contain \"%s\"\n", args[0])
	} else {
		fmt.Printf("⚠️ The reference text doesn't contain \"%s\"\n", args[0])
	}
}

func logParser() {
	in := bufio.NewScanner(os.Stdin)

	visitTable := make(map[string]int)
	errorTable := make(map[int]error)

	visitTableKeys := make([]string, 0)
	errorTableKeys := make([]int, 0)

	line := 0
	for in.Scan() {
		line++

		d, v, err := lineParser(in.Text())
		if err != nil {
			errorTableKeys = append(errorTableKeys, line)
			errorTable[line] = err
			continue
		}

		visitTableKeys = append(visitTableKeys, d)
		visitTable[d] += v
	}

	sort.Strings(visitTableKeys)
	sort.Ints(errorTableKeys)

	printVisitTable(visitTableKeys, visitTable)
	printErrorTable(errorTableKeys, errorTable)

	if err := in.Err(); err != nil {
		fmt.Println("Errors for Scanner:", err)
	}
}

func lineParser(line string) (domain string, visit int, err error) {
	record := strings.Fields(line) // record := strings.Split(line, " ")

	if len(record) != 2 {
		return "", 0, errors.New("A line does not have two values.")
	} else if n, err := strconv.Atoi(record[1]); err != nil {
		return "", 0, errors.New("The second variable (visit) is not a number.")
	} else if n < 0 {
		return "", 0, errors.New("The second variable (visit) is a negative number.")
	} else {
		return record[0], n, nil
	}
}

func printVisitTable(keys []string, records map[string]int) {
	fmt.Printf(
		"\n\n%-30s %s\n%s\n",
		"DOMAIN",
		"VISIT",
		strings.Repeat("-", 36),
	)
	total := 0
	for _, k := range keys {
		total += records[k]
		fmt.Printf("%-30s %d\n", k, records[k])
	}
	fmt.Printf("\n%-30s %d\n\n", "TOTAL", total)
}

func printErrorTable(keys []int, errs map[int]error) {
	fmt.Printf(
		"\n\n%-6s %-50s\n%s\n",
		"LINE",
		"MESSAGE",
		strings.Repeat("-", 60),
	)

	for _, k := range keys {
		fmt.Printf("%-6d %-50s\n", k, errs[k])
	}
	fmt.Println()
}
