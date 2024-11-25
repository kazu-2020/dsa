package main

import (
	"fmt"
)

func main() {
	var A, B, C, D, E, F int
	fmt.Scan(&A, &B, &C, &D, &E, &F)

	// DPテーブル: dp[i][j]は砂糖数の総量i、砂糖の総量jが作れるかどうか
	dp := make([][]bool, F+1)
	for i := range dp {
		dp[i] = make([]bool, F+1)
	}
	dp[0][0] = true

	var bestWater, bestSugar int
	maxConcentration := 0.0

	// DP更新
	for total := 0; total <= F; total++ {
		for sugar := 0; sugar <= total; sugar++ {
			water := total - sugar
			if !dp[total][sugar] {
				continue
			}

			// 水を追加
			if total+100*A <= F {
				dp[total+100*A][sugar] = true
			}
			if total+100*B <= F {
				dp[total+100*B][sugar] = true
			}

			// 砂糖を追加
			if sugar+C <= total && total+C <= F {
				dp[total+C][sugar+C] = true
			}
			if sugar+D <= total && total+D <= F {
				dp[total+D][sugar+D] = true
			}

			// 現在の濃度を計算
			if water > 0 && sugar*100 <= water*E {
				concentration := float64(sugar) * 100 / float64(total)
				if concentration > maxConcentration {
					maxConcentration = concentration
					bestWater = total
					bestSugar = sugar
				}
			}
		}
	}

	if bestWater == 0 {
		bestWater = 100 * A
	}

	fmt.Printf("%d %d\n", bestWater, bestSugar)
}
