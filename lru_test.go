package main

import "testing"

func TestCacheHas(t *testing.T) {
	c := NewCache(1)
	if c.Has("test") {
		t.Error("expected to have nothing in the cache")
	}
}

func TestCachePut(t *testing.T) {
	c := NewCache(1)

	t.Run("TestEmptyCache", func(t *testing.T) {
		// put new item in cache
		c.Put("test")

		// check item has been added to cache
		if !c.Has("test") {
			t.Error("expected to have 'test' item in cache")
		}
	})

	t.Run("TestFullCache", func(t *testing.T) {
		// put another new item in cache
		c.Put("test-2")

		// check item has been added to cache
		if !c.Has("test-2") {
			t.Error("expected to have 'test' item in cache")
		}

		// check original item was deleted from cache
		if c.Has("test") {
			t.Error("expected 'test' item to be deleted from cache")
		}
	})
}

func TestCacheDelete(t *testing.T) {
	c := NewCache(1)

	// put new item in cache
	c.Put("test")

	// check item has been added to cache
	if !c.Has("test") {
		t.Error("expected to have 'test' item in cache")
	}

	// delete item from cache
	c.Delete("test")

	// check item has been deleted from cache
	if c.Has("test") {
		t.Error("expected to have nothing in the cache")
	}
}
