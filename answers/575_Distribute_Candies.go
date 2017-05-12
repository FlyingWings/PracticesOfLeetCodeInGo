package main

import "fmt"

func main(){
	data := []int{1,1,2,3}
	can := (distributeCandies(data))
	for _,v := range(can){
	fmt.Println(v)
	}
}

func distributeCandies(candies []int) []int {
    candy_kind := make(map[int] int)
    value_list := make([]int,0)
    for _,v := range(candies) {
	value,ok := candy_kind[v]
	if ok {
	candy_kind[v] = value+1
	}else{
	candy_kind[v]=1
	}
    }
    for _,v := range(candy_kind){
    value_list = append(value_list,v)
    }
    return value_list
}

