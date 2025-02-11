package pictures

import (
	"github.com/zxy248/nechego/handlers"

	tele "gopkg.in/zxy248/telebot.v3"
)

type Sima struct{}

func (h *Sima) Match(c tele.Context) bool {
	return handlers.HasPrefix(c.Text(), "!сима")
}

func (h *Sima) Handle(c tele.Context) error {
	s, err := randomSticker(c, "catsima_vk")
	if err != nil {
		return err
	}
	return c.Send(s)
}
