package form

import (
	"bytes"
	"reflect"
	"strings"
)

// DecodeErrors is a map of errors encountered during form decoding
type DecodeErrors map[string]error

func (d DecodeErrors) Error() string {
	buff := bytes.NewBufferString(blank)

	for k, err := range d {
		buff.WriteString(fieldNS)
		buff.WriteString(k)
		buff.WriteString(errorText)
		buff.WriteString(err.Error())
		buff.WriteString("\n")
	}

	return strings.TrimSpace(buff.String())
}

// An InvalidDecoderError describes an invalid argument passed to Decode.
// (The argument passed to Decode must be a non-nil pointer.)
type InvalidDecoderError struct {
	Type reflect.Type
}

func (e *InvalidDecoderError) Error() string {

	if e.Type == nil {
		return "form: Decode(nil)"
	}

	if e.Type.Kind() != reflect.Ptr {
		return "form: Decode(non-pointer " + e.Type.String() + ")"
	}

	return "form: Decode(nil " + e.Type.String() + ")"
}
