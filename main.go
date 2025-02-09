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
	fmt.Println("Expression: ", expr.String())

	head := core.CreateFunction("swap",
		core.CreateFunction("pair",
			core.CreateSymbol("x"),
			core.CreateSymbol("y")))

	body := core.CreateFunction("pair",
		core.CreateSymbol("y"),
		core.CreateSymbol("x"))

	rule := core.Rule{
		Head: head,
		Body: body,
	}

	fmt.Println("Rule: ", rule.String())
}
