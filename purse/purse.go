package purse

import (
     "fmt"
     "math"
)

var coins float64
var cash float64
var targetCashPercentage float64
var tradingFeePercentage float64

const minimum_transaction_amount = 10.00

type Purse struct {
     coins     float64
     cash      float64
     targetCashPercentage     float64
     tradingFeePercentage     float64
     minimum_transaction_amt  float64
}

func NewPurse(target float64, fee float64) *Purse {
     p := Purse{
          targetCashPercentage: target,
          tradingFeePercentage: fee,
          minimum_transaction_amt: 10.0,
     }
     return &p
}

func (p *Purse) Fund(funds float64, spot float64) {
     p.coins = (funds * (1 - p.targetCashPercentage)) / spot
     p.cash = funds * p.targetCashPercentage
}

func (p *Purse) SetTargetCashPercentage(target float64) float64 {
     p.targetCashPercentage = target
     return p.targetCashPercentage
}

/*
     The "adjustment" value is either positive or negative reflecting a BUY or a SELL.
*/
func (p *Purse) CashRequiredToAlignWithTarget(spot float64) float64 {
     cashTarget := p.valueAt(spot) * p.targetCashPercentage
     adjustment := cashTarget - p.cash

     if math.Abs(adjustment) < p.minimum_transaction_amt {
          return 0
     }

     fee := p.tradingFee(adjustment)
     if feeCausesCostOverrun(fee, adjustment, p.cash) {
          return adjustment + fee
     }

     return adjustment
}

func (p *Purse) valueAt(spot float64) float64 {
     return p.cash + (p.coins * spot)
}

func feeCausesCostOverrun(fee float64, adjustment float64, holdings float64) bool {
     if adjustment >= holdings {
          return false
     }
     return (fee + math.Abs(adjustment)) > holdings 
}

/*
     The "amount" is either positive or negative reflecting a BUY or a SELL.
     A BUY will "add" a negative value to Cash, etc.
 */
func (p *Purse) ReflectOrderFill(amount float64, spot float64) {
     p.addCash(amount)
     p.subCash(p.tradingFee(amount))
     p.subCoins(amount / spot)     
}

func (p *Purse) addCash(c float64) {
     p.cash = p.cash + c
}

func (p *Purse) subCash(c float64) {
     p.cash = p.cash - c
}

func (p *Purse) subCoins(c float64) {
     p.coins = p.coins - c
}

func (p *Purse) tradingFee(amt float64) float64 {
     return math.Abs(amt) * p.tradingFeePercentage
}

func  (p *Purse) String(spot float64) string {
     return fmt.Sprintf("Spot: %f\tCash: %f\tCoins: %f\t\tTotal Purse: %f", spot, p.cash, p.coins, p.valueAt(spot))
}
