package handlers

import (
	"fmt"
	"nechego/game"
	"nechego/teleutil"

	tele "gopkg.in/telebot.v3"
)

type DailyEblan struct {
	Universe *game.Universe
}

var dailyEblanRe = Regexp("^![ие][б6п]?л[ап]н[а-я]*")

func (h *DailyEblan) Match(s string) bool {
	return dailyEblanRe.MatchString(s)
}

func (h *DailyEblan) Handle(c tele.Context) error {
	world, _ := teleutil.Lock(c, h.Universe)
	defer world.Unlock()

	eblan, ok := world.DailyEblan()
	if !ok {
		return c.Send("😸")
	}
	who := teleutil.Link(c, teleutil.Member(c, tele.ChatID(eblan.TUID)))
	out := fmt.Sprintf("<b>Еблан дня</b> — %s 😸", who)
	return c.Send(out, tele.ModeHTML)
}

type DailyAdmin struct {
	Universe *game.Universe
}

var dailyAdminRe = Regexp("^!админ")

func (h *DailyAdmin) Match(s string) bool {
	return dailyAdminRe.MatchString(s)
}

func (h *DailyAdmin) Handle(c tele.Context) error {
	world, _ := teleutil.Lock(c, h.Universe)
	defer world.Unlock()

	admin, ok := world.DailyAdmin()
	if !ok {
		return c.Send("👑")
	}
	m := teleutil.Link(c, teleutil.Member(c, tele.ChatID(admin.TUID)))
	out := fmt.Sprintf("<b>Админ дня</b> — %s 👑", m)
	return c.Send(out, tele.ModeHTML)
}

type DailyPair struct {
	Universe *game.Universe
}

var dailyPairRe = Regexp("^!пара")

func (h *DailyPair) Match(s string) bool {
	return dailyPairRe.MatchString(s)
}

func (h *DailyPair) Handle(c tele.Context) error {
	world, _ := teleutil.Lock(c, h.Universe)
	defer world.Unlock()

	pair, ok := world.DailyPair()
	if !ok {
		return c.Send("💔")
	}
	m0 := teleutil.Link(c, tele.ChatID(pair[0].TUID))
	m1 := teleutil.Link(c, tele.ChatID(pair[1].TUID))
	out := fmt.Sprintf("<b>✨ Пара дня</b> — %s 💘 %s", m0, m1)
	return c.Send(out, tele.ModeHTML)
}
