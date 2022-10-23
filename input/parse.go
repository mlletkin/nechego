package input

import (
	"regexp"
	"strings"
)

// TODO: Parseable is a string, []string, or *regexp.Regexp.
type Parseable interface{}

var (
	eblanRe    = regexp.MustCompile("(?i)^![ие][б6п*]?л[ап]н[А-я]* дня")
	masyunyaRe = regexp.MustCompile("(?i)^(!ма[нс]ю[нс][а-я]*[пая])")
	helloRe    = regexp.MustCompile(
		constructHelloRe("п[рл]ив[а-я]*", "хай", "зд[ао]ров[а-я]*", "ку", "здрав[а-я]*"))
	weatherRe     = regexp.MustCompile("(?i)^!погода ([-А-я]+)")
	probabilityRe = regexp.MustCompile("(?i)^!инфа *(.*)")
	whoRe         = regexp.MustCompile("(?i)^!кто *(.*)")
	listRe        = regexp.MustCompile("(?i)^!список *(.*)")
	topRe         = regexp.MustCompile("(?i)^!топ[- ]*(\\d*) *(.*)")
)

// ParseCommand returns a command corresponding to the input string.
func ParseCommand(s string) Command {
	switch {
	case probabilityRe.MatchString(s):
		return CommandProbability
	case whoRe.MatchString(s):
		return CommandWho
	case startsWith(s, "!имя"):
		return CommandTitle
	case startsWith(s, "!аним", "!мульт"):
		return CommandAnime
	case startsWith(s, "!фур"):
		return CommandFurry
	case startsWith(s, "!флаг"):
		return CommandFlag
	case startsWith(s, "!чел"):
		return CommandPerson
	case startsWith(s, "!лошадь", "!конь"):
		return CommandHorse
	case startsWith(s, "!арт"):
		return CommandArt
	case startsWith(s, "!авто", "!тачк", "!машин"):
		return CommandCar
	case startsWith(s, "!пара дня"):
		return CommandPair
	case eblanRe.MatchString(s):
		return CommandEblan
	case startsWith(s, "!админ дня"):
		return CommandAdmin
	case startsWith(s, "!драка", "!дуэль", "!поединок", "!бой", "!сражение", "!борьба", "!атака", "!битва", "!схватка"):
		return CommandFight
	case startsWith(s, "!баланс", "!деньги"):
		return CommandBalance
	case startsWith(s, "!перевод"):
		return CommandTransfer
	case startsWith(s, "!профиль", "!статистика", "!обо мне"):
		return CommandProfile
	case startsWith(s, "!топ богат", "!топ богач"):
		return CommandTopRich
	case startsWith(s, "!топ нищих", "!топ бедн"):
		return CommandTopPoor
	case startsWith(s, "!капитал"):
		return CommandCapital
	case startsWith(s, "!энерг", "!стамина"):
		return CommandEnergy
	case startsWith(s, "!рейтинг", "!ммр", "!эло"):
		return CommandRating
	case startsWith(s, "!удочка"):
		return CommandFishingRod
	case startsWith(s, "!рыбалка", "!ловля рыб", "!закинуть удочку"):
		return CommandFishing
	case startsWith(s, "!топ сил"):
		return CommandTopStrong
	case startsWith(s, "!поесть", "!еда"):
		return CommandEatFish
	case startsWith(s, "!депозит"):
		return CommandDeposit
	case startsWith(s, "!обнал", "!снять"):
		return CommandWithdraw
	case startsWith(s, "!банк"):
		return CommandBank
	case startsWith(s, "!кредит"):
		return CommandDebt
	case startsWith(s, "!погасить"):
		return CommandRepay
	case startsWith(s, "!топ слабых"):
		return CommandTopWeak
	case startsWith(s, "!парламент"):
		return CommandParliament
	case startsWith(s, "!импичмент"):
		return CommandImpeachment
	case startsWith(s, "!рыба", "!улов"):
		return CommandFishList
	case startsWith(s, "!замороз"):
		return CommandFreezeFish
	case startsWith(s, "!холодильн"):
		return CommandFreezer
	case startsWith(s, "!размороз"):
		return CommandUnfreezeFish
	case startsWith(s, "!продажа", "!продать"):
		return CommandSellFish
	case startsWith(s, "!сила"):
		return CommandStrength
	case masyunyaRe.MatchString(s) || startsWith(s, "Масюня 🎀"):
		return CommandMasyunya
	case startsWith(s, "!паппи", "Паппи 🦊"):
		return CommandPoppy
	case startsWith(s, "!сима", "Сима 💖"):
		return CommandSima
	case helloRe.MatchString(s):
		return CommandHello
	case startsWith(s, "!мыш"):
		return CommandMouse
	case weatherRe.MatchString(s):
		return CommandWeather
	case startsWith(s, "!тикток"):
		return CommandTikTok
	case listRe.MatchString(s):
		return CommandList
	case topRe.MatchString(s):
		return CommandTop
	case startsWith(s, "!кот василия", "!кошка василия", "!марс", "!муся"):
		return CommandBasili
	case startsWith(s, "!касп", "!кот касп"):
		return CommandCasper
	case startsWith(s, "!зевс"):
		return CommandZeus
	case startsWith(s, "!кот", "!кош"):
		return CommandCat
	case startsWith(s, "!пик"):
		return CommandPic
	case startsWith(s, "!кости"):
		return CommandDice
	case startsWith(s, "!игр"):
		return CommandGame
	case startsWith(s, "!клав", "!открыт"):
		return CommandKeyboardOpen
	case startsWith(s, "!закрыт", "!скрыт"):
		return CommandKeyboardClose
	case startsWith(s, "!вкл", "!подкл", "!подруб"):
		return CommandTurnOn
	case startsWith(s, "!выкл", "!откл"):
		return CommandTurnOff
	case startsWith(s, "!бан"):
		return CommandBan
	case startsWith(s, "!разбан"):
		return CommandUnban
	case startsWith(s, "!инфо"):
		return CommandInfo
	case startsWith(s, "!помощь", "!команды"):
		return CommandHelp
	case startsWith(s, "!запретить"):
		return CommandForbid
	case startsWith(s, "!разрешить"):
		return CommandPermit
	case startsWith(s, "!питомец"):
		return CommandPet
	case startsWith(s, "!взять"):
		return CommandBuyPet
	case startsWith(s, "!назвать"):
		return CommandNamePet
	case startsWith(s, "!выкинуть", "!прогнать", "!выбросить", "!выгнать"):
		return CommandDropPet
	}
	return CommandUnknown
}

// startsWith returns true if the input string starts with one of the specified prefixes, false otherwise.
func startsWith(s string, prefix ...string) bool {
	s = strings.ToLower(s)
	for _, p := range prefix {
		p = strings.ToLower(p)
		if strings.HasPrefix(s, p) {
			return true
		}
	}
	return false
}

const (
	helloPrefix = "((^|[^а-я])"
	helloSuffix = "([^а-я]|$))"
)

// constructHelloRe combines the given hello regexps.
func constructHelloRe(hello ...string) string {
	var l []string
	for _, h := range hello {
		l = append(l, helloPrefix+h+helloSuffix)
	}
	return "(?i)" + strings.Join(l, "|")
}
