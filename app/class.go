package app

// Class represents a class
type Class struct {
	Week     int
	Students []*Student
}

// ByName is a helper type
// to sort classes by name
type ByName []*Class

func (b ByName) Len() int {
	return len(b)
}

func (b ByName) Less(i, j int) bool {
	return b[i].Week < b[j].Week
}

func (b ByName) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
