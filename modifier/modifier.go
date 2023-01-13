package modifier

type Mod struct {
	Emoji       string
	Multiplier  float64
	Description string
	Title       string
}

var (
	Admin         = &Mod{"👑", +0.20, "Вы ощущаете власть над остальными.", "администратор"}
	Eblan         = &Mod{"😸", -0.20, "Вы чувствуете себя оскорбленным.", "еблан"}
	MuchEnergy    = &Mod{"🍥", +0.20, "Вы хорошо поели.", ""}
	FullEnergy    = &Mod{"⚡️", +0.10, "Вы полны сил.", ""}
	NoEnergy      = &Mod{"😣", -0.25, "Вы чувствуете себя уставшим.", ""}
	TerribleLuck  = &Mod{"☠️", -0.50, "Вас преследуют неудачи.", ""}
	BadLuck       = &Mod{"", -0.10, "Вам не везет.", ""}
	GoodLuck      = &Mod{"🤞", +0.10, "Вам везет.", ""}
	ExcellentLuck = &Mod{"🍀", +0.30, "Сегодня ваш день.", ""}
	Rich          = &Mod{"🎩", +0.05, "Вы богаты.", "магнат"}
	Poor          = &Mod{"", -0.05, "Вы бедны.", ""}
	Fisher        = &Mod{"🎣", +0.05, "Вы можете рыбачить.", ""}
	Debtor        = &Mod{"💳", -0.25, "У вас есть кредит.", ""}
	Heavy         = &Mod{"🪨", -0.25, "Ваш инвентарь переполнен.", ""}
	RatingFirst   = &Mod{"🥇", +0.03, "Вы на 1-м месте в рейтинге.", ""}
	RatingSecond  = &Mod{"🥈", +0.02, "Вы на 2-м месте в рейтинге.", ""}
	RatingThird   = &Mod{"🥉", +0.01, "Вы на 3-м месте в рейтинге.", ""}
)

type Set map[*Mod]bool

func (s Set) Active(m *Mod) bool {
	return s[m]
}

func (s Set) Add(m *Mod) {
	s[m] = true
}

func (s Set) List() []*Mod {
	r := []*Mod{}
	for m := range s {
		r = append(r, m)
	}
	return r
}

func (s Set) Sum() float64 {
	r := 0.0
	for m := range s {
		r += m.Multiplier
	}
	return r
}
