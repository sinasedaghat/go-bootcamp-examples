package bookCity

import "fmt"

type Price float64

func (p Price) String() string {
	return fmt.Sprintf("$%.2f (USD)", p)
}
