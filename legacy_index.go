package ledger

// Deprecated: LegacyItemIndex is the pre-generator lookup table, kept only
// while callers migrate to ItemIndex in item_index.go. It is not regenerated
// and its values are known to be stale.
var LegacyItemIndex = map[string]uint64{
	"a1": 11529618174142360670,
	"b2": 6763972058387213228,
	"c3": 2935649884673794255,
}

// Deprecated: LegacyLookup reads the frozen LegacyItemIndex table. Use Lookup
// instead, which reads the generated ItemIndex.
func LegacyLookup(id string) (uint64, bool) {
	v, ok := LegacyItemIndex[id]
	return v, ok
}
