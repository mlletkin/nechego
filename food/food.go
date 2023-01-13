package food

import "math/rand"

type Type int

const (
	Bread Type = iota
	ChickenLeg
	BigTasty
	BigMac
	Fries
	PizzaFourCheese
	PizzaPepperoni
	PizzaCheeseChicken
	Toast
	Shawarma

	typeN
)

func (t Type) String() string {
	return data[t].Description
}

func (t Type) Emoji() string {
	return data[t].Emoji
}

func (t Type) Nutrition() float64 {
	return data[t].Nutrition
}

func (t Type) Description() string {
	return data[t].Description
}

var data = map[Type]struct {
	Emoji       string
	Nutrition   float64
	Description string
}{
	Bread:              {"🍞", 0.08, "Хлеб"},
	ChickenLeg:         {"🍗", 0.12, "Куриная ножка"},
	BigTasty:           {"🍔", 0.16, "Биг Тейсти"},
	BigMac:             {"🍔", 0.14, "Биг Мак"},
	Fries:              {"🍟", 0.08, "Картофель фри"},
	PizzaFourCheese:    {"🍕", 0.16, "Пицца (4 сыра)"},
	PizzaPepperoni:     {"🍕", 0.16, "Пицца (пеперони)"},
	PizzaCheeseChicken: {"🍕", 0.16, "Пицца (сырный цыпленок)"},
	Toast:              {"🥪", 0.10, "Бутерброд"},
	Shawarma:           {"🌯", 0.16, "Шаурма"},
}

type Food struct {
	Type
}

func (f Food) String() string {
	return f.Type.Emoji() + " " + f.Type.Description()
}

func Random() *Food {
	return &Food{Type: Type(rand.Intn(int(typeN)))}
}