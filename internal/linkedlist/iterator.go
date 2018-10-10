package linkedlist

type listIterator struct {
	list             *List
	currentTailIndex int
	nextNode         *Node
}

func (list *List) Iterator() listIterator {
	return listIterator{
		list:             list,
		currentTailIndex: 0,
		nextNode:         list.allTails[0],
	}
}

func (iterator *listIterator) HasNext() bool {
	return iterator.nextNode != nil
}

func (iterator *listIterator) Next() NodeData {
	nextNode := iterator.nextNode
	iterator.goToNextNode()
	return nextNode.data
}

func (iterator *listIterator) goToNextNode() {
	if iterator.nextNode != nil {
		if iterator.nextNode.HasPrevious() {
			iterator.nextNode = iterator.nextNode.GetPrevious()
		} else if iterator.currentTailIndex < len(iterator.list.allTails) {
			iterator.currentTailIndex++
			iterator.nextNode = iterator.list.allTails[iterator.currentTailIndex]
		}
	}
}
