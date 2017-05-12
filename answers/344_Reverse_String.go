package main
import "fmt"

func main(){
	v := "hello"
	fmt.Println(reverseString(v))
}

func reverseString(s string) string {
    length := len(s)
    temp := make([]rune,length)
    for i,v := range(s){
        temp[length-i-1] = v
    }
    return string(temp)
}
