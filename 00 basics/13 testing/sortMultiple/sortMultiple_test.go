package sortMultiple

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

var kerouac = Book{
	Title:  "The Darma Bums",
	Author: "Kerouac, Jack",
	Year:   1958,
}

var hemingway = Book{
	Title:  "The Old Man and the Sea",
	Author: "Hemingway, Ernest",
	Year:   1952,
}

var onan = Book{
	Title:  "Last Night at the Lobster",
	Author: "O’Nan, Stewart",
	Year:   2007,
}

var palahniuk = Book{
	Title:  "Fight Club",
	Author: "Palahniuk, Chuck",
	Year:   1996,
}

var irving = Book{
	Title:  "Last Night in Twisted River",
	Author: "Irving, John",
	Year:   2009,
}

type test struct {
	input  []Book
	output []Book
}

func Test_sortByYear(t *testing.T) {
	tests := []test{
		{[]Book{kerouac, hemingway}, []Book{hemingway, kerouac}},
		{[]Book{kerouac, onan, palahniuk, hemingway}, []Book{hemingway, kerouac, palahniuk, onan}},
		{[]Book{irving, kerouac, onan, palahniuk, hemingway}, []Book{hemingway, kerouac, palahniuk, onan, irving}},
	}
	executeTests(t, tests, BookSort{less: sortByYear})
}

func Test_sortByTitle(t *testing.T) {
	tests := []test{
		{[]Book{kerouac, hemingway}, []Book{kerouac, hemingway}},
		{[]Book{kerouac, onan, palahniuk, hemingway}, []Book{palahniuk, onan, kerouac, hemingway}},
		{[]Book{irving, kerouac, onan, palahniuk, hemingway}, []Book{palahniuk, onan, irving, kerouac, hemingway}},
	}
	executeTests(t, tests, BookSort{less: sortByTitle})
}

func Test_sortByAuthor(t *testing.T) {
	tests := []test{
		{[]Book{kerouac, hemingway}, []Book{hemingway, kerouac}},
		{[]Book{kerouac, onan, palahniuk, hemingway}, []Book{hemingway, kerouac, onan, palahniuk}},
		{[]Book{irving, kerouac, onan, palahniuk, hemingway}, []Book{hemingway, irving, kerouac, onan, palahniuk}},
	}
	executeTests(t, tests, BookSort{less: sortByAuthor})
}

func executeTests(t *testing.T, tests []test, f BookSort) {
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			sort.Sort(BookSort{tt.input, f.less})
			if !reflect.DeepEqual(tt.input, tt.output) {
				t.Errorf("%d = %v, want %v", i, tt.input, tt.output)
			}
		})
	}
}
