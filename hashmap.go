package csfundamentals

import "hash/fnv"

// HashFunc represents a key based hashing function
type HashFunc func(string) int

// FNV32Hash uses the FNV hashing algorithm to calculate an index
// based of a given byte slice and HashMap size
var FNV32Hash = func(s string) int {
	h := fnv.New32()
	h.Write([]byte(s))
	return int(h.Sum32())
}

type kv struct {
	key   string
	value interface{}
}

// HashMap is used to store key/value pairs in an associative array
type HashMap struct {
	hash   HashFunc
	size   int
	values [][]*kv
}

// Get checks if a given key exists in the HashMap
func (hm *HashMap) Get(key string) (interface{}, bool) {
	index := hm.index(key)
	for _, kv := range hm.values[index] {
		if kv.key == key {
			return kv.value, true
		}
	}
	return nil, false
}

// Add checks if a given key exists in the HashMap
// If so, it will update the value
// If not, it will create a new value
// Returns true if item was created, false if not
func (hm *HashMap) Add(key string, value interface{}) bool {
	index := hm.index(key)

	// check if key exists
	for _, kv := range hm.values[index] {
		if kv.key == key {
			kv.value = value
			return false
		}
	}

	values := hm.values[index]
	kv := &kv{
		key:   key,
		value: value,
	}
	values = append(values, kv)
	hm.values[index] = values

	return true
}

// Delete wil attempt to delete a value from the HashMap
// Returns true if deleted, false is not
func (hm *HashMap) Delete(key string) bool {
	index := hm.index(key)

	// check if key exists
	for i, kv := range hm.values[index] {
		if kv.key == key {
			hm.values[index] = append(hm.values[index][:i], hm.values[index][i+1:]...)
			return true
		}
	}

	return false
}

func (hm *HashMap) index(s string) int {
	hash := hm.hash(s)
	return hash % hm.size
}

// NewHashMap creates a HashMap of the given length
func NewHashMap(size int, hash HashFunc) *HashMap {
	if size < 1 {
		size = 1
	}

	if hash == nil {
		hash = FNV32Hash
	}

	hm := &HashMap{
		hash:   hash,
		size:   size,
		values: make([][]*kv, size),
	}

	for i := range hm.values {
		hm.values[i] = make([]*kv, 0)
	}
	return hm
}
