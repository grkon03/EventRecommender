package main

import (
	"fmt"
	"recsys/engineV1"
	"recsys/model"
	"recsys/samples"
)

func main() {
	sample := samples.Sample1()
	sense := model.DefaultSense(1.0, 0.5)
	sol := model.Recommend(sample, sense, engineV1.FullSearchEngine())
	fmt.Println("solution: ", sol.Events)
	fmt.Println("score: ", sense.Score(sol))
}
