package main

import (
"fmt"
"sort"
)

func main(){
   g := []int { 1,2,3}
   s := []int { 3}
   fmt.Println(findContentChildren(g,s))
}
func findContentChildren(g []int, s []int) int {
    //sort g ,sort s
    //start from max element in g , walk through s from min
    //found equal :++ , break, if > max s , fail 
    sort.Ints(g)
    sort.Ints(s)
    leng := 0
    
    for i:=len(g)-1;i>=0;i-- {
        if len(s) == 0 {
        return leng
        }
        if g[i] > s[len(s)-1]{
            continue
        }
        if g[i] < s[0] {
            leng++
            if len(s) > 1{
                s = s[1:] 
            }else {
                s = []int{}
            }
            continue
        }
        for j:=0;j<len(s);j++{
            if g[i] <= s[j] {
                leng++
                if len(s) > 1{
                    s = append(s[:j],s[j+1:]...)
                }else{
                    s = []int{}
                }
                
                break
            }
             
        }
    }
    return leng
    
    
}
