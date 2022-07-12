package app

import (
	"errors"
	"fmt"
	"nechego/input"
	"nechego/model"

	tele "gopkg.in/telebot.v3"
)

const bank = "🏦 *Банк:* на вашем счете %s\n\n" +
	"_Снять средства: `!обнал`\\._\n" +
	"_Пополнить счет: `!депозит`\\._\n\n" +
	"_%s_\n\n" +
	"_Взять кредит: `!кредит`\\._\n" +
	"_Погасить кредит: `!погасить`\\._\n" +
	"_Процентная ставка: %s_\n" +
	"_Кредитный лимит: %s_\n" +
	"_Комиссия за пополнение: %s_\n"

func (a *App) handleBank(c tele.Context) error {
	user := getUser(c)
	return c.Send(fmt.Sprintf(bank,
		formatMoney(user.Account),
		debtStatus(user),
		formatRatio(debtFee),
		formatMoney(user.DebtLimit),
		formatMoney(bankFee)),
		tele.ModeMarkdownV2)
}

func debtStatus(u model.User) string {
	if !u.Debtor() {
		return "У вас нет кредитов\\."
	}
	return fmt.Sprintf("Вы должны банку %s", formatMoney(u.Debt))
}

const deposit = "💳 Вы оплатили налог и положили %s в банк\\.\n\n_Теперь на счету %s_"

// !депозит
func (a *App) handleDeposit(c tele.Context) error {
	user := getUser(c)
	amount, err := getMessage(c).MoneyArgument()
	if errors.Is(err, input.ErrAllIn) {
		amount = user.Balance
	} else if err != nil {
		return userError(c, specifyAmount)
	}
	amount = amountAfterBankFee(amount)
	if ok := a.model.Deposit(user, amount, bankFee); !ok {
		return userError(c, notEnoughMoney)
	}
	out := fmt.Sprintf(deposit, formatMoney(amount), formatMoney(user.Account+amount))
	return c.Send(out, tele.ModeMarkdownV2)
}

func amountAfterBankFee(amount int) int {
	return amount - bankFee
}

const (
	withdraw       = "💳 Вы сняли %s со счета\\.\n\n_Теперь в кошельке %s_"
	withdrawDebtor = "Вы не можете снимать средства со счета, пока у вас есть непогашенные кредиты.\n"
)

// !обнал
func (a *App) handleWithdraw(c tele.Context) error {
	user := getUser(c)
	if user.Debtor() {
		return userError(c, withdrawDebtor)
	}
	amount, err := getMessage(c).MoneyArgument()
	if errors.Is(err, input.ErrAllIn) {
		amount = user.Account
	} else if err != nil {
		return userError(c, specifyAmount)
	}
	if ok := a.model.Withdraw(user, amount, 0); !ok {
		return userError(c, notEnoughMoney)
	}
	out := fmt.Sprintf(withdraw, formatMoney(amount), formatMoney(user.Balance+amount))
	return c.Send(out, tele.ModeMarkdownV2)
}

const (
	debtorCannotLoan = "Вы не можете взять средства в долг, пока у вас есть непогашенные кредиты."
	debtTooLow       = "Минимальный кредит — %s"
	limitTooLow      = "Ваш кредитный лимит — %s"
	debtSuccess      = "💳 Вы взяли в кредит %s\n\n_Вам необходимо вернуть %s_"
)

// !долг, !кредит
func (a *App) handleDebt(c tele.Context) error {
	user := getUser(c)
	if user.Debtor() {
		return userError(c, debtorCannotLoan)
	}
	amount, err := getMessage(c).MoneyArgument()
	if errors.Is(err, input.ErrAllIn) {
		amount = user.DebtLimit
	} else if err != nil {
		return userError(c, specifyAmount)
	}
	if amount < minDebt {
		return userErrorMarkdown(c, fmt.Sprintf(debtTooLow, formatMoney(minDebt)))
	}
	fee := calculateDebtFee(amount)
	if ok := a.model.Loan(user, amount, fee); !ok {
		return userErrorMarkdown(c, fmt.Sprintf(limitTooLow, formatMoney(user.DebtLimit)))
	}
	out := fmt.Sprintf(debtSuccess, formatMoney(amount), formatMoney(amount+fee))
	return c.Send(out, tele.ModeMarkdownV2)
}

func calculateDebtFee(amount int) int {
	return int(float64(amount) * debtFee)
}

const (
	notDebtor              = "У вас нет непогашенных кредитов."
	notEnoughOnBankAccount = "Недостаточно средств на банковском счете."
	repayFullSuccess       = "💳 Вы погасили кредит."
	repayPartialSuccess    = "💳 Вы погасили %s\n\n_Осталось погасить: %s_"
)

func (a *App) handleRepay(c tele.Context) error {
	user := getUser(c)
	if !user.Debtor() {
		return userError(c, notDebtor)
	}
	amount, err := getMessage(c).MoneyArgument()
	if errors.Is(err, input.ErrAllIn) {
		amount = user.Account
	} else if err != nil {
		return userError(c, specifyAmount)
	}
	if user.Debt <= amount {
		amount = user.Debt
	}
	if ok := a.model.Repay(user, amount); !ok {
		return userError(c, notEnoughOnBankAccount)
	}
	if amount == user.Debt {
		return c.Send(repayFullSuccess)
	}
	out := fmt.Sprintf(repayPartialSuccess, formatMoney(amount), formatMoney(user.Debt-amount))
	return c.Send(out, tele.ModeMarkdownV2)
}
