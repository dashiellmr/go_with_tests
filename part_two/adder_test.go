package integers

import "testing"

func TestAdder(t *testing.T) {
	t.Run("testing barebones additions", func(t *testing.T) {
		got := Add(2, 2)
		want := 4
		assertCorrectSum(t, got, want)
	})
}

func assertCorrectSum(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("want %q, got %q", want, got)
	}
}
