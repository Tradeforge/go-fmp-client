package types

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"strconv"
	"time"
)

type Date string

var (
	errInvalidDateFormat     = errors.New("invalid date format")
	errInvalidDateTimeFormat = errors.New("invalid date time format")
	errInvalidTimeFormat     = errors.New("invalid time format")
	errInvalidUnixTimestamp  = errors.New("invalid unix timestamp")
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

func TimeHHMMFromTime(t time.Time) TimeHHMM {
	return TimeHHMM(t.Format(time.TimeOnly))
}

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
	//nolint:perfsprint
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

// UnixTimestamp is a Unix epoch expressed in whole seconds.
//
// FMP serialises epoch fields as a bare JSON number. That number used to be
// integral (1768692032) and is now sometimes fractional, carrying sub-second
// precision (1768692032.131). Both forms decode here; the sub-second part is
// dropped, because every consumer of these fields reads whole seconds.
//
// A JSON null leaves the zero value and a JSON string is rejected, both
// matching the plain int64 this type replaced.
type UnixTimestamp int64

func UnixTimestampFromTime(t time.Time) UnixTimestamp {
	return UnixTimestamp(t.Unix())
}

func (u UnixTimestamp) Time() time.Time {
	return time.Unix(int64(u), 0)
}

func (u UnixTimestamp) Int64() int64 {
	return int64(u)
}

func (u *UnixTimestamp) UnmarshalJSON(data []byte) error {
	literal := string(data)
	if literal == "null" {
		return nil
	}
	if len(literal) > 0 && literal[0] == '"' {
		return fmt.Errorf("%w: expected a number, got a string: %s", errInvalidUnixTimestamp, literal)
	}

	var number json.Number
	if err := json.Unmarshal(data, &number); err != nil {
		return fmt.Errorf("%w: %s: %w", errInvalidUnixTimestamp, literal, err)
	}
	// The integral form is both the common case and exact for every int64,
	// so it never goes near a float64.
	if seconds, err := number.Int64(); err == nil {
		*u = UnixTimestamp(seconds)
		return nil
	}
	seconds, err := number.Float64()
	if err != nil {
		return fmt.Errorf("%w: %s: %w", errInvalidUnixTimestamp, literal, err)
	}
	if math.IsNaN(seconds) || math.IsInf(seconds, 0) || math.Abs(seconds) >= math.MaxInt64 {
		return fmt.Errorf("%w: out of range: %s", errInvalidUnixTimestamp, literal)
	}
	// Floor rather than truncate: the decoded value is the whole second the
	// instant falls in on either side of the epoch, where truncating toward
	// zero would push a pre-1970 fraction into the following second.
	*u = UnixTimestamp(math.Floor(seconds))
	return nil
}

func (u UnixTimestamp) String() string {
	return strconv.FormatInt(int64(u), 10)
}
