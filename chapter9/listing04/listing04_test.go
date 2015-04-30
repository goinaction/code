// go test -run=XXX -bench=.

// This sample test demonstrates how to write a benchmark test.
package listing04

import "testing"

// BenchmarkRowTraverse capture the time it takes to perform
// a row traversal.
func BenchmarkRowTraverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rowTraverse()
	}
}

// BenchmarkColTraverse capture the time it takes to perform
// a column traversal.
func BenchmarkColTraverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		colTraverse()
	}
}
