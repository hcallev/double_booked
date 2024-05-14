package events

type Event struct {
	Start uint32
	End   uint32
}

func Overlaps(a, b Event) bool {
	return a.Start < b.End && a.End > b.Start
}
