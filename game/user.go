package game

import (
	"nechego/fishing"
	"time"
)

type Gender int

const (
	GenderUnknown Gender = iota
	GenderMale
	GenderFemale
	GenderTrans
)

type User struct {
	TUID      int64
	Energy    int
	EnergyCap int
	Rating    float64
	Messages  int
	Banned    bool
	Birthday  time.Time
	Gender    Gender
	Status    string
	Inventory *Items
}

func NewUser(tuid int64) *User {
	return &User{
		TUID:      tuid,
		EnergyCap: 5,
		Rating:    1500,
		Inventory: NewItems(),
	}
}

func (u *User) Ban() {
	u.Banned = true
}

func (u *User) Unban() {
	u.Banned = false
}

func (u *User) IncrementMessages() {
	u.Messages++
}

func (u *User) AddRating(r float64) {
	u.Rating += r
}

func (u *User) SpendEnergy(e int) bool {
	if u.Energy < e {
		return false
	}
	u.Energy -= e
	return true
}

func (u *User) RestoreEnergy(e int) {
	u.Energy += e
	if u.Energy > u.EnergyCap {
		u.Energy = u.EnergyCap
	}
}

func (u *User) SpendMoney(n int) bool {
	w, ok := u.Wallet()
	if !ok {
		return false
	}
	if w.Money < n {
		return false
	}
	w.Money -= n
	return true
}

func (u *User) AddMoney(n int) {
	w, ok := u.Wallet()
	if !ok {
		return
	}
	w.Money += n
}

func (u *User) Total() int {
	t := 0
	for _, v := range u.Inventory.list() {
		switch x := v.Value.(type) {
		case *Wallet:
			t += x.Money
		case *CreditCard:
			t += x.Money
		case *Debt:
			t -= x.Money
		}
	}
	return t
}

func (u *User) InDebt() bool {
	for _, v := range u.Inventory.list() {
		switch v.Value.(type) {
		case *Debt:
			return true
		}
	}
	return false
}

func (u *User) IsEblan() bool {
	for _, v := range u.Inventory.list() {
		switch v.Value.(type) {
		case *EblanToken:
			return true
		}
	}
	return false
}

func (u *User) IsAdmin() bool {
	for _, v := range u.Inventory.list() {
		switch v.Value.(type) {
		case *AdminToken:
			return true
		}
	}
	return false
}

func (u *User) IsPair() bool {
	for _, v := range u.Inventory.list() {
		switch v.Value.(type) {
		case *PairToken:
			return true
		}
	}
	return false
}

func (u *User) Eat(i *Item) bool {
	switch x := i.Value.(type) {
	case *fishing.Fish:
		u.Inventory.Remove(i)
		e := 1
		if x.Heavy() {
			e = 2
		}
		u.RestoreEnergy(e)
		return true
	}
	return false
}
