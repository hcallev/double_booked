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

func TestOverlappingEvents(t *testing.T) {
	tests := []struct {
		name     string
		events   []Event
		overlaps [][]Event
	}{
		{
			name: "overlapping event",
			events: []Event{
				{Start: 1, End: 3},
				{Start: 2, End: 4},
				{Start: 4, End: 6},
			},
			overlaps: [][]Event{{{Start: 1, End: 3}, {Start: 2, End: 4}}},
		},
		{
			name: "multiple overlapping events",
			events: []Event{
				{Start: 1, End: 3},
				{Start: 2, End: 4},
				{Start: 4, End: 6},
				{Start: 5, End: 7},
			},
			overlaps: [][]Event{
				{{Start: 1, End: 3}, {Start: 2, End: 4}},
				{{Start: 4, End: 6}, {Start: 5, End: 7}},
			},
		},
		{
			name: "event overlapping multiple times",
			events: []Event{
				{Start: 1, End: 5},
				{Start: 2, End: 4},
				{Start: 4, End: 6},
			},
			overlaps: [][]Event{
				{{Start: 1, End: 5}, {Start: 2, End: 4}},
				{{Start: 1, End: 5}, {Start: 4, End: 6}},
			},
		},
		{
			name: "no overlaps",
			events: []Event{
				{Start: 5, End: 6},
				{Start: 1, End: 2},
				{Start: 2, End: 4},
			},
			overlaps: [][]Event{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := OverlappingEvents(tt.events)

			assert.Len(t, got, len(tt.overlaps))
			for _, want := range tt.overlaps {
				assert.Contains(t, got, want)
			}
		})
	}
}
