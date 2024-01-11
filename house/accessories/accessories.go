package accessories

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

func PrintBathroomAccessories() Accessory {
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
	fmt.Println("\t\t\tАксессуары для ванной комнаты")
	fmt.Println("\tНазвание предмета:", brush.Name, "\nВысота предмета:", brush.Height, "\nВес предмета:", brush.Weight, "\nЦвет предмета:", brush.Colour, "\nМатериал:", brush.Material)
	fmt.Println("\tНазвание предмета:", soapbox.Name, "\nВысота предмета:", soapbox.Height, "\nВес предмета:", soapbox.Weight, "\nЦвет предмета:", soapbox.Colour, "\nМатериал:", soapbox.Material)
	fmt.Println("\tНазвание предмета:", organizer.Name, "\nВысота предмета:", organizer.Height, "\nВес предмета:", organizer.Weight, "\nШирина предмета:", organizer.Width, "\nЦвет предмета:", organizer.Colour, "\nМатериал:", organizer.Material)
	fmt.Println("\tНазвание предмета:", shelf.Name, "\nВысота предмета:", shelf.Height, "\nВес предмета:", shelf.Weight, "\nЦвет предмета:", shelf.Colour, "\nМатериал:", shelf.Material)
	return Accessory{}
}
