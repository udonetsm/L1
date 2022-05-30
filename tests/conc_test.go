package tests

import (
	"L1/concatenate"
	"testing"
)

func BenchmarkCopy(b *testing.B) {
	b.ResetTimer()
	concatenate.ConcCopy()
	b.StopTimer()
}

func BenchmarkBuffer(b *testing.B) {
	b.ResetTimer()
	concatenate.ConcBuffer()
	b.StopTimer()
}
