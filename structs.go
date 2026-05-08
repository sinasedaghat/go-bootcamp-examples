package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type V2 struct {
	f21 string
	f22 string
	f23 bool
}

type (
	song struct {
		name     string
		artist   string
		duration int
	}

	album struct {
		name     string
		artist   string
		count    int
		bestSong song
	}

	playList struct {
		name  string
		genre string
		songs []song
	}
)

type permission map[string]bool

type user struct {
	Name       string     `json:"nameName"`
	Password   string     `json:"-"`
	Permission permission `json:"perms,omitempty"`
	Department string     `json:",omitempty"`
}

func startStruct() {
	type V1 struct {
		f11 string
		f12 int
	}

	// anonymous struct
	var sV0 struct {
		fs  string
		fi  int
		ff  float64
		fb  bool
		fss []string
		fa  [2]int
		fm  map[string]int
	}
	fmt.Printf("sV0: %+v\n", sV0)

	sV1 := V1{}
	fmt.Printf("sV1: %+v\n", sV1)

	sV1.f11 = "name field"
	fmt.Printf("\nAFTER\nsV1.f11 = \"name field\"\nsV1: %+v\n", sV1)
	// f13 := sV1.f13 // ERROR: sV1.f13 undefined (type V1 has no field or method f13)

	sV20 := V2{
		"f21",
		"f22",
		true,
	}
	fmt.Printf("sV20: %+v\n", sV20)

	sV21 := V2{
		f23: true,
		f21: "f21 field",
	}
	fmt.Printf("sV21: %+v\n", sV21)

	fmt.Printf("\nsV20 == sV21: %t\n", sV20 == sV21)
	// fmt.Printf("\nsV1 == sV21: %t\n", sV1 == sV21) // ERROR: invalid operation: sV1 == sV21 (mismatched types V1 and V2)
}

func assignStruct() {
	song1 := song{
		name:     "No Me",
		artist:   "Kensington",
		duration: 362,
	}

	song2 := song{
		name:     "Nothing Else Matters",
		artist:   "Metallica",
		duration: 389,
	}

	song3 := song{
		name:     "Send Me An Angel",
		artist:   "Scorpions",
		duration: 272,
	}

	song4 := song3
	fmt.Printf("\nAFTER\nsong4 := song3\nsong3: %+v\nsong4: %+v\n", song3, song4)

	song4.name = "Still Loving You"
	song4.duration = 402
	fmt.Printf("\nAFTER\nsong4.name = \"Still Loving You\"; song4.duration = 402\nsong3: %+v\nsong4: %+v\nsong3 == song4: %t\n", song3, song4, song3 == song4)

	// album1 := album{
	// 	name:   "Crazy World",
	// 	artist: "Scorpions",
	// 	count:  11,
	// }

	album2 := album{
		name:   "Metallica",
		artist: "Metallica",
		count:  12,
	}
	fmt.Printf("\nalbum2: %+v\n", album2)
	// fmt.Printf("\nCOMPARE\nalbum1 == album2: %t\nalbum1 == song3: %t\n", album1 == album2, album1 == song3) // ERROR: invalid operation: album1 == song3 (mismatched types album and song)

	album2.bestSong = song2
	fmt.Printf("\nAFTER\nalbum2.bestSong = song2\nalbum2: %+v\n", album2)

	album2.bestSong.duration = 400
	fmt.Printf("\nAFTER\nalbum2.bestSong.duration = 400\nalbum2: %+v\nsong2: %+v\n", album2, song2)

	playList1 := playList{
		name:  "best of the world",
		genre: "Heavy Metal",
	}
	fmt.Printf("\nplayList1: %+v\nplayList1.songs == nil: %t\ncap(playList1.songs): %d\n", playList1, playList1.songs == nil, cap(playList1.songs))

	playList2 := playList{
		name:  "Wow",
		genre: "Rock",
		songs: []song{song1},
	}
	fmt.Printf("\nplayList2: %+v\nplayList2.songs == nil: %t\ncap(playList2.songs): %d\n", playList2, playList2.songs == nil, cap(playList2.songs))

	playList1.songs = append(playList1.songs, song2, song3, song4)
	fmt.Printf("\nAFTER\nplayList1.songs = append(playList1.songs, song1, song2, song3, song4)\nplayList1: %+v\n", playList1)

	// fmt.Printf("\nCOMPARE\nplayList1 == playList2: %t\n", playList1 == playList2) // ERROR: playList1 == playList2 (struct containing []song cannot be compared)

	// playList22 := playList2
	// fmt.Printf("\nCOMPARE\n2 == playList22: %t\n", playList2 == playList22) // ERROR: playList2 == playList22 (struct containing []song cannot be compared)

	songs := []song{
		{
			name:     "No Me",
			artist:   "Kensington",
			duration: 362,
		},
		{
			name:     "Nothing Else Matters",
			artist:   "Metallica",
			duration: 389,
		},
		{
			name:     "Send Me An Angel",
			artist:   "Scorpions",
			duration: 272,
		},
	}
	playList3 := playList{
		name:  "best of rock",
		genre: "Rock",
		songs: songs,
	}
	fmt.Printf("\nplayList3: %+v\n", playList3)

	fmt.Printf("\n\nfor i, v := range playList1.songs {\n")
	for i, v := range playList3.songs {
		fmt.Printf("\n\tplayList3[%d] = %+v\n", i, v)

		if i == 1 {
			v.duration = 1200
		}

		if i == 2 {
			playList3.songs[i].duration = 1200
		}

		fmt.Printf("\tAFTER SOME LINE: playList3[%d] = %+v\n", i, v)
	}
	fmt.Println("}")

	fmt.Printf("\nAFTER\nv.duration = 1200(for second song)\nplayList3.songs[i].duration = 1200(for third song)\nplayList3: %+v\n", playList3)
}

func bareNotBareTypes() {
	structSongs := []song{
		{
			name:     "No Me",
			artist:   "Kensington",
			duration: 362,
		},
		{
			name:     "Nothing Else Matters",
			artist:   "Metallica",
			duration: 389,
		},
		{
			name:     "Send Me An Angel",
			artist:   "Scorpions",
			duration: 272,
		},
	}

	fmt.Printf("\n\nfor i, v := range structSongs {\n")
	for i, v := range structSongs {
		fmt.Printf("\n\tstructSongs[%d] = %+v\n", i, v)

		if i == 1 {
			v.duration = 1200
		}

		if i == 2 {
			structSongs[i].duration = 1200
		}

		fmt.Printf("\tAFTER SOME LINE: structSongs[%d] = %+v\n", i, v)
	}
	fmt.Println("}")
	fmt.Printf("\nAFTER\nv.duration = 1200(for second song)\nstructSongs[i].duration = 1200(for third song)\nstructSongs: %+v\n", structSongs)

	mapSongs := map[string]int{
		"noMe":               306,
		"nothingElseMatters": 389,
		"sendMeAnAngle":      272,
	}

	fmt.Printf("\n\nfor i, v := range mapSongs {\n")
	for k, v := range mapSongs {
		fmt.Printf("\n\tmapSongs[%s] = %+v\n", k, v)

		if k == "nothingElseMatters" {
			v = 1200
		}

		if k == "sendMeAnAngle" {
			mapSongs[k] = 1200
		}

		fmt.Printf("\tAFTER SOME LINE: mapSongs[%s] = %+v\n", k, v)
	}
	fmt.Println("}")
	fmt.Printf("\nAFTER\nv = 1200(for second song)\nmapSongs[i] = 1200(for third song)\nmapSongs: %v\n", mapSongs)

	arraySongs := [3][2]string{
		{"noMe", "306"},
		{"nothingElseMatters", "389"},
		{"sendMeAnAngle", "272"},
	}

	fmt.Printf("\n\nfor i, v := range arraySongs {\n")
	for i, v := range arraySongs {
		fmt.Printf("\n\tarraySongs[%d] = %+v\n", i, v)

		if i == 1 {
			v[1] = "1200"
		}

		if i == 2 {
			arraySongs[i][1] = "1200"
		}

		fmt.Printf("\tAFTER SOME LINE: arraySongs[%d] = %v\n", i, v)
	}
	fmt.Println("}")
	fmt.Printf("\nAFTER\nv = 1200(for second song)\narraySongs[i] = 1200(for third song)\narraySongs: %v\n", arraySongs)

	sliceSongs := [][]string{
		{"noMe", "306"},
		{"nothingElseMatters", "389"},
		{"sendMeAnAngle", "272"},
	}

	fmt.Printf("\n\nfor i, v := range sliceSongs {\n")
	for i, v := range sliceSongs {
		fmt.Printf("\n\tsliceSongs[%d] = %+v\n", i, v)

		if i == 1 {
			v[1] = "1200"
		}

		if i == 2 {
			sliceSongs[i][1] = "1200"
		}

		fmt.Printf("\tAFTER SOME LINE: sliceSongs[%d] = %v\n", i, v)
	}
	fmt.Println("}")
	fmt.Printf("\nAFTER\nv = 1200(for second song)\nsliceSongs[i] = 1200(for third song)\nsliceSongs: %v\n", sliceSongs)
}

func embedStruct() {
	type text struct {
		title string
		words int
	}

	type text1 struct {
		title string
		isbn  string
	}

	type book1 struct {
		title string
		words int
		isbn  string
	}

	type book2 struct {
		text text
		isbn string
	}

	type book3 struct {
		text
		isbn string
	}

	type book4 struct {
		text
		text1
		isbn string
	}

	firstBook := book1{
		title: "book1",
		words: 1,
		isbn:  "11",
	}

	secondBook := book2{
		text: text{
			title: "book2",
			words: 2,
		},
		isbn: "22",
	}

	thirdBook := book3{
		text: text{
			title: "book3",
			words: 3,
		},
		isbn: "33",
	}

	forthBook := book4{
		text: text{
			title: "book4",
			words: 4,
		},
		text1: text1{
			title: "book44",
			isbn:  "444",
		},
		isbn: "44",
	}

	fmt.Println("Structs Type")
	fmt.Printf("\ntype text struct {\n\ttitle string\n\twords int\n}\n")
	fmt.Printf("\ntype text1 struct {\n\ttitle string\n\tisbn string\n}\n")
	fmt.Printf("\ntype book1 struct {\n\ttitle string\n\twords int\n\tisbn  string\n}\n")
	fmt.Printf("\ntype book2 struct {\n\ttext text\n\tisbn string\n}\n")
	fmt.Printf("\ntype book3 struct {\n\ttext\n\tisbn string\n}\n")
	fmt.Printf("\ntype book4 struct {\n\ttext\n\ttext1\n\tisbn string\n}\n")

	fmt.Println("Literal Embed Structs")
	fmt.Printf("\nfirstBook := book1{\n\ttitle: \"book1\",\n\twords: 1,\n\tisbn: \"11\",\n}\nfirstBook: %+v\n", firstBook)
	fmt.Printf("\nsecondBook := book2{\n\ttext: text{\n\t\ttitle: \"book2\",\n\t\twords: 2,\n\t}\n\tisbn:  \"22\",\n}\nsecondBook: %+v\n", secondBook)
	fmt.Printf("\nthirdBook := book3{\n\ttext: text{\n\t\ttitle: \"book3\",\n\t\twords: 3,\n\t}\n\tisbn:  \"33\",\n}\nthirdBook: %+v\n", thirdBook)
	fmt.Printf(
		"\nforthBook := book4{\n\ttext: text{\n\t\ttitle: \"book3\",\n\t\twords: 3,\n\t}\n\ttext1: text1{\n\t\ttitle: \"book44\"\n\t\tisbn: \"444\",\n\t}\n\tisbn: \"44\",\n}\nfourthBook: %+v\n",
		forthBook,
	)

	fmt.Println("Read Field Embed Structs")
	fmt.Printf(
		"\nsecondBook.text == %+v\nsecondBook.text.title == %+v\nthirdBook == %+v",
		secondBook.text, secondBook.text.title, thirdBook,
	)
	fmt.Printf(
		"\nthirdBook.text == %+v\nthirdBook.text.title == %+v\nthirdBook.title == %+v\nthirdBook == %+v\n",
		thirdBook.text, thirdBook.text.title, thirdBook.title, thirdBook,
	)
	fmt.Printf(
		"\nforthBook.text == %+v\nforthBook.text1 == %+v\nforthBook.title == ERROR (have error: ambiguous selector forthBook.title)\nforthBook.isbn == %+v\nforthBook.text1.isbn == %+v\n",
		forthBook.text, forthBook.text1, forthBook.isbn, forthBook.text1.isbn,
	)

	forthBook.isbn = "new 44"
	fmt.Printf(
		"\nAFTER: forthBook.isbn = \"new 44\"\nforthBook.isbn == %+v\nforthBook.text1.isbn == %+v\n",
		forthBook.isbn, forthBook.text1.isbn,
	)
}

func logParserStructVersion() {
	type record struct {
		domain string
		visit  int
	}
	type logger struct {
		recordCount int
		errors      map[int]error
		errorLine   []int
		records     map[string]record
		domains     []string
	}

	in := bufio.NewScanner(os.Stdin)
	fmt.Printf("%+v", in)

	l := logger{
		errors:    make(map[int]error),
		errorLine: make([]int, 0),
		records:   make(map[string]record),
		domains:   make([]string, 0),
	}

	for in.Scan() {
		l.recordCount++
		r := strings.Fields(in.Text())

		if len(r) != 2 {
			l.errorLine = append(l.errorLine, l.recordCount)
			l.errors[l.recordCount] = errors.New("A line does not have two values.")
		} else if n, err := strconv.Atoi(r[1]); err != nil {
			l.errorLine = append(l.errorLine, l.recordCount)
			l.errors[l.recordCount] = errors.New("The second variable (visit) is not a number.")
		} else if n < 0 {
			l.errorLine = append(l.errorLine, l.recordCount)
			l.errors[l.recordCount] = errors.New("The second variable (visit) is a negative number.")
		} else {
			key := strings.ToLower(r[0])
			// l.records[strings.ToLower(r[0])].visit += n // ERROR: cannot assign to struct field l.records[strings.ToLower(r[0])].visit in map
			l.records[key] = record{
				domain: r[0],
				visit:  l.records[strings.ToLower(r[0])].visit + n,
			}
			l.domains = append(l.domains, strings.ToLower(r[0]))
		}
	}

	if err := in.Err(); err != nil {
		fmt.Println("Errors for Scanner:", err)
	}

	fmt.Printf("%+v", l)
}

func marshalStructs() {
	users := []user{
		{"Sina", "S1234", nil, "A"},
		{"Reza", "R4321", permission{"admin": true}, ""},
		{"Hamed", "H12344321", permission{"staff": true}, "B"},
	}

	// j, err := json.Marshal(users)
	j, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println(string(j))
}

func unmarshalJSONs() {
	var input []byte
	for in := bufio.NewScanner(os.Stdin); in.Scan(); {
		input = append(input, in.Bytes()...)
	}

	var users []user

	err := json.Unmarshal(input, &users)
	if err != nil {
		fmt.Println("Error from unmarshal target JSON file", err)
	}

	fmt.Printf("%#v", users)
}
