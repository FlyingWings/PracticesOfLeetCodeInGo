package main
import "fmt"

func main(){
	s := "I'm ok"
	fmt.Println(readWords(s))
}

func reverseWords(s string) string {
	t := ""
	for _,v := range(s){
	t = string(v)+t
	}
	return t
}

func readWords(s string) string {
	temp,total := "",""
	for _,v := range(s){
	if string(v) != " " {
	temp += string(v)
	}else{
	total += reverseWords(temp)+" "
	temp = ""
	}
    }
    total += reverseWords(temp)
    return total
}
