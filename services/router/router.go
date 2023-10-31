package router

import (
	"nechego/chat"
	"nechego/game"
	"nechego/services"
	"nechego/services/command"

	tele "gopkg.in/telebot.v3"
)

type Service struct {
	Universe *game.Universe
}

func (h *Service) Match(tele.Context) bool {
	return true
}

func (h *Service) Handle(c tele.Context) error {
	m := chat.Convert(c)
	handlers := []services.Handler{
		&command.Add{Universe: h.Universe},
		&command.Remove{Universe: h.Universe},
		&command.Use{Universe: h.Universe},
	}
	for _, h := range handlers {
		r := h.Match(m)
		if r != nil {
			return r.Process()
		}
	}
	return nil
}