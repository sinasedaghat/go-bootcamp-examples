package products

import "fmt"

type price float64

func (p price) String() string {
	return fmt.Sprintf("%.2f USD", p)
}
