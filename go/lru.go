package csfundamentals

import (
	"container/list"
	"errors"
	"sync"
)

// LRUCache is an implementation of an LRU cache.
type LRUCache struct {
	// size represents the size of the cache.
	size int

	// items is a doubly linked list of elements in the cache.
	items *list.List

	// exists contains an indexed list of elements in the cache.
	exists map[interface{}]*list.Element

	// protect against concurrent reads/writes.
	mutex sync.Mutex
}

// item represents an item saved in the cache
type item struct {
	key interface{}
	val interface{}
}

// NewLRUCache create a new LRU of a given size.
func NewLRUCache(size int) (*LRUCache, error) {
	if size < 1 {
		return nil, errors.New("cache size must be > 0")
	}

	lru := &LRUCache{
		size:   size,
		items:  list.New(),
		exists: make(map[interface{}]*list.Element, size),
	}
	return lru, nil
}

// Get retrieves a given element from the LRU Cache.
func (lru *LRUCache) Get(key interface{}) interface{} {
	lru.mutex.Lock()
	defer lru.mutex.Unlock()

	if e, ok := lru.exists[key]; ok {
		lru.items.MoveToFront(e)
		if i, ok := e.Value.(item); ok {
			return i.val
		}
	}
	return nil
}

// Put creates a new item in the LRU Cache.
// It will return true if an item was evicted.
func (lru *LRUCache) Put(key, value interface{}) bool {
	lru.mutex.Lock()
	defer lru.mutex.Unlock()

	if e, ok := lru.exists[key]; ok {
		if i, ok := e.Value.(item); ok {
			i.val = value
			e.Value = i
		}
		lru.items.MoveToFront(e)
		return false
	}

	evicted := false

	// if the cache is full, remove last element
	if lru.items.Len() == lru.size {
		last := lru.items.Back()
		lru.items.Remove(last)
		if e, ok := last.Value.(item); ok {
			delete(lru.exists, e.key)
		}
		evicted = true
	}

	i := item{
		key: key,
		val: value,
	}

	lru.exists[key] = lru.items.PushFront(i)
	return evicted
}

// Delete deletes an item from the LRU.
func (lru *LRUCache) Delete(key interface{}) {
	lru.mutex.Lock()
	defer lru.mutex.Unlock()

	// if item is in cache, remove it
	if e, ok := lru.exists[key]; ok {
		lru.items.Remove(e)
		delete(lru.exists, key)
	}
}
