package main

import (
	"fmt"
)

func main() {

	fmt.Println(STRING)
	fmt.Println(NUMBER)
	fmt.Println(HASHMAP)
	fmt.Println(ARRAY)
}

type TokenUnion int

const (
	STRING TokenUnion = iota
	NUMBER
	HASHMAP
	ARRAY
)

var types map[TokenUnion]string

var toks = [...]string{
	STRING:  "STRING",
	NUMBER:  "NUMBER",
	HASHMAP: "HASHMAP",
	ARRAY:   "ARRAY",
}

func (t TokenUnion) String() string {
	return toks[t]
}
