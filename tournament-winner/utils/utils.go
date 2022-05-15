package utils

func GetScore(competition [][]string, result []int) (scoreCards map[string]int) {
	var matchWinner string
	scoreCards = map[string]int{}
	for index, val := range competition {
		if result[index] == 1 {
			matchWinner = val[0]
		} else {
			matchWinner = val[1]
		}

		if _, ok := scoreCards[matchWinner]; ok {
			scoreCards[matchWinner] += 3
		} else {
			scoreCards[matchWinner] = 3
		}
	}
	return
}

func GetWinner(scoreCards map[string]int) (winner string) {
	max := 0
	for team, score := range scoreCards {
		if score > max {
			max, winner = score, team
		}
	}
	return
}
