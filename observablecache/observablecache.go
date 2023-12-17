package observablecache

import (
	"log"
	"strconv"
	"sync"
	"time"
)

var a = time.Now

const (
	DefaultGCInterval = 1
	DefaultTTL        = 1
)

type value struct {
	data   string
	expiry time.Time
}
type LocalCache struct {
	store      map[string]value
	ttl        int
	gcInterval int
	mutex      *sync.RWMutex
}

func (c *LocalCache) Get(key string) (string, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if !c.store[key].expiry.After(time.Now()) {
		return "", false
	}
	return c.store[key].data, true
}

func (c *LocalCache) Set(key string, val string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.store[key] = value{
		data:   val,
		expiry: a().Add(time.Minute * time.Duration(c.ttl)),
	}
}

func New(duration ...int) LocalCache {
	var expiryTime int
	if len(duration) > 0 {
		expiryTime = duration[0]
	} else {
		expiryTime = DefaultTTL
	}
	cache := LocalCache{
		store:      make(map[string]value),
		ttl:        expiryTime,
		gcInterval: DefaultGCInterval,
		mutex:      &sync.RWMutex{},
	}
	// call routine to purge
	go cache.Purge()
	return cache
}

func (c *LocalCache) Purge() {
	for {
		for key, value := range c.store {
			if (value.expiry).Before(time.Now()) {
				c.mutex.Lock()
				delete(c.store, key)
				c.mutex.Unlock()
			}
		}
		time.Sleep(time.Duration(c.gcInterval) * time.Minute)
		log.Println("Garbage collected for every " + strconv.Itoa(c.gcInterval) + " minute(s).")
	}
}
