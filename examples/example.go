package main

import (
	"github.com/typerandom/cleo"
	"strconv"
)

func main() {
	index := cleo.NewIndex()

	index.Add("1941", "pizza")
	index.Add("1210", "pizzaz")
	index.Add("2544", "pizzarello")
	index.Add("59120", "pizzeuto")
	index.Add("53200", "pizzaria")
	index.Add("53320", "pizzuto")

	results := index.Search("pizza")

	print("Found " + strconv.Itoa(len(results)) + " results:\n")

	for i := 0; i < len(results); i++ {
		print(strconv.Itoa(i+1) + ") " + results[i].Value + " (" + results[i].Id + ")\n")
	}
}
