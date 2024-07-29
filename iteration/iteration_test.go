package iteration

import (
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a", 5)
	want := "aaaaa"
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
		for range(b.N) {
				Repeat("a", 5)
		}
}
func BenchmarkRepeat2(b *testing.B) {
		for range(b.N) {
				Repeat2("a", 5)
		}
}
