package main

import "fmt"

func main() {
        fmt.Println(hammingDistance(4,1));
}
func hammingDistance(x int, y int) int {
    big := 0
    if x > y {
    big = x
    }else{
    big = y
    }
    small := x+y-big
    big_list := [32]int{0}
    small_list := [32]int{0}
    c := 0
    for big > 0 {
        big_list[c] = big%2
        small_list[c] = small % 2
        big/=2
        small/=2
        c++
     }
    count := 0
    for i,v := range(big_list) {
        if v != small_list[i] {
        count++;
        }
    }
    return count ;
}

