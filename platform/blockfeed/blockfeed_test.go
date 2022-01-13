package blockfeed

import "testing"

func TestAdd(t *testing.T) {
	feed := New()
	feed.Add(Block{})
	if len(feed.Blocks) != 1 {
		t.Errorf("Block was not added")
	}
}
func TestGetAll(t *testing.T) {
	feed := New()
	feed.Add(Block{})
	results := feed.GetAll()
	if len(results) != 1 {
		t.Errorf("Blocks not listed")
	}
}
