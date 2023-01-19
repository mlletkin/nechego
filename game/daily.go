package game

import (
	"nechego/dates"
	"nechego/item"
	"nechego/token"
)

func (w *World) DailyEblan() (u *User, ok bool) {
	for _, u = range w.Users {
		if u.Eblan() {
			return u, true
		}
	}
	return w.rollDailyEblan(), true
}

func (w *World) rollDailyEblan() *User {
	u := w.RandomUser()
	u.Inventory.Add(&item.Item{
		Type:   item.TypeEblan,
		Value:  &token.Eblan{},
		Expire: dates.Tomorrow(),
	})
	return u
}

func (w *World) DailyAdmin() (u *User, ok bool) {
	for _, u = range w.Users {
		if u.Admin() {
			return u, true
		}
	}
	return w.rollDailyAdmin(), true
}

func (w *World) rollDailyAdmin() *User {
	u := w.RandomUser()
	u.Inventory.Add(&item.Item{
		Type:         item.TypeAdmin,
		Value:        &token.Admin{},
		Expire:       dates.Tomorrow(),
		Transferable: true,
	})
	return u
}

func (w *World) DailyPair() (pair []*User, ok bool) {
	if len(w.Users) < 2 {
		return nil, false
	}
	for _, u := range w.Users {
		if u.Pair() {
			pair = append(pair, u)
		}
		if len(pair) == 2 {
			break
		}
	}
	if len(pair) != 2 {
		return w.rollDailyPair()
	}
	return pair, true
}

func (w *World) rollDailyPair() (pair []*User, ok bool) {
	pair = w.RandomUsers(2)
	if len(pair) != 2 {
		return nil, false
	}
	pair[0].Inventory.Add(pairToken())
	pair[1].Inventory.Add(pairToken())
	return pair, true
}

func pairToken() *item.Item {
	return &item.Item{
		Type:   item.TypePair,
		Value:  &token.Pair{},
		Expire: dates.Tomorrow(),
	}
}
