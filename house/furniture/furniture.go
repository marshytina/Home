package furniture

import "fmt"

type Furniture struct {
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

func PrintHallFirnuture() Furniture {
	table := Furniture{
		Name:     "Table",
		Height:   75,
		Width:    80,
		Depth:    0,
		Weight:   13.1,
		Colour:   "white",
		Material: "wood",
		Shape:    "square"}
	chair := Furniture{
		Name:     "Chair",
		Height:   83,
		Width:    45,
		Depth:    0,
		Weight:   3.6,
		Colour:   "black",
		Material: "plastic",
		Shape:    "rectangular",
	}
	dresser := Furniture{
		Name:     "Dresser",
		Height:   80,
		Width:    47,
		Depth:    40,
		Weight:   28,
		Colour:   "white",
		Material: "wood",
		Shape:    "rectangular",
	}
	sofa := Furniture{
		Name:     "Sofa",
		Height:   70,
		Width:    130,
		Depth:    50,
		Weight:   45,
		Colour:   "black",
		Material: "flock",
		Shape:    "rectangular",
	}
	closet := Furniture{
		Name:     "Closet",
		Height:   185,
		Width:    100,
		Depth:    70,
		Weight:   35,
		Colour:   "brown",
		Material: "wood",
		Shape:    "rectangular",
	}
	fmt.Println("\t\t\tМебель в основной комнате")
	fmt.Println("\tНазвание предмета:", closet.Name, "\nВес предмета:", closet.Weight, "\nЦвет предмета:", closet.Colour, "\nМатериал:", closet.Material, "\nДлина предмета:", closet.Height, "\nШирина предмета:", closet.Width, "\nГлубина предмета:", closet.Depth)
	fmt.Println("\tНазвание предмета:", table.Name, "\nВес предмета:", table.Weight, "\nЦвет предмета:", table.Colour, "\nМатериал:", table.Material, "\nДлина предмета:", table.Height, "\nШирина предмета:", table.Width, "\nГлубина предмета:", table.Depth)
	fmt.Println("\tНазвание предмета:", chair.Name, "\nВес предмета:", chair.Weight, "\nЦвет предмета:", chair.Colour, "\nМатериал:", chair.Material, "\nДлина предмета:", chair.Height, "\nШирина предмета:", chair.Width, "\nГлубина предмета:", chair.Depth)
	fmt.Println("\tНазвание предмета:", dresser.Name, "\nВес предмета:", dresser.Weight, "\nЦвет предмета:", dresser.Colour, "\nМатериал:", dresser.Material, "\nДлина предмета:", dresser.Height, "\nШирина предмета:", dresser.Width, "\nГлубина предмета:", dresser.Depth)
	fmt.Println("\tНазвание предмета:", sofa.Name, "\nВес предмета:", sofa.Weight, "\nЦвет предмета:", sofa.Colour, "\nМатериал:", sofa.Material, "\nДлина предмета:", sofa.Height, "\nШирина предмета:", sofa.Width, "\nГлубина предмета:", sofa.Depth)
	return Furniture{}
}
func PrintBedroomFurniture() Furniture {
	worktable := Furniture{
		Name:     "Table",
		Height:   50,
		Width:    90,
		Depth:    0,
		Weight:   14.5,
		Colour:   "white",
		Material: "wood",
		Shape:    "square"}
	armchair := Furniture{
		Name:     "Armhair",
		Height:   70,
		Width:    50,
		Depth:    0,
		Weight:   5,
		Colour:   "white",
		Material: "wood",
		Shape:    "rectangular",
	}
	bed := Furniture{
		Name:     "Bed",
		Height:   75,
		Width:    150,
		Depth:    60,
		Weight:   45,
		Colour:   "white",
		Material: "wood",
		Shape:    "rectangular",
	}
	wardrobe := Furniture{
		Name:     "Wardrobe",
		Height:   180,
		Width:    90,
		Depth:    40,
		Weight:   60,
		Colour:   "white",
		Material: "wood",
		Shape:    "rectangular",
	}
	fmt.Println("\t\t\tМебель в спальне")
	fmt.Println("\tНазвание предмета:", armchair.Name, "\nВес предмета:", armchair.Weight, "\nЦвет предмета:", armchair.Colour, "\nМатериал:", armchair.Material, "\nДлина предмета:", armchair.Height, "\nШирина предмета:", armchair.Width, "\nГлубина предмета:", armchair.Depth)
	fmt.Println("\tНазвание предмета:", wardrobe.Name, "\nВес предмета:", wardrobe.Weight, "\nЦвет предмета:", wardrobe.Colour, "\nМатериал:", wardrobe.Material, "\nДлина предмета:", wardrobe.Height, "\nШирина предмета:", wardrobe.Width, "\nГлубина предмета:", wardrobe.Depth)
	fmt.Println("\tНазвание предмета:", worktable.Name, "\nВес предмета:", worktable.Weight, "\nЦвет предмета:", worktable.Colour, "\nМатериал:", worktable.Material, "\nДлина предмета:", worktable.Height, "\nШирина предмета:", worktable.Width, "\nГлубина предмета:", worktable.Depth)
	fmt.Println("\tНазвание предмета:", bed.Name, "\nВес предмета:", bed.Weight, "\nЦвет предмета:", bed.Colour, "\nМатериал:", bed.Material, "\nДлина предмета:", bed.Height, "\nШирина предмета:", bed.Width, "\nГлубина предмета:", bed.Depth)
	return Furniture{}
}
func PrintBathroomFurniture() Furniture {
	bath := Furniture{
		Name:     "Bath",
		Length:   150,
		Height:   59,
		Width:    70,
		Depth:    41,
		Weight:   16.5,
		Colour:   "white",
		Material: "acrylic",
		Shape:    "rectangular",
	}
	sink := Furniture{
		Name:     "Sink",
		Length:   100,
		Height:   16,
		Depth:    50,
		Weight:   22.39,
		Colour:   "white",
		Material: "marble",
		Shape:    "rectangular",
	}
	wallCabinet := Furniture{
		Name:     "WallCabinet",
		Height:   70,
		Width:    50,
		Depth:    19,
		Weight:   13,
		Colour:   "white",
		Material: "wood",
		Shape:    "rectangular",
	}
	toilet := Furniture{
		Name:     "Toilet",
		Height:   78,
		Width:    35,
		Depth:    62,
		Weight:   28,
		Colour:   "white",
		Material: "ceramic",
	}
	fmt.Println("\t\t\tМебель в ванной")
	fmt.Println("\tНазвание предмета:", toilet.Name, "\nВес предмета:", toilet.Weight, "\nЦвет предмета:", toilet.Colour, "\nМатериал:", toilet.Material, "\nДлина предмета:", toilet.Height, "\nШирина предмета:", toilet.Width, "\nГлубина предмета:", toilet.Depth)
	fmt.Println("\tНазвание предмета:", sink.Name, "\nВес предмета:", sink.Weight, "\nЦвет предмета:", sink.Colour, "\nМатериал:", sink.Material, "\nДлина предмета:", sink.Height, "\nШирина предмета:", sink.Width, "\nГлубина предмета:", sink.Depth)
	fmt.Println("\tНазвание предмета:", wallCabinet.Name, "\nВес предмета:", wallCabinet.Weight, "\nЦвет предмета:", wallCabinet.Colour, "\nМатериал:", wallCabinet.Material, "\nДлина предмета:", wallCabinet.Height, "\nШирина предмета:", wallCabinet.Width, "\nГлубина предмета:", wallCabinet.Depth)
	fmt.Println("\tНазвание предмета:", bath.Name, "\nВес предмета:", bath.Weight, "\nЦвет предмета:", bath.Colour, "\nМатериал:", bath.Material, "\nДлина предмета:", bath.Height, "\nШирина предмета:", bath.Width, "\nГлубина предмета:", bath.Depth)
	return Furniture{}
}
