package gache

import (
	"time"
)

// Options is a struct that contains options for the cache.
type Options struct {
	// Path to the file to store the cache.
	// If not specified (empty string), the cache will be stored in memory.
	Path string

	// Lifetime is the time duration after which the cache expires.
	// If nil is set, the cache will never expire.
	Lifetime *time.Duration

	// ExpirationHook is a function that is called when the cache expires.
	// If nil is set, the cache will not call the function.
	ExpirationHook func()

	// FileSystem is a filesystem that is used to store the cache file.
	FileSystem FileSystem

	// Encoder is the encoder to use for the cache.
	Encoder Encoder

	// Decoder is the decoder to use for the cache.
	Decoder Decoder
}
