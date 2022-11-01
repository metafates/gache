package gache

import (
	"os"
	"testing"
	"time"
)

func TestLoad(t *testing.T) {
	_ = os.WriteFile(testpath, []byte(`{"Internal":"test","Time":"2022-11-01T16:07:55.833489+03:00"}`), 0644)

	// given
	g := New[string](&Options{Path: testpath})

	if g.data.Internal != "" {
		t.Fatalf("Cache.data.Internal = %v, want %v", g.data.Internal, "")
	}

	// when
	err := g.load()

	// then
	if err != nil {
		t.Errorf("Cache.load() error = %v, wantErr %v", err, nil)
	}

	if g.data.Internal != "test" {
		t.Errorf("Cache.data.Internal = %v, want %v", g.data.Internal, "test")
	}

	clear()
}

func TestLoadMalformed(t *testing.T) {
	_ = os.WriteFile(testpath, []byte(`{"aaaaa":"test",!"oooo":-11-01T16:07:55.833489+03:00"`), 0644)

	// given
	g := New[string](&Options{Path: testpath})

	if g.data.Internal != "" {
		t.Fatalf("Cache.data.Internal = %v, want %v", g.data.Internal, "")
	}

	// when
	err := g.load()

	// then
	if err != nil {
		t.Errorf("Cache.load() error = %v, wantErr %v", err, nil)
	}

	if g.data.Internal != "" {
		t.Errorf("Cache.data.Internal = %v, want %v", g.data.Internal, "")
	}

	clear()
}

func TestLoadExpired(t *testing.T) {
	_ = os.WriteFile(testpath, []byte(`{"Internal":"test","Time":"2000-11-01T16:07:55.833489+03:00"}`), 0644)

	// given
	g := New[string](&Options{Path: testpath, Lifetime: time.Hour})

	if g.data.Internal != "" {
		t.Fatalf("Cache.data.Internal = %v, want %v", g.data.Internal, "")
	}

	// when
	err := g.load()

	// then
	if err != nil {
		t.Errorf("Cache.load() error = %v, wantErr %v", err, nil)
	}

	if g.data.Internal != "" {
		t.Errorf("Cache.data.Internal = %v, want %v", g.data.Internal, "")
	}

	clear()
}
