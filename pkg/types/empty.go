package types

import "encoding/json"

type EmptyOr[T any] struct {
	hasValue bool
	value    *T
}

func (e EmptyOr[T]) Value() *T {
	return e.value
}

func (e EmptyOr[T]) IsEmpty() bool {
	return !e.hasValue
}

func (e EmptyOr[T]) MarshalJSON() ([]byte, error) {
	if e.hasValue {
		return json.Marshal(e.value)
	}
	return nil, nil
}

func (e *EmptyOr[T]) UnmarshalJSON(data []byte) error {
	if string(data) == `""` || len(data) == 0 {
		e.hasValue = false
		return nil
	}
	e.hasValue = true
	return json.Unmarshal(data, &e.value)
}
