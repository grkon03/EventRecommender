package main

import (
	"fmt"
	"recsys/engineV1"
	"recsys/model"
	"recsys/samples"
	"time"
)

func main() {
	sample := samples.Sample6()
	sense := model.DefaultSense(1.0, 0.5)
	start := time.Now()
	engine := engineV1.CustomBinaryPartitionEngine(engineV1.RecursiveExclusionEngine(), 20)
	sol := model.Recommend(sample, sense, engine)
	processTime := time.Since(start)
	fmt.Println("solution: ", sol.Events)
	fmt.Println("score: ", sense.Score(sol))
	fmt.Println("time: ", processTime)
}
