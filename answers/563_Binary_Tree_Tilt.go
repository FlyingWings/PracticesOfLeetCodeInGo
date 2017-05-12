package main

import "fmt"

func main(){
	left := TreeNode{2,nil,nil}
	right:= TreeNode{3,nil,nil}
	root := TreeNode{1,&left,&right}
	fmt.Println(findTilt(&root))
	
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func findTilt(root *TreeNode) int {
_,tilt :=  viewTilt(root)
    return tilt
}

func viewTilt(node *TreeNode) (int,int) {
        if node == nil {
            return 0,0
        }
        left, right, lt,rt := 0,0,0,0
        left,lt = viewTilt(node.Left)
        right,rt = viewTilt(node.Right)
        res := left - right 
        if res < 0 {
            res = -res
        }
        
        return left+right+node.Val, res+lt+rt
}

