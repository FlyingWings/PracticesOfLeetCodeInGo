package main

import "fmt"

func main(){
	res := fizzBuzz(15)
	for _,v := range(res){
	fmt.Println(v)
	}
}

func fizzBuzz(n int) []string {
	res := make([]string,0)
	for i:=1;i <=n;i++ {
	temp := ""
	if  i % 3 == 0 {
	temp = "Fizz"
	}
	if i % 5 == 0 {
	temp +="Buzz"
	}
	if temp  == "" {
	temp = fmt.Sprintf("%d", i)
	}
	res = append(res,temp)
	}
	
	return res  
}
