# Event Recommender

## エンジンの回し方

1. `model.Problem` に全イベントとイベントを主催する全クラブの情報を保存します
2. `model.Sense` にユーザーの価値基準を追加します
3. `model.Engine` にエンジンを指定します
4. `model.Recommend` 関数にこれらの情報を渡すことで、`model.Solution` として解を受け取れます
5. `model.Solution.Events` から価値基準に適したイベント一覧を取得します

```go:main.go
package main

import (
	"fmt"
	"recsys/engineV1"
	"recsys/model"
	"recsys/samples"
)

func main() {
    // 1. 問題として、問題サンプル3を使います。
    sample := samples.Sample3()
    // 2. 価値基準として、DefaultSense(後述)を用います
    sense := model.DefaultSense(1.0, 0.5)
    // 3. エンジンとして、BinaryPartitionEngine(後述)を用います
    engine := engineV1.BinaryPartitionEngine()
    // 4. Recommend 関数に情報を渡して解を得ます
    sol := model.Recommend(sample, sense, engine)
    // 5. 解から参加するイベント情報を取得します
    fmt.Println("solution: ", sol.Events)
}

```

## サンプル

`samples`パッケージには`model.Problem` のサンプルがいくつか置いてあります。

## 標準エンジン

あらかじめ用意したエンジンです

- `engineV1.FullSearchEngine()`
  - 全探索をします
  - 常に最適解を出します
  - イベント数がおおよそ 20 を超えたあたりから時間がかかり始めます
- `engineV1.BinaryPartitionEngine()`
  - Quick Sort のように Binary Partitioning を用いて実装しています
  - 最適解が出ないですが、おおよそ近いスコアの解を計算します
  - "〜に 3 個以上参加する" や "〜を 1 回以下に抑える" といった絶対数で指定する条件におそらく弱いです
  - イベント個数が 100 個でも 2 秒程度で計算します

## 価値基準について

デフォルトの価値基準として、`model.DefaultSense(w1, w2 float64)`関数が用意されています。
この価値基準では、"参加するイベントが多いほど良い"と"参加するイベントを主催するクラブの種類が多いほど良い"という価値基準を採用しており、`w1`, `w2` にそれらの重みづけをします。

`model.Sense` を設定することによって、そのほかにもユーザーの価値基準に対応できます。
価値基準を追加するためには新しい価値基準 `model.Evaluation` を作成し、`model.Sense.Add` 関数により重みと共に追加します。
`model.Evaluation` は `Score(model.Solution) float64` メソッドを持つインターフェイスです。この値が大きいほどユーザーが好むスケジュールであることを意味しています。

```go:newsense.go
package newsense

import (
    "math"
    "recsys/model"
)

// 新しい価値基準を定義する

/*
価値基準:
ClubID が 1 であるクラブのイベントにたくさん参加したいが、参加するイベントが多い場合はそれ以上増やしてもそこまで嬉しさは変わらない
*/
type NewEval struct {}
func (NewEval) Score(s model.Solution) float64 {
    count := 0
    for _, e := range s.Events {
        if e.ClubID == 1 {
            count++
        }
    }

    return math.Sqrt(float64(count))
}

// この価値基準を考慮に入れた model.Sense を返す関数

func NewSense(w1, w2, w3 float64) model.Sense {
    sense := model.DefaultSense(w1, w2)
    sense.Add(NewEval{}, w3)

    return sense
}
```

## エンジンを作成する

エンジンは、`model.Engine`インターフェイスのインスタンスです。新たに作成する場合には、

```go
Run(model.Problem, model.Sense) model.Solution
```

メソッドを持つ型を定義してください。
