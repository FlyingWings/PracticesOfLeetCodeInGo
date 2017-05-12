package main
import "fmt"
func main(){
	fmt.Println(addDigits(19))
}

func addDigits1(num int) int {
    if num / 10 == 0 {
        return num
    }
    sum := 0
    for num  > 0 {
        sum += num % 10
        num /= 10
    }
    return addDigits(sum)
    
}

func addDigits(num int) int {
    if num / 10 == 0 {
        return num
    }
    return addDigits(addDigits(num /10)+ num % 10)
    
}
