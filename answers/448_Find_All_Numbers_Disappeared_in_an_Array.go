package main
import "fmt"
func main(){
	nums := []int{4,3,2,7,8,2,3,1}
	res := findDisappearedNumbers(nums)
	for _,v := range(res){
	fmt.Println(v)
	}	
}
func findDisappearedNumbers(nums []int) []int {
    n := len(nums)
    res_map := make(map[int]int)
    for n>0 {
        res_map[n] = 0
        n--
    }
    for _,v := range(nums){
        res_map[v] += 1
    }
    res := make([]int,0)
    for k,v := range(res_map){
        if v == 0 {
            res = append(res,k)
        }
    }
    return res
} 
