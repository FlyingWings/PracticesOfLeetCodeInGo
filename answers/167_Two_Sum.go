package main
import "fmt"
func main() {
numbers,ta := []int{2,3,4,6},7
	fmt.Println(twoSum(numbers,ta))
}

func twoSum(numbers []int, target int) []int {
    res := []int{0,0}
    pos  := 0
    for k,v := range(numbers) {
        pos = binarySearch(numbers[k+1:],target-v)
            if pos != 0 {
                res[0],res[1] = k+1,pos+k+2
                return res
            }

    }
    return res
}

func binarySearch( numbers []int, target int) int{
    pos := 0
    if len(numbers) == 0 {
        return 0
        }
    if len(numbers) == 1 {
        return 1
        }
    if target <= numbers[len(numbers)/2-1] {
        return binarySearch(numbers[:len(numbers)/2],target)
    }else {
        pos = binarySearch(numbers[len(numbers)/2:],target)
        return pos+len(numbers)/2-1
    }
}

