package engineV1

import (
	"recsys/model"
	"recsys/util"
)

type FullSearch struct{}

// it seems to be able to calculate quikly only in the case up to 20 events
func FullSearchEngine() FullSearch { return FullSearch{} }

func (e FullSearch) Run(p model.Problem, sense model.Sense) model.Solution {
	var optimal, consider model.Solution
	optimal = model.EmptySolution(p)
	var optscore, conscore float64 = 0, 0

	eventset := util.MakeSetFromSlice(p.Events)

	for subset := range eventset.ForEachSubset() {
		consider = model.EmptySolution(p)
		consider.AddEvents(subset.Elements())
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
