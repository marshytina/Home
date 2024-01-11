package dishes

import "fmt"

type Dishes struct {
	Length    float32
	Height    int
	Name      string
	Weight    float32
	Colour    string
	Shape     string
	Material  string
	Volume    float32
	Diameter  float32
	Country   string
	Guarantee int
	Quantity  int
}

func PrintKitchenDishes() Dishes {
	fmt.Println("Вся посуда имеющаяся в доме:")
	pan := Dishes{
		Length:    27,
		Name:      "Pan",
		Weight:    0.52,
		Colour:    "black",
		Shape:     "round",
		Material:  "aluminium",
		Volume:    3,
		Diameter:  24,
		Country:   "China",
		Guarantee: 2,
		Quantity:  2,
	}
	pot := Dishes{
		Length:    30,
		Name:      "Pot",
		Weight:    2,
		Colour:    "silver",
		Shape:     "round",
		Material:  "steel",
		Volume:    4.5,
		Diameter:  24,
		Country:   "China",
		Guarantee: 2,
		Quantity:  2,
	}
	flatPlate := Dishes{
		Name:     "FlatPlate",
		Weight:   0.35,
		Colour:   "black",
		Shape:    "round",
		Material: "porcelain",
		Diameter: 20,
		Quantity: 5,
	}
	soupPlate := Dishes{
		Name:     "SoupPlate",
		Weight:   0.38,
		Colour:   "white",
		Shape:    "round",
		Material: "porcelain",
		Volume:   0.4,
		Diameter: 20,
		Quantity: 3,
	}
	mug := Dishes{
		Name:     "Mug",
		Weight:   0.34,
		Colour:   "white",
		Material: "porcelain",
		Volume:   0.25,
		Diameter: 12,
		Quantity: 4,
	}
	wineGlass := Dishes{
		Name:     "Wineglass",
		Weight:   0.38,
		Colour:   "transparent",
		Shape:    "round",
		Material: "glass",
		Volume:   0.4,
		Diameter: 9,
		Quantity: 4,
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
		Length:   17.5,
		Weight:   22.2,
		Colour:   "silver",
		Material: "steel",
		Quantity: 10,
		Name:     "Spoon",
	}
	dessertSpoon := Dishes{
		Length:   13.5,
		Weight:   27,
		Colour:   "silver",
		Material: "steel",
		Quantity: 5,
		Name:     "DessertSpoon",
	}
	knife := Dishes{
		Length:   21.8,
		Weight:   68,
		Colour:   "silver",
		Material: "steel",
		Quantity: 10,
		Name:     "Knife",
	}
	spatula := Dishes{
		Length:   25,
		Weight:   15,
		Colour:   "brown",
		Material: "wood",
		Quantity: 1,
		Name:     "Spatula",
	}
	cuttingBoard := Dishes{
		Length:   36.8,
		Weight:   20,
		Colour:   "beige",
		Material: "bamboo",
		Quantity: 3,
		Name:     "CuttingBoard",
	}
	fmt.Println("\t\t\tПосуда и кухонные приборы")
	fmt.Println("\tНазвание предмета:", pan.Name, "\nДлина предмета:", pan.Length, "\nВес предмета:", pan.Weight, "\nЦвет предмета:", pan.Colour, "\nМатериал:", pan.Material, "\nФорма предмета:", pan.Shape, "\nОбъем:", pan.Volume, "\nДиаметр предмета:", pan.Diameter, "\nСтрана производитель:", pan.Country, "\nГарантия:", pan.Guarantee, "\nКоличество:", pan.Quantity)
	fmt.Println("\tНазвание предмета:", pot.Name, "\nДлина предмета:", pot.Length, "\nВес предмета:", pot.Weight, "\nЦвет предмета:", pot.Colour, "\nМатериал:", pot.Material, "\nФорма предмета:", pot.Shape, "\nОбъем:", pot.Volume, "\nДиаметр предмета:", pot.Diameter, "\nСтрана производитель:", pot.Country, "\nГарантия:", pot.Guarantee, "\nКоличество:", pot.Quantity)
	fmt.Println("\tНазвание предмета:", flatPlate.Name, "\nВес предмета:", flatPlate.Weight, "\nЦвет предмета:", flatPlate.Colour, "\nМатериал:", flatPlate.Material, "\nФорма предмета:", flatPlate.Shape, "\nДиаметр предмета:", flatPlate.Diameter, "\nКоличество:", pot.Quantity)
	fmt.Println("\tНазвание предмета:", soupPlate.Name, "\nВес предмета:", soupPlate.Weight, "\nЦвет предмета:", soupPlate.Colour, "\nМатериал:", soupPlate.Material, "\nФорма предмета:", soupPlate.Shape, "\nОбъем:", soupPlate.Volume, "\nДиаметр предмета:", soupPlate.Diameter, "\nКоличество:", soupPlate.Quantity)
	fmt.Println("\tНазвание предмета:", mug.Name, "\nВес предмета:", mug.Weight, "\nЦвет предмета:", mug.Colour, "\nМатериал:", mug.Material, "\nОбъем:", mug.Volume, "\nДиаметр предмета:", mug.Diameter, "\nКоличество:", mug.Quantity)
	fmt.Println("\tНазвание предмета:", soupPlate.Name, "\nВес предмета:", soupPlate.Weight, "\nЦвет предмета:", soupPlate.Colour, "\nМатериал:", soupPlate.Material, "\nФорма предмета:", soupPlate.Shape, "\nОбъем:", soupPlate.Volume, "\nДиаметр предмета:", soupPlate.Diameter, "\nКоличество:", soupPlate.Quantity)
	fmt.Println("\tНазвание предмета:", wineGlass.Name, "\nВес предмета:", wineGlass.Weight, "\nЦвет предмета:", wineGlass.Colour, "\nМатериал:", wineGlass.Material, "\nФорма предмета:", wineGlass.Shape, "\nОбъем:", wineGlass.Volume, "\nДиаметр предмета:", wineGlass.Diameter, "\nКоличество:", wineGlass.Quantity)
	fmt.Println("\tНазвание предмета:", fork.Name, "\nДлина предмета:", fork.Length, "\nВес предмета:", fork.Weight, "\nЦвет предмета:", fork.Colour, "\nМатериал:", fork.Material, "\nКоличество:", fork.Quantity)
	fmt.Println("\tНазвание предмета:", spoon.Name, "\nВес предмета:", spoon.Weight, "\nЦвет предмета:", spoon.Colour, "\nМатериал:", spoon.Material, "\nДлина предмета:", spoon.Length, "\nКоличество:", spoon.Quantity)
	fmt.Println("\tНазвание предмета:", dessertSpoon.Name, "\nВес предмета:", dessertSpoon.Weight, "\nЦвет предмета:", dessertSpoon.Colour, "\nМатериал:", dessertSpoon.Material, "\nДлина предмета:", dessertSpoon.Length, "\nКоличество:", dessertSpoon.Quantity)
	fmt.Println("\tНазвание предмета:", knife.Name, "\nВес предмета:", knife.Weight, "\nЦвет предмета:", knife.Colour, "\nМатериал:", knife.Material, "\nДлина предмета:", knife.Length, "\nКоличество:", knife.Quantity)
	fmt.Println("\tНазвание предмета:", spatula.Name, "\nВес предмета:", spatula.Weight, "\nЦвет предмета:", spatula.Colour, "\nМатериал:", spatula.Material, "\nДлина предмета:", spatula.Length, "\nКоличество:", spatula.Quantity)
	fmt.Println("\tНазвание предмета:", cuttingBoard.Name, "\nВес предмета:", cuttingBoard.Weight, "\nЦвет предмета:", cuttingBoard.Colour, "\nМатериал:", cuttingBoard.Material, "\nДлина предмета:", cuttingBoard.Length, "\nКоличество:", cuttingBoard.Quantity)
	return Dishes{}
}
