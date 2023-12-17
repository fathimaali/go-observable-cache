package observablecache

import (
	"testing"
	"time"
	"strconv"
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
	key, _ := a.Get("key")
	if key != "value" || a.ttl != 10 {
		t.Errorf("Function Set is not working.")
	}
}

func addOneToCache(cache *LocalCache, t *testing.T) {
	value, found := cache.Get("key")
	if found == false {
		t.Errorf("Key not found in cache")
	}
	valueInt, err := strconv.Atoi(value)
	if err != nil {
		t.Errorf("Couldn't convert %s to integer", value)
	}
	value = strconv.Itoa(valueInt+1)
	cache.Set("key", value)
}

func TestIncrementingValues(t *testing.T) {
	cache := New(10)
	cache.Set("key", "0")
	for i:=0 ; i < 5000; i++ {
		go addOneToCache(&cache, t)
	}
	value, found := cache.Get("key")
	if found == false {
		t.Errorf("Key not found in cache")
	}
	if value != "5000" {
		t.Errorf("Function Set is not concurrent safe. Wanted 5000, received %s", value)
	}
}
func TestGet(t *testing.T) {
	a := New(10)
	a.Set("key", "value")
	key, _ := a.Get("key")
	if key != "value" || a.ttl != 10 {
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
