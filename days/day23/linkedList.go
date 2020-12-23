package days

import (
	"fmt"
	"strings"
)

// Node -
type Node struct {
	value int
	next  *Node
}

// LinkedList -
type LinkedList struct {
	head *Node
}

// Add -
func (ll *LinkedList) Add(toAdd int) {
	if ll.head == nil {
		head := &Node{value: toAdd}
		ll.head = head
		ll.head.next = head
	} else {

		currentNode := ll.head
		for currentNode.next != ll.head {
			currentNode = currentNode.next
		}
		newNode := Node{value: toAdd, next: ll.head}
		currentNode.next = &newNode
	}
}

// AddAllAfter - Add's a int to the list after the given int, if the number is not found, does not insert
func (ll *LinkedList) AddAllAfter(toAdd []int, after int) bool {

	currentNode := ll.head

	isFirst := true

	for currentNode != ll.head || isFirst {

		if currentNode.value == after {
			temp := currentNode.next

			for _, val := range toAdd {
				currentNode.next = &Node{value: val}
				currentNode = currentNode.next
			}

			currentNode.next = temp
			return true
		}

		currentNode = currentNode.next
		isFirst = false
	}
	return false
}

// AddAfter - Add's a int to the list after the given int, if the number is not found, does not insert
func (ll *LinkedList) AddAfter(toAdd, after int) bool {

	currentNode := ll.head
	isFirst := true
	for currentNode != ll.head || isFirst {

		if currentNode.value == after {
			temp := Node{value: toAdd, next: currentNode.next}
			currentNode.next = &temp
			return true
		}

		currentNode = currentNode.next
		isFirst = false
	}
	return false
}

// RemoveNextN -
func (ll *LinkedList) RemoveNextN(prev, n int) []*Node {

	toReturn := make([]*Node, 0)

	isFirst := true

	currentNode := ll.head
	for currentNode != ll.head || isFirst {

		if currentNode.value == prev {

			temp := currentNode

			for i := 0; i < n; i++ {
				currentNode = currentNode.next
				if currentNode == nil {
					return toReturn
				}
				if currentNode == ll.head {
					ll.head = temp
				}
				toReturn = append(toReturn, currentNode)
			}

			temp.next = currentNode.next
			return toReturn
		}
		currentNode = currentNode.next
		isFirst = false
	}
	return nil
}

// Remove -
func (ll *LinkedList) Remove(toRemove int) *Node {

	currentNode := ll.head
	for currentNode.next != nil {

		if currentNode.next.value == toRemove {
			temp := currentNode.next
			currentNode.next = currentNode.next.next
			return temp
		}
		currentNode = currentNode.next

	}
	return nil
}

// ToString -
func (ll *LinkedList) ToString(currentVal int) string {
	var result strings.Builder
	isFirst := true
	currentNode := ll.head
	for currentNode != ll.head || isFirst {
		if currentVal == currentNode.value {
			result.WriteString(fmt.Sprintf("(%v)", currentNode.value))
		} else {
			result.WriteString(fmt.Sprintf("%v", currentNode.value))
		}
		if currentNode.next != ll.head || isFirst {
			result.WriteString(" -> ")
		}
		currentNode = currentNode.next
		isFirst = false
	}
	return result.String()
}
