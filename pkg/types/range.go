package types

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"slices"

	"github.com/shopspring/decimal"
)

const RangePattern = `(?P<min>[-+]?\d*[\.,]?\d+([eE][-+]?\d+)?)-(?P<max>[-+]?\d*[\.,]?\d+([eE][-+]?\d+)?)`

var rangeRegexp = regexp.MustCompile(RangePattern)

type Range52w struct {
	Min decimal.Decimal
	Max decimal.Decimal
	Sep string
}

func (r Range52w) String() string {
	return fmt.Sprintf("%s%s%s", r.Min.String(), r.Sep, r.Max.String())
}

func (r Range52w) IsEmpty() bool {
	return r.Min.IsZero() && r.Max.IsZero()
}

func (r Range52w) MarshalJSON() ([]byte, error) {
	if r.IsEmpty() {
		return []byte(`""`), nil
	}
	return json.Marshal(r.String())
}

func (r *Range52w) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("unmarshalling range: %w", err)
	}
	rng, err := ParseRange52w(s, "-")
	if err != nil {
		return fmt.Errorf("parsing range: %w", err)
	}
	r.Min = rng.Min
	r.Max = rng.Max
	r.Sep = rng.Sep
	return nil
}

func ParseRange52w(r string, separator string) (Range52w, error) {
	m, err := regexp.MatchString(RangePattern, r)
	if err != nil {
		return Range52w{}, fmt.Errorf("invalid range: %w", err)
	}
	if !m {
		return Range52w{}, fmt.Errorf("invalid range: %s", r)
	}

	rMin, rMax, err := parseRangeCaptureGroups(rangeRegexp, r)
	if err != nil {
		return Range52w{}, fmt.Errorf("parsing range: %w", err)
	}
	return Range52w{
		Min: rMin,
		Max: rMax,
		Sep: separator,
	}, nil
}

func parseRangeCaptureGroups(r *regexp.Regexp, str string) (decimal.Decimal, decimal.Decimal, error) {
	rMin, rMax := decimal.Zero, decimal.Zero
	matches := r.FindStringSubmatch(str)
	if len(matches) == 0 {
		return rMin, rMax, errors.New("no matches found")
	}
	if !slices.Contains(r.SubexpNames(), "min") || !slices.Contains(r.SubexpNames(), "max") {
		return rMin, rMax, errors.New("missing submatch names")
	}
	rMin, err := decimal.NewFromString(matches[r.SubexpIndex("min")])
	if err != nil {
		return rMin, rMax, fmt.Errorf("parsing min: %w", err)
	}
	rMax, err = decimal.NewFromString(matches[r.SubexpIndex("max")])
	if err != nil {
		return rMin, rMax, fmt.Errorf("parsing max: %w", err)
	}
	return rMin, rMax, nil
}
