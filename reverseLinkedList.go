package main

type ListNode struct {
	Val int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	// return nil if the list is empty
    if head == nil {
        return nil
    }
    
    var revHead *ListNode
	// while head exists
    for head != nil {
		// temp variable that we need
        tmp := head.Next
		// move head forward, we still have acopy of it in temp
        head.Next = revHead
		// set revhead to the head
        revHead = head
		// set head to the temp
        head = tmp
    }
    return revHead
}