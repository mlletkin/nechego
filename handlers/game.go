package handlers

import (
	"errors"
	"fmt"
	"html"
	"math/rand"
	"nechego/format"
	"nechego/game"
	"nechego/teleutil"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"

	tele "gopkg.in/telebot.v3"
)

type Save struct {
	Universe *game.Universe
}

var saveRe = regexp.MustCompile("^!сохран")

func (h *Save) Match(s string) bool {
	return saveRe.MatchString(s)
}

func (h *Save) Handle(c tele.Context) error {
	if err := h.Universe.SaveAll(); err != nil {
		return err
	}
	return c.Send("💾 Игра сохранена.")
}

type Name struct{}

var nameRe = regexp.MustCompile("!имя (.*)")

func (h *Name) Match(s string) bool {
	return nameRe.MatchString(s)
}

func (h *Name) Handle(c tele.Context) error {
	name := html.EscapeString(teleutil.Args(c, nameRe)[1])
	const max = 16
	if utf8.RuneCountInString(name) > max {
		return c.Send(fmt.Sprintf("⚠️ Максимальная длина имени %d символов.", max))
	}

	if err := teleutil.Promote(c, teleutil.Member(c, c.Sender())); err != nil {
		return err
	}
	if err := c.Bot().SetAdminTitle(c.Chat(), c.Sender(), name); err != nil {
		return err
	}
	return c.Send(fmt.Sprintf("Имя <b>%s</b> установлено ✅", name), tele.ModeHTML)
}

type Inventory struct {
	Universe *game.Universe
}

var inventoryRe = regexp.MustCompile("^!инвентарь")

func (h *Inventory) Match(s string) bool {
	return inventoryRe.MatchString(s)
}

func (h *Inventory) Handle(c tele.Context) error {
	world := h.Universe.MustWorld(c.Chat().ID)
	world.Lock()
	defer world.Unlock()

	u, ok := world.UserByID(c.Sender().ID)
	if !ok {
		return errors.New("user not found")
	}
	items := u.Inventory.List()
	mention := teleutil.Mention(c, teleutil.Member(c, c.Sender()))
	head := fmt.Sprintf("<b>🗄 Инвентарь: %s</b>", mention)
	lines := append([]string{head}, format.Items(items)...)
	return c.Send(strings.Join(lines, "\n"), tele.ModeHTML)
}

type Drop struct {
	Universe *game.Universe
}

var dropRe = regexp.MustCompile("^!(выкинуть|выбросить) (.*)")

func (h *Drop) Match(s string) bool {
	return dropRe.MatchString(s)
}

func (h *Drop) Handle(c tele.Context) error {
	world := h.Universe.MustWorld(c.Chat().ID)
	world.Lock()
	defer world.Unlock()

	user, ok := world.UserByID(c.Sender().ID)
	if !ok {
		return errors.New("user not found")
	}
	key, err := strconv.Atoi(teleutil.Args(c, dropRe)[2])
	if err != nil {
		return c.Send("#⃣ Укажите номер предмета.")
	}
	item, ok := user.Inventory.ByKey(key)
	if !ok {
		return c.Send("🗄 Такого предмета нет в инвентаре.")
	}
	if ok := user.Inventory.Move(world.Floor, item); !ok {
		return c.Send("♻ Вы не можете выбросить этот предмет.")
	}
	out := fmt.Sprintf("🚮 Вы выбросили %s.", format.Item(item))
	return c.Send(out, tele.ModeHTML)
}

type Pick struct {
	Universe *game.Universe
}

var pickRe = regexp.MustCompile("^!взять (.*)")

func (h *Pick) Match(s string) bool {
	return pickRe.MatchString(s)
}

func (h *Pick) Handle(c tele.Context) error {
	world := h.Universe.MustWorld(c.Chat().ID)
	world.Lock()
	defer world.Unlock()

	user, ok := world.UserByID(c.Sender().ID)
	if !ok {
		return errors.New("user not found")
	}
	key, err := strconv.Atoi(teleutil.Args(c, pickRe)[1])
	if err != nil {
		return c.Send("#⃣ Укажите номер предмета.")
	}
	item, ok := world.Floor.ByKey(key)
	if !ok {
		return c.Send("🗄 Такого предмета нет на полу.")
	}
	if ok := world.Floor.Move(user.Inventory, item); !ok {
		return c.Send("♻ Вы не можете взять этот предмет.")
	}
	out := fmt.Sprintf("🫳 Вы взяли %s.", format.Item(item))
	return c.Send(out, tele.ModeHTML)
}

type Floor struct {
	Universe *game.Universe
}

var floorRe = regexp.MustCompile("^!пол")

func (h *Floor) Match(s string) bool {
	return floorRe.MatchString(s)
}

func (h *Floor) Handle(c tele.Context) error {
	world := h.Universe.MustWorld(c.Chat().ID)
	world.Lock()
	defer world.Unlock()

	items := world.Floor.List()
	head := "<b>🗑 Пол</b>"
	lines := append([]string{head}, format.Items(items)...)
	return c.Send(strings.Join(lines, "\n"), tele.ModeHTML)
}

type Market struct {
	Universe *game.Universe
}

var marketRe = regexp.MustCompile("^!магазин")

func (h *Market) Match(s string) bool {
	return marketRe.MatchString(s)
}

func (h *Market) Handle(c tele.Context) error {
	world := h.Universe.MustWorld(c.Chat().ID)
	world.Lock()
	defer world.Unlock()

	products := world.Market.Products()
	head := "<b>🏪 Магазин</b>"
	lines := append([]string{head}, format.Products(products)...)
	return c.Send(strings.Join(lines, "\n"), tele.ModeHTML)
}

type Buy struct {
	Universe *game.Universe
}

var buyRe = regexp.MustCompile("^!купить (.*)")

func (h *Buy) Match(s string) bool {
	return buyRe.MatchString(s)
}

func (h *Buy) Handle(c tele.Context) error {
	world := h.Universe.MustWorld(c.Chat().ID)
	world.Lock()
	defer world.Unlock()

	key, err := strconv.Atoi(teleutil.Args(c, buyRe)[1])
	if err != nil {
		return c.Send("#⃣ Укажите номер предмета.")
	}
	user, ok := world.UserByID(c.Sender().ID)
	if !ok {
		return errors.New("user not found")
	}
	product, ok := user.Buy(world.Market, key)
	if !ok {
		return c.Send("💵 Недостаточно средств.")
	}
	out := fmt.Sprintf("🛒 Вы приобрели %s за %s.", format.Item(product.Item), format.Money(product.Price))
	return c.Send(out, tele.ModeHTML)
}

type Eat struct {
	Universe *game.Universe
}

var eatRe = regexp.MustCompile("^!с[ъь]есть (.*)")

func (h *Eat) Match(s string) bool {
	return eatRe.MatchString(s)
}

func (h *Eat) Handle(c tele.Context) error {
	world := h.Universe.MustWorld(c.Chat().ID)
	world.Lock()
	defer world.Unlock()

	user, ok := world.UserByID(c.Sender().ID)
	if !ok {
		return c.Send("user not found")
	}
	if user.Energy == user.EnergyCap {
		return c.Send("🍊 Вы не хотите есть.")
	}
	key, err := strconv.Atoi(teleutil.Args(c, eatRe)[1])
	if err != nil {
		return c.Send("#⃣ Укажите номер предмета.")
	}
	item, ok := user.Inventory.ByKey(key)
	if !ok {
		return c.Send("🗄 Такого предмета нет в инвентаре.")
	}
	if ok := user.Eat(item); !ok {
		return c.Send("🤮")
	}
	out := fmt.Sprintf("🍊 Вы съели %s.\n\n<i>Энергии осталось: %s</i>",
		format.Item(item),
		format.Energy(user.Energy),
	)
	return c.Send(out, tele.ModeHTML)
}

type Fish struct {
	Universe *game.Universe
}

var fishRe = regexp.MustCompile("^!рыбалка")

func (h *Fish) Match(s string) bool {
	return fishRe.MatchString(s)
}

func (h *Fish) Handle(c tele.Context) error {
	world := h.Universe.MustWorld(c.Chat().ID)
	world.Lock()
	defer world.Unlock()

	user, ok := world.UserByID(c.Sender().ID)
	if !ok {
		return errors.New("user not found")
	}
	rod, ok := user.FishingRod()
	if !ok {
		return c.Send("🎣 Приобретите удочку, прежде чем рыбачить.")
	}
	if ok := user.SpendEnergy(1); !ok {
		return c.Send("⚡ Недостаточно энергии.")
	}
	fish := user.Fish(rod)
	if rod.Durability < 0 {
		c.Send("🎣 Ваша удочка сломалась.")
	}
	if rand.Float64() < 0.5 {
		outcomes := [...]string{
			"Вы не смогли выудить рыбу.",
			"Рыба сорвалась с крючка.",
			"Рыба сорвала леску.",
			"Рыба скрылась в водорослях.",
			"Рыба выскользнула из рук.",
			"Вы отпустили рыбу обратно в воду.",
		}
		return c.Send("🎣 " + outcomes[rand.Intn(len(outcomes))])
	}
	user.Inventory.Add(&game.Item{
		Type:         game.ItemTypeFish,
		Transferable: true,
		Value:        fish,
	})
	mention := teleutil.Mention(c, teleutil.Member(c, c.Sender()))
	out := fmt.Sprintf("🎣 %s получает рыбу: %s", mention, format.Fish(fish))
	return c.Send(out, tele.ModeHTML)
}

type Status struct {
	Universe *game.Universe
}

var statusRe = regexp.MustCompile("^!статус (.*)")

func (h *Status) Match(s string) bool {
	return statusRe.MatchString(s)
}

func (h *Status) Handle(c tele.Context) error {
	world := h.Universe.MustWorld(c.Chat().ID)
	world.Lock()
	defer world.Unlock()

	user, ok := world.UserByID(c.Sender().ID)
	if !ok {
		return errors.New("user not found")
	}
	status := teleutil.Args(c, statusRe)[1]
	const maxlen = 120
	if utf8.RuneCountInString(status) > maxlen {
		return c.Send("💬 Максимальная длина статуса %d символов.", maxlen)
	}
	user.Status = status
	return c.Send("✅ Статус установлен.")
}
