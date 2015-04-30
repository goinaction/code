// go test -run=XXX -bench=.

// This sample test demonstrates how to write a benchmark test.
package listing04

import "testing"

func BenchmarkRowTraverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rowTraverse()
	}
}

func BenchmarkColTraverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		colTraverse()
	}
}
