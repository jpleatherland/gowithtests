package integers

import (
	"testing"
	"fmt"
)

func TestAdder(t *testing.T) {
	got := Add(2, 2)
	want := 4
	assertCorrectResult(t, got, want)
}

func assertCorrectResult(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func ExampleAdd(){
	got := Add(10, 12)
	fmt.Println(got)
	//Output: 22
}
