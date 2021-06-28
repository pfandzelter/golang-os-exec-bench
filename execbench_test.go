package execbench

import "testing"

var result error

func benchmarkYes(lookPath bool, b *testing.B) {
	var err error
	for n := 0; n < b.N; n++ {
		// always record the result of Fib to prevent
		// the compiler eliminating the function call.
		err = Yes(lookPath)
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = err
}

func BenchmarkYesLook(b *testing.B)  { benchmarkYes(true, b) }
func BenchmarkYesNoLook(b *testing.B)  { benchmarkYes(false, b) }