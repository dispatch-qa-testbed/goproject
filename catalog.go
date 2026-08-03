package ledger

// Lookup returns the catalog checksum recorded for id and whether the id is
// known. Unknown ids return (0, false).
//
// The backing table lives in item_index.go, which is generated from
// data/items.json by gen/gen_index.py.
func Lookup(id string) (uint64, bool) {
	v, ok := ItemIndex[id]
	return v, ok
}
