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

// CQueue 队列
type CQueue struct {
	stack    []*TreeNode
	tmpStack []*TreeNode
}

func Constructor() CQueue {
	return CQueue{
		stack:    make([]*TreeNode, 0),
		tmpStack: make([]*TreeNode, 0),
	}
}

func (this *CQueue) AppendTail(value *TreeNode) {
	this.stack = append(this.stack, value) // 追加到数组最右边，即入栈
}

func (this *CQueue) DeleteHead() *TreeNode {
	if len(this.stack) == 0 {
		return nil
	}
	for len(this.stack) > 0 {
		//result := this.stack[len(this.stack)-1]// 表示最右边元素，这里表示取出，并未出栈
		this.tmpStack = append(this.tmpStack, this.stack[len(this.stack)-1])
		this.stack = this.stack[:len(this.stack)-1] // 移除 this.stack 最右边元素，即出栈
	}
	result := this.tmpStack[len(this.tmpStack)-1] // this.tmpStack 最右边元素
	this.tmpStack = this.tmpStack[:len(this.tmpStack)-1]
	for len(this.tmpStack) > 0 {
		this.stack = append(this.stack, this.tmpStack[len(this.tmpStack)-1])
		this.tmpStack = this.tmpStack[:len(this.tmpStack)-1]
	}
	return result
}

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
