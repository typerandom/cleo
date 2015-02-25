package cleo

import "testing"

func TestLevenshtein(t *testing.T) {
	if levenshteinDistance("abcdefghij", "abcdefghix") != 1 {
		t.Fail()
	}

	if levenshteinDistance("abcdefghij", "abcdefghijk") != 1 {
		t.Fail()
	}

	if levenshteinDistance("abcdefghij", "abcdefghi") != 1 {
		t.Fail()
	}
}
