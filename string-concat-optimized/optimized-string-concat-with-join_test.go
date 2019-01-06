package string_concat_optimized

import (
	"testing"
)

func BenchmarkStringConcatWithJoin(b *testing.B) {
	for n := 0; n < b.N; n++ {
		StringConcatWithJoin()
	}
}