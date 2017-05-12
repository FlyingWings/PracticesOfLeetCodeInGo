package main
import "fmt"
func main(){
   value := []int {1,1,0,1,1,1}
	fmt.Println(findMaxConsecutiveOnes(value))	
}

func findMaxConsecutiveOnes(nums []int) int {
    max,length,flag,lens := 0,0,false,len(nums)-1
    for k,v := range(nums) {
        
        if flag == false && v == 1 {
            length = 1
            flag = true
            
        }else if flag == true && v == 1 {
            length++
        }else if flag == true && v == 0 {
            if length > max {
                max = length
            }
            length = 0
            flag = false
        }
        
        if k == lens {
            if length > max {
                max = length
            }
            return max
        }
    }
    return max
}
