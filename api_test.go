package gache

import (
	"os"
	"testing"
)

const testpath = "test/api_test.json"

func init() {
	_ = os.Remove(testpath)
}

func testEmptyGet(t *testing.T, g *Gache[string]) {
	t.Helper()

	// when get value from empty string cache
	value, err := g.Get()

	// error should be nil
	if err != nil {
		t.Fatalf("Gache.Get() error = %v, wantErr %v", err, nil)
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
	value, err := g.Get()

	// error should be nil
	if err != nil {
		t.Fatalf("Gache.Get() error = %v, wantErr %v", err, nil)
	}

	// value should be "test"
	if value != "test" {
		t.Errorf("Gache.Get() = %v, want %v", value, "test")
	}
}

func TestGache_EmptyGet(t *testing.T) {
	testEmptyGet(t, New[string](&Options{}))
	testEmptyGet(t, New[string](&Options{Path: testpath}))
}

func TestGache_Set(t *testing.T) {
	testSet(t, New[string](&Options{}))
	testSet(t, New[string](&Options{Path: testpath}))
}
