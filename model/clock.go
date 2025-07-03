package model

import (
	"encoding/json"
	"fmt"
	"regexp"
	"time"

	"go.tradeforge.dev/fmp/pkg/types"
)

const (
	tradingHoursRegexp     = `(1[0-2]|0?[1-9]):[0-5][0-9]\s*(AM|PM)`
	tradingHoursTimeFormat = "3:04 PM"
)

type GetAllExchangesTradingHoursResponse []ExchangeTradingHours

type ExchangeTradingHours struct {
	Exchange     string         `json:"name"`
	OpeningHour  types.TimeHHMM `json:"open"`
	ClosingHour  types.TimeHHMM `json:"close"`
	TimeZone     string         `json:"timezone"`
	IsMarketOpen bool           `json:"isMarketOpen"`
}

func (e *ExchangeTradingHours) UnmarshalJSON(data []byte) error {
	type marshallable struct {
		ExchangeName string `json:"name"`
		OpeningHour  string `json:"open"`
		ClosingHour  string `json:"close"`
		TimeZone     string `json:"timezone"`
		IsMarketOpen bool   `json:"isMarketOpen"`
	}

	var m marshallable
	if err := json.Unmarshal(data, &m); err != nil {
		return fmt.Errorf("unmarshalling exchange trading hours: %w", err)
	}
	e.Exchange = m.ExchangeName
	e.TimeZone = m.TimeZone

	r := regexp.MustCompile(tradingHoursRegexp)
	openingHours := r.FindString(m.OpeningHour)
	closingHours := r.FindString(m.ClosingHour)

	loc, err := time.LoadLocation(m.TimeZone)
	if err != nil {
		return fmt.Errorf("loading location: %w", err)
	}
	t, err := time.ParseInLocation(tradingHoursTimeFormat, openingHours, loc)
	if err != nil {
		return fmt.Errorf("parsing opening hours: %w", err)
	}
	e.OpeningHour = types.TimeHHMMFromTime(t)
	t, err = time.ParseInLocation(tradingHoursTimeFormat, closingHours, loc)
	if err != nil {
		return fmt.Errorf("parsing closing hours: %w", err)
	}
	e.ClosingHour = types.TimeHHMMFromTime(t)

	return nil
}

type GetExchangeHolidaysParams struct {
	Exchange string `query:"exchange"`
	Year     int
}

type GetExchangeHolidaysResponse []Holiday

type OriginalGetExchangeHolidaysResponse struct {
	HolidaysByYear []mappedHolidays `json:"stockMarketHolidays"`
}

type mappedHolidays map[string]any

func (h mappedHolidays) Year() int {
	return h["year"].(int)
}

func (h mappedHolidays) Get() []Holiday {
	holidays := make([]Holiday, 0, len(h)-1) // -1 to skip the "year" key
	for k, v := range h {
		if k == "year" {
			continue
		}
		holidays = append(holidays, Holiday{
			Date: types.Date(k),
			Name: v.(string),
		})
	}
	return holidays
}

type Holiday struct {
	Date types.Date `json:"date"`
	Name string     `json:"name"`
}
