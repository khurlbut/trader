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
}

func NewPurse(target float64, fee float64) *Purse {
     p := Purse{
          targetCashPercentage: target,
          tradingFeePercentage: fee,
     }
     return &p
}

func Init(target float64, fee float64) {
     targetCashPercentage = target
     tradingFeePercentage = fee
}

func (p *Purse) FundPurse(funds float64, spot float64) {
     p.coins = (funds * (1 - p.targetCashPercentage)) / spot
     p.cash = funds * p.targetCashPercentage
}
// func Fund(funds float64, spot float64) {
//      coins = (funds * (1 - targetCashPercentage)) / spot
//      cash = funds * targetCashPercentage
// }

func (p *Purse) Coins() float64 {
     return p.coins
}

func (p *Purse) CashHoldings() float64 {
     return p.cash
}

func (p *Purse) CoinValue(spot float64) float64 {
     return p.coins * spot
}

func (p *Purse) ValueAt(spot float64) float64 {
     return p.CashHoldings() + p.CoinValue(spot)
}

/*
     The "adjustment" value is either positive or negative reflecting a BUY or a SELL.
*/
func (p *Purse) CashRequiredToAlignWithTarget(spot float64) float64 {
     cashTarget := p.ValueAt(spot) * p.targetCashPercentage
     holdings := p.CashHoldings() 
     adjustment := cashTarget - holdings
     if math.Abs(adjustment) < p.minimum_transaction_amount {
          return 0
     }
     fee := p.tradingFee(adjustment)
     if feeCausesCostOverrun(fee, adjustment, holdings) {
          return adjustment + fee
     }
     return adjustment
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
     p.subCash(tradingFee(amount))
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
     return fmt.Sprintf("Spot: %f\tCashHoldings %f\tCoins: %f\t\tTotal Purse: %f", spot, p.CashHoldings(), p.Coins(), p.ValueAt(spot))
}

func (p *Purse) Properties() string {
     return fmt.Sprintf("Purse Properties:\nTarget Cash Percentage: %f\nTradingFeePercentage: %f\n", p.targetCashPercentage, p.tradingFeePercentage)
}

func (p *Purse) Holdings(spot float64) string {
     return fmt.Sprintf("Spot: %f\tCashHoldings %f\tCoins: %f\n", spot, p.cash, p.coins)
}
