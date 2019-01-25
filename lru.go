package main

import (
	"container/list"
	"sync"
)

// Cache is an implementation of an LRU cache.
type Cache struct {
	// s represents the size of the cache.
	s int

	// l is a doubly linked list of elements in the cache.
	l *list.List

	// e contains an indexed list of elements in the cache.
	e map[interface{}]*list.Element

	// protect against concurrent reads/writes.
	sync.Mutex
}

// Has checks if a given element is in the Cache.
func (c *Cache) Has(v interface{}) bool {
	c.Lock()
	defer c.Unlock()

	if e, ok := c.e[v]; ok {
		c.l.MoveToFront(e)
		return true
	}
	return false
}

// Put creates a new item in the Cache.
func (c *Cache) Put(v interface{}) {
	c.Lock()
	defer c.Unlock()

	if e, ok := c.e[v]; ok {
		c.l.MoveToFront(e)
		return
	}

	// if the cache is full, remove last element
	if c.l.Len() == c.s {
		last := c.l.Back()
		c.l.Remove(last)
		delete(c.e, last.Value)
	}

	c.e[v] = c.l.PushFront(v)
}

// Delete deletes an item from the Cache.
func (c *Cache) Delete(v interface{}) {
	c.Lock()
	defer c.Unlock()

	e, ok := c.e[v]
	if !ok {
		// item not in cache, do nothing
		return
	}

	c.l.Remove(e)
	delete(c.e, v)
}

// NewCache create a new Cache of a given size.
func NewCache(size int) *Cache {
	s := &Cache{
		s: size,
		l: list.New(),
		e: make(map[interface{}]*list.Element, size),
	}
	return s
}
