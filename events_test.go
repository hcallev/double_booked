package events

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOverlaps(t *testing.T) {
	tests := []struct {
		name   string
		events []Event
		want   bool
	}{
		{
			name: "partial overlap",
			events: []Event{
				{Start: 2, End: 4},
				{Start: 3, End: 5},
			},
			want: true,
		},
		{
			name: "full overlap",
			events: []Event{
				{Start: 1, End: 4},
				{Start: 2, End: 3},
			},
			want: true,
		},
		{
			name: "no overlap",
			events: []Event{
				{Start: 1, End: 2},
				{Start: 3, End: 4},
			},
			want: false,
		},
		{
			name: "no overlap edge",
			events: []Event{
				{Start: 1, End: 2},
				{Start: 2, End: 3},
			},
			want: false,
		},
		{
			name: "no overlap edge inverted",
			events: []Event{
				{Start: 2, End: 3},
				{Start: 1, End: 2},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Overlaps(tt.events[0], tt.events[1])
			assert.Equal(t, tt.want, got)
		})
	}
}
