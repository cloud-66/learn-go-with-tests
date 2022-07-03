package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func BenchmarkRepeatSB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatSB("a", 5)
	}
}

func ExampleRepeat() {
	result := Repeat("b", 6)
	fmt.Println(result)
	// Output: bbbbbb
}
