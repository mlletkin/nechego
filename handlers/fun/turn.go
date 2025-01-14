package fun

import (
	"math/rand/v2"

	"github.com/zxy248/nechego/game"
	"github.com/zxy248/nechego/handlers"
	tu "github.com/zxy248/nechego/teleutil"

	tele "gopkg.in/zxy248/telebot.v3"
)

type TurnOn struct {
	Universe *game.Universe
}

var turnOnRe = handlers.NewRegexp("^!(вкл|подкл|подруб)")

func (h *TurnOn) Match(c tele.Context) bool {
	return turnOnRe.MatchString(c.Text())
}

func (h *TurnOn) Handle(c tele.Context) error {
	world := tu.Lock(c, h.Universe)
	defer world.Unlock()

	world.Inactive = false
	es := [...]string{"🔈", "🔔", "✅", "🆗", "▶️"}
	e := es[rand.N(len(es))]
	return c.Send(e + " Робот включен.")
}

type TurnOff struct {
	Universe *game.Universe
}

var turnOffRe = handlers.NewRegexp("^!(выкл|откл|отруб)")

func (h *TurnOff) Match(c tele.Context) bool {
	return turnOffRe.MatchString(c.Text())
}

func (h *TurnOff) Handle(c tele.Context) error {
	world := tu.Lock(c, h.Universe)
	defer world.Unlock()

	world.Inactive = true
	es := [...]string{"🔇", "🔕", "💤", "❌", "⛔️", "🚫", "⏹"}
	e := es[rand.N(len(es))]
	return c.Send(e + " Робот выключен.")
}
