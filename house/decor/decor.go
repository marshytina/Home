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
	Style    bool
}

func (d Decor) StyleMatch() {
	if d.Style {
		fmt.Println(d.Name, "подходит к общему стилю дома")
	} else {
		fmt.Println(d.Name, "не подходит к общему стилю дома")
	}
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
		Style:    true,
	}
	carpet := Decor{
		Name:     "Carpet",
		Length:   100,
		Width:    50,
		Depth:    0,
		Weight:   6.4,
		Colour:   "grey",
		Material: "pile",
		Style:    true,
	}
	bookcase := Decor{
		Name:     "Bookcase",
		Length:   104,
		Width:    70,
		Depth:    31,
		Weight:   21.74,
		Colour:   "white",
		Material: "wood",
		Style:    true,
	}
	poof := Decor{
		Name:     "Poof",
		Length:   46,
		Width:    25,
		Weight:   4.9,
		Colour:   "biege",
		Diameter: 40,
		Shape:    "circle",
		Style:    false,
	}
	pillow := Decor{
		Name:     "Pillow",
		Length:   35,
		Width:    35,
		Weight:   0.129,
		Colour:   "silver",
		Material: "polyester",
		Shape:    "square",
		Style:    false,
	}
	fmt.Println("\t\t\tДекор в основной комнате")
	fmt.Println("\tНазвание предмета:", pillow.Name, "\nДлина предмета:", pillow.Length, "\nШирина предмета:", pillow.Width, "\nВес предмета:", pillow.Weight, "\nЦвет предмета:", pillow.Colour, "\nМатериал:", pillow.Material, "\nФорма предмета:", pillow.Shape)
	Decor.StyleMatch(pillow)
	fmt.Println("\tНазвание предмета:", poof.Name, "\nДлина предмета:", poof.Length, "\nШирина предмета:", poof.Width, "\nВес предмета:", poof.Weight, "\nЦвет предмета:", poof.Colour, "\nМатериал:", poof.Material, "\nФорма предмета:", poof.Shape)
	fmt.Println("\tНазвание предмета:", floorLamp.Name, "\nДлина предмета:", floorLamp.Length, "\nШирина предмета:", floorLamp.Width, "\nВес предмета:", floorLamp.Weight, "\nЦвет предмета:", floorLamp.Colour, "\nМатериал:", floorLamp.Material, "\nФорма предмета:", floorLamp.Shape)
	Decor.StyleMatch(floorLamp)
	fmt.Println("\tНазвание предмета:", carpet.Name, "\nДлина предмета:", carpet.Length, "\nШирина предмета:", carpet.Width, "\nВес предмета:", carpet.Weight, "\nЦвет предмета:", carpet.Colour, "\nМатериал:", carpet.Material, "\nФорма предмета:", carpet.Shape)
	fmt.Println("\tНазвание предмета:", bookcase.Name, "\nДлина предмета:", bookcase.Length, "\nШирина предмета:", bookcase.Width, "\nВес предмета:", bookcase.Weight, "\nЦвет предмета:", bookcase.Colour, "\nМатериал:", bookcase.Material, "\nФорма предмета:", bookcase.Shape)
	Decor.StyleMatch(bookcase)
	return Decor{}
}
func PrintBedroomDecor() Decor {
	lamp := Decor{
		Name:     "Lamp",
		Height:   40,
		Width:    15,
		Depth:    10,
		Weight:   2,
		Colour:   "transparent",
		Material: "glass",
		Style:    true,
	}
	picture := Decor{
		Name:     "Picture",
		Length:   50,
		Width:    80,
		Depth:    0,
		Weight:   3,
		Colour:   "black",
		Material: "wood",
		Style:    true,
	}
	statuette := Decor{
		Name:     "Statuette",
		Length:   30,
		Width:    50,
		Depth:    20,
		Weight:   21.74,
		Colour:   "white",
		Material: "gypsum",
		Style:    true,
	}
	photo := Decor{
		Name:     "Photo",
		Height:   20,
		Width:    15,
		Depth:    0,
		Weight:   1,
		Colour:   "multicoloured",
		Material: "glass",
		Style:    true,
	}
	candles := Decor{
		Name:     "Candles",
		Height:   10,
		Width:    8,
		Depth:    0,
		Weight:   1,
		Colour:   "multicoloured",
		Material: "wax",
		Style:    true,
	}
	fmt.Println("\t\t\tДекор в спальне")
	fmt.Println("\tНазвание предмета:", photo.Name, "\nДлина предмета:", photo.Length, "\nШирина предмета:", photo.Width, "\nВес предмета:", photo.Weight, "\nЦвет предмета:", photo.Colour, "\nМатериал:", photo.Material, "\nФорма предмета:", photo.Shape)
	fmt.Println("\tНазвание предмета:", picture.Name, "\nДлина предмета:", picture.Length, "\nШирина предмета:", picture.Width, "\nВес предмета:", picture.Weight, "\nЦвет предмета:", picture.Colour, "\nМатериал:", picture.Material, "\nФорма предмета:", picture.Shape)
	Decor.StyleMatch(picture)
	fmt.Println("\tНазвание предмета:", statuette.Name, "\nДлина предмета:", statuette.Length, "\nШирина предмета:", statuette.Width, "\nВес предмета:", statuette.Weight, "\nЦвет предмета:", statuette.Colour, "\nМатериал:", statuette.Material, "\nФорма предмета:", statuette.Shape)
	fmt.Println("\tНазвание предмета:", lamp.Name, "\nДлина предмета:", lamp.Length, "\nШирина предмета:", lamp.Width, "\nВес предмета:", lamp.Weight, "\nЦвет предмета:", lamp.Colour, "\nМатериал:", lamp.Material, "\nФорма предмета:", lamp.Shape)
	Decor.StyleMatch(lamp)
	fmt.Println("\tНазвание предмета:", candles.Name, "\nДлина предмета:", candles.Length, "\nШирина предмета:", candles.Width, "\nВес предмета:", candles.Weight, "\nЦвет предмета:", candles.Colour, "\nМатериал:", candles.Material, "\nФорма предмета:", candles.Shape)
	return Decor{}
}
