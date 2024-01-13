package dishes

import "fmt"

type Dishes struct {
	Length      float32
	Height      int
	Name        string
	Weight      float32
	Colour      string
	Shape       string
	Material    string
	Volume      float32
	Diameter    float32
	Country     string
	Guarantee   int
	Quantity    int
	Fragile     bool
	Maintenance int
}

func (d Dishes) IsBrittle() {
	if d.Fragile {
		fmt.Println(d.Name, " хрупкий предмет посуды")
	} else {
		fmt.Println(d.Name, " не хрупкий предмет посуды")
	}

}
func (d Dishes) Explu() {
	if d.Maintenance <= 1 {
		fmt.Println(d.Name, "куплен(а) недавно")
	} else {
		fmt.Println(d.Name, "используется давно")
	}
}
func PrintKitchenDishes() Dishes {
	fmt.Println("Вся посуда имеющаяся в доме:")
	pan := Dishes{
		Length:      27,
		Name:        "Pan",
		Weight:      0.52,
		Colour:      "black",
		Shape:       "round",
		Material:    "aluminium",
		Volume:      3,
		Diameter:    24,
		Country:     "China",
		Guarantee:   2,
		Quantity:    2,
		Fragile:     false,
		Maintenance: 2,
	}
	pot := Dishes{
		Length:      30,
		Name:        "Pot",
		Weight:      2,
		Colour:      "silver",
		Shape:       "round",
		Material:    "steel",
		Volume:      4.5,
		Diameter:    24,
		Country:     "China",
		Guarantee:   2,
		Quantity:    2,
		Fragile:     false,
		Maintenance: 4,
	}
	soupPlate := Dishes{
		Name:        "SoupPlate",
		Weight:      0.38,
		Colour:      "white",
		Shape:       "round",
		Material:    "porcelain",
		Volume:      0.4,
		Diameter:    20,
		Quantity:    3,
		Fragile:     true,
		Maintenance: 2,
	}
	mug := Dishes{
		Name:        "Mug",
		Weight:      0.34,
		Colour:      "white",
		Material:    "porcelain",
		Volume:      0.25,
		Diameter:    12,
		Quantity:    4,
		Fragile:     true,
		Maintenance: 3,
	}
	fork := Dishes{ //вилка
		Length:   17.7,
		Weight:   20,
		Colour:   "silver",
		Material: "steel",
		Quantity: 10,
		Name:     "Fork",
	}
	spoon := Dishes{
		Length:      17.5,
		Weight:      22.2,
		Colour:      "silver",
		Material:    "steel",
		Quantity:    10,
		Name:        "Spoon",
		Fragile:     false,
		Maintenance: 5,
	}
	knife := Dishes{
		Length:      21.8,
		Weight:      68,
		Colour:      "silver",
		Material:    "steel",
		Quantity:    10,
		Name:        "Knife",
		Fragile:     false,
		Maintenance: 2,
	}
	fmt.Println("\t\t\tПосуда и кухонные приборы")
	fmt.Println("\tНазвание предмета:", pan.Name, "\nДлина предмета:", pan.Length, "\nВес предмета:", pan.Weight, "\nЦвет предмета:", pan.Colour, "\nМатериал:", pan.Material, "\nФорма предмета:", pan.Shape, "\nОбъем:", pan.Volume, "\nДиаметр предмета:", pan.Diameter, "\nСтрана производитель:", pan.Country, "\nГарантия:", pan.Guarantee, "\nКоличество:", pan.Quantity)
	Dishes.IsBrittle(pan)
	fmt.Println("\tНазвание предмета:", pot.Name, "\nДлина предмета:", pot.Length, "\nВес предмета:", pot.Weight, "\nЦвет предмета:", pot.Colour, "\nМатериал:", pot.Material, "\nФорма предмета:", pot.Shape, "\nОбъем:", pot.Volume, "\nДиаметр предмета:", pot.Diameter, "\nСтрана производитель:", pot.Country, "\nГарантия:", pot.Guarantee, "\nКоличество:", pot.Quantity)
	Dishes.Explu(pot)
	fmt.Println("\tНазвание предмета:", mug.Name, "\nВес предмета:", mug.Weight, "\nЦвет предмета:", mug.Colour, "\nМатериал:", mug.Material, "\nОбъем:", mug.Volume, "\nДиаметр предмета:", mug.Diameter, "\nКоличество:", mug.Quantity)
	Dishes.IsBrittle(mug)
	fmt.Println("\tНазвание предмета:", soupPlate.Name, "\nВес предмета:", soupPlate.Weight, "\nЦвет предмета:", soupPlate.Colour, "\nМатериал:", soupPlate.Material, "\nФорма предмета:", soupPlate.Shape, "\nОбъем:", soupPlate.Volume, "\nДиаметр предмета:", soupPlate.Diameter, "\nКоличество:", soupPlate.Quantity)
	fmt.Println("\tНазвание предмета:", fork.Name, "\nДлина предмета:", fork.Length, "\nВес предмета:", fork.Weight, "\nЦвет предмета:", fork.Colour, "\nМатериал:", fork.Material, "\nКоличество:", fork.Quantity)
	Dishes.IsBrittle(fork)
	fmt.Println("\tНазвание предмета:", spoon.Name, "\nВес предмета:", spoon.Weight, "\nЦвет предмета:", spoon.Colour, "\nМатериал:", spoon.Material, "\nДлина предмета:", spoon.Length, "\nКоличество:", spoon.Quantity)
	Dishes.Explu(spoon)
	fmt.Println("\tНазвание предмета:", knife.Name, "\nВес предмета:", knife.Weight, "\nЦвет предмета:", knife.Colour, "\nМатериал:", knife.Material, "\nДлина предмета:", knife.Length, "\nКоличество:", knife.Quantity)
	return Dishes{}
}
