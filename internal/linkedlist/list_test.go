//go:generate genny -in=$GOFILE -out=../linkedlist-gen/$GOFILE -pkg=linkedlist gen "NodeData=block.Block"

package linkedlisttpl

import "testing"

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
	t.Run("new list", testAddNodeWithNewList())
	//t.Run("xxx", testAddNode())
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
