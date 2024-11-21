package types

import (
	"fmt"
	"strconv"
)

type Bool struct {
	value bool
}

func (b Bool) BoolValue() bool {
	return b.value
}

func (b Bool) StringValue() string {
	return strconv.FormatBool(b.value)
}

func (b Bool) MarshalText() ([]byte, error) {
	return []byte(b.StringValue()), nil
}

func (b *Bool) UnmarshalText(data []byte) error {
	parsedValue, err := strconv.ParseBool(string(data))
	if err != nil {
		return err
	}
	*b = Bool{value: parsedValue}
	return nil
}

func (b *Bool) Scan(src interface{}) error {
	if src == nil {
		*b = Bool{}
		return nil
	}

	switch src := src.(type) {
	case bool:
		*b = Bool{value: src}
		return nil
	case string:
		parsedValue, err := strconv.ParseBool(src)
		if err != nil {
			return err
		}
		*b = Bool{value: parsedValue}
		return nil
	case []byte:
		parsedValue, err := strconv.ParseBool(string(src))
		if err != nil {
			return err
		}
		*b = Bool{value: parsedValue}
		return nil
	}
	return fmt.Errorf("cannot convert %T to Bool", src)
}
