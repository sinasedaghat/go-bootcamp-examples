package products

import (
	"sort"
	"strings"
)

type list []*product

func (l list) String() string {
	var str strings.Builder
	for _, p := range l {
		str.WriteString("* ")
		str.WriteString(p.String())
		str.WriteRune('\n')
	}
	return str.String()
}

func (l list) discount(dp float64) {
	for _, p := range l {
		p.discount(dp)
	}
}

func (l list) Len() int {
	return len(l)
}

func (l list) Less(i, j int) bool {
	return l[i].Title < l[j].Title
}

func (l list) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

type byPrice struct {
	list
}

func (bp byPrice) Less(i, j int) bool {
	return bp.list[i].Price < bp.list[j].Price
}

func sortByPrice(l list) sort.Interface {
	return &byPrice{l}
}

type reverse struct {
	sort.Interface
}

func (r reverse) Less(i, j int) bool {
	return r.Interface.Less(j, i)
}

func reverseSort(i sort.Interface) sort.Interface {
	return reverse{i}
}
