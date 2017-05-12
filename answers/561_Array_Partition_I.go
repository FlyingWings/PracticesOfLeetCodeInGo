package main

import "fmt"

func main(){
	nums := []int{7,2,6,1,5,3,4}
	nums = qsort(nums)
	sum := 0
	for i,v := range(nums){
	if i%2 == 0 {
	sum += v
	}
	}
	fmt.Println(sum)
}
func qsort(nums []int) []int {
    if  len(nums) == 1 || len(nums) == 0 {
        return nums
        }
    sub := nums[0]
    right := make([]int,10000)
    left := make([]int,10000)
    r_c,l_c := 0,0
    for i,v := range(nums){
        if i == 0 {
        continue;
        }
        if v >= sub {
            right[r_c] = v
            r_c++
        }else{
            left[l_c] = v
            l_c++
        }
    }
    left[l_c] = sub
    return append(qsort(left),qsort(right)...)
}

