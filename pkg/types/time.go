package types

import (
	"errors"
	"fmt"
	"time"
)

type Date string

var (
	errInvalidDateFormat     = fmt.Errorf("invalid date format")
	errInvalidDateTimeFormat = fmt.Errorf("invalid date time format")
	errInvalidTimeFormat     = fmt.Errorf("invalid time format")
)

func DateFromTime(t time.Time) Date {
	return Date(t.Format(time.DateOnly))
}

func (d Date) Time() time.Time {
	t, err := time.Parse(time.DateOnly, string(d))
	if err != nil {
		panic(errors.Join(errInvalidTimeFormat, err))
	}
	return t
}

func (d *Date) Scan(data any) error {
	if data == nil {
		return nil
	}
	t, ok := data.(time.Time)
	if !ok {
		return errInvalidDateFormat
	}
	*d = Date(t.Format(time.DateOnly))
	return nil
}

func (d *Date) MarshalText() ([]byte, error) {
	t, err := time.Parse(time.DateOnly, string(*d))
	if err != nil {
		return nil, errors.Join(errInvalidTimeFormat, err)
	}
	return []byte(fmt.Sprintf("%04d-%02d-%02d", t.Year(), t.Month(), t.Day())), nil
}

func (d *Date) UnmarshalText(data []byte) error {
	t, err := time.Parse(time.DateOnly, string(data))
	if err != nil {
		return errors.Join(errInvalidTimeFormat, err)
	}
	*d = Date(t.Format(time.DateOnly))
	return nil
}

func (d Date) String() string {
	b, err := d.MarshalText()
	if err != nil {
		panic(fmt.Errorf("marshalling date: %w", err))
	}
	return string(b)
}

type TimeHHMM string

func (o TimeHHMM) Time() time.Time {
	t, err := time.Parse(time.TimeOnly, string(o))
	if err != nil {
		panic(errors.Join(errInvalidTimeFormat, err))
	}
	return t
}

func (o TimeHHMM) Duration() time.Duration {
	t, err := time.Parse(time.TimeOnly, string(o))
	if err != nil {
		panic(errors.Join(errInvalidTimeFormat, err))
	}
	return time.Duration(t.Hour())*time.Hour + time.Duration(t.Minute())*time.Minute
}

func (o *TimeHHMM) Scan(data any) error {
	if data == nil {
		return nil
	}
	var t time.Time
	switch typed := data.(type) {
	case string:
		parsed, err := time.Parse(time.TimeOnly, typed)
		if err != nil {
			panic(errors.Join(errInvalidTimeFormat, err))
		}
		t = parsed
	case time.Time:
		t = typed
	}
	*o = TimeHHMM(t.Format(time.TimeOnly))
	return nil
}

func (o *TimeHHMM) MarshalText() ([]byte, error) {
	t, err := time.Parse(time.TimeOnly, *(*string)(o))
	if err != nil {
		return nil, errors.Join(errInvalidTimeFormat, err)
	}
	return []byte(fmt.Sprintf("%02d:%02d", t.Hour(), t.Minute())), nil
}

func (o *TimeHHMM) UnmarshalText(data []byte) error {
	t, err := time.Parse(time.TimeOnly, fmt.Sprintf("%s:00", string(data)))
	if err != nil {
		panic(errors.Join(errInvalidTimeFormat, err))
	}
	*o = TimeHHMM(t.Format(time.TimeOnly))
	return nil
}

func (o TimeHHMM) String() string {
	b, err := o.MarshalText()
	if err != nil {
		panic(fmt.Errorf("marshalling time: %w", err))
	}
	return string(b)
}

type DateTime string

func DateTimeFromTime(t time.Time) DateTime {
	return DateTime(t.Format(time.DateTime))
}

func (d DateTime) Time() time.Time {
	t, err := time.Parse(time.DateTime, string(d))
	if err != nil {
		panic(errors.Join(errInvalidTimeFormat, err))
	}
	return t
}

func (d *DateTime) Scan(data any) error {
	if data == nil {
		return nil
	}
	t, ok := data.(time.Time)
	if !ok {
		return errInvalidDateTimeFormat
	}
	*d = DateTime(t.Format(time.DateTime))
	return nil
}

func (d *DateTime) MarshalText() ([]byte, error) {
	t, err := time.Parse(time.DateTime, string(*d))
	if err != nil {
		return nil, errors.Join(errInvalidDateTimeFormat, err)
	}
	return []byte(t.Format(time.DateTime)), nil
}

func (d *DateTime) UnmarshalText(data []byte) error {
	t, err := time.Parse(time.DateTime, string(data))
	if err != nil {
		return errors.Join(errInvalidDateTimeFormat, err)
	}
	*d = DateTime(t.Format(time.DateTime))
	return nil
}

func (d DateTime) String() string {
	b, err := d.MarshalText()
	if err != nil {
		panic(fmt.Errorf("marshalling date time: %w", err))
	}
	return string(b)
}
