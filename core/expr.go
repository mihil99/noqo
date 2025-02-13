package core

// Basic Expression type which is either a function or a symbol
// Comments:
// - Enum definition is required to make the Expression interface be defined by a specific definition (maybe for now, maybe other functionality will be added in the future so this won't be an issue)
// - Enum definition in go is weird for me, but not the worst I guess
// - `String()` isn't any sort of override, so I'll still need to specifically call the function to get the string representation of an expression
// - I can't access a specific Enum value in a <EnumClass>::<EnumValue> kind of way, which is annoying, I could switch to prepending the class to any value if I can't handle it lol (something like <EnumClass>_<EnumValue>)

type ExprType int

const (
	ExprType_Sym ExprType = iota
	ExprType_Func
)

type Expression interface {
	String() string
	Type() ExprType
}

type Function struct {
	Name Symbol
	Args []Expression
}

type Symbol struct {
	Name string
}

func (f Function) String() string {
	result := f.Name.String() + "("
	for i, arg := range f.Args {
		if i > 0 {
			result += ", "
		}
		result += arg.String()
	}
	result += ")"
	return result
}

func (s Symbol) String() string {
	return s.Name
}

func (f Function) Type() ExprType {
	return ExprType_Func
}

func (s Symbol) Type() ExprType {
	return ExprType_Sym
}

func CreateFunction(name string, args ...Expression) Function {
	return Function{Name: Symbol{Name: name}, Args: args}
}

func CreateSymbol(name string) Symbol {
	return Symbol{Name: name}
}

func match(lhs Expression, rhs Expression, bindings *map[string]Expression) bool {
	switch lhs.(type) {

	}
	return false
}

func PatternMatch(lhs Expression, rhs Expression) map[string]Expression {
	bindings := map[string]Expression{}
	return bindings
}
