package ledger

import (
	"sort"
	"testing"
)

// TestStoreKeysEmpty verifies that a freshly constructed store returns a
// non-nil, empty slice from Keys().
func TestStoreKeysEmpty(t *testing.T) {
	s := NewStore()

	keys := s.Keys()
	if keys == nil {
		t.Fatal("Keys() returned nil for an empty store; want a non-nil empty slice")
	}
	if len(keys) != 0 {
		t.Fatalf("Keys() returned %d keys for an empty store; want 0", len(keys))
	}
}

// TestStoreKeysReturnsAllSetKeys verifies that every key passed to Set is
// returned by Keys(), with no duplicates when a key is set more than once.
func TestStoreKeysReturnsAllSetKeys(t *testing.T) {
	s := NewStore()
	s.Set("alpha", 1)
	s.Set("beta", 2)
	s.Set("gamma", 3)
	// Setting an existing key again must not produce a duplicate entry.
	s.Set("beta", 20)

	want := []string{"alpha", "beta", "gamma"}

	got := s.Keys()
	if len(got) != len(want) {
		t.Fatalf("Keys() returned %d keys (%v); want %d (%v)", len(got), got, len(want), want)
	}

	// Keys() order is unspecified, so compare as sorted sets.
	sort.Strings(got)
	for i, k := range want {
		if got[i] != k {
			t.Errorf("Keys() = %v; want %v", got, want)
			break
		}
	}
}

// TestStoreKeysIsSnapshot verifies that the slice returned by Keys() is a
// snapshot: mutating it must not affect the store's internal state.
func TestStoreKeysIsSnapshot(t *testing.T) {
	s := NewStore()
	s.Set("one", 1)
	s.Set("two", 2)
	s.Set("three", 3)

	keys := s.Keys()
	if len(keys) != 3 {
		t.Fatalf("Keys() returned %d keys; want 3", len(keys))
	}

	// Mutate the returned slice: overwrite an element and truncate it.
	keys[0] = "mutated"
	keys = keys[:0]

	got := s.Keys()
	if len(got) != 3 {
		t.Fatalf("mutating the returned slice changed the store: Keys() now returns %d keys (%v); want 3", len(got), got)
	}

	want := map[string]bool{"one": true, "two": true, "three": true}
	for _, k := range got {
		if !want[k] {
			t.Errorf("Keys() returned unexpected key %q after mutation; want one of %v", k, want)
		}
	}
}
