package main

import (
  "log"
  "fmt"
)

type Money struct {
  amount uint
  currency string
}

func NewMoney(amount uint) *Money {
  return &Money{amount, "yen"}
}

func (this *Money) Format() string {
  return fmt.Sprintf("%d %s", this.amount, this.currency)
}

func (this *Money) Add(that *Money) {
  this.amount += that.amount
}

func main() {
  money := NewMoney(120)
  log.Printf(money.Format())
  money.Add(NewMoney(180))
  log.Printf(money.Format())
}
