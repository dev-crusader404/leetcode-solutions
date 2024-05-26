package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeInBetween2(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	var start *ListNode
	end := list1

	for i := 0; i < b; end = end.Next {
		i++
		if i == a {
			start = end
		}
	}
	start.Next = list2

	for list2.Next != nil {
		list2 = list2.Next
	}
	list2.Next = end.Next
	return list1
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	toJoin := list2
	toInsert := list1
	head := list1
	runner := list1
	count := 0
	var inserted bool

	for toJoin.Next != nil {
		toJoin = toJoin.Next
	}
	var end *ListNode
	for runner != nil {
		if (a - 1) == count {
			end = runner.Next
			toInsert.Next = list2
			runner = end
			inserted = true
			count++
			continue
		}

		if count == b {
			toJoin.Next = runner.Next
			break
		}

		count++
		runner = runner.Next
		if !inserted {
			toInsert = toInsert.Next
		}
	}
	return head
}

func RunLC1669() {
	// [10,1,13,6,9,5], a = 3, b = 4, list2 = [1000000,1000001,1000002]
	l1 := Node{}
	for _, v := range []int{10, 1, 13, 6, 9, 5} {
		l1.Add(v)
	}

	l2 := Node{}
	for _, v := range []int{1000000, 1000001, 1000002} {
		l2.Add(v)
	}

	mergeInBetween(l1.head, 3, 4, l2.head)
	mergeInBetween2(l1.head, 3, 4, l2.head)
}
