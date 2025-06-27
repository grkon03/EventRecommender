package model

import "recsys/util"

type MostNumberOfEventsAsPossible struct{}

func (MostNumberOfEventsAsPossible) Score(s Solution) (score float32) {
	return float32(len(s.Events))
}

type MostKindOfClubsAsPossible struct{}

func (MostKindOfClubsAsPossible) Score(s Solution) (score float32) {
	clubIDs := util.NewSet[int]()
	for _, e := range s.Events {
		clubIDs.Add(e.ClubID)
	}
	return float32(clubIDs.Count())
}

func DefaultSense(weightOfMostNum, weightOfMostKind float32) (s Sense) {
	s = MakeSense()
	s.Add(MostNumberOfEventsAsPossible{}, weightOfMostNum)
	s.Add(MostKindOfClubsAsPossible{}, weightOfMostKind)

	return
}
