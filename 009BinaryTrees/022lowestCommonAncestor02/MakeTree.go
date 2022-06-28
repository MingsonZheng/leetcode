package main

func MakeTree(array []int) *TreeNode {
	if array == nil || len(array) == 0 {
		return nil
	}
	index := 0
	length := len(array)
	root := &TreeNode{Val: array[0]}
	nodeQueue := Constructor()
	nodeQueue.AppendTail(root)
	var curTreeNode *TreeNode
	for index < length {
		index++
		if index >= length {
			return root
		}
		curTreeNode = nodeQueue.DeleteHead()
		leftChild := array[index]
		if leftChild != -1 {
			curTreeNode.Left = &TreeNode{Val: leftChild}
			nodeQueue.AppendTail(curTreeNode.Left)
		}
		index++
		if index >= length {
			return root
		}
		rightChild := array[index]
		if rightChild != -1 {
			curTreeNode.Right = &TreeNode{Val: rightChild}
			nodeQueue.AppendTail(curTreeNode.Right)
		}
	}
	return root
}
