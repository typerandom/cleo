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

type InvertedIndex map[string][]document

func NewInvertedIndex() *InvertedIndex {
	i := make(InvertedIndex)
	return &i
}

func (x *InvertedIndex) Size() int {
	return len(map[string][]document(*x))
}

func (x *InvertedIndex) Add(id string, value string) {
	bloom := ComputeBloomFilter(value)
	for _, word := range strings.Fields(value) {
		word = getPrefix(word)

		ref, ok := (*x)[word]

		if !ok {
			ref = nil
		}

		(*x)[word] = append(ref, document{id: id, bloom: bloom})
	}
}

func (x *InvertedIndex) Search(query string) []document {
	q := getPrefix(query)

	if ref, ok := (*x)[q]; ok {
		return ref
	}

	return nil
}

//Forward Index - Maps the document id to the document

type ForwardIndex map[string]string

func NewForwardIndex() *ForwardIndex {
	i := make(ForwardIndex)
	return &i
}

func (x *ForwardIndex) Add(id string, value string) {
	for _, word := range strings.Fields(value) {
		if _, ok := (*x)[id]; !ok {
			(*x)[id] = word
		}
	}
}

func (x *ForwardIndex) itemAt(i string) string {
	return (*x)[i]
}
