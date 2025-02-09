package main

import (
	"fmt"
	"mihil99/noqo/core"
)

func main() {

	expr := core.Function{
		Name: "f",
		Args: []core.Expr{
			core.Symbol{Name: "x"},
			core.Function{
				Name: "g",
				Args: []core.Expr{
					core.Symbol{Name: "y"},
				},
			},
		},
	}

	fmt.Println(expr.String())
}
