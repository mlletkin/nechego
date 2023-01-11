package app

import (
	"errors"
	"fmt"
	"nechego/model"
	"nechego/service"
	"nechego/statistics"

	tele "gopkg.in/telebot.v3"
)

const profile = Response(`📇 <b>%s %s</b>
<code>%s  %s  %s  %s  %s</code>

💵 Денег в кошельке: %s
💳 На счету в банке: %s

%s

%s

%s`)

// !профиль
func (a *App) handleProfile(c tele.Context) error {
	user, ok := maybeGetReplyUser(c)
	if !ok {
		user = getUser(c)
	}
	strength, err := a.stat.Strength(user)
	if err != nil {
		return respondInternalError(c, err)
	}
	modset, err := a.stat.UserModset(user)
	if err != nil {
		return respondInternalError(c, err)
	}
	response := profile.Fill(
		formatTitles(modset.Titles()...),
		a.mention(user),
		formatEnergy(user.Energy),
		formatElo(user.Elo),
		formatStrength(strength),
		formatMessages(user.Messages),
		formatFood(user.Fishes),
		formatMoney(user.Balance),
		formatMoney(user.Account),
		formatStatus(modset.Descriptions()...),
		formatIcons(modset.Icons()...),
		formatStatus(getStatus(user.UID)),
	)
	if ava, ok := loadAvatar(user.UID); ok {
		ava.Caption = string(response)
		return c.Send(ava, tele.ModeHTML)
	}
	return respond(c, response)
}

const (
	topStrong = Response(`🏋️‍♀️ <b>Самые сильные пользователи</b>
%s`)
	topWeak = Response(`🤕 <b>Самые слабые пользователи</b>
%s`)
)

// !топ сильных
func (a *App) handleTopStrong(c tele.Context) error {
	users, err := a.stat.SortedUsers(getGroup(c), a.stat.ByStrengthDesc)
	if err != nil {
		return respondInternalError(c, err)
	}
	n := clampTopNumber(len(users))
	top, err := a.topStrength(users[:n])
	if err != nil {
		return respondInternalError(c, err)
	}
	return respond(c, topStrong.Fill(top))
}

// !топ слабых
func (a *App) handleTopWeak(c tele.Context) error {
	users, err := a.stat.SortedUsers(getGroup(c), a.stat.ByStrengthAsc)
	if err != nil {
		return respondInternalError(c, err)
	}
	n := clampTopNumber(len(users))
	top, err := a.topStrength(users[:n])
	if err != nil {
		return respondInternalError(c, err)
	}
	return respond(c, topWeak.Fill(top))
}

func (a *App) topStrength(u []model.User) (string, error) {
	s := []string{}
	for _, uu := range u {
		str, err := a.stat.Strength(uu)
		if err != nil {
			return "", err
		}
		s = append(s, fmt.Sprintf("%s %s", a.mention(uu), formatStrength(str)))
	}
	return enumerate(s...), nil
}

// !сила
func (a *App) handleStrength(c tele.Context) error {
	str, err := a.stat.Strength(getUser(c))
	if err != nil {
		return respondInternalError(c, err)
	}
	return respond(c, Response("Ваша сила: %s").Fill(formatStrength(str)))
}

const topRating = Response("🏆 <b>Боевой рейтинг</b>\n%s")

func (a *App) handleTopElo(c tele.Context) error {
	users, err := a.stat.SortedUsers(getGroup(c), statistics.ByEloDesc)
	if err != nil {
		return respondInternalError(c, err)
	}
	users = users[:clampTopNumber(len(users))]
	return respond(c, topRating.Fill(a.topElo(users)))
}

func (a *App) topElo(u []model.User) string {
	s := []string{}
	for _, uu := range u {
		s = append(s, fmt.Sprintf("%s %s", a.mention(uu), formatElo(uu.Elo)))
	}
	return enumerate(s...)
}
