package model

type Timeframe string

const (
	Timeframe1Min   Timeframe = "1min"
	Timeframe5Min   Timeframe = "5min"
	Timeframe15Min  Timeframe = "15min"
	Timeframe30Min  Timeframe = "30min"
	Timeframe1Hour  Timeframe = "1hour"
	Timeframe4Hour  Timeframe = "4hour"
	Timeframe1Day   Timeframe = "1day"
	Timeframe1Week  Timeframe = "1week"
	Timeframe1Month Timeframe = "1month"
	Timeframe1Year  Timeframe = "1year"
)
