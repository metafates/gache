package gache

import (
	"errors"
	"time"
)

// Set sets the value of the cache.
// If initialization or marshalling fails, it will return an error.
// In memory-only mode it will never fail.
func (g *Gache[T]) Set(value T) error {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	err := g.init()
	if err != nil {
		return err
	}

	// update time
	g.data.Internal = value
	if g.data.Time != nil {
		now := time.Now()
		g.data.Time = &now
	}

	if g.options.Path != "" {
		err = g.save()
	}

	return err
}

// Get returns the value of the cache.
// If initialization fails, it will return an error.
// In memory-only mode it will never fail.
func (g *Gache[T]) Get() (T, error) {
	g.mutex.RLock()
	defer g.mutex.RUnlock()

	err := g.init()
	if err != nil {
		var t T
		return t, err
	}

	// Do not use tryExpire() here, because we can't write to the cache
	// since we RLocked mutex.
	if g.isExpired() {
		var t T
		return t, ErrExpired(errors.New("cache expired"))
	}

	return g.data.Internal, nil
}
