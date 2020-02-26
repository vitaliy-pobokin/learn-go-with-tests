package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeater(t *testing.T) {
	repeated := Repeat("B", 6)
	expected := strings.Repeat("B", 6)

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("A", 5)
	fmt.Println(repeated)
	// Output: AAAAA
}
