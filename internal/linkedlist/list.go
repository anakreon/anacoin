package linkedlist

type List struct {
	head     *Node
	mainTail *Node
	allTails []*Node
}

func NewList(initialData NodeData) *List {
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

func (list *List) CreateNewNodeAndLinkWithPreviousNode(previousNode *Node, newData NodeData) {
	if !list.containsNode(previousNode) {
		return
	}
	newNode := NewNode(previousNode, newData)
	list.addTail(newNode)
	if !previousNode.HasNext() {
		list.linkAndCheckIfNewMainTail(newNode, previousNode)
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

func (list *List) GetMainTailData() NodeData {
	return list.mainTail.data
}

func (list *List) relinkMainChain() {
	for iteratorNode := list.mainTail; iteratorNode.HasPrevious(); iteratorNode = iteratorNode.GetPrevious() {
		iteratorNode.GetPrevious().SetNext(iteratorNode)
	}
}

func (list *List) containsNode(node *Node) bool {
	return nil != list.FindNodeInList(func(loopNode *Node) bool {
		return loopNode == node
	})
}

type CorrectNodeCallback func(node *Node) bool

func (list *List) FindNodeInList(isCorrectNode CorrectNodeCallback) *Node {
	var resultNode *Node = nil
Mainloop:
	for _, currentTail := range list.allTails {
		for node := currentTail; node.HasFullLinkWithPreviousNode(); node = node.GetPrevious() {
			if isCorrectNode(node) {
				resultNode = node
				break Mainloop
			}
		}
	}
	return resultNode
}