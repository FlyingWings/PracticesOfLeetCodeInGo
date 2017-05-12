package main
import "fmt"
func main(){
	
}



/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {
    res := midOrderWalk(root)
    small := abs(res[0]-res[1])
    length := len(res)
    temp := 0
    for i:=1;i<length-1;i++{
        temp = abs(res[i]-res[i+1])
        if temp < small {
            small = temp
        }
    }
    return small
    
}

func midOrderWalk(node *TreeNode)[]int {
    if node == nil {
        return nil
    }
    
    left, right := midOrderWalk(node.Left),midOrderWalk(node.Right)
    left = append(left,node.Val)
    return append(left,right...)
    
}

func abs(num int) int {
    if num < 0 {
        return -num
    }
    return num
}
