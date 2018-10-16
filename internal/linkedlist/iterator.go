//go:generate genny -in=$GOFILE -out=../linkedlist-gen/$GOFILE -pkg=linkedlist gen "NodeData=block.Block"

package linkedlisttpl

type listIterator struct {
	list             *List
	currentTailIndex int
	nextNode         *Node
	goToNextNode     func()
}

func (list *List) Iterator() listIterator {
	iterator := listIterator{
		list:     list,
		nextNode: list.mainTail,
	}
	iterator.goToNextNode = goToNextNodeInThisTail(&iterator)
	return iterator
}

func (list *List) AllTailsIterator() listIterator {
	currentListTailIndex := 0
	iterator := listIterator{
		list:             list,
		currentTailIndex: currentListTailIndex,
		nextNode:         list.allTails[currentListTailIndex],
	}
	iterator.goToNextNode = goToNextNodeIncludingOtherTails(&iterator)
	return iterator
}

func (iterator *listIterator) HasNext() bool {
	return iterator.nextNode != nil
}

func (iterator *listIterator) Next() *NodeData {
	nextNode := iterator.nextNode
	iterator.goToNextNode()
	return nextNode.GetData()
}

func goToNextNodeIncludingOtherTails(iterator *listIterator) func() {
	return func() {
		if iterator.nextNode != nil {
			if iterator.nextNode.HasPrevious() {
				iterator.nextNode = iterator.nextNode.GetPrevious()
			} else if iterator.currentTailIndex < (len(iterator.list.allTails) - 1) {
				iterator.currentTailIndex++
				iterator.nextNode = iterator.list.allTails[iterator.currentTailIndex]
			}
		}
	}
}

func goToNextNodeInThisTail(iterator *listIterator) func() {
	return func() {
		if iterator.nextNode != nil {
			if iterator.nextNode.HasPrevious() {
				iterator.nextNode = iterator.nextNode.GetPrevious()
			}
		}
	}
}
