package products

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
)

var products list = list{
	{"Fundamentals of Physics", 104.99, toTimestamp(118281600)},
	{"Introduction to Classical Mechanics", 13.32, toTimestamp("733622400")},
	{Title: "Statistical Mechanics", Price: 70.39},
}

func Store() {
	fmt.Printf("Discount of each products\n")
	products.discount(.5)
	fmt.Printf("List of Products\n")
	fmt.Print(products)
	fmt.Printf("\n\nList of Product after sort by title\n")
	sort.Sort(products)
	fmt.Print(products)
	fmt.Printf("\n\nList of Product after sort by title (reverse)\n")
	sort.Sort(reverseSort(products))
	fmt.Print(products)
	fmt.Printf("\n\nList of Product after sort by ascending price\n")
	sort.Sort(sortByPrice(products)) // sort.Sort(byPrice{products})
	fmt.Print(products)
	fmt.Printf("\n\nList of Product after sort by descending price\n")
	sort.Sort(reverseSort(sortByPrice(products)))
	fmt.Print(products)
}

func MarshalInventoryStore() {
	i, err := json.MarshalIndent(products, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	err = writeJsonFile("inventory", i)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Store inventory created")
}

func UnmarshalInventoryStore() {
	content, err := readJsonFile()
	if err != nil {
		log.Fatal(err)
	}
	var l list

	err = json.Unmarshal(content, &l)
	if err != nil {
		log.Fatal(err)
	}

	sort.Sort(reverseSort(sortByPrice(l)))
	fmt.Print(l)
}

func writeJsonFile(name string, content []byte) error {
	path := filepath.Join("temp", "products", fmt.Sprintf("%s.json", name))
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	return os.WriteFile(path, content, 0644)
}

func readJsonFile() ([]byte, error) {
	args := os.Args[1:]
	if len(args) != 1 {
		return nil, fmt.Errorf("enter directory location")
	}

	content, err := os.ReadFile(args[0])
	if err != nil {
		return nil, err
	}
	return content, nil
}
