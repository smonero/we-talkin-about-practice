func convertBST(root *TreeNode) *TreeNode {

    stack := []*TreeNode{}
    node := root
    runningSum := 0
    
    for (len(stack) != 0 || node != nil) {
        for node != nil {
            stack = append(stack, node)
            node = node.Right
        }
        
        // pop the last item off the stack
        lastIndex := len(stack) - 1
        node = stack[lastIndex]
        stack = stack[:lastIndex]
        runningSum += node.Val
        node.Val = runningSum
        // we are at the hrightmost, go leeft once
        node = node.Left
    }
    
    
    
    
    return root
    
}
