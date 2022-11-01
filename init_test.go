package gache

import "testing"

func TestGache_init(t *testing.T) {
	cache := New[string](&Options{})

	if cache.initialized {
		t.Fatalf("Cache.initialized = %v, want %v", cache.initialized, false)
	}

	err := cache.init()
	if err != nil {
		t.Fatalf("Cache.init() error = %v, wantErr %v", err, nil)
	}

	if !cache.initialized {
		t.Fatalf("Cache.initialized = %v, want %v", cache.initialized, true)
	}
}
