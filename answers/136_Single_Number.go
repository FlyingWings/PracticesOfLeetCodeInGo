package main

import "fmt"

func main(){

	req := []int { 1,1,3,2,3,2,6}
	fmt.Println(singleNumber(req))
}

func singleNumber( nums []int) int {
	res := 0
	for _,v := range(nums){
	res ^= v
	}
	return res
}
