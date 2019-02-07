package csfundamentals

import "testing"

func TestNewHashMap(t *testing.T) {
	t.Run("Test0Size", func(t *testing.T) {
		hm := NewHashMap(0, FNV32Hash)
		if len(hm.values) != 1 {
			t.Errorf("expected length: 1, got: %d", len(hm.values))
		}
	})

	t.Run("Test20Size", func(t *testing.T) {
		hm := NewHashMap(20, FNV32Hash)
		if len(hm.values) != 20 {
			t.Errorf("expected length: 20, got: %d", len(hm.values))
		}
	})
}

func TestGetAndAdd(t *testing.T) {
	hm := NewHashMap(1, FNV32Hash)
	key := "test-key"
	value := "test-value"

	// get value from empty HashMap
	if _, ok := hm.Get(key); ok {
		t.Errorf("did not expect to get a value for key: %s", key)
	}

	// add value to HashMap
	if created := hm.Add(key, value); !created {
		t.Error("expected value to be created, not updated")
	}

	// get new value from HashMap
	v, ok := hm.Get(key)
	if !ok {
		t.Errorf("expected to get a value for key: %s", key)
	}
	if v != value {
		t.Errorf("expected value: %s for key: %s, got: %d", value, key, v)
	}

	// add the same value to HashMap
	if created := hm.Add(key, value); created {
		t.Error("expected value to be updated, not created")
	}

	// get new value from HashMap
	v, ok = hm.Get(key)
	if !ok {
		t.Errorf("expected to get a value for key: %s", key)
	}
	if v != value {
		t.Errorf("expected value: %s for key: %s, got: %d", value, key, v)
	}
}

func TestDelete(t *testing.T) {
	hm := NewHashMap(1, FNV32Hash)
	key := "test-key"
	value := "test-value"

	// add value to HashMap
	hm.Add(key, value)

	// delete recently added value
	if deleted := hm.Delete(key); !deleted {
		t.Error("expected key/value to be deleted")
	}

	// get value from empty HashMap
	if _, ok := hm.Get(key); ok {
		t.Errorf("did not expect to get a value for key: %s", key)
	}

	// attempt to delete non-existent value
	if deleted := hm.Delete(key); deleted {
		t.Error("did not expect key/value to be deleted")
	}
}
