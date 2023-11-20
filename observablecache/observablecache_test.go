package observablecache

import "testing"

func TestNew(t *testing.T) {
	a := New()
	if a.store == nil {
		t.Errorf("Function New has empty store")
	}
}

func TestSet(t *testing.T) {
	a := New()
	a.Set("key", "value")
	if a.store["key"] != "value" {
		t.Errorf("Function Set is not working.")
	}
}
func TestGet(t *testing.T) {
	a := New()
	a.store["key"] = "value"
	if a.Get("key") != "value" {
		t.Errorf("Function Get is returning nil")
	}
}
