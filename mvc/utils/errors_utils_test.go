package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	elms := BubbleSort([]int{4, 5, 6, 7, 1, 2})
	assert.NotNil(t, elms)
	assert.EqualValues(t, 1, elms[0])
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleSort(GenerateArr(1000))
	}
}

func BenchmarkInbuiltSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InbuiltSort(GenerateArr(1000))
	}
}
