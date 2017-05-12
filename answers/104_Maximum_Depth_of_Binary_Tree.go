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
func maxDepth(root *TreeNode) int {
    return countDepth(root)
}

func countDepth(node *TreeNode) int {
    if node == nil {
        return 0
    }
    left_len,right_len := 0,0
    if node.Left == nil {
        left_len = 0
    }else{
        left_len = countDepth(node.Left)
    }
    if node.Right == nil {
        right_len = 0
    }else{
        right_len = countDepth(node.Right)
    }
    if (left_len > right_len){
        return left_len+1
    }
    return right_len+1
     
}
