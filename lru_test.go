package csfundamentals

import "testing"

func TestNewLRUCache(t *testing.T) {
	t.Run("TestError", func(t *testing.T) {
		_, err := NewLRUCache(0)
		if err == nil {
			t.Error("expected error due to invalid cache size")
		}
	})

	t.Run("TestSuccess", func(t *testing.T) {
		_, err := NewLRUCache(10)
		if err != nil {
			t.Errorf("expected error: nil, got: %v", err)
		}
	})
}

func TestLRUCacheGet(t *testing.T) {
	c, _ := NewLRUCache(1)
	if val := c.Get("test"); val != nil {
		t.Errorf("expected to have nothing in the cache, got: %v", val)
	}
}

func TestLRUCachePut(t *testing.T) {
	c, _ := NewLRUCache(1)

	t.Run("TestEmptyLRU", func(t *testing.T) {
		// put new item in cache
		if evicted := c.Put("test", "test-value"); evicted {
			t.Error("expected 0 items to have been evicted")
		}

		// check item has been added to cache
		if val := c.Get("test"); val != "test-value" {
			t.Errorf("expected to have value: 'test-value' for item: 'test', got: '%v'", val)
		}
	})

	t.Run("TestFullLRU", func(t *testing.T) {
		// put another new item in cache
		if evicted := c.Put("test-2", "test-value-2"); !evicted {
			t.Error("expected 1 item to have been evicted")
		}

		// check item has been added to cache
		if val := c.Get("test-2"); val != "test-value-2" {
			t.Errorf("expected to have value: 'test-value-2' for item: 'test-2', got: '%v'", val)
		}

		// update a value in cache
		if evicted := c.Put("test-2", "new-test-value-2"); evicted {
			t.Error("expected 0 items to have been evicted")
		}

		// check new value is as expected
		if val := c.Get("test-2"); val != "new-test-value-2" {
			t.Errorf("expected to have value: 'new-test-value-2' for item: 'test-2', got: '%v'", val)
		}
	})
}

func TestLRUCacheDelete(t *testing.T) {
	c, _ := NewLRUCache(1)

	// put new item in cache
	c.Put("test", "test-value")

	// check item has been added to cache
	if val := c.Get("test"); val != "test-value" {
		t.Errorf("expected to have value: 'test-value' for item: 'test', got: '%v'", val)
	}

	// delete item from cache
	c.Delete("test")

	// check item has been deleted from cache
	if val := c.Get("test"); val != nil {
		t.Errorf("expected to have value: nil for item: 'test', got: '%v'", val)
	}
}
