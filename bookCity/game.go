package bookCity

type Game struct {
	Product
}

func newGames() []Game {
	COD := Game{
		Product{
			Name:  "Call of Duty",
			Price: 59.99,
		},
	}

	rayman := Game{
		Product{
			Name:  "Rayman® Legends",
			Price: 29.99,
		},
	}

	return []Game{
		COD,
		rayman,
	}
}
