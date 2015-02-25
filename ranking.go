package cleo

type rankedResults []rankedResult

type rankedResult struct {
	Id    string  `json:"id"`
	Value string  `json:"value"`
	Score float64 `json:"score"`
}

type ScoreFn func(value, query string) (score float64)
type byScore struct{ rankedResults }

func (s rankedResults) Len() int      { return len(s) }
func (s rankedResults) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s byScore) Less(i, j int) bool  { return s.rankedResults[i].Score > s.rankedResults[j].Score }

//Levenshtein distance is the number of inserts, deletions,
//and substitutions that differentiate one word from another.
//This algorithm is dynamic programming found at
//http://en.wikipedia.org/wiki/Levenshtein_distance
func levenshteinDistance(s, t string) int {
	m := len(s)
	n := len(t)
	width := n - 1
	d := make([]int, m*n)
	//y * w + h for position in array
	for i := 1; i < m; i++ {
		d[i*width+0] = i
	}

	for j := 1; j < n; j++ {
		d[0*width+j] = j
	}

	for j := 1; j < n; j++ {
		for i := 1; i < m; i++ {
			if s[i] == t[j] {
				d[i*width+j] = d[(i-1)*width+(j-1)]
			} else {
				d[i*width+j] = min(d[(i-1)*width+j]+1, //deletion
					d[i*width+(j-1)]+1,     //insertion
					d[(i-1)*width+(j-1)]+1) //substitution
			}
		}
	}
	return d[m*(width)+0]
}

//http://en.wikipedia.org/wiki/Jaccard_index
func JacaardScoring(query, candidate string) float64 {
	lev := levenshteinDistance(query, candidate)
	length := max(len(candidate), len(query))
	return float64(length-lev) / float64(length+lev) //Jacard score
}
