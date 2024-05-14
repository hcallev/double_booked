package events

type Event struct {
	Start uint32
	End   uint32
}

func Overlaps(a, b Event) bool {
	return a.Start < b.End && a.End > b.Start
}

func OverlappingEvents(events []Event) [][]Event {
	overlaps := make([][]Event, 0)
	for i, ce := range events {
		for _, ne := range events[i+1:] {
			if Overlaps(ce, ne) {
				overlaps = append(overlaps, []Event{ce, ne})
			}
		}
	}
	return overlaps
}
