package gache

import (
	"os"
	"testing"
	"time"
)

const testpath = "test/api_test.json"

func init() {
	clear()
}

func clear() {
	_ = os.Remove(testpath)
}

func testEmptyGet(t *testing.T, g *Gache[string]) {
	t.Helper()

	// when get value from empty string cache
	value, _, err := g.Get()

	// error should be nil
	if err != nil {
		t.Fatalf("Gache.Get() error = %v, wantErr %v", err, nil)
	}

	// value should be empty string
	if value != "" {
		t.Errorf("Gache.Get() = %v, want %v", value, "")
	}
}

func testSetExpire(t *testing.T, g *Gache[string]) {
	t.Helper()

	// given a cache with lifetime of
	lifetime := g.options.Lifetime

	const v = "test"

	// when setting value to cache
	err := g.Set(v)

	// error should be nil
	if err != nil {
		t.Fatalf("Gache.Set() error = %v, wantErr %v", err, nil)
	}

	// when getting value from cache
	value, _, err := g.Get()

	// error should be nil
	if err != nil {
		t.Fatalf("Gache.Get() error = %v, wantErr %v", err, nil)
	}

	// value should be "test"
	if value != v {
		t.Errorf("Gache.Get() = %v, want %v", value, v)
	}

	// if lifetime is not infinite (values below 0 are considered infinite lifetime)
	if lifetime >= 0 {
		// wait for the cache to expire
		time.Sleep(lifetime)
	}

	// when getting value from cache
	value, expired, err := g.Get()

	// err should be nil
	if err != nil {
		t.Fatalf("Gache.Get() error = %v, wantErr %v", err, nil)
	}

	// expired should be true
	if !expired {
		t.Errorf("Gache.Get() expired = %v, want %v", expired, true)
	}

	// value should be empty string
	if value != "" {
		t.Errorf("Gache.Get() = %v, want %v", value, "")
	}
}

func testSet(t *testing.T, g *Gache[string]) {
	t.Helper()

	// when set value to empty string cache
	err := g.Set("test")

	// error should be nil
	if err != nil {
		t.Fatalf("Gache.Set() error = %v, wantErr %v", err, nil)
	}

	// when getting value from cache
	value, _, err := g.Get()

	// error should be nil
	if err != nil {
		t.Fatalf("Gache.Get() error = %v, wantErr %v", err, nil)
	}

	// value should be "test"
	if value != "test" {
		t.Errorf("Gache.Get() = %v, want %v", value, "test")
	}
}

func testCrossSession(t *testing.T) {
	t.Helper()

	mkCache := func() *Gache[string] {
		return New[string](&Options{Path: testpath})
	}

	cache := mkCache()

	err := cache.Set("hello")
	if err != nil {
		t.Fatalf("Gache.Set() error = %v, wantErr %v", err, nil)
	}

	// reset cache
	cache = mkCache()

	value, _, err := cache.Get()
	if err != nil {
		t.Fatalf("Gache.Get() error = %v, wantErr %v", err, nil)
	}

	if value != "hello" {
		t.Errorf("Gache.Get() = %v, want %v", value, "hello")
	}

	clear()
}

func TestGache_EmptyGet(t *testing.T) {
	testEmptyGet(t, New[string](&Options{}))
	testEmptyGet(t, New[string](&Options{Path: testpath}))
}

func TestGache_Set(t *testing.T) {
	testSet(t, New[string](&Options{}))
	testSetExpire(t, New[string](&Options{Lifetime: time.Second}))

	testSet(t, New[string](&Options{Path: testpath}))
	clear()
	testSetExpire(t, New[string](&Options{Path: testpath, Lifetime: time.Second}))
	clear()

	testCrossSession(t)
}
