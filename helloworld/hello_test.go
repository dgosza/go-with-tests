package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people - english", func(t *testing.T) {
		got := Hello("Diego", "english")
		wanted := "Hello, Diego"

		assertCorrectMessage(t, got, wanted)

	})

	t.Run("saying hello to people - default(no name)", func(t *testing.T) {
		got := Hello("", "english")
		wanted := "Hello, World"

		assertCorrectMessage(t, got, wanted)
	})

	t.Run("saying hello to people - default(no language)", func(t *testing.T) {
		got := Hello("Diego", "")
		wanted := "Hello, Diego"

		assertCorrectMessage(t, got, wanted)
	})

	t.Run("saying hello to people - default(no name and no language)", func(t *testing.T) {
		got := Hello("", "")
		wanted := "Hello, World"

		assertCorrectMessage(t, got, wanted)
	})

	t.Run("saying hello to people - russian", func(t *testing.T) {
		got := Hello("Diego", "russian")
		wanted := "Привет, Diego"

		assertCorrectMessage(t, got, wanted)

	})

	t.Run("saying hello to people - spanish", func(t *testing.T) {
		got := Hello("Diego", "spanish")
		wanted := "Hola, Diego"

		assertCorrectMessage(t, got, wanted)

	})

}

func assertCorrectMessage(t testing.TB, got, wanted string) {
	t.Helper()
	if got != wanted {
		t.Errorf("got: %s wanted: %s", got, wanted)
	}
}
