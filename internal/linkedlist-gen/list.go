// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package linkedlist

import "github.com/anakreon/anacoin/internal/block"

type List struct {
	head     *Node
	mainTail *Node
	allTails []*Node
}

func NewList(initialData block.Block) *List {
	headNode := &Node{
		index:        0,
		previousNode: nil,
		nextNode:     nil,
		data:         initialData,
	}
	return &List{
		head:     headNode,
		mainTail: headNode,
		allTails: []*Node{headNode},
	}
}

func (list *List) AddNode(previousData *block.Block, nodeData block.Block) {
	previousNode := list.getNodeByData(previousData)
	if previousNode != nil {
		newNode := NewNode(previousNode, nodeData)
		list.addTail(newNode)
		if !previousNode.HasNext() {
			list.linkAndCheckIfNewMainTail(newNode, previousNode)
		}
	}
}

func (list *List) addTail(node *Node) {
	list.allTails = append(list.allTails, node)
}

func (list *List) removeTail(node *Node) {
	for tailIndex, tail := range list.allTails {
		if tail == node {
			list.allTails[tailIndex] = list.allTails[len(list.allTails)-1]
			list.allTails = list.allTails[:len(list.allTails)-1]
			break
		}
	}
}

func (list *List) linkAndCheckIfNewMainTail(node *Node, previousNode *Node) {
	previousNode.SetNext(node)
	list.removeTail(previousNode)
	list.setAsMainTailIfIndexLarger(node)
}

func (list *List) setAsMainTailIfIndexLarger(node *Node) {
	if node.HasIndexLargerThan(list.mainTail) {
		list.setMainTail(node)
		list.relinkMainChain()
	}
}

func (list *List) setMainTail(node *Node) {
	list.mainTail = node
}

func (list *List) GetMainTailData() *block.Block {
	return list.mainTail.GetData()
}

func (list *List) relinkMainChain() {
	for iteratorNode := list.mainTail; iteratorNode.HasPrevious(); iteratorNode = iteratorNode.GetPrevious() {
		iteratorNode.GetPrevious().SetNext(iteratorNode)
	}
}

func (list *List) getNodeByData(nodeData *block.Block) *Node {
	return list.findNodeInList(func(loopNode *Node) bool {
		return &(loopNode.data) == nodeData
	})
}

type correctNodeCallback func(node *Node) bool

func (list *List) findNodeInList(isCorrectNode correctNodeCallback) (resultNode *Node) {
Mainloop:
	for _, currentTail := range list.allTails {
		for node := currentTail; node.HasFullLinkWithPreviousNode(); node = node.GetPrevious() {
			if isCorrectNode(node) {
				resultNode = node
				break Mainloop
			}
		}
	}
	return
}