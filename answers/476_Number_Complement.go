package main

import "fmt"

func main(){
	fmt.Println(findComplement(5))
}

func findComplement(num int) int {
    sum := 1;
    source := num
    for num > 0 {
        num/=2
        sum*=2
    }
    return sum-source-1
    
}
