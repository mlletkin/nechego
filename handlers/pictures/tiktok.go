package pictures

import (
	"nechego/handlers"

	tele "gopkg.in/telebot.v3"
)

type Tiktok struct {
	Path string
}

func (h *Tiktok) Match(s string) bool {
	return handlers.MatchPrefix("!тикток", s)
}

func (h *Tiktok) Handle(c tele.Context) error {
	f, err := randomFile(h.Path)
	if err != nil {
		return err
	}
	return c.Send(&tele.Video{File: tele.FromDisk(f)})
}