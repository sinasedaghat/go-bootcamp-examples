package products

import (
	"fmt"
	"sort"
)

func Store() {
	list := list{
		{"Fundamentals of Physics", 104.99, toTimestamp(118281600)},
		{"Introduction to Classical Mechanics", 13.32, toTimestamp("733622400")},
		{title: "Statistical Mechanics", price: 70.39},
	}

	fmt.Printf("Discount of each products\n")
	list.discount(.5)
	fmt.Printf("List of Products\n")
	fmt.Print(list)
	fmt.Printf("\n\nList of Product after sort by title\n")
	sort.Sort(list)
	fmt.Print(list)
	fmt.Printf("\n\nList of Product after sort by title (reverse)\n")
	sort.Sort(reverseSort(list))
	fmt.Print(list)
	fmt.Printf("\n\nList of Product after sort by ascending price\n")
	sort.Sort(sortByPrice(list)) // sort.Sort(byPrice{list})
	fmt.Print(list)
	fmt.Printf("\n\nList of Product after sort by descending price\n")
	sort.Sort(reverseSort(sortByPrice(list)))
	fmt.Print(list)
}
