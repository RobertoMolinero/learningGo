package sortMultiple

type Book struct {
	Title  string
	Author string
	Year   int
}

type BookSort struct {
	set  []Book
	less func(i, j Book) bool
}

func (b BookSort) Len() int {
	return len(b.set)
}

func (b BookSort) Less(i, j int) bool {
	return b.less(b.set[i], b.set[j])
}

func (b BookSort) Swap(i, j int) {
	b.set[i], b.set[j] = b.set[j], b.set[i]
}

func sortByYear(i, j Book) bool   { return i.Year < j.Year }
func sortByTitle(i, j Book) bool  { return i.Title < j.Title }
func sortByAuthor(i, j Book) bool { return i.Author < j.Author }
