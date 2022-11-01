package gache

import (
	"encoding/json"
	"io"
)

type Encoder interface {
	// Encode encodes the given data and writes it to the given writer.
	Encode(w io.Writer, data any) error
}

type Decoder interface {
	// Decode decodes the given data and writes it to the given writer.
	Decode(r io.Reader, data any) error
}

type defaultJSONEncoderDecoder struct {
}

func (defaultJSONEncoderDecoder) Encode(w io.Writer, data any) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(data)
}

func (defaultJSONEncoderDecoder) Decode(r io.Reader, data any) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(data)
}
