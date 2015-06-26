package simple

import "math"
import "fmt"

// contains a set of tokens and can recombinate them
type combination struct {
	size int64
	tokens []string
	counts []int64
	position int64
}

// I +1 MOD given M, also tells if I+1 MOD M == 0
func moduloTick(i int64, m int64) (value int64, didRollOver bool) {
	i++
	return i%m, i%m==0
}
// iterates counts and position in comb
func (comb combination) Next() {
	length := int64(len(comb.tokens))
	posIndex := comb.size-1
	rollover := true
	for {
		comb.counts[posIndex], rollover = moduloTick(comb.counts[posIndex], length)
		if rollover && posIndex > 0 { posIndex-- } else { break }
	}
}
// calculates how many different strings can be produced
func (comb combination) MaxVal() int64 {
	var max int64 = 0
	tokenCount := len(comb.tokens)
	for i := comb.size; i > 0; i-- {
		max += int64(math.Pow(float64(tokenCount),float64(i)))
	}
	return max
}

func (comb combination) Skip(n int64) combination {
	size := comb.size
	tokenCount := int64(len(comb.tokens))
	for i,j := int64(0), size-1; i < size; i,j = i+1, j-1 {
		comb.counts[j] = (n / int64(math.Pow(float64(tokenCount),float64(i)))) % tokenCount
	}
	return comb
}

func (comb combination) String() string {
	value := ""
	for i := int64(0); i < comb.size; i++ {
		value += comb.tokens[comb.counts[i]]
	}
	return value
}

func NewCombinator(tokens []string, size int) combination {
	return combination{int64(size), tokens, make([]int64, size), 0 }
}
