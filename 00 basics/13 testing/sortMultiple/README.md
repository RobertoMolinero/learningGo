# Sorting an object according to different attributes

In this example a list of books should be sorted.

If you only want to sort by one attribute you can implement the interface 'Interface' from 'sort/sort.go' for your own type.

```
type bookSortByYear []Book

func (b bookSortByYear) Len() int {
	return len(b)
}

func (b bookSortByYear) Less(i, j int) bool {
	return b[i].Year < b[j].Year
}

func (b bookSortByYear) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
```

If you now sort the same list for another attribute you will quickly see that much code just doubles.

```
type bookSortByTitle []Book

func (b bookSortByTitle) Len() int {
	return len(b)
}

func (b bookSortByTitle) Less(i, j int) bool {
	return b[i].Title < b[j].Title
}

func (b bookSortByTitle) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
```

The only part that differs is the method 'Less'. To remove the duplication one creates another type which has a list of objects and the definition of the method to compare. 

```
type BookSort struct {
	set  []Book
	less func(i, j Book) bool
}

func (b BookSort) Less(i, j int) bool {
	return b.less(b.set[i], b.set[j])
}
```

This method is then passed as a parameter.

```
sort.Sort(BookSort{tt.input, f.less})
```

To make sure everything worked, I wrote a test class. To run the tests type 'go test'. For more information on the test run, use the parameter '-v'.