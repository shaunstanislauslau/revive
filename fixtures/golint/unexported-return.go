// Test for unexported return types.

// Package foo ...
package foo

type hidden struct{}

// Exported returns a hidden type, which is annoying.
func Exported() hidden { // MATCH /exported func Exported returns unexported type foo.hidden, which can be annoying to use/
	return hidden{}
}

// ExpErr returns a builtin type.
func ExpErr() error { // ok
}

func (hidden) ExpOnHidden() hidden { // ok
}

// T is another test type.
type T struct{}

// MethodOnT returns a hidden type, which is annoying.
func (T) MethodOnT() hidden { // MATCH /exported method MethodOnT returns unexported type foo.hidden, which can be annoying to use/
	return hidden{}
}

// ExpT returns a T.
func ExpT() T { // ok
	return T{}
}

func unexp() hidden { // ok
	return hidden{}
}

// This is slightly sneaky: we shadow the builtin "int" type.

type int struct{}

// ExportedIntReturner returns an unexported type from this package.
func ExportedIntReturner() int { // MATCH /exported func ExportedIntReturner returns unexported type foo.int, which can be annoying to use/
	return int{}
}
