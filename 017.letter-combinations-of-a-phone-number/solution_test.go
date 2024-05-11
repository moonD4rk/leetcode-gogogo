package solution

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_letterCombinations(t *testing.T) {
	var testCases = []struct {
		digits      string
		wantedArray []string
	}{
		{
			digits:      "23",
			wantedArray: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			digits:      "",
			wantedArray: []string{},
		},
		{
			digits:      "2",
			wantedArray: []string{"a", "b", "c"},
		},
	}
	for _, tc := range testCases {
		actual := letterCombinations(tc.digits)
		assert.Equal(t, tc.wantedArray, actual)
	}
}

func Test_Combine(t *testing.T) {
	s := []string{"1", "2", "3"}
	s2 := []string{"4", "5"}
	s3 := []string{"q", "w", "e", "r"}
	result := combine([][]string{s, s2, s3})
	fmt.Println(result)
}

func Benchmark_Combine1(b *testing.B) {
	s := []string{"1", "2", "3"}
	s2 := []string{"4", "5"}
	s3 := []string{"q", "w", "e", "r"}
	for i := 0; i < b.N; i++ {
		combine([][]string{s, s2, s3})
	}
}

func Benchmark_Combine2(b *testing.B) {
	s := []string{"1", "2", "3"}
	s2 := []string{"4", "5"}
	s3 := []string{"q", "w", "e", "r"}
	for i := 0; i < b.N; i++ {
		combine2([][]string{s, s2, s3})
	}
}

func Benchmark_Combine3(b *testing.B) {
	s := []string{"1", "2", "3"}
	s2 := []string{"4", "5"}
	s3 := []string{"q", "w", "e", "r"}
	for i := 0; i < b.N; i++ {
		combine3([][]string{s, s2, s3})
	}
}

func Test_CombineThird(t *testing.T) {
	s := []string{"1", "2", "3"}
	s2 := []string{"4", "5"}
	s3 := []string{"q", "w", "e", "r"}
	s4 := []string{"a", "s", "d"}
	s5 := []string{"j", "k"}
	s6 := []string{"v", "b", "n", "m"}
	raw := [][][]string{
		[][]string{s, s4},
		[][]string{s2, s5},
		[][]string{s3, s6},
	}
	result := combineThird(raw)
	fmt.Println(result)
}
