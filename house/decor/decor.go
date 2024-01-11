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
	fmt.Println("\t\t\tДекор в основной комнате")
	fmt.Println("\tНазвание предмета:", pillow.Name, "\nДлина предмета:", pillow.Length, "\nШирина предмета:", pillow.Width, "\nВес предмета:", pillow.Weight, "\nЦвет предмета:", pillow.Colour, "\nМатериал:", pillow.Material, "\nФорма предмета:", pillow.Shape)
	fmt.Println("\tНазвание предмета:", poof.Name, "\nДлина предмета:", poof.Length, "\nШирина предмета:", poof.Width, "\nВес предмета:", poof.Weight, "\nЦвет предмета:", poof.Colour, "\nМатериал:", poof.Material, "\nФорма предмета:", poof.Shape)
	fmt.Println("\tНазвание предмета:", floorLamp.Name, "\nДлина предмета:", floorLamp.Length, "\nШирина предмета:", floorLamp.Width, "\nВес предмета:", floorLamp.Weight, "\nЦвет предмета:", floorLamp.Colour, "\nМатериал:", floorLamp.Material, "\nФорма предмета:", floorLamp.Shape)
	fmt.Println("\tНазвание предмета:", carpet.Name, "\nДлина предмета:", carpet.Length, "\nШирина предмета:", carpet.Width, "\nВес предмета:", carpet.Weight, "\nЦвет предмета:", carpet.Colour, "\nМатериал:", carpet.Material, "\nФорма предмета:", carpet.Shape)
	fmt.Println("\tНазвание предмета:", bookcase.Name, "\nДлина предмета:", bookcase.Length, "\nШирина предмета:", bookcase.Width, "\nВес предмета:", bookcase.Weight, "\nЦвет предмета:", bookcase.Colour, "\nМатериал:", bookcase.Material, "\nФорма предмета:", bookcase.Shape)
	return Decor{}
}
