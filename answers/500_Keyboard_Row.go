package main

import "fmt"

func main(){
t := []string{"fdsfa","alska"}
res := findWords(t)
for _,v := range(res){
	fmt.Println(v)
}
}
func findWords(words []string) []string {
        alpha := map[string]int {
            "q":1,"w":1,"e":1,"r":1,"t":1,"y":1,"u":1,"i":1,"o":1,"p":1,
	    "a":2,"s":2,"d":2,"f":2,"g":2,"h":2,"j":2,"k":2,"l":2,
	    "z":3,"x":3,"c":3,"v":3,"b":3,"n":3,"m":3,
        }
  
        res := make([]string,0)
        for _,v := range(words) {
                flag := false
                pos_flag := alpha[string(v[0])]
                for _,va := range(v){
                if pos_flag != alpha[string(va)] {
                    flag = true
                }
                }
                if (!flag) {
                    res = append(res,v)
                }
        }
        return res
                
}               
                
func toLower(a string) string {
        res := ""
        for _,v := range(a){
                if (v >= 65 && v <= 91) {
                res+= string(v+32)
                }else{
                res += string(v)
                }
        }
        return res

}      

