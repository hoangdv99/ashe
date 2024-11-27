package linkedlist

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
}

// Insert at the end of the linked list
func (ll *LinkedList) InsertAtEnd(value int) {
	newNode := &Node{value: value}
	if ll.head == nil {
		ll.head = newNode
		return
	}

	current := ll.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

// Insert at the beginning
func (ll *LinkedList) InsertAtBegin(value int) {
	newNode := &Node{value: value}
	newNode.next = ll.head
	ll.head = newNode
}
