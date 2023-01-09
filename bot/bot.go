package main

import (
	"log"
	"nechego/game"
	"nechego/handlers"
	"os"
	"time"

	tele "gopkg.in/telebot.v3"
)

type Router struct {
	Handlers   []handlers.Handler
	Middleware []Wrapper
}

func (r *Router) OnText(c tele.Context) error {
	for _, h := range r.Handlers {
		if h.Match(c.Message().Text) {
			f := h.Handle
			for _, w := range r.Middleware {
				f = w.Wrap(f)
			}
			return f(c)
		}
	}
	return nil
}

func main() {
	pref := tele.Settings{
		Token:  os.Getenv("NECHEGO_TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}
	bot, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	universe := game.NewUniverse("universe")
	router := &Router{}
	router.Handlers = []handlers.Handler{
		&handlers.Mouse{Path: "data/mouse.mp4"},
		&handlers.Tiktok{Path: "data/tiktok/"},
		&handlers.Game{},
		&handlers.Infa{},
		&handlers.Who{Universe: universe},
		&handlers.Top{Universe: universe},
		&handlers.List{Universe: universe},
		&handlers.Save{Universe: universe},
		&handlers.Weather{},
		&handlers.Cat{},
		&handlers.Anime{},
		&handlers.Furry{},
		&handlers.Flag{},
		&handlers.Person{},
		&handlers.Horse{},
		&handlers.Art{},
		&handlers.Car{},
		&handlers.Name{},
		&handlers.DailyEblan{Universe: universe},
		&handlers.DailyAdmin{Universe: universe},
		&handlers.DailyPair{Universe: universe},
		&handlers.Inventory{Universe: universe},
		&handlers.Drop{Universe: universe},
		&handlers.Pick{Universe: universe},
		&handlers.Floor{Universe: universe},
		&handlers.Calculator{},
		&handlers.Market{Universe: universe},
		&handlers.Buy{Universe: universe},
		&handlers.Eat{Universe: universe},
		&handlers.Fish{Universe: universe},
		&handlers.Masyunya{},
		&handlers.Poppy{},
		&handlers.Sima{},
		&handlers.Hello{Path: "data/hello.json"},
		&handlers.Status{Universe: universe},
	}
	router.Middleware = []Wrapper{
		&MessageIncrementer{Universe: universe},
		&UserAdder{Universe: universe},
		&RequireSupergroup{},
	}
	go func() {
		for range time.NewTicker(time.Second * 30).C {
			universe.ForEachWorld(func(w *game.World) {
				w.RestoreEnergy()
				w.Market.Refill()
			})
		}
	}()
	bot.Handle(tele.OnText, router.OnText)
	bot.Start()
}
