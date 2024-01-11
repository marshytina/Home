package decor

import "fmt"

type Decor struct {
	Name     string
	Length   int
	Height   int
	Width    int
	Depth    int
	Weight   float32
	Colour   string
	Material string
	Shape    string
	Diameter int
}

func PrintHallDecor() Decor {
	floorLamp := Decor{
		Name:     "FloorLamp",
		Height:   170,
		Width:    84,
		Depth:    84,
		Weight:   6.4,
		Colour:   "transparent",
		Material: "glass",
	}
	carpet := Decor{
		Name:     "Carpet",
		Length:   100,
		Width:    50,
		Depth:    0,
		Weight:   6.4,
		Colour:   "grey",
		Material: "pile",
	}
	bookcase := Decor{
		Name:     "Bookcase",
		Length:   104,
		Width:    70,
		Depth:    31,
		Weight:   21.74,
		Colour:   "white",
		Material: "wood",
	}
	poof := Decor{
		Name:     "Poof",
		Length:   46,
		Width:    25,
		Weight:   4.9,
		Colour:   "biege",
		Diameter: 40,
		Shape:    "circle",
	}
	pillow := Decor{
		Name:     "Pillow",
		Length:   35,
		Width:    35,
		Weight:   0.129,
		Colour:   "silver",
		Material: "polyester",
		Shape:    "square",
	}
	fmt.Println("Название предмета:", pillow.Name, "Длина предмета:", pillow.Length, "Ширина предмета:", pillow.Width, "Вес предмета:", pillow.Weight, "Цвет предмета:", pillow.Colour, "Материал:", pillow.Material, "Форма предмета:", pillow.Shape)
	fmt.Println("Название предмета:", poof.Name, "Длина предмета:", poof.Length, "Ширина предмета:", poof.Width, "Вес предмета:", poof.Weight, "Цвет предмета:", poof.Colour, "Материал:", poof.Material, "Форма предмета:", poof.Shape)
	fmt.Println("Название предмета:", floorLamp.Name, "Длина предмета:", floorLamp.Length, "Ширина предмета:", floorLamp.Width, "Вес предмета:", floorLamp.Weight, "Цвет предмета:", floorLamp.Colour, "Материал:", floorLamp.Material, "Форма предмета:", floorLamp.Shape)
	fmt.Println("Название предмета:", carpet.Name, "Длина предмета:", carpet.Length, "Ширина предмета:", carpet.Width, "Вес предмета:", carpet.Weight, "Цвет предмета:", carpet.Colour, "Материал:", carpet.Material, "Форма предмета:", carpet.Shape)
	fmt.Println("Название предмета:", bookcase.Name, "Длина предмета:", bookcase.Length, "Ширина предмета:", bookcase.Width, "Вес предмета:", bookcase.Weight, "Цвет предмета:", bookcase.Colour, "Материал:", bookcase.Material, "Форма предмета:", bookcase.Shape)
	return Decor{}
}
