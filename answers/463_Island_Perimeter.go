package main

import "fmt"

func main(){
	island := [][]int{ {0,1,0,0},{1,1,1,0},{0,1,0,0},{1,1,0,0}}
	island = [][]int { {0},{1}}
	fmt.Println(islandPerimeter(island))
}
func islandPerimeter(grid [][]int) int {
    count,side,r_size,c_size := 0,0,len(grid),len(grid[0])
   
    for r_k,r_v := range(grid){
        for c_k,c_v := range(r_v){
            if c_v == 1 {
                count++
                if c_k - 1 >= 0 {
                    side += r_v[c_k-1]
                }
                if c_k + 1 < c_size {
                    side += r_v[c_k+1]
                }
                if r_k - 1 >= 0 {
                    side += grid[r_k-1][c_k]
                }
                if r_k + 1 < r_size {
                    side += grid[r_k+1][c_k]
                }
            }
        }
    }
    return 4*count-side
}

