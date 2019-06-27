package saying

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	expected := "Welcome my dear James"
	s := Greet("James")

	if s != expected {
		t.Errorf("got: %s expected: %s", s, expected)
	}
}

func ExampleGreet() {
	fmt.Println(Greet("James"))
	// Output:
	// Welcome my dear James
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("James")
	}
}
