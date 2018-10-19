package linkedlisttpl

import (
	"testing"
)

func TestNewList(t *testing.T) {
	var initialData NodeData = "asdf"
	list := NewList(initialData)
	if list.head != list.mainTail {
		t.Error("list head not equal to mainTail, is ", list.head, list.mainTail)
	}
	if len(list.allTails) != 1 {
		t.Error("has more than one tail", list.allTails)
	}
	if list.head.index != 0 {
		t.Error("list head index not 0, is ", list.head.index)
	}
	if list.head.previousNode != nil {
		t.Error("list head previousNode not nil, is ", list.head.previousNode)
	}
	if list.head.nextNode != nil {
		t.Error("list head nextNode not nil, is ", list.head.nextNode)
	}
	if list.head.data != initialData {
		t.Error("list head data not initialData, is ", list.head.data)
	}
}

func TestAddNode(t *testing.T) {
	t.Run("new list, add node", testAddNodeWithNewList())
	t.Run("add 3 nodes, 1 tail", testAddThreeNodes())
	t.Run("add 3 nodes, 2 tails", testAddThreeNodesTwoTails())
	t.Run("add 3 nodes, 2 tails, branching late", testAddThreeNodesTwoTailsLaterBranch())
	t.Run("add 3 nodes, 3 tails", testAddThreeNodesThreeTails())
	t.Run("add 5 nodes, switch between 2 tails", testAddSevenNodesThreeTails())
}

func testAddNodeWithNewList() func(*testing.T) {
	return func(t *testing.T) {
		initialData := "im data"
		list := NewList(initialData)
		previousData := &list.head.data
		newData := "im new data"

		list.AddNode(previousData, newData)

		if tailLen := len(list.allTails); tailLen != 1 {
			t.Error("allTails len != 1, but ", tailLen)
		}
		if list.allTails[0] != list.mainTail {
			t.Error("allTails[0] should be mainTail")
		}
		if list.head == list.mainTail {
			t.Error("head equal to mainTail")
		}
		if list.head.index != 0 {
			t.Error("list head index not 0, is ", list.head.index)
		}
		if list.head.previousNode != nil {
			t.Error("list head previousNode not nil, is ", list.head.previousNode)
		}
		if list.head.data != initialData {
			t.Error("list head data not initialData, is ", list.head.data)
		}
		if list.head.nextNode != list.mainTail {
			t.Error("list head nextNode not mainTail, is ", list.head.nextNode)
		}
		if list.mainTail.nextNode != nil {
			t.Error("list tail nextNode not nil, is ", list.mainTail.nextNode)
		}
		if list.mainTail.previousNode != list.head {
			t.Error("list tail previousNode not head, is ", list.mainTail.previousNode)
		}
		if list.mainTail.index != 1 {
			t.Error("list tail index not 1, is ", list.mainTail.index)
		}
		if list.mainTail.data != newData {
			t.Error("list tail data not newData, is ", list.mainTail.data)
		}
	}
}

func testAddThreeNodes() func(*testing.T) {
	return func(t *testing.T) {
		initialData := "im data"
		list := NewList(initialData)

		listHeadData := &list.head.data
		nodeOneData := "im nodeone data"
		nodeTwoData := "im nodetwo data"
		nodeThreeData := "im nodethree data"

		nodeOne := list.AddNode(listHeadData, nodeOneData)
		nodeTwo := list.AddNode(&nodeOne.data, nodeTwoData)
		nodeThree := list.AddNode(&nodeTwo.data, nodeThreeData)

		if tailLen := len(list.allTails); tailLen != 1 {
			t.Error("allTails len != 1, but ", tailLen)
		}
		if list.allTails[0] != list.mainTail {
			t.Error("allTails[0] should be mainTail")
		}
		if list.head == list.mainTail {
			t.Error("head equal to mainTail")
		}
		if list.head.index != 0 {
			t.Error("list head index not 0, is ", list.head.index)
		}
		if nodeOne.index != 1 {
			t.Error("node one index not 1, is ", nodeOne.index)
		}
		if nodeTwo.index != 2 {
			t.Error("node two index not 2, is ", nodeTwo.index)
		}
		if nodeThree.index != 3 {
			t.Error("node three index not 3, is ", nodeThree.index)
		}
		if list.head.previousNode != nil {
			t.Error("list head previousNode not nil, is ", list.head.previousNode)
		}
		if nodeOne.previousNode != list.head {
			t.Error("node one previousNode not list.head, is ", nodeOne.previousNode)
		}
		if nodeTwo.previousNode != nodeOne {
			t.Error("node two previousNode not nodeOne, is ", nodeTwo.previousNode)
		}
		if nodeThree.previousNode != nodeTwo {
			t.Error("node three previousNode not nodeTwo, is ", nodeThree.previousNode)
		}
		if list.head.data != initialData {
			t.Error("list head data not initialData, is ", list.head.data)
		}
		if nodeOne.data != nodeOneData {
			t.Error("node one data not nodeOneData, is ", nodeOne.data)
		}
		if nodeTwo.data != nodeTwoData {
			t.Error("node two data not nodeTwoData, is ", nodeTwo.data)
		}
		if nodeThree.data != nodeThreeData {
			t.Error("node three data not nodeThreeData, is ", nodeThree.data)
		}
		if list.head.nextNode != nodeOne {
			t.Error("list head nextNode not node one, is ", list.head.nextNode)
		}
		if nodeOne.nextNode != nodeTwo {
			t.Error("node one nextNode not node two, is ", nodeOne.nextNode)
		}
		if nodeTwo.nextNode != nodeThree {
			t.Error("node two nextNode not node three, is ", nodeTwo.nextNode)
		}
		if nodeThree.nextNode != nil {
			t.Error("node three nextNode not nil, is ", nodeThree.nextNode)
		}
		if nodeThree != list.mainTail {
			t.Error("node three not main tail, is ", nodeThree, list.mainTail)
		}
	}
}

func testAddThreeNodesTwoTails() func(*testing.T) {
	return func(t *testing.T) {
		initialData := "im data"
		list := NewList(initialData)

		listHeadData := &list.head.data
		nodeOneData := "im nodeone data"
		nodeTwoData := "im nodetwo data"
		nodeThreeData := "im nodethree data"

		nodeOne := list.AddNode(listHeadData, nodeOneData)
		nodeTwo := list.AddNode(listHeadData, nodeTwoData)
		nodeThree := list.AddNode(&(nodeTwo.data), nodeThreeData)

		if tailLen := len(list.allTails); tailLen != 2 {
			t.Error("allTails len != 2, but ", tailLen)
		}
		if list.allTails[0] != list.mainTail && list.allTails[1] != list.mainTail {
			t.Error("allTails should contain mainTail")
		}
		if list.head == list.mainTail {
			t.Error("head should not equal to mainTail")
		}
		if list.head.index != 0 {
			t.Error("list head index not 0, is ", list.head.index)
		}
		if nodeOne.index != 1 {
			t.Error("node one index not 1, is ", nodeOne.index)
		}
		if nodeTwo.index != 1 {
			t.Error("node two index not 1, is ", nodeTwo.index)
		}
		if nodeThree.index != 2 {
			t.Error("node three index not 2, is ", nodeThree.index)
		}
		if list.head.previousNode != nil {
			t.Error("list head previousNode not nil, is ", list.head.previousNode)
		}
		if nodeOne.previousNode != list.head {
			t.Error("node one previousNode not list.head, is ", nodeOne.previousNode)
		}
		if nodeTwo.previousNode != list.head {
			t.Error("node two previousNode not list.head, is ", nodeTwo.previousNode)
		}
		if nodeThree.previousNode != nodeTwo {
			t.Error("node three previousNode not nodeTwo, is ", nodeThree.previousNode)
		}
		if list.head.data != initialData {
			t.Error("list head data not initialData, is ", list.head.data)
		}
		if nodeOne.data != nodeOneData {
			t.Error("node one data not nodeOneData, is ", nodeOne.data)
		}
		if nodeTwo.data != nodeTwoData {
			t.Error("node two data not nodeTwoData, is ", nodeTwo.data)
		}
		if nodeThree.data != nodeThreeData {
			t.Error("node three data not nodeThreeData, is ", nodeThree.data)
		}
		if list.head.nextNode != nodeTwo {
			t.Error("list head nextNode not node two, is ", list.head.nextNode)
		}
		if nodeOne.nextNode != nil {
			t.Error("node one nextNode not nil, is ", nodeOne.nextNode)
		}
		if nodeTwo.nextNode != nodeThree {
			t.Error("node two nextNode not node three, is ", nodeTwo.nextNode)
		}
		if nodeThree.nextNode != nil {
			t.Error("node three nextNode not nil, is ", nodeThree.nextNode)
		}
		if nodeThree != list.mainTail {
			t.Error("node three not main tail, is ", nodeThree, list.mainTail)
		}
	}
}

func testAddThreeNodesTwoTailsLaterBranch() func(*testing.T) {
	return func(t *testing.T) {
		initialData := "im data"
		list := NewList(initialData)

		listHeadData := &list.head.data
		nodeOneData := "im nodeone data"
		nodeTwoData := "im nodetwo data"
		nodeThreeData := "im nodethree data"

		nodeOne := list.AddNode(listHeadData, nodeOneData)
		nodeTwo := list.AddNode(&(nodeOne.data), nodeTwoData)
		nodeThree := list.AddNode(&(nodeOne.data), nodeThreeData)

		if tailLen := len(list.allTails); tailLen != 2 {
			t.Error("allTails len != 2, but ", tailLen)
		}
		if list.allTails[0] != list.mainTail && list.allTails[1] != list.mainTail {
			t.Error("allTails should contain mainTail")
		}
		if list.head == list.mainTail {
			t.Error("head should not equal to mainTail")
		}
		if list.head.index != 0 {
			t.Error("list head index not 0, is ", list.head.index)
		}
		if nodeOne.index != 1 {
			t.Error("node one index not 1, is ", nodeOne.index)
		}
		if nodeTwo.index != 2 {
			t.Error("node two index not 2, is ", nodeTwo.index)
		}
		if nodeThree.index != 2 {
			t.Error("node three index not 2, is ", nodeThree.index)
		}
		if list.head.previousNode != nil {
			t.Error("list head previousNode not nil, is ", list.head.previousNode)
		}
		if nodeOne.previousNode != list.head {
			t.Error("node one previousNode not list.head, is ", nodeOne.previousNode)
		}
		if nodeTwo.previousNode != nodeOne {
			t.Error("node two previousNode not node one, is ", nodeTwo.previousNode)
		}
		if nodeThree.previousNode != nodeOne {
			t.Error("node three previousNode not node one, is ", nodeThree.previousNode)
		}
		if list.head.data != initialData {
			t.Error("list head data not initialData, is ", list.head.data)
		}
		if nodeOne.data != nodeOneData {
			t.Error("node one data not nodeOneData, is ", nodeOne.data)
		}
		if nodeTwo.data != nodeTwoData {
			t.Error("node two data not nodeTwoData, is ", nodeTwo.data)
		}
		if nodeThree.data != nodeThreeData {
			t.Error("node three data not nodeThreeData, is ", nodeThree.data)
		}
		if list.head.nextNode != nodeOne {
			t.Error("list head nextNode not node one, is ", list.head.nextNode)
		}
		if nodeOne.nextNode != nodeTwo {
			t.Error("node one nextNode not node two, is ", nodeOne.nextNode)
		}
		if nodeTwo.nextNode != nil {
			t.Error("node two nextNode not nil, is ", nodeTwo.nextNode)
		}
		if nodeThree.nextNode != nil {
			t.Error("node three nextNode not nil, is ", nodeThree.nextNode)
		}
		if nodeTwo != list.mainTail {
			t.Error("node two not main tail, is ", nodeTwo, list.mainTail)
		}
	}
}

func testAddThreeNodesThreeTails() func(*testing.T) {
	return func(t *testing.T) {
		initialData := "im data"
		list := NewList(initialData)

		listHeadData := &list.head.data
		nodeOneData := "im nodeone data"
		nodeTwoData := "im nodetwo data"
		nodeThreeData := "im nodethree data"

		nodeOne := list.AddNode(listHeadData, nodeOneData)
		nodeTwo := list.AddNode(listHeadData, nodeTwoData)
		nodeThree := list.AddNode(listHeadData, nodeThreeData)

		if tailLen := len(list.allTails); tailLen != 3 {
			t.Error("allTails len != 3, but ", tailLen)
		}
		if list.allTails[0] != list.mainTail && list.allTails[1] != list.mainTail && list.allTails[2] != list.mainTail {
			t.Error("allTails should contain mainTail")
		}
		if list.head == list.mainTail {
			t.Error("head should not equal to mainTail")
		}
		if list.head.index != 0 {
			t.Error("list head index not 0, is ", list.head.index)
		}
		if nodeOne.index != 1 {
			t.Error("node one index not 1, is ", nodeOne.index)
		}
		if nodeTwo.index != 1 {
			t.Error("node two index not 1, is ", nodeTwo.index)
		}
		if nodeThree.index != 1 {
			t.Error("node three index not 1, is ", nodeThree.index)
		}
		if list.head.previousNode != nil {
			t.Error("list head previousNode not nil, is ", list.head.previousNode)
		}
		if nodeOne.previousNode != list.head {
			t.Error("node one previousNode not list.head, is ", nodeOne.previousNode)
		}
		if nodeTwo.previousNode != list.head {
			t.Error("node two previousNode not list.head, is ", nodeTwo.previousNode)
		}
		if nodeThree.previousNode != list.head {
			t.Error("node three previousNode not list.head, is ", nodeThree.previousNode)
		}
		if list.head.data != initialData {
			t.Error("list head data not initialData, is ", list.head.data)
		}
		if nodeOne.data != nodeOneData {
			t.Error("node one data not nodeOneData, is ", nodeOne.data)
		}
		if nodeTwo.data != nodeTwoData {
			t.Error("node two data not nodeTwoData, is ", nodeTwo.data)
		}
		if nodeThree.data != nodeThreeData {
			t.Error("node three data not nodeThreeData, is ", nodeThree.data)
		}
		if list.head.nextNode != nodeOne {
			t.Error("list head nextNode not node one, is ", list.head.nextNode)
		}
		if nodeOne.nextNode != nil {
			t.Error("node one nextNode not nil, is ", nodeOne.nextNode)
		}
		if nodeTwo.nextNode != nil {
			t.Error("node two nextNode not nil, is ", nodeTwo.nextNode)
		}
		if nodeThree.nextNode != nil {
			t.Error("node three nextNode not nil, is ", nodeThree.nextNode)
		}
		if nodeOne != list.mainTail {
			t.Error("node one not main tail, is ", nodeOne, list.mainTail)
		}
	}
}

func testAddSevenNodesThreeTails() func(*testing.T) {
	return func(t *testing.T) {
		initialData := "im data"
		list := NewList(initialData)

		listHeadData := &list.head.data
		nodeOneData := "im nodeone data"
		nodeTwoData := "im nodetwo data"
		nodeThreeData := "im nodethree data"
		nodeFourData := "im nodefour data"
		nodeFiveData := "im nodefive data"
		nodeSixData := "im nodesix data"
		nodeSevenData := "im nodeseven data"

		nodeOne := list.AddNode(listHeadData, nodeOneData)
		nodeTwo := list.AddNode(listHeadData, nodeTwoData)
		nodeThree := list.AddNode(&(nodeTwo.data), nodeThreeData)
		nodeFour := list.AddNode(&(nodeOne.data), nodeFourData)
		nodeFive := list.AddNode(&(nodeFour.data), nodeFiveData)
		nodeSix := list.AddNode(listHeadData, nodeSixData)
		nodeSeven := list.AddNode(&(nodeSix.data), nodeSevenData)

		if tailLen := len(list.allTails); tailLen != 3 {
			t.Error("allTails len != 2, but ", tailLen)
		}
		if list.allTails[0] != list.mainTail && list.allTails[1] != list.mainTail && list.allTails[2] != list.mainTail {
			t.Error("allTails should contain mainTail")
		}
		if list.head == list.mainTail {
			t.Error("head should not equal to mainTail")
		}
		if list.head.index != 0 {
			t.Error("list head index not 0, is ", list.head.index)
		}
		if nodeOne.index != 1 {
			t.Error("node one index not 1, is ", nodeOne.index)
		}
		if nodeTwo.index != 1 {
			t.Error("node two index not 1, is ", nodeTwo.index)
		}
		if nodeThree.index != 2 {
			t.Error("node three index not 2, is ", nodeThree.index)
		}
		if nodeFour.index != 2 {
			t.Error("node four index not 2, is ", nodeFour.index)
		}
		if nodeFive.index != 3 {
			t.Error("node five index not 3, is ", nodeFive.index)
		}
		if nodeSix.index != 1 {
			t.Error("node six index not 1, is ", nodeSix.index)
		}
		if nodeSeven.index != 2 {
			t.Error("node seven index not 2, is ", nodeSeven.index)
		}
		if list.head.previousNode != nil {
			t.Error("list head previousNode not nil, is ", list.head.previousNode)
		}
		if nodeOne.previousNode != list.head {
			t.Error("node one previousNode not list.head, is ", nodeOne.previousNode)
		}
		if nodeTwo.previousNode != list.head {
			t.Error("node two previousNode not list.head, is ", nodeTwo.previousNode)
		}
		if nodeThree.previousNode != nodeTwo {
			t.Error("node three previousNode not node two, is ", nodeThree.previousNode)
		}
		if nodeFour.previousNode != nodeOne {
			t.Error("node four previousNode not node one, is ", nodeFour.previousNode)
		}
		if nodeFive.previousNode != nodeFour {
			t.Error("node five previousNode not node five, is ", nodeFive.previousNode)
		}
		if nodeSix.previousNode != list.head {
			t.Error("node six previousNode not list.head, is ", nodeSix.previousNode)
		}
		if nodeSeven.previousNode != nodeSix {
			t.Error("node seven previousNode not node six, is ", nodeSeven.previousNode)
		}
		if list.head.nextNode != nodeOne {
			t.Error("list head nextNode not node one, is ", list.head.nextNode)
		}
		if nodeOne.nextNode != nodeFour {
			t.Error("node one nextNode not node four, is ", nodeOne.nextNode)
		}
		if nodeTwo.nextNode != nodeThree {
			t.Error("node two nextNode not node three, is ", nodeTwo.nextNode)
		}
		if nodeThree.nextNode != nil {
			t.Error("node three nextNode not nil, is ", nodeThree.nextNode)
		}
		if nodeFour.nextNode != nodeFive {
			t.Error("node four nextNode not node five, is ", nodeFour.nextNode)
		}
		if nodeFive.nextNode != nil {
			t.Error("node five nextNode not nil, is ", nodeFive.nextNode)
		}
		if nodeSix.nextNode != nodeSeven {
			t.Error("node six nextNode not node seven, is ", nodeSix.nextNode)
		}
		if nodeSeven.nextNode != nil {
			t.Error("node seven nextNode not nil, is ", nodeSeven.nextNode)
		}
		if nodeFive != list.mainTail {
			t.Error("node five not main tail, is ", nodeFive, list.mainTail)
		}
	}
}

func TestGetMainTailData(t *testing.T) {
	t.Run("new list, one node", testGetMainTailDataWithNewList())
	t.Run("3 nodes, 1 tail", testGetMainTailDataWithThreeNodes())
	t.Run("3 nodes, 2 tails", testGetMainTailDataWithThreeNodesTwoTails())
	t.Run("3 nodes, 2 tails, branching late", testGetMainTailDataWithThreeNodesTwoTailsLaterBranch())
	t.Run("3 nodes, 3 tails", testGetMainTailDataWithThreeNodesThreeTails())
}

func testGetMainTailDataWithNewList() func(*testing.T) {
	return func(t *testing.T) {
		initialData := "im data"
		list := NewList(initialData)
		previousData := &list.head.data
		newData := "im new data"

		node := list.AddNode(previousData, newData)

		if list.GetMainTailData() != &node.data {
			t.Error("list tail data not node data, is ", list.GetMainTailData())
		}
	}
}

func testGetMainTailDataWithThreeNodes() func(*testing.T) {
	return func(t *testing.T) {
		initialData := "im data"
		list := NewList(initialData)

		listHeadData := &list.head.data
		nodeOneData := "im nodeone data"
		nodeTwoData := "im nodetwo data"
		nodeThreeData := "im nodethree data"

		nodeOne := list.AddNode(listHeadData, nodeOneData)
		nodeTwo := list.AddNode(&nodeOne.data, nodeTwoData)
		nodeThree := list.AddNode(&nodeTwo.data, nodeThreeData)

		if list.GetMainTailData() != &nodeThree.data {
			t.Error("list tail data not nodeTwo data, is ", list.GetMainTailData())
		}
	}
}

func testGetMainTailDataWithThreeNodesTwoTails() func(*testing.T) {
	return func(t *testing.T) {
		initialData := "im data"
		list := NewList(initialData)

		listHeadData := &list.head.data
		nodeOneData := "im nodeone data"
		nodeTwoData := "im nodetwo data"
		nodeThreeData := "im nodethree data"

		list.AddNode(listHeadData, nodeOneData)
		nodeTwo := list.AddNode(listHeadData, nodeTwoData)
		nodeThree := list.AddNode(&(nodeTwo.data), nodeThreeData)

		if list.GetMainTailData() != &nodeThree.data {
			t.Error("list tail data not nodeThree data, is ", list.GetMainTailData())
		}
	}
}

func testGetMainTailDataWithThreeNodesTwoTailsLaterBranch() func(*testing.T) {
	return func(t *testing.T) {
		initialData := "im data"
		list := NewList(initialData)

		listHeadData := &list.head.data
		nodeOneData := "im nodeone data"
		nodeTwoData := "im nodetwo data"
		nodeThreeData := "im nodethree data"

		nodeOne := list.AddNode(listHeadData, nodeOneData)
		nodeTwo := list.AddNode(&(nodeOne.data), nodeTwoData)
		list.AddNode(&(nodeOne.data), nodeThreeData)

		if list.GetMainTailData() != &nodeTwo.data {
			t.Error("list tail data not nodeTwo data, is ", list.GetMainTailData())
		}
	}
}

func testGetMainTailDataWithThreeNodesThreeTails() func(*testing.T) {
	return func(t *testing.T) {
		initialData := "im data"
		list := NewList(initialData)

		listHeadData := &list.head.data
		nodeOneData := "im nodeone data"
		nodeTwoData := "im nodetwo data"
		nodeThreeData := "im nodethree data"

		nodeOne := list.AddNode(listHeadData, nodeOneData)
		list.AddNode(listHeadData, nodeTwoData)
		list.AddNode(listHeadData, nodeThreeData)

		if list.GetMainTailData() != &nodeOne.data {
			t.Error("list tail data not nodeOne data, is ", list.GetMainTailData())
		}
	}
}
