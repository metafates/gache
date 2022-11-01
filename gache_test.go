package gache

import "testing"

func TestNew(t *testing.T) {
	// given options
	options := &Options{
		Path: "test",
	}

	// when create new gache with options
	g := New[string](options)

	// then
	if g.options.Path != "test" {
		t.Errorf("New().options.Path = %v, want %v", g.options.Path, "test")
	}
}
