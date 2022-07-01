package main

import (
	"testing"
)

func BenchmarkStringAllocation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = stringAllocation()
	}
}
