package iteration

import "testing"

func TestIteration(t *testing.T) {
	t.Run("testing barebones additions", func(t *testing.T) {
		got := Iterate("a")
		want := "aaaaa"
		assertEquality(t, got, want)
	})
}
func assertEquality(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
func BenchmarkIterate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Iterate("a")
	}
}
