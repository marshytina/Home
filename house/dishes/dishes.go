package deshes

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
	fmt.Println("Название предмета:", pan.Name, "Длина предмета:", pan.Length, "Вес предмета:", pan.Weight, "Цвет предмета:", pan.Colour, "Материал:", pan.Material, "Форма предмета:", pan.Shape, "Объем:", pan.Volume, "Диаметр предмета:", pan.Diameter, "Страна производитель:", pan.Country, "Гарантия:", pan.Guarantee, "Количество:", pan.Quantity)
	fmt.Println("Название предмета:", pot.Name, "Длина предмета:", pot.Length, "Вес предмета:", pot.Weight, "Цвет предмета:", pot.Colour, "Материал:", pot.Material, "Форма предмета:", pot.Shape, "Объем:", pot.Volume, "Диаметр предмета:", pot.Diameter, "Страна производитель:", pot.Country, "Гарантия:", pot.Guarantee, "Количество:", pot.Quantity)
	fmt.Println("Название предмета:", flatPlate.Name, "Вес предмета:", flatPlate.Weight, "Цвет предмета:", flatPlate.Colour, "Материал:", flatPlate.Material, "Форма предмета:", flatPlate.Shape, "Диаметр предмета:", flatPlate.Diameter, "Количество:", pot.Quantity)
	fmt.Println("Название предмета:", soupPlate.Name, "Вес предмета:", soupPlate.Weight, "Цвет предмета:", soupPlate.Colour, "Материал:", soupPlate.Material, "Форма предмета:", soupPlate.Shape, "Объем:", soupPlate.Volume, "Диаметр предмета:", soupPlate.Diameter, "Количество:", soupPlate.Quantity)
	fmt.Println("Название предмета:", mug.Name, "Вес предмета:", mug.Weight, "Цвет предмета:", mug.Colour, "Материал:", mug.Material, "Объем:", mug.Volume, "Диаметр предмета:", mug.Diameter, "Количество:", mug.Quantity)
	fmt.Println("Название предмета:", soupPlate.Name, "Вес предмета:", soupPlate.Weight, "Цвет предмета:", soupPlate.Colour, "Материал:", soupPlate.Material, "Форма предмета:", soupPlate.Shape, "Объем:", soupPlate.Volume, "Диаметр предмета:", soupPlate.Diameter, "Количество:", soupPlate.Quantity)
	fmt.Println("Название предмета:", wineGlass.Name, "Вес предмета:", wineGlass.Weight, "Цвет предмета:", wineGlass.Colour, "Материал:", wineGlass.Material, "Форма предмета:", wineGlass.Shape, "Объем:", wineGlass.Volume, "Диаметр предмета:", wineGlass.Diameter, "Количество:", wineGlass.Quantity)
	fmt.Println("Название предмета:", fork.Name, "Длина предмета:", fork.Length, "Вес предмета:", fork.Weight, "Цвет предмета:", fork.Colour, "Материал:", fork.Material, "Количество:", fork.Quantity)
	fmt.Println("Название предмета:", spoon.Name, "Вес предмета:", spoon.Weight, "Цвет предмета:", spoon.Colour, "Материал:", spoon.Material, "Длина предмета:", spoon.Length, "Количество:", spoon.Quantity)
	fmt.Println("Название предмета:", dessertSpoon.Name, "Вес предмета:", dessertSpoon.Weight, "Цвет предмета:", dessertSpoon.Colour, "Материал:", dessertSpoon.Material, "Длина предмета:", dessertSpoon.Length, "Объем:", "Количество:", dessertSpoon.Quantity)
	fmt.Println("Название предмета:", knife.Name, "Вес предмета:", knife.Weight, "Цвет предмета:", knife.Colour, "Материал:", knife.Material, "Длина предмета:", knife.Length, "Количество:", knife.Quantity)
	fmt.Println("Название предмета:", spatula.Name, "Вес предмета:", spatula.Weight, "Цвет предмета:", spatula.Colour, "Материал:", spatula.Material, "Длина предмета:", spatula.Length, "Объем:", "Количество:", spatula.Quantity)
	fmt.Println("Название предмета:", cuttingBoard.Name, "Вес предмета:", cuttingBoard.Weight, "Цвет предмета:", cuttingBoard.Colour, "Материал:", cuttingBoard.Material, "Длина предмета:", cuttingBoard.Length, "Количество:", cuttingBoard.Quantity)
	return Dishes{}
}
