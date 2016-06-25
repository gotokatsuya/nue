package nue

import (
	"testing"
)

func BenchmarkSplitURLPath(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = splitURLPath("/user/show")
	}
}
