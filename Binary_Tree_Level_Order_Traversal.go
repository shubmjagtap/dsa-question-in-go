func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }

    ans := [][]int{}
    queue := list.New()
    queue.PushBack(root)

    for queue.Len() != 0 {
        count := queue.Len()
        subAns := []int{}

        for i := 0; i < count; i++ {
            frontElement := queue.Front().Value.(*TreeNode)
            queue.Remove(queue.Front())

            if frontElement.Left != nil {
                queue.PushBack(frontElement.Left)
            }

            if frontElement.Right != nil {
                queue.PushBack(frontElement.Right)
            }

            subAns = append(subAns, frontElement.Val)
        }

        ans = append(ans, subAns)
    }

    return ans
}
