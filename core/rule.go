package core

// Basic implementation of a rule, this takes in a head and a body which are both expressions
// The rule defines the way a head expression can be transformed into a body expression

// Comments:
// Mostly same as in expr.go

type Rule struct {
	Head Expr
	Body Expr
}

func (r Rule) String() string {
	return r.Head.String() + " => " + r.Body.String()
}
