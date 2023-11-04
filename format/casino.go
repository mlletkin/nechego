package format

import (
	"fmt"
	"nechego/game"
	"time"
)

const GameGoing = "🎲 Игра уже идет."

func DiceGame(who string, bet int, timeout time.Duration) string {
	sec := int(timeout / time.Second)
	c := NewConnector("\n")
	c.Add(fmt.Sprintf("🎲 %s играет на %s", Name(who), Money(bet)))
	c.Add(fmt.Sprintf("У вас <code>%d секунд</code> на то, чтобы бросить кости!", sec))
	return c.String()
}

func DiceGameResult(r game.DiceGameResult) string {
	if r.Outcome == game.Win {
		return fmt.Sprintf("💥 Вы выиграли %s", Money(r.Prize))
	}
	if r.Outcome == game.Lose {
		return "😵 Вы проиграли."
	}
	return "🎲 Ничья."
}

func DiceTimeout(bet int) string {
	return fmt.Sprintf("<i>⏰ Время вышло: вы потеряли %s</i>", Money(bet))
}

func MinBet(n int) string {
	return fmt.Sprintf("💵 Минимальная ставка %s", Money(n))
}

func SlotWin(who string, prize int) string {
	return fmt.Sprintf("🎰 %s выигрывает %s 💥", Name(who), Money(prize))
}

func SlotRoll(who string, bet int) string {
	return fmt.Sprintf("🎰 %s прокручивает слоты на %s", Name(who), Money(bet))
}

func BetSet(who string, n int) string {
	return fmt.Sprintf("🎰 %s устанавливает ставку %s", Name(who), Money(n))
}
