package model

type A struct {
	Name string
	Age int
}

type B struct {
	Name string
	Score float64
}

type C struct {
	A 
	B
}

type D struct {
	*A 
	*B
}
