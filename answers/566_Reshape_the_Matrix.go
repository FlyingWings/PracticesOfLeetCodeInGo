package main
import "fmt"

func main(){
	temp := [][]int {{1,2},{3,4}}
	
	res := matrixReshape(temp,4,1)
	for _,v := range(res) {
		fmt.Print("[")
                for _,v1 := range(v) {
                        fmt.Print(v1,",")
                }
		fmt.Println("]")
        }

}

func matrixReshape(nums [][]int, r int, c int) [][]int {
	temp := make([]int,0)
	c_count, r_count, temp_count := 0,0,0
	for _,v := range(nums) {
		for _,v1 := range(v) {
			temp = append(temp,v1)
		}
	}
	fmt.Println(len(temp))
	if r*c != len(temp) {
	return nums
	}else{

	result := make([][]int,r,r)
	for r_count < r {
		c_count = 0;
		for c_count < c {
			result[r_count] = append(result[r_count],temp[temp_count])
			temp_count++
			c_count++
		}
		r_count++;
	}
	return result
	}	
}
