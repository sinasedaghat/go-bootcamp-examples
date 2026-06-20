package products

type list []*product

func (l list) print() {
	for _, p := range l {
		p.print()
	}
}

func (l list) discount(dp float64) {
	for _, p := range l {
		p.discount(dp)
	}
}

func Store() {
	list := list{
		{"Fundamentals of Physics", 104.99, toTimestamp(118281600)},
		{"Introduction to Classical Mechanics", 13.32, toTimestamp("733622400")},
		{title: "Statistical Mechanics", price: 70.39},
	}

	list.discount(.5)
	list.print()
}
