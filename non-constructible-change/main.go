package main

import (
	"fmt"
	"sort"
)

func NonConstructibleChange(coins []int) int {
	// Write your code here.
	sort.Ints(coins)
	if len(coins)<1{
	  return 1
	}
	if len(coins)==1{
	  if coins [0]>1{
	    return 1 
	  }else{
	    return 2
	  }
	}
	sum :=coins[0]
	minChange:=-1
	for index, v := range coins {
    if v<=sum+1 {
     if index>0{
       sum+=v
       minChange=sum+1
     }
    	}
    }
	  return minChange
}

func main(){
  coins:=[]int{1, 5, 1, 1, 1, 10, 15, 20, 100}
  sort.Ints(coins)
  fmt.Println(" the sorted value is ",coins)

  fmt.Println("the minimum change that we can't make is ",NonConstructibleChange(coins))
}
