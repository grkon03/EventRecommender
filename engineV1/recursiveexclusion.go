package engineV1

import (
	"recsys/model"
	"recsys/util"
)

type RecursiveExclusion struct {
	additionalEvents []model.Event
}

func RecursiveExclusionEngine() RecursiveExclusion {
	return RecursiveExclusion{
		additionalEvents: []model.Event{},
	}
}
func (origin RecursiveExclusion) childREE(additional model.Event) RecursiveExclusion {
	newAddiEvents := make([]model.Event, len(origin.additionalEvents))
	copy(newAddiEvents, origin.additionalEvents)
	newAddiEvents = append(newAddiEvents, additional)

	return RecursiveExclusion{additionalEvents: newAddiEvents}
}

func (engine RecursiveExclusion) fullSearch(p model.Problem, sense model.Sense) model.Solution {
	var optimal, consider model.Solution
	optimal = model.EmptySolution(p)
	var optscore, conscore float64 = 0, 0

	eventset := util.MakeSetFromSlice(p.Events)

	for subset := range eventset.ForEachSubset() {
		consider = model.EmptySolution(p)
		consider.AddEvents(subset.Elements())
		consider.AddEvents(engine.additionalEvents)
		if len(consider.Events) == 0 {
			continue
		}
		if consider.IsThereConflict() {
			continue
		}

		conscore = sense.Score(consider)
		if optscore < conscore {
			optimal = consider.DeepCopy()
			optscore = conscore
		}
	}

	return optimal
}

func (engine RecursiveExclusion) Run(p model.Problem, sense model.Sense) model.Solution {
	if len(p.Events) <= 14 {
		return engine.fullSearch(p, sense)
	}

	excludeSection := p.Events[0].SectionStart

	var excludes, remains []model.Event

	for _, ev := range p.Events {
		if ev.SectionStart <= excludeSection && excludeSection < ev.SectionEnd {
			excludes = append(excludes, ev)
		} else {
			remains = append(remains, ev)
		}
	}

	var optimal, consider model.Solution
	var optscore, conscore float64
	var restricted []model.Event

	optimal = engine.Run(p.SubProblem(remains), sense)
	optscore = sense.Score(optimal)

	for _, ev := range excludes {
		restricted = make([]model.Event, 0)
		for _, rev := range remains {
			if rev.SectionStart >= ev.SectionEnd || rev.SectionEnd <= ev.SectionStart {
				restricted = append(restricted, rev)
			}
		}

		consider = engine.childREE(ev).Run(p.SubProblem(restricted), sense)
		conscore = sense.Score(consider)

		if optscore < conscore {
			optimal = consider.DeepCopy()
			optscore = conscore
		}
	}

	return optimal
}
