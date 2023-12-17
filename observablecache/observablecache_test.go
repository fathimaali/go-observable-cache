package observablecache

import (
	"testing"
	"time"
)

func TestNewTTLNotSpecified(t *testing.T) {
	a := New()
	if a.store == nil || a.ttl == 0 {
		t.Errorf("Function New has empty store and or ttl is blank")
	}
}

func TestNewTTLSpecified(t *testing.T) {
	a := New(10)
	if a.store == nil || a.ttl != 10 {
		t.Errorf("Function New has empty store and or the ttl is not set properly")
	}
}

func TestSet(t *testing.T) {
	a := New(10)
	a.Set("key", "value")
	if a.Get("key") != "value" || a.ttl != 10 {
		t.Errorf("Function Set is not working.")
	}
}
func TestGet(t *testing.T) {
	a := New(10)
	a.Set("key", "value")
	if a.Get("key") != "value" || a.ttl != 10 {
		t.Errorf("Function Get is returning nil")
	}
}

func TestPurge(t *testing.T) {
	a := New(1)
	a.Set("key", "value")
	time.Sleep(2)
	if a.store == nil {
		t.Errorf("Function Purge is not purging")
	}
}

