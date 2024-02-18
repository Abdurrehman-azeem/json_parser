package jsonparser

type Value struct {
	Value interface{}
}

type Pair struct {
	Pair map[string]Value
}

type Object struct {
	Value Value
}

type Array struct {
	Array []Value
}
