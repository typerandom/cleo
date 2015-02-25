package main

import (
	"github.com/typerandom/cleo"
)

func main() {
	index := cleo.NewIndex()

	index.LoadFromFile("w1_fixed.txt")

	cleo.ListenAndServe(index, 8080)
}
