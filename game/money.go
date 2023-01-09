package game

import (
	"fmt"
	"time"
)

type Wallet struct {
	Money int
}

func (w Wallet) String() string {
	return fmt.Sprintf("💰 Кошелек (%d ₽)", w.Money)
}

func (u *User) Wallet() (w *Wallet, ok bool) {
	for _, v := range u.Inventory.list() {
		switch x := v.Value.(type) {
		case *Wallet:
			return x, true
		}
	}
	return nil, false
}

type CreditCard struct {
	Bank    int
	Number  int
	Expires time.Time
	Money   int
}

func (c CreditCard) String() string {
	return fmt.Sprintf("💳 Кредитная карта (%d ₽)", c.Money)
}

type Debt struct {
	CreditorID int
	Money      int
	Percent    int
}

func (d Debt) String() string {
	return fmt.Sprintf("💵 Долг (%d ₽, %d%%)", d.Money, d.Percent)
}
