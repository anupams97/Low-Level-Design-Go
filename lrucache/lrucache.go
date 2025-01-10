package main

import (
	"fmt"
	"sync"
)

type Node struct {
	Key   string
	Value string
	next  *Node
	prev  *Node
}

type LRUCache struct {
	capacity int
	cache    map[string]*Node
	head     *Node
	tail     *Node
	mu       sync.Mutex
}

func (c *LRUCache) Get(key string) (string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if node, ok := c.cache[key]; ok {
		c.moveToHead(node)
		return node.Value, true
	}
	return "", false
}

func (c *LRUCache) Put(k, v string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if node, exists := c.cache[k]; exists {
		node.Value = v
		// put the node in front
		c.moveToHead(node)
	}
	newNode := &Node{Key: k, Value: v}
	c.cache[k] = newNode
	c.addToHead(newNode)

	if len(c.cache) > c.capacity {
		lastNode := c.removeTail()
		delete(c.cache, lastNode.Key)
	}
}

func (c *LRUCache) removeTail() *Node {
	lastNode := c.tail.prev
	c.removeNode(lastNode)
	return lastNode
}

func (c *LRUCache) moveToHead(node *Node) {
	c.removeNode(node)
	c.addToHead(node)
}
func (c *LRUCache) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}
func (c *LRUCache) addToHead(node *Node) {
	node.prev = c.head
	node.next = c.head.next
	c.head.next.prev = node
	c.head.next = node
}
func NewLRUCache(capacity int) *LRUCache {
	cache := &LRUCache{capacity: capacity,
		cache: make(map[string]*Node),
	}
	cache.head = &Node{}
	cache.tail = &Node{}
	cache.head.next = cache.tail
	cache.tail.next = cache.head
	return cache
}

func main() {
	lruCache := NewLRUCache(3)
	lruCache.Put("1", "Value 1")
	lruCache.Put("2", "Value 2")
	lruCache.Put("3", "Value 3")

	// Get values and print them
	if val, exists := lruCache.Get("1"); exists {
		fmt.Println(val) // Output: Value 1
	}
	if val, exists := lruCache.Get("2"); exists {
		fmt.Println(val) // Output: Value 2
	}

	// Add a new value that should evict the least recently used one
	lruCache.Put("4", "Value 4")

	// Try to get the evicted value
	if val, exists := lruCache.Get("3"); exists {
		fmt.Println(val)
	} else {
		fmt.Println("Value 3 was evicted") // Output: Value 3 was evicted
	}

	// Get the newly added value
	if val, exists := lruCache.Get("4"); exists {
		fmt.Println(val) // Output: Value 4
	}

	// Update an existing value
	lruCache.Put("2", "Updated Value 2")

	// Get the values again
	if val, exists := lruCache.Get("1"); exists {
		fmt.Println(val) // Output: Value 1
	}
	if val, exists := lruCache.Get("2"); exists {
		fmt.Println(val) // Output: Updated Value 2
	}

}
