/*
 * Copyright (c) 2011 jamra.source@gmail.com, 2015 Robin Orheden
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not
 * use this file except in compliance with the License. You may obtain a copy of
 * the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations under
 * the License.
 */

// Package cleo provides a fast search algorithm for prefix matching large amounts of text.

package cleo

import (
	"sort"
)

type Index struct {
	forward  *ForwardIndex
	inverted *InvertedIndex
	scoreFn  ScoreFn
}

func NewIndex() *Index {
	return &Index{
		forward:  NewForwardIndex(),
		inverted: NewInvertedIndex(),
		scoreFn:  JacaardScoring,
	}
}

func (i *Index) Add(id string, value string) {
	i.inverted.Add(id, value)
	i.forward.Add(id, value)
}

//Iterates through all of the 8 bytes (64 bits) and tests
//each bit that is set to 1 in the query's filter against
//the bit in the comparison's filter.  If the bit is not
// also 1, you do not have a match.
func testBytesFromQuery(bf int, qBloom int) bool {
	for i := uint(0); i < 64; i++ {
		//a & (1 << idx) == b & (1 << idx)
		if (bf&(1<<i) != (1 << i)) && qBloom&(1<<i) == (1<<i) {
			return false
		}
	}
	return true
}

func (ix *Index) Search(query string) []rankedResult {
	results := make([]rankedResult, 0, 0)

	candidates := ix.inverted.Search(query) //First get candidates from Inverted Index
	qBloom := ComputeBloomFilter(query)

	for _, i := range candidates {
		if testBytesFromQuery(i.bloom, qBloom) { //Filter using Bloom Filter
			c := ix.forward.itemAt(i.id)  //Get whole document from Forward Index
			score := ix.scoreFn(query, c) //Score the Forward Index between 0-1
			results = append(results, rankedResult{i.id, c, score})
		}
	}

	sort.Sort(ByScore{results})

	return results
}
