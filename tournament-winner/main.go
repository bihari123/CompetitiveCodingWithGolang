package main

import (
	"fmt"

	utils "github.com/bihari123/CompetitiveCodingWithGolang/tournament-winner/utils"
)

var competitions=[][]string{{"HTML","C#"},{"C#","Python"},{"Python","HTML"}}
var results=[]int{0,0,1}

func main(){
	scoreCard:=utils.GetScore(competitions,results)
  fmt.Println("the score card is as follows: ",utils.GetScore(competitions,results))
  fmt.Println(" the winner is ",utils.GetWinner(scoreCard))
}
