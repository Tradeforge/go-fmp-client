package types

import (
	"encoding/json"
	"strings"
)

type ThousandSeparatedNumeric[T comparable] struct {
	value T
}

func (t ThousandSeparatedNumeric[T]) Value() T {
	return t.value
}

func (t ThousandSeparatedNumeric[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Value())
}

func (t *ThousandSeparatedNumeric[T]) UnmarshalJSON(data []byte) error {
	s := strings.ReplaceAll(string(data), ",", "")
	if len(s) == 0 {
		return nil
	}
	return json.Unmarshal([]byte(s), &t.value)
}
