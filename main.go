package main

import (
	"fmt"
	"mihil99/noqo/core"
)

func main() {

	expr := core.CreateFunction(
		"f",
		core.CreateSymbol("x"),
		core.CreateFunction(
			"x",
			core.CreateSymbol("y"),
		),
	)
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
