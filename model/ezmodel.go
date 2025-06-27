package model

import (
	"recsys/util"
	"sort"
)

// will not be used...?
type Club struct {
	ID int
}

type Event struct {
	ID           int
	ClubID       int
	SectionStart int // user cannot participate the events with the same section
	SectionEnd   int
}

func NewEvent(id, clubID, sectionStart, sectionEnd int) Event {
	return Event{ID: id, ClubID: clubID, SectionStart: sectionStart, SectionEnd: sectionEnd}
}

type Problem struct {
	Clubs  []Club
	Events []Event
}

// caution: assuming subevents is subset of p.Events (this method does not assert)
func (p Problem) SubProblem(subevents []Event) Problem {
	return Problem{Clubs: p.Clubs, Events: subevents}
}

func (p Problem) DeepCopy() Problem {
	var new Problem
	util.DeepCopy(p, &new)
	return new
}

type solution struct {
	Events  []Event
	Problem Problem
}

type Solution = *solution

func EmptySolution(p Problem) Solution {
	return &solution{Events: []Event{}, Problem: p}
}

func (s Solution) AddEvent(e Event) {
	s.Events = append(s.Events, e)
}

func (s Solution) AddEvents(es []Event) {
	s.Events = append(s.Events, es...)
}

func (s Solution) DeepCopy() Solution {
	var new Solution
	util.DeepCopy(s, &new)
	return new
}

func (s Solution) SortBySectionStart() {
	sort.Slice(s.Events, func(i, j int) bool {
		return s.Events[i].SectionStart < s.Events[j].SectionStart
	})
}

func (s Solution) IsThereConflict() bool {
	s.SortBySectionStart()

	lastEnd := s.Events[0].SectionEnd

	for _, e := range s.Events[1:] {
		if e.SectionStart < lastEnd {
			return true
		}
		lastEnd = e.SectionEnd
	}

	return false
}
