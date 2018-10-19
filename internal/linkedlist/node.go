//go:generate genny -in=$GOFILE -out=../linkedlist-gen/$GOFILE -pkg=linkedlist gen "NodeData=block.Block"

package linkedlisttpl

import (
	"github.com/cheekybits/genny/generic"
)

type Node struct {
	index        uint64
	previousNode *Node
	nextNode     *Node
	data         NodeData
}

type NodeData generic.Type

func NewNode(previousNode *Node, data NodeData) *Node {
	return &Node{
		index:        previousNode.index + 1,
		previousNode: previousNode,
		nextNode:     nil,
		data:         data,
	}
}

func (node *Node) HasNext() bool {
	return node.nextNode != nil
}

func (node *Node) SetNext(nextNode *Node) {
	node.nextNode = nextNode
}

func (node *Node) HasPrevious() bool {
	return node.previousNode != nil
}

func (node *Node) GetPrevious() *Node {
	return node.previousNode
}

func (node *Node) GetData() *NodeData {
	return &node.data
}

func (node *Node) HasFullLinkWithPreviousNode() bool {
	return node.previousNode == nil || node.previousNode.nextNode == node
}

func (node *Node) HasIndexLargerThan(otherNode *Node) bool {
	return node.index > otherNode.index
}
