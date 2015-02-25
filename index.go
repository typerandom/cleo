package cleo

import (
	"strings"
)

type document struct {
	id    string
	bloom int
}

func getPrefix(query string) string {
	qLen := min(len(query), 4)
	q := query[0:qLen]
	return strings.ToLower(q)
}

//Inverted Index - Maps the query prefix to the matching documents

type invertedIndex map[string][]document

func NewInvertedIndex() *invertedIndex {
	i := make(invertedIndex)
	return &i
}

func (x *invertedIndex) Size() int {
	return len(map[string][]document(*x))
}

func (x *invertedIndex) Add(id string, value string) {
	bloom := computeBloomFilter(value)
	for _, word := range strings.Fields(value) {
		word = getPrefix(word)

		ref, ok := (*x)[word]

		if !ok {
			ref = nil
		}

		(*x)[word] = append(ref, document{id: id, bloom: bloom})
	}
}

func (x *invertedIndex) Search(query string) []document {
	q := getPrefix(query)

	if ref, ok := (*x)[q]; ok {
		return ref
	}

	return nil
}

//Forward Index - Maps the document id to the document

type forwardIndex map[string]string

func NewForwardIndex() *forwardIndex {
	i := make(forwardIndex)
	return &i
}

func (x *forwardIndex) Add(id string, value string) {
	for _, word := range strings.Fields(value) {
		if _, ok := (*x)[id]; !ok {
			(*x)[id] = word
		}
	}
}

func (x *forwardIndex) itemAt(i string) string {
	return (*x)[i]
}
