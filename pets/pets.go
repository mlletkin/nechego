package pets

import (
	"fmt"
	"math/rand"
	"nechego/modifier"
	"nechego/valid"
	"strings"
	"time"
)

type Gender int

const (
	Male Gender = iota
	Female
)

func (g Gender) Emoji() string  { return genders[g].Emoji }
func (g Gender) String() string { return genders[g].Description }

type Pet struct {
	Name    string
	Species Species
	Gender  Gender
	Birth   time.Time
}

func Random() *Pet {
	return &Pet{
		Species: randomSpecies(),
		Gender:  Gender(rand.Intn(2)),
		Birth:   time.Now(),
	}
}

func (p *Pet) Age() time.Duration {
	return time.Since(p.Birth)
}

func (p *Pet) String() string {
	name := p.Name
	if name != "" {
		name = name + " "
	}
	return fmt.Sprintf("%s %s %s(%s)", p.Species.Emoji(),
		strings.Title(p.Species.String()), name, p.Gender.Emoji())
}

func (p *Pet) SetName(s string) bool {
	if !valid.Name(s) {
		return false
	}
	p.Name = strings.Title(s)
	return true
}

func (p *Pet) Mod() (m *modifier.Mod, ok bool) {
	if p == nil {
		return nil, false
	}
	var multiplier float64
	switch p.Species.Quality() {
	case Common:
		multiplier = 0.05
	case Rare:
		multiplier = 0.10
	case Exotic:
		multiplier = 0.15
	case Legendary:
		multiplier = 0.20
	}
	prefix := ""
	if p.Species.Quality() != Common {
		prefix = fmt.Sprintf("%s ", p.Species.Quality())
	}
	return &modifier.Mod{
		Emoji:       "🐱",
		Multiplier:  multiplier,
		Description: fmt.Sprintf("У вас есть %sпитомец: <code>%s</code>", prefix, p),
	}, true
}

func (p *Pet) Nutrition() float64 {
	switch p.Species.Size() {
	case Small:
		return 0.06
	case Medium:
		return 0.12
	case Big:
		return 0.18
	}
	panic(fmt.Errorf("unexpected pet size %v", p.Species.Size()))
}
