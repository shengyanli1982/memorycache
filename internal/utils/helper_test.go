package utils

import (
	"github.com/stretchr/testify/assert"
	"hash/fnv"
	"testing"
)

func TestNewFnv32(t *testing.T) {
	for i := 0; i < 10; i++ {
		key := AlphabetNumeric.Generate(16)
		h := fnv.New32()
		h.Write(key)
		assert.Equal(t, h.Sum32(), Fnv32(string(key)))
	}
}

func TestToBinaryNumber(t *testing.T) {
	assert.Equal(t, 8, ToBinaryNumber(7))
	assert.Equal(t, 1, ToBinaryNumber(0))
	assert.Equal(t, 128, ToBinaryNumber(120))
	assert.Equal(t, 1024, ToBinaryNumber(1024))
}

func TestUniq(t *testing.T) {
	assert.ElementsMatch(t, Uniq([]int{1, 3, 5, 7, 7, 9}), []int{1, 3, 5, 7, 9})
	assert.ElementsMatch(t, Uniq([]string{"ming", "ming", "shi"}), []string{"ming", "shi"})
}

func TestRandomString(t *testing.T) {
	assert.Less(t, Numeric.Intn(10), 10)
	Numeric.Uint32()
	Numeric.Uint64()
}
