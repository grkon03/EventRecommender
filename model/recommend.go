package model

type Evaluation interface {
	Score(Solution) float32
}

type EvaluationWithWeight struct {
	EvaluationWay Evaluation
	Weight        float32
}

func MakeEvaluationWithWeight(eval Evaluation, weight float32) EvaluationWithWeight {
	return EvaluationWithWeight{eval, weight}
}

type sense struct {
	Evals []EvaluationWithWeight
}

type Sense = *sense

func MakeSense(evals ...EvaluationWithWeight) Sense {
	return &sense{Evals: evals}
}

func (s Sense) AddEWW(eww EvaluationWithWeight) {
	s.Evals = append(s.Evals, eww)
}

func (s Sense) Add(eval Evaluation, weight float32) {
	s.AddEWW(MakeEvaluationWithWeight(eval, weight))
}

func (s Sense) Score(sol Solution) (score float32) {
	for _, eval := range s.Evals {
		score += eval.EvaluationWay.Score(sol) * eval.Weight
	}

	return score
}

type Engine interface {
	Run(Problem, Sense) Solution
}

func Recommend(p Problem, s Sense, e Engine) Solution {
	return e.Run(p, s)
}
