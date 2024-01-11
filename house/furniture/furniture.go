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

func PrintFirnuture() Furniture {
	fmt.Println("Вся мебель имеющаяся в доме:")
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
	chandelier := Furniture{
		Name:     "Chandelier",
		Height:   65,
		Width:    48,
		Depth:    48,
		Weight:   2.2,
		Colour:   "white",
		Material: "glass",
		Shape:    "author's",
	}
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
	armchair := Furniture{
		Name:     "Armchair",
		Height:   83,
		Width:    45,
		Depth:    0,
		Weight:   9,
		Colour:   "black",
		Material: "flock",
	}
	fmt.Println("Название предмета:", armchair.Name, "Вес предмета:", armchair.Weight, "Цвет предмета:", armchair.Colour, "Материал:", armchair.Material, "Длина предмета:", armchair.Height, "Ширина предмета:", armchair.Width, "Глубина предмета:", armchair.Depth)
	fmt.Println("Название предмета:", closet.Name, "Вес предмета:", closet.Weight, "Цвет предмета:", closet.Colour, "Материал:", closet.Material, "Длина предмета:", closet.Height, "Ширина предмета:", closet.Width, "Глубина предмета:", closet.Depth)
	fmt.Println("Название предмета:", toilet.Name, "Вес предмета:", toilet.Weight, "Цвет предмета:", toilet.Colour, "Материал:", toilet.Material, "Длина предмета:", toilet.Height, "Ширина предмета:", toilet.Width, "Глубина предмета:", toilet.Depth)
	fmt.Println("Название предмета:", table.Name, "Вес предмета:", table.Weight, "Цвет предмета:", table.Colour, "Материал:", table.Material, "Длина предмета:", table.Height, "Ширина предмета:", table.Width, "Глубина предмета:", table.Depth)
	fmt.Println("Название предмета:", chair.Name, "Вес предмета:", armchair.Weight, "Цвет предмета:", chair.Colour, "Материал:", chair.Material, "Длина предмета:", chair.Height, "Ширина предмета:", chair.Width, "Глубина предмета:", chair.Depth)
	fmt.Println("Название предмета:", sink.Name, "Вес предмета:", sink.Weight, "Цвет предмета:", sink.Colour, "Материал:", sink.Material, "Длина предмета:", sink.Height, "Ширина предмета:", sink.Width, "Глубина предмета:", sink.Depth)
	fmt.Println("Название предмета:", wallCabinet.Name, "Вес предмета:", wallCabinet.Weight, "Цвет предмета:", wallCabinet.Colour, "Материал:", wallCabinet.Material, "Длина предмета:", wallCabinet.Height, "Ширина предмета:", wallCabinet.Width, "Глубина предмета:", wallCabinet.Depth)
	fmt.Println("Название предмета:", dresser.Name, "Вес предмета:", dresser.Weight, "Цвет предмета:", dresser.Colour, "Материал:", dresser.Material, "Длина предмета:", dresser.Height, "Ширина предмета:", dresser.Width, "Глубина предмета:", dresser.Depth)
	fmt.Println("Название предмета:", sofa.Name, "Вес предмета:", sofa.Weight, "Цвет предмета:", sofa.Colour, "Материал:", sofa.Material, "Длина предмета:", sofa.Height, "Ширина предмета:", sofa.Width, "Глубина предмета:", sofa.Depth)
	fmt.Println("Название предмета:", chandelier.Name, "Вес предмета:", chandelier.Weight, "Цвет предмета:", chandelier.Colour, "Материал:", chandelier.Material, "Длина предмета:", chandelier.Height, "Ширина предмета:", chandelier.Width, "Глубина предмета:", chandelier.Depth)
	fmt.Println("Название предмета:", bath.Name, "Вес предмета:", bath.Weight, "Цвет предмета:", bath.Colour, "Материал:", bath.Material, "Длина предмета:", bath.Height, "Ширина предмета:", bath.Width, "Глубина предмета:", bath.Depth)
	return Furniture{}
}
