import "container/heap"

/**
23. Merge k Sorted Lists
将k个有序链表合并，仔细想想很有意思的一题。

  看到题重直觉上最简单的解就是每次将链表的第一个数字排序，选出最小的一个，将它放在新的链表中，再如此迭代。一般基于交换的排序算法复杂度为O(klogk）,所以整个过程的复杂度就为O(nklogk)。
而在这个题目中k是从1开始增长的，所以再排序上我们可以使用桶排，复杂度就为O(nk)。
  回看这道题的复杂度是hard，最优解肯定没有那么简单。"直觉"告诉我，在这里排序还可以做做文章。
  直接对链表头的数字进行排序。。。每个排序都是重新排的，但前一次排好的序列明显可以好好利用一番。每次选出最小值的算法里，最小堆就是一个比快排快的算法，这样复杂度就降为O(nlogk)
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

import "container/heap"

func mergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)

	// min heap
	h := &Heap{}
	heap.Init(h)

	head := &ListNode{}
	node := head
	var min *ListNode

	var flag bool
	for !flag {
		flag = true
		for i, list := range lists {
			if list != nil {
				// add new node to the heap
				heap.Push(h, list)
				lists[i] = list.Next

				flag = false
			}
		}

		if h.Len() > length {
			min = heap.Pop(h).(*ListNode)
			min.Next = nil
			node.Next = min
			node = min
		}
	}

	// pop all data in the heap
	for h.Len() != 0 {
		// merge the node
		min = heap.Pop(h).(*ListNode)
		min.Next = nil
		node.Next = min
		node = min
	}

	return head.Next
}

// An Heap is a min-heap of ListNode.
type Heap []*ListNode

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h Heap) Swap(i, j int)      { h[i].Val, h[j].Val = h[j].Val, h[i].Val }

func (h *Heap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*ListNode))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}