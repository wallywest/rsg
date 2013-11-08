package generator

import (
	"github.com/nu7hatch/gouuid"
)

type Range struct {
	days  string
	start int
	end   int
}

const (
	MINUTES_IN_DAY = 1440
)

func SplitRanges(split int, days string) []Segment {
	var e, s int
	segments := make([]Segment, split)
	max := MINUTES_IN_DAY / split
	s = 0
	e = (max - 1)
	u, _ := uuid.NewV4()
	r := Segment{Id: u.String(), Start: s, End: e, Days: days}
	segments[0] = r
	for i := 1; i < split; i++ {
		s = s + max
		e = (s + max) - 1
		u, _ := uuid.NewV4()
		r := Segment{Id: u.String(), Start: s, End: e, Days: days}
		segments[i] = r
	}
	return segments
}
