package iteration

import (
	"testing"
)

func TestRepeatString(t *testing.T) {
	got := RepeatString("ab", 4)
	expected := "abababab"

	if got != expected {
		t.Errorf("got %q want %q", got, expected)
	}
}

func BenchmarkRepeatString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatString("a", i)
	}
}
