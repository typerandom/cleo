package cleo

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func min(a ...int) int {
	m := int(^uint(0) >> 1) // largest int
	for _, i := range a {
		if i < m {
			m = i
		}
	}
	return m
}

func max(a ...int) int {
	m := int(0)
	for _, i := range a {
		if i > m {
			m = i
		}
	}
	return m
}

var index *Index

func searchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.FormValue("query")

	searchResult := index.Search(query)
	resultJson, _ := json.Marshal(searchResult)

	fmt.Fprintf(w, string(resultJson))
}

func ListenAndServe(i *Index, port int) error {
	index = i

	http.HandleFunc("/search", searchHandler)

	print("Server running at http://localhost:" + strconv.Itoa(port) + "/\n")
	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)

	if err != nil {
		return err
	}

	return nil
}

func (i *Index) LoadFromFile(filePath string) error {
	file, err := os.Open(filePath)

	defer file.Close()

	if err != nil {
		return err
	}

	id := 0
	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			return err
		}

		i.Add(strconv.Itoa(id), line)
		id++
	}

	return nil
}
