package appliances

import (
	"fmt"
)

type Appliances struct { //бытовая техника
	Length    int
	Height    int
	Width     int
	Depth     int
	Name      string
	Weight    float32
	Colour    string
	Guarantee int
	Country   string
}

func PrintBathroomAppliances() Appliances {
	fmt.Println("Вся бытовая техника имеющаяся в доме:")
	washingMachine := Appliances{
		Height:    85,
		Width:     59,
		Depth:     40,
		Name:      "WashingMachine",
		Weight:    58,
		Colour:    "white",
		Guarantee: 1,
		Country:   "China"}
	electricToothBrush := Appliances{
		Length:    22,
		Name:      "ElectricToothbrush",
		Weight:    0.208,
		Colour:    "black",
		Guarantee: 2,
		Country:   "Germany"}
	hairDryer := Appliances{
		Length:    35,
		Name:      "Hairdryer",
		Weight:    0.84,
		Colour:    "black",
		Guarantee: 2,
		Country:   "China"}
	electricShaver := Appliances{
		Length:    17,
		Name:      "ElectricShaver",
		Weight:    0.330,
		Colour:    "black",
		Guarantee: 2,
		Country:   "China"}
	curlingIron := Appliances{
		Length:    32,
		Name:      "CurlingIron",
		Weight:    0.5,
		Colour:    "black",
		Guarantee: 2,
		Country:   "China"}
	fmt.Println("Название предмета:", washingMachine.Name, "Высота предмета:", washingMachine.Height, "Ширина предмета:", washingMachine.Width, "Глубина предмета:", washingMachine.Depth, "Вес предмета:", washingMachine.Weight, "Цвет предмета:", washingMachine.Colour, "Гарантия:", washingMachine.Guarantee, "Страна производитель:", washingMachine.Country)
	fmt.Println("Название предмета:", electricShaver.Name, "Высота предмета:", electricShaver.Height, "Ширина предмета:", electricShaver.Width, "Глубина предмета:", electricShaver.Depth, "Вес предмета:", electricShaver.Weight, "Цвет предмета:", electricShaver.Colour, "Гарантия:", electricShaver.Guarantee, "Страна производитель:", electricShaver.Country)
	fmt.Println("Название предмета:", electricToothBrush.Name, "Высота предмета:", electricToothBrush.Height, "Ширина предмета:", electricToothBrush.Width, "Глубина предмета:", electricToothBrush.Depth, "Вес предмета:", electricToothBrush.Weight, "Цвет предмета:", electricToothBrush.Colour, "Гарантия:", electricToothBrush.Guarantee, "Страна производитель:", electricToothBrush.Country)
	fmt.Println("Название предмета:", hairDryer.Name, "Высота предмета:", hairDryer.Height, "Ширина предмета:", hairDryer.Width, "Глубина предмета:", hairDryer.Depth, "Вес предмета:", hairDryer.Weight, "Цвет предмета:", hairDryer.Colour, "Гарантия:", hairDryer.Guarantee, "Страна производитель:", hairDryer.Country)
	fmt.Println("Название предмета:", curlingIron.Name, "Высота предмета:", curlingIron.Height, "Ширина предмета:", curlingIron.Width, "Глубина предмета:", curlingIron.Depth, "Вес предмета:", curlingIron.Weight, "Цвет предмета:", curlingIron.Colour, "Гарантия:", curlingIron.Guarantee, "Страна производитель:", curlingIron.Country)
	return Appliances{}
}
func PrintHallAppliances() Appliances {
	tv := Appliances{
		Height:    71,
		Width:     122,
		Depth:     8,
		Name:      "TV",
		Weight:    14.68,
		Colour:    "black",
		Guarantee: 1,
		Country:   "China"}
	ps5 := Appliances{
		Height:    39,
		Width:     16,
		Depth:     9,
		Name:      "PS5",
		Weight:    3.9,
		Colour:    "white",
		Guarantee: 1,
		Country:   "Japan"}
	airConditioner := Appliances{
		Height:    77,
		Width:     49,
		Depth:     29,
		Name:      "AirConditioner",
		Weight:    26,
		Colour:    "white",
		Guarantee: 1,
		Country:   "China"}
	musicCentre := Appliances{
		Height:    72,
		Width:     56,
		Depth:     37,
		Name:      "MusicCentre",
		Weight:    7,
		Colour:    "black",
		Guarantee: 1,
		Country:   "China",
	}
	laptop := Appliances{
		Height:    38,
		Width:     27,
		Depth:     15,
		Name:      "Laptop",
		Weight:    1.9,
		Colour:    "black",
		Guarantee: 1,
		Country:   "China",
	}
	fmt.Println("Название предмета:", musicCentre.Name, "Высота предмета:", musicCentre.Height, "Ширина предмета:", musicCentre.Width, "Глубина предмета:", musicCentre.Depth, "Вес предмета:", musicCentre.Weight, "Цвет предмета:", musicCentre.Colour, "Гарантия:", musicCentre.Guarantee, "Страна производитель:", musicCentre.Country)
	fmt.Println("Название предмета:", laptop.Name, "Высота предмета:", laptop.Height, "Ширина предмета:", laptop.Width, "Глубина предмета:", laptop.Depth, "Вес предмета:", laptop.Weight, "Цвет предмета:", laptop.Colour, "Гарантия:", laptop.Guarantee, "Страна производитель:", laptop.Country)
	fmt.Println("Название предмета:", airConditioner.Name, "Высота предмета:", airConditioner.Height, "Ширина предмета:", airConditioner.Width, "Глубина предмета:", airConditioner.Depth, "Вес предмета:", airConditioner.Weight, "Цвет предмета:", airConditioner.Colour, "Гарантия:", airConditioner.Guarantee, "Страна производитель:", airConditioner.Country)
	fmt.Println("Название предмета:", tv.Name, "Высота предмета:", tv.Height, "Ширина предмета:", tv.Width, "Глубина предмета:", tv.Depth, "Вес предмета:", tv.Weight, "Цвет предмета:", tv.Colour, "Гарантия:", tv.Guarantee, "Страна производитель:", tv.Country)
	fmt.Println("Название предмета:", ps5.Name, "Высота предмета:", ps5.Height, "Ширина предмета:", ps5.Width, "Глубина предмета:", ps5.Depth, "Вес предмета:", ps5.Weight, "Цвет предмета:", ps5.Colour, "Гарантия:", ps5.Guarantee, "Страна производитель:", ps5.Country)
	return Appliances{}
}
func PrintKitchenAppliances() Appliances {
	dishwasher := Appliances{
		Height:    89,
		Width:     64,
		Depth:     65,
		Name:      "Dishwasher",
		Weight:    40.58,
		Colour:    "silver",
		Guarantee: 3,
		Country:   "Italy",
	}
	kettle := Appliances{
		Height:    25,
		Width:     21,
		Depth:     18,
		Name:      "Kettle",
		Weight:    1.25,
		Colour:    "black",
		Guarantee: 1,
		Country:   "Italy",
	}
	coffemachine := Appliances{
		Height:    37,
		Width:     25,
		Depth:     43,
		Name:      "Coffeemachine",
		Weight:    7.5,
		Colour:    "black",
		Guarantee: 1,
		Country:   "Netherlands",
	}
	electricStove := Appliances{
		Height:    85,
		Width:     50,
		Depth:     59,
		Name:      "ElectricStove",
		Weight:    38.8,
		Colour:    "black",
		Guarantee: 1,
		Country:   "Czech",
	}
	fridge := Appliances{
		Height:    185,
		Width:     60,
		Depth:     64,
		Name:      "Fridge",
		Weight:    65.4,
		Colour:    "black",
		Guarantee: 1,
		Country:   "Italy",
	}
	mixer := Appliances{
		Height:    38,
		Width:     25,
		Depth:     36,
		Name:      "Mixer",
		Weight:    6.9,
		Colour:    "black",
		Guarantee: 1,
		Country:   "China",
	}
	multicooker := Appliances{
		Height:    28,
		Width:     24,
		Depth:     37,
		Name:      "Multicooker",
		Weight:    3.4,
		Colour:    "black",
		Guarantee: 1,
		Country:   "China",
	}
	fmt.Println("Название предмета:", electricStove.Name, "Высота предмета:", electricStove.Height, "Ширина предмета:", electricStove.Width, "Глубина предмета:", electricStove.Depth, "Вес предмета:", electricStove.Weight, "Цвет предмета:", electricStove.Colour, "Гарантия:", electricStove.Guarantee, "Страна производитель:", electricStove.Country)
	fmt.Println("Название предмета:", mixer.Name, "Высота предмета:", mixer.Height, "Ширина предмета:", mixer.Width, "Глубина предмета:", mixer.Depth, "Вес предмета:", mixer.Weight, "Цвет предмета:", mixer.Colour, "Гарантия:", mixer.Guarantee, "Страна производитель:", mixer.Country)
	fmt.Println("Название предмета:", multicooker.Name, "Высота предмета:", multicooker.Height, "Ширина предмета:", multicooker.Width, "Глубина предмета:", multicooker.Depth, "Вес предмета:", multicooker.Weight, "Цвет предмета:", multicooker.Colour, "Гарантия:", multicooker.Guarantee, "Страна производитель:", multicooker.Country)
	fmt.Println("Название предмета:", coffemachine.Name, "Высота предмета:", coffemachine.Height, "Ширина предмета:", coffemachine.Width, "Глубина предмета:", coffemachine.Depth, "Вес предмета:", coffemachine.Weight, "Цвет предмета:", coffemachine.Colour, "Гарантия:", coffemachine.Guarantee, "Страна производитель:", coffemachine.Country)
	fmt.Println("Название предмета:", dishwasher.Name, "Высота предмета:", dishwasher.Height, "Ширина предмета:", dishwasher.Width, "Глубина предмета:", dishwasher.Depth, "Вес предмета:", dishwasher.Weight, "Цвет предмета:", dishwasher.Colour, "Гарантия:", dishwasher.Guarantee, "Страна производитель:", dishwasher.Country)
	fmt.Println("Название предмета:", kettle.Name, "Высота предмета:", kettle.Height, "Ширина предмета:", kettle.Width, "Глубина предмета:", kettle.Depth, "Вес предмета:", kettle.Weight, "Цвет предмета:", kettle.Colour, "Гарантия:", kettle.Guarantee, "Страна производитель:", kettle.Country)
	fmt.Println("Название предмета:", fridge.Name, "Высота предмета:", fridge.Height, "Ширина предмета:", fridge.Width, "Глубина предмета:", fridge.Depth, "Вес предмета:", fridge.Weight, "Цвет предмета:", fridge.Colour, "Гарантия:", fridge.Guarantee, "Страна производитель:", fridge.Country)
	return Appliances{}
}
