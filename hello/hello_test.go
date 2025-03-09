package main

import "testing"

/*
Let's go over the cycle again

Write a test
Make the compiler pass
Run the test, see that it fails and check the error message is meaningful
Write enough code to make the test pass
Refactor
*/

func TestHello(t *testing.T) {
	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("Yogi", "")
		want := "Hello, Yogi"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello, world when given empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("yogi", "spanish")
		want := "Hola, yogi"
		assertCorrectMessage(t, got, want)
	})
}

// this func was refactored from assert in functions above
func assertCorrectMessage(t testing.TB, got, want string) {
	// t.Helper() for  is needed to tell the test suite that this method is a helper.
	// By doing this, when it fails, the line number reported will be in our function call
	// rather than inside our test helper.
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
