package main

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Run("add two integers", func(t *testing.T) {
		got := Add(5, 6)
		want := 11
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

// testable example, untuk expect nya di taruh di comment bawah
func ExampleAdd() {
	fmt.Println(Add(5, 6))
	// Output: 11
}
