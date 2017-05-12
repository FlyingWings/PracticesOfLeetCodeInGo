package main
import "fmt"
func main(){
	fmt.Println(detectCapitalUse("USA"))	
}

func detectCapitalUse(word string) bool {
    res,length := 0,len(word)
    for _,v := range(word){
        if v >= 65 && v <= 91 {
            res++
        }
    }
    if res == length || res == 0 || res == 1 && word[0] <= 91 && word[0] >= 65{
        return true
    }
    return false
}
