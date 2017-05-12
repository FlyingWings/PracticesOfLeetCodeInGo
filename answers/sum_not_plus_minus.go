package main
import "fmt"
func main(){
	fmt.Println(getSum(-1,1))
}

func getSum(a int, b int) int {
    res,i := 1.0,0
    if ( a < 0){
        for i=a;i< 0;i++ {
        res = res >> 1
        }
    }else{
        for i=0;i< a;i++ {
        res = res << 1
        }
    }
	return res
/*
    if ( b < 0){
        for i=b;i< 0;i++ {
        res >>= 1
        }
    }else{
        for i=0;i< b;i++ {
        res <<= 1
        }
    }

    i= 0
    if (res > 1){
        for res > 1 {
        res >>= 1
        i++
        }
    }else if res < 1{
        for res < 1 {
        res <<= 1
        i--
        }
    }else{
	i = 0
	}
	

    return i
*/
}
