package ledger

import (
	"encoding/json"
	"fmt"
	"hash/fnv"
	"os"
	"testing"
)

type catalogItem struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Qty  int    `json:"qty"`
}

func loadCatalog(t *testing.T) []catalogItem {
	t.Helper()
	raw, err := os.ReadFile("data/items.json")
	if err != nil {
		t.Fatalf("read data/items.json: %v", err)
	}
	var items []catalogItem
	if err := json.Unmarshal(raw, &items); err != nil {
		t.Fatalf("parse data/items.json: %v", err)
	}
	return items
}

// TestLookupCoversCatalog asserts every item in data/items.json resolves
// through Lookup with the checksum the generator would produce.
func TestLookupCoversCatalog(t *testing.T) {
	for _, item := range loadCatalog(t) {
		got, ok := Lookup(item.ID)
		if !ok {
			t.Errorf("Lookup(%q) missing from ItemIndex", item.ID)
			continue
		}
		h := fnv.New64a()
		fmt.Fprintf(h, "%s:%s:%d", item.ID, item.Name, item.Qty)
		if want := h.Sum64(); got != want {
			t.Errorf("Lookup(%q) = %d, want %d", item.ID, got, want)
		}
	}
}

// TestItemIndexHasNoStaleEntries asserts the table carries nothing that is not
// in the catalog any more.
func TestItemIndexHasNoStaleEntries(t *testing.T) {
	known := make(map[string]bool)
	for _, item := range loadCatalog(t) {
		known[item.ID] = true
	}
	for id := range ItemIndex {
		if !known[id] {
			t.Errorf("ItemIndex has stale entry %q not present in data/items.json", id)
		}
	}
}
