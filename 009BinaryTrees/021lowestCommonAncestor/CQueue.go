package main

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
