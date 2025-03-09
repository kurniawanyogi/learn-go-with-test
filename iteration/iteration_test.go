package main

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("print iteration", func(t *testing.T) {
		got := Repeat("a")
		want := "aaaaa"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

}

// referensi how to create benchmark https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go
// to run benchmarking `go test -bench . -v -benchmem`
func BenchmarkRepeat(b *testing.B) {
	b.Log("start benchmarking")
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
	b.Log("finish benchmarking")
}
