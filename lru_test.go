package csfundamentals

import "testing"

func TestNewLRU(t *testing.T) {
	t.Run("TestError", func(t *testing.T) {
		_, err := NewLRU(0)
		if err == nil {
			t.Error("expected error due to invalid cache size")
		}
	})

	t.Run("TestSuccess", func(t *testing.T) {
		_, err := NewLRU(10)
		if err != nil {
			t.Errorf("expected error: nil, got: %v", err)
		}
	})
}

func TestLRUContains(t *testing.T) {
	c, _ := NewLRU(1)
	if c.Contains("test") {
		t.Error("expected to have nothing in the cache")
	}
}

func TestLRUAdd(t *testing.T) {
	c, _ := NewLRU(1)

	t.Run("TestEmptyLRU", func(t *testing.T) {
		// put new item in cache
		evicted := c.Add("test")

		// check item has been added to cache
		if !c.Contains("test") {
			t.Error("expected to have 'test' item in cache")
		}
		if evicted {
			t.Error("expected 0 items to have been evicted")
		}
	})

	t.Run("TestFullLRU", func(t *testing.T) {
		// put another new item in cache
		evicted := c.Add("test-2")

		// check item has been added to cache
		if !c.Contains("test-2") {
			t.Error("expected to have 'test' item in cache")
		}

		// check original item was deleted from cache
		if c.Contains("test") || !evicted {
			t.Error("expected 'test' item to be deleted from cache")
		}
	})
}

func TestLRUDelete(t *testing.T) {
	c, _ := NewLRU(1)

	// put new item in cache
	c.Add("test")

	// check item has been added to cache
	if !c.Contains("test") {
		t.Error("expected to have 'test' item in cache")
	}

	// delete item from cache
	c.Delete("test")

	// check item has been deleted from cache
	if c.Contains("test") {
		t.Error("expected to have nothing in the cache")
	}
}
