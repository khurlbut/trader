package cash_percentage_adjuster


import (
)

type  CashPercentageAdjuster interface {
	CashPercentageTarget(spot float64) float64
}
