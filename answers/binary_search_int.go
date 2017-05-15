package main

import "fmt"

func main(){
	num,ta := []int{1,2,3,4,5,6,7,8},4
	fmt.Println(binarySearch(num,ta))
}

func binarySearch( numbers []int, target int) int{ 
    pos := 0	
    if len(numbers) == 0 {
	return -1
	}
    if target < numbers[len(numbers)/2-1] {
        return binarySearch(numbers[:len(numbers)/2],target)
    }else if target > numbers[len(numbers)/2-1] {
        pos = binarySearch(numbers[len(numbers)/2:],target)
        return pos+len(numbers)/2-1
    }else {
        return len(numbers)/2-1
    }

}
