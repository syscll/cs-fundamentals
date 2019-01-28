package main

import (
	"container/list"
	"errors"
	"sync"
)

// LRU is an implementation of an LRU cache.
type LRU struct {
	// size represents the size of the cache.
	size int

	// items is a doubly linked list of elements in the cache.
	items *list.List

	// exists contains an indexed list of elements in the cache.
	exists map[interface{}]*list.Element

	// protect against concurrent reads/writes.
	mutex sync.Mutex
}

// Contains checks if a given element is in the LRU.
func (lru *LRU) Contains(v interface{}) bool {
	lru.mutex.Lock()
	defer lru.mutex.Unlock()

	if e, ok := lru.exists[v]; ok {
		lru.items.MoveToFront(e)
		return true
	}
	return false
}

// Add creates a new item in the LRU.
// It will return true if an item was evicted.
func (lru *LRU) Add(v interface{}) bool {
	lru.mutex.Lock()
	defer lru.mutex.Unlock()

	evicted := false

	if e, ok := lru.exists[v]; ok {
		lru.items.MoveToFront(e)
		return evicted
	}

	// if the cache is full, remove last element
	if lru.items.Len() == lru.size {
		last := lru.items.Back()
		lru.items.Remove(last)
		delete(lru.exists, last.Value)
		evicted = true
	}

	lru.exists[v] = lru.items.PushFront(v)
	return evicted
}

// Delete deletes an item from the LRU.
func (lru *LRU) Delete(v interface{}) {
	lru.mutex.Lock()
	defer lru.mutex.Unlock()

	e, ok := lru.exists[v]
	if !ok {
		// item not in cache, do nothing
		return
	}

	lru.items.Remove(e)
	delete(lru.exists, v)
}

// NewLRU create a new LRU of a given size.
func NewLRU(size int) (*LRU, error) {
	if size < 1 {
		return nil, errors.New("cache size must be > 0")
	}

	lru := &LRU{
		size:   size,
		items:  list.New(),
		exists: make(map[interface{}]*list.Element, size),
	}
	return lru, nil
}
