package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestCat(t *testing.T) {
	s := "Shaken not stirred"
	xs := strings.Split(s, " ")
	s = Cat(xs)
	if s != "Shaken not stirred" {
		t.Error("got", s, "want", "Shaken not stirred")
	}
}

func TestJoin(t *testing.T) {
	s := "Shaken not stirred"
	xs := strings.Split(s, " ")
	s = Join(xs)
	if s != "Shaken not stirred" {
		t.Error("got", s, "want", "Shaken not stirred")
	}
}

func ExampleCat() {
	s := "Shaken not stirred"
	xs := strings.Split(s, " ")
	fmt.Println(Cat(xs))
	// Output:
	// Shaken not stirred
}

func ExampleJoin() {
	s := "Shaken not stirred"
	xs := strings.Split(s, " ")
	fmt.Println(Join(xs))
	// Output:
	// Shaken not stirred
}

const s = "I went to the woods because I wished to live deliberately, to front only the essential facts of life, and " +
	"see if I could not learn what it had to teach, and not, when I came to die, discover that I had not lived. I did " +
	"not wish to live what was not life, living is so dear; nor did I wish to practise resignation, unless it was " +
	"quite necessary. I wanted to live deep and suck out all the marrow of life, to live so sturdily and Spartan-like " +
	"as to put to rout all that was not life, to cut a broad swath and shave close, to drive life into a corner, and " +
	"reduce it to its lowest terms, and, if it proved to be mean, why then to get the whole and genuine meanness of " +
	"it, and publish its meanness to the world; or if it were sublime, to know it by experience, and be able to give " +
	"a true account of it in my next excursion."

var xs []string

func BenchmarkCat(b *testing.B) {
	xs = strings.Split(s, " ")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Cat(xs)
	}
}

func BenchmarkJoin(b *testing.B) {
	xs = strings.Split(s, " ")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Join(xs)
	}
}