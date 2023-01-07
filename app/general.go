package app

import (
	"math/rand"

	tele "gopkg.in/telebot.v3"
)

// !инфа
func (a *App) handleProbability(c tele.Context) error {
	m := getMessage(c).Argument()
	return respond(c, probabilityResponse(m))
}

var probabilityTemplates = []string{
	"Здравый смысл говорит мне о том, что %s с вероятностью %d%%",
	"Благодаря чувственному опыту я определил, что %s с вероятностью %d%%",
	"Я думаю, что %s с вероятностью %d%%",
	"Используя диалектическую логику, я пришел к выводу, что %s с вероятностью %d%%",
	"Проведя некие изыскания, я высяснил, что %s с вероятностью %d%%",
	"Я провел мысленный экперимент и выяснил, что %s с вероятностью %d%%",
	"Мои интеллектуальные потуги привели меня к тому, что %s с вероятностью %d%%",
	"С помощью фактов и логики я доказал, что %s с вероятностью %d%%",
	"Как показывает практика, %s с вероятностью %d%%",
	"Прикинув раз на раз, я определился с тем, что %s с вероятностью %d%%",
	"Уверяю вас в том, что %s с вероятностью %d%%",
}

func randomProbabilityTemplate() string {
	return probabilityTemplates[rand.Intn(len(probabilityTemplates))]
}

func probabilityResponse(message string) Response {
	return Response(randomProbabilityTemplate()).Fill(message, rand.Intn(100+1))
}

const who = Response("%s %s")

// !кто
func (a *App) handleWho(c tele.Context) error {
	message := getMessage(c).Argument()
	u, err := a.service.Who(getGroup(c), message)
	if err != nil {
		return respondInternalError(c, err)
	}
	return respond(c, who.Fill(a.mention(u), message))
}

const list = Response("Список %s 📝\n%s")

// !список
func (a *App) handleList(c tele.Context) error {
	users, err := a.service.List(getGroup(c), a.pref.ListLength.Random())
	if err != nil {
		return respondInternalError(c, err)
	}
	return respond(c, list.Fill(
		getMessage(c).Argument(),
		a.itemizeUsers(users...)))
}

const (
	numberedTopTemplate   = Response("Топ %d %s 🏆\n%s")
	unnumberedTopTemplate = Response("Топ %s 🏆\n%s")
	badTopNumber          = UserError("Некорректное число.")
)

// !топ
func (a *App) handleTop(c tele.Context) error {
	argument, err := getMessage(c).TopArgument()
	if err != nil {
		return respondInternalError(c, err)
	}
	var number int
	if argument.Number != nil {
		number = *argument.Number
	} else {
		number = maxTopNumber
	}
	if number <= 0 || number > maxTopNumber {
		return respondUserError(c, badTopNumber)
	}
	users, err := a.service.List(getGroup(c), number)
	if err != nil {
		return respondInternalError(c, err)
	}
	if argument.Number != nil {
		return respond(c, numberedTopTemplate.Fill(
			number,
			argument.String,
			a.enumerateUsers(users...)))
	}
	return respond(c, unnumberedTopTemplate.Fill(
		argument.String,
		a.enumerateUsers(users...)))
}
