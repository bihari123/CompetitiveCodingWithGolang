package main

import (
	"sort"
)
func getThreeNumberSum( input []int,target int)(result [][]int){
result =[][]int{}
sort.Ints(input)

for i:=0;i<len(input)-2;i++{
  left,right:= i+1,len(input)-1

  for left <right {
    currentSum:=input[i] +input[left] +input[right]

    if currentSum==target{
      result=append(result, []int{input[i],input[left],input[right]})
      left+=1
      right -=1
    }else if currentSum <target {
      left +=1
    }else if target>currentSum{
      right -=1
    }
  }
}

return
}
