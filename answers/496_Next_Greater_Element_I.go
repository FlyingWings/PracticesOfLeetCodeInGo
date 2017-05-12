package main
import "fmt"
func main(){
	//skip
}

func nextGreaterElement(findNums []int, nums []int) []int {
    res := make([]int,len(findNums))
    num_len := len(nums)
    for i,v := range(findNums) {
        if v == nums[num_len-1] {
            res[i] = -1
            continue
        }
        for k,va := range(nums){
            val :=-1
            if k == num_len -1 {
                res[i] = -1
                break
            }
            
            if va == v {
                for pos:= k+1;pos < num_len;pos++ {
                    if nums[pos] > va {
                        val = nums[pos]
                        break
                    }
                }
                res[i] = val
                break
            }
            
        }
    }
    return res
}
