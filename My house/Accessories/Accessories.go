package main

import (
	"fmt"
)

type Accessory struct {
	Name     string
	Length   int
	Height   int
	Width    int
	Depth    int
	Weight   float32
	Colour   string
	Material string
	Shape    string
}

func PrintAccessories() Accessory {
	fmt.Println("Все аксессуары имеющиеся в доме:")
	brush := Accessory{
		Name:     "Brush",
		Height:   36,
		Weight:   0.09,
		Colour:   "white",
		Material: "steel",
	}
	soapbox := Accessory{
		Name:     "SoapBox",
		Height:   2,
		Weight:   0.144,
		Colour:   "black",
		Material: "polyresin",
	}
	organizer := Accessory{
		Name:     "Organizer",
		Height:   8,
		Width:    22,
		Weight:   0.3,
		Colour:   "transparent",
		Material: "acrylic",
	}
	shelf := Accessory{
		Name:     "Shelf",
		Height:   58,
		Weight:   1.405,
		Colour:   "silver",
		Material: "steel",
	}
	fmt.Println("Название предмета:", brush.Name, "Высота предмета:", brush.Height, "Вес предмета:", brush.Weight, "Цвет предмета:", brush.Colour, "Материал:", brush.Material)
	fmt.Println("Название предмета:", soapbox.Name, "Высота предмета:", soapbox.Height, "Вес предмета:", soapbox.Weight, "Цвет предмета:", soapbox.Colour, "Материал:", soapbox.Material)
	fmt.Println("Название предмета:", organizer.Name, "Высота предмета:", organizer.Height, "Вес предмета:", "Ширина предмета:", organizer.Width, organizer.Weight, "Цвет предмета:", organizer.Colour, "Материал:", organizer.Material)
	fmt.Println("Название предмета:", shelf.Name, "Высота предмета:", shelf.Height, "Вес предмета:", shelf.Weight, "Цвет предмета:", shelf.Colour, "Материал:", shelf.Material)
	return Accessory{}
}
