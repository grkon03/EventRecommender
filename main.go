package main

import (
	"fmt"
	"recsys/engineV1"
	"recsys/model"
	"recsys/samples"
	"time"
)

func main() {
	sample := samples.Sample1()
	sense := model.DefaultSense(1.0, 0.5)
	start := time.Now()
	sol := model.Recommend(sample, sense, engineV1.FullSearchEngine())
	processTime := time.Since(start)
	fmt.Println("solution: ", sol.Events)
	fmt.Println("score: ", sense.Score(sol))
	fmt.Println("time: ", processTime)
}
