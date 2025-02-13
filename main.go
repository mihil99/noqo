package main

import (
	"fmt"
	"mihil99/noqo/core"
)

func main() {

	// Basic function creation and printing
	expr := core.CreateFunction(
		"f",
		core.CreateSymbol("g"),
		core.CreateFunction(
			"x",
			core.CreateSymbol("y"),
		),
	)
	fmt.Println("Expression: ", expr.String())

	// Basic rule creation and pringing
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

	// Basic application
	// todo: create pattern matching
}
