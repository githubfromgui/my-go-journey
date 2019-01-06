package string_concat

import (
	"testing"
)

func BenchmarkStringConcatWithoutJoin(b *testing.B) {
	for n := 0; n < b.N; n++ {
		StringConcatWithoutJoin()
	}
}