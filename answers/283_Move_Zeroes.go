package main
import "fmt"
func main(){
req := []int {0,1,0,3,12}
moveZeroes(req)
for _,v := range(req){
	fmt.Println(v)
	}	
}

func moveZeroes(nums []int)  {
    leng := len(nums)
    i := 0
    for k, v:= range(nums) {
        if v== 0 {
            
            i = k+1
            for  i< leng {
                if nums[i] != 0 {
                    nums[k],nums[i] = nums[i],nums[k]
                    break
                }
                i++
            }
            if i == leng {
                break
            }
            i = 0
        }
    }
}
