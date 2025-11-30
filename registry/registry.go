package registry

type PartFunc func() any

type Entry struct {
	Day   int
	Part  int
	Label string
	Fn    PartFunc
}

var entries []Entry

func Register(day, part int, label string, fn PartFunc) {
	entries = append(entries, Entry{
		Day:   day,
		Part:  part,
		Label: label,
		Fn:    fn,
	})
}

func All() []Entry {
	return entries
}
