pacakge main
import "fmt"
func main(){
	fmt.Println(canWinNim(4))	
}

func canWinNim(n int) bool {
    if n % 4 == 0 {
        return false
    }
    return true
    
}
