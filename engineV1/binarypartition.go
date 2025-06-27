package engineV1

import (
	"recsys/model"
	"sort"
)

type BinaryPartition struct {
	start int
	end   int
	child bool
}

/*
Calculate solution by binary partioning

Caution 1: The ptimality is not guaranteed

Caution 2: There are a few types of "Sense" for which this engine will submit too bad solution
e.g.) The case like an user doesn't want to participate more than 5 events
*/
func BinaryPartitionEngine() BinaryPartition { return BinaryPartition{child: false} }
func childBPE(start, end int) BinaryPartition {
	return BinaryPartition{start: start, end: end, child: true}
}

func (engine BinaryPartition) Run(p model.Problem, sense model.Sense) model.Solution {
	num := len(p.Events)

	// to find the best number here is good issue to consider
	if num <= 8 {
		return FullSearchEngine().Run(p.DeepCopy(), sense)
	}

	events := make([]model.Event, num)
	copy(events, p.Events)

	sort.Slice(events, func(i, j int) bool {
		return p.Events[i].SectionStart < p.Events[j].SectionStart
	})

	// start, end must be determined only in the branch
	if !engine.child {
		engine.start = events[0].SectionStart
		for _, ev := range events {
			if engine.end < ev.SectionEnd {
				engine.end = ev.SectionEnd
			}
		}
	}

	midindex := num / 2
	midsection := (events[midindex].SectionStart + events[midindex].SectionEnd) / 2

	var eventsFormer, eventsMiddle, eventsLatter []model.Event
	for _, ev := range events {
		if ev.SectionEnd <= midsection {
			eventsFormer = append(eventsFormer, ev)
		} else if ev.SectionStart >= midsection {
			eventsLatter = append(eventsLatter, ev)
		} else {
			eventsMiddle = append(eventsMiddle, ev)
		}
	}

	var optimal, consider model.Solution
	var optscore, conscore float64 = 0, 0

	var formerP, latterP model.Problem
	var formerS, latterS model.Solution

	// first, we consider the case without any middle events

	formerP = p.SubProblem(eventsFormer)
	latterP = p.SubProblem(eventsLatter)
	formerS = childBPE(engine.start, midsection).Run(formerP, sense)
	latterS = childBPE(midsection, engine.end).Run(latterP, sense)

	optimal = model.EmptySolution(p)
	optimal.AddEvents(formerS.Events)
	optimal.AddEvents(latterS.Events)
	optscore = sense.Score(optimal)

	// second, we consider the case with each middle event

	var apprEvents []model.Event

	for _, midev := range eventsMiddle {
		// solve the former problem

		start := engine.start
		end := midev.SectionStart

		apprEvents = make([]model.Event, 0)
		for _, ev := range eventsFormer {
			if ev.SectionEnd <= end {
				apprEvents = append(apprEvents, ev)
			}
		}
		formerP = p.SubProblem(apprEvents)
		formerS = childBPE(start, end).Run(formerP, sense)

		// solve the latter problem

		start = midev.SectionEnd
		end = engine.end

		apprEvents = make([]model.Event, 0)
		for _, ev := range eventsLatter {
			if ev.SectionStart >= start {
				apprEvents = append(apprEvents, ev)
			}
		}
		latterP = p.SubProblem(apprEvents)
		latterS = childBPE(start, end).Run(latterP, sense)

		// combine

		consider = model.EmptySolution(p)
		consider.AddEvents(formerS.Events)
		consider.AddEvent(midev)
		consider.AddEvents(latterS.Events)
		conscore = sense.Score(consider)

		if optscore < conscore {
			optimal = consider.DeepCopy()
		}
	}

	return optimal
}
