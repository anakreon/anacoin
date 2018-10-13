package linkedlist

import (
	"testing"
)

func TestNewNode(t *testing.T) {
	var previousNode *Node = &Node{
		index:        0,
		previousNode: nil,
		nextNode:     nil,
		data:         nil,
	}
	var data NodeData = "asdf"
	node := NewNode(previousNode, data)
	if node.index != 1 {
		t.Error("node index not 1, is ", node.index)
	}
	if node.previousNode != previousNode {
		t.Error("previous node invalid, is ", node.previousNode)
	}
	if node.nextNode != nil {
		t.Error("next node not nil, is ", node.nextNode)
	}
	if node.data != "asdf" {
		t.Error("node data not asdf, is ", node.data)
	}
}

func TestHasNext(t *testing.T) {
	nodeWithoutNext := &Node{
		nextNode: nil,
	}
	nodeWithNext := &Node{
		nextNode: nodeWithoutNext,
	}
	t.Run("has next node", testHasNext(nodeWithNext, true))
	t.Run("no next node", testHasNext(nodeWithoutNext, false))
}

func testHasNext(node *Node, expected bool) func(*testing.T) {
	return func(t *testing.T) {
		if node.HasNext() != expected {
			t.Fail()
		}
	}
}

func TestSetNext(t *testing.T) {
	nodeOne := &Node{}
	nodeTwo := &Node{}
	nodeOne.SetNext(nodeTwo)
	if nodeOne.nextNode != nodeTwo {
		t.Error("next node not nodeTwo, is ", nodeOne.nextNode)
	}
}

func TestHasPrevious(t *testing.T) {
	nodeWithoutPrevious := &Node{
		previousNode: nil,
	}
	nodeWithPrevious := &Node{
		previousNode: nodeWithoutPrevious,
	}
	t.Run("has previous node", testHasPrevious(nodeWithPrevious, true))
	t.Run("no previous node", testHasPrevious(nodeWithoutPrevious, false))
}

func testHasPrevious(node *Node, expected bool) func(*testing.T) {
	return func(t *testing.T) {
		if node.HasPrevious() != expected {
			t.Fail()
		}
	}
}

func TestGetPrevious(t *testing.T) {
	nodeWithoutPrevious := &Node{
		previousNode: nil,
	}
	nodeWithPrevious := &Node{
		previousNode: nodeWithoutPrevious,
	}
	t.Run("has previous node", testGetPrevious(nodeWithPrevious, nodeWithoutPrevious))
	t.Run("no previous node", testGetPrevious(nodeWithoutPrevious, nil))
}

func testGetPrevious(node *Node, expected *Node) func(*testing.T) {
	return func(t *testing.T) {
		if node.GetPrevious() != expected {
			t.Fail()
		}
	}
}

func TestGetData(t *testing.T) {
	nodeWithoutData := &Node{
		data: nil,
	}
	nodeWithData := &Node{
		data: "im data",
	}
	t.Run("has data", testGetData(nodeWithData, &nodeWithData.data))
	t.Run("no data", testGetData(nodeWithoutData, &nodeWithoutData.data))
}

func testGetData(node *Node, expected NodeData) func(*testing.T) {
	return func(t *testing.T) {
		if result := node.GetData(); result != expected {
			t.Error(result)
		}
	}
}

func TestHasFullLinkWithPreviousNode(t *testing.T) {
	nodeWithNoLink := &Node{
		previousNode: nil,
		nextNode:     nil,
	}
	nodeWithPartialLinkOne := &Node{
		previousNode: nodeWithNoLink,
		nextNode:     nodeWithNoLink,
	}
	nodeWithPartialLinkTwo := &Node{
		previousNode: nodeWithPartialLinkOne,
	}
	nodeWithFullLink := &Node{
		previousNode: nodeWithPartialLinkTwo,
		nextNode:     nil,
	}
	nodeWithPartialLinkTwo.nextNode = nodeWithFullLink
	t.Run("has no link", testHasFullLinkWithPreviousNode(nodeWithNoLink, true))
	t.Run("has partial link, no nextNode", testHasFullLinkWithPreviousNode(nodeWithPartialLinkOne, false))
	t.Run("has partial link, wrong nextNode", testHasFullLinkWithPreviousNode(nodeWithPartialLinkTwo, false))
	t.Run("has full link", testHasFullLinkWithPreviousNode(nodeWithFullLink, true))
}

func testHasFullLinkWithPreviousNode(node *Node, expected bool) func(*testing.T) {
	return func(t *testing.T) {
		if result := node.HasFullLinkWithPreviousNode(); result != expected {
			t.Error(result)
		}
	}
}

func TestHasIndexLargerThan(t *testing.T) {
	nodeOne := &Node{
		index: 1,
	}
	nodeTwo := &Node{
		index: 2,
	}
	t.Run("true", testHasIndexLargerThana(nodeTwo, nodeOne, true))
	t.Run("false", testHasIndexLargerThana(nodeOne, nodeTwo, false))
}

func testHasIndexLargerThana(node *Node, otherNode *Node, expected NodeData) func(*testing.T) {
	return func(t *testing.T) {
		if result := node.HasIndexLargerThan(otherNode); result != expected {
			t.Error(result)
		}
	}
}