package main

// 146 lru-cache 2023-04-10 14:53:43

// leetcode submit region begin(Prohibit modification and deletion)

type LRUCache struct {
	head, tail *Node
	Keys       map[int]*Node
	Cap        int
}

type Node struct {
	Key, Val   int
	Prev, Next *Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{Keys: make(map[int]*Node), Cap: capacity}
}

func (lc *LRUCache) Get(key int) int {
	if node, ok := lc.Keys[key]; ok {
		lc.Remove(node)
		lc.Add(node)
		return node.Val
	}
	return -1
}

func (lc *LRUCache) Put(key int, value int) {
	if node, ok := lc.Keys[key]; ok {
		node.Val = value
		lc.Remove(node)
		lc.Add(node)
		return
	} else {
		node = &Node{Key: key, Val: value}
		lc.Keys[key] = node
		lc.Add(node)
	}
	if len(lc.Keys) > lc.Cap {
		delete(lc.Keys, lc.tail.Key)
		lc.Remove(lc.tail)
	}
}

func (lc *LRUCache) Add(node *Node) {
	node.Prev = nil
	node.Next = lc.head
	if lc.head != nil {
		lc.head.Prev = node
	}
	lc.head = node
	if lc.tail == nil {
		lc.tail = node
		lc.tail.Next = nil
	}
}

func (lc *LRUCache) Remove(node *Node) {
	if node == lc.head {
		lc.head = node.Next
		node.Next = nil
		return
	}
	if node == lc.tail {
		lc.tail = node.Prev
		node.Prev.Next = nil
		node.Prev = nil
		return
	}
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
