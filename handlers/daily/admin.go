package daily

import (
	"fmt"
	"github.com/zxy248/nechego/game"
	"github.com/zxy248/nechego/handlers"
	tu "github.com/zxy248/nechego/teleutil"

	tele "gopkg.in/zxy248/telebot.v3"
)

type Admin struct {
	Universe *game.Universe
}

var adminRe = handlers.NewRegexp("^!админ")

func (h *Admin) Match(c tele.Context) bool {
	return adminRe.MatchString(c.Text())
}

func (h *Admin) Handle(c tele.Context) error {
	world := tu.Lock(c, h.Universe)
	defer world.Unlock()

	u := world.DailyAdmin()
	l := tu.Link(c, u)
	s := fmt.Sprintf("<b>Админ дня</b> — %s 👑", l)
	return c.Send(s, tele.ModeHTML)
}
