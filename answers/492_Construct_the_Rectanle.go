package main
import "fmt"
import "math"
func main(){
        fmt.Println(constructRectangle(63))
}

func constructRectangle(area int) []int {
	res := make([]int,2)
	sq := int(math.Sqrt(float64(area)))
	if sq * sq == area {
	res[0],res[1] = sq,sq
	return res
	}
	L,W := sq+1,0
	for L<=area {
	  W = area / L
	  if L * W == area {
	  res[0],res[1] = L,W
		return res
	}
	L++
	}
	res[0],res[1] = area,1
	return res
}
