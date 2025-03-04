package golang

type Node struct {
	Key  int
	Val  int
	Prev *Node
	Next *Node
}

type LRUCache struct {
	capacity int
	cache    map[int]*Node
	head     *Node
	tail     *Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{Key: -1, Val: -1, Prev: nil, Next: nil}
	tail := &Node{Key: -1, Val: -1, Prev: nil, Next: nil}
	head.Next = tail
	tail.Prev = head
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node, capacity),
		head:     head,
		tail:     tail,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.remove(node)
		this.add(node)
		return node.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.cache[key]
	if ok {
		this.remove(node)
	}
	node = &Node{
		Key:  key,
		Val:  value,
		Next: nil,
		Prev: nil,
	}
	this.add(node)
	this.cache[key] = node

	if len(this.cache) > this.capacity {
		next := this.head.Next
		this.remove(next)
		delete(this.cache, next.Key)
	}

}

func (this *LRUCache) remove(node *Node) {
	prev := node.Prev
	next := node.Next
	prev.Next = next
	next.Prev = prev
}

func (this *LRUCache) add(node *Node) {
	p := this.tail.Prev
	p.Next = node
	this.tail.Prev = node
	node.Prev = p
	node.Next = this.tail
}
