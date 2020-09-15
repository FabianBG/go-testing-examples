package main

import (
	"testing"
)

func BenchmarkMain(b *testing.B) {
	// ns/op means how many time per operation or for loop
	// allocs/op means how many distinct memory allocations occurred per op (single iteration)
	// B/op is how many bytes were allocated per op

	b.Run("operation", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			operation()
		}
	})
}
