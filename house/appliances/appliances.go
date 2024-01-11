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
	fmt.Println("\t\t\tЭлектроприборы в ванной комнате")
	fmt.Println("\tНазвание предмета:", washingMachine.Name, "\nВысота предмета:", washingMachine.Height, "\nШирина предмета:", washingMachine.Width, "\nГлубина предмета:", washingMachine.Depth, "\nВес предмета:", washingMachine.Weight, "\nЦвет предмета:", washingMachine.Colour, "\nГарантия:", washingMachine.Guarantee, "\nСтрана производитель:", washingMachine.Country)
	fmt.Println("\tНазвание предмета:", electricShaver.Name, "\nВысота предмета:", electricShaver.Height, "\nШирина предмета:", electricShaver.Width, "\nГлубина предмета:", electricShaver.Depth, "\nВес предмета:", electricShaver.Weight, "\nЦвет предмета:", electricShaver.Colour, "\nГарантия:", electricShaver.Guarantee, "\nСтрана производитель:", electricShaver.Country)
	fmt.Println("\tНазвание предмета:", electricToothBrush.Name, "\nВысота предмета:", electricToothBrush.Height, "\nШирина предмета:", electricToothBrush.Width, "\nГлубина предмета:", electricToothBrush.Depth, "\nВес предмета:", electricToothBrush.Weight, "\nЦвет предмета:", electricToothBrush.Colour, "\nГарантия:", electricToothBrush.Guarantee, "\nСтрана производитель:", electricToothBrush.Country)
	fmt.Println("\tНазвание предмета:", hairDryer.Name, "\nВысота предмета:", hairDryer.Height, "\nШирина предмета:", hairDryer.Width, "\nГлубина предмета:", hairDryer.Depth, "\nВес предмета:", hairDryer.Weight, "\nЦвет предмета:", hairDryer.Colour, "\nГарантия:", hairDryer.Guarantee, "\nСтрана производитель:", hairDryer.Country)
	fmt.Println("\tНазвание предмета:", curlingIron.Name, "\nВысота предмета:", curlingIron.Height, "\nШирина предмета:", curlingIron.Width, "\nГлубина предмета:", curlingIron.Depth, "\nВес предмета:", curlingIron.Weight, "\nЦвет предмета:", curlingIron.Colour, "\nГарантия:", curlingIron.Guarantee, "\nСтрана производитель:", curlingIron.Country)
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
	fmt.Println("\t\t\tЭлектроприборы в основной комнате")
	fmt.Println("\tНазвание предмета:", musicCentre.Name, "\nВысота предмета:", musicCentre.Height, "\nШирина предмета:", musicCentre.Width, "\nГлубина предмета:", musicCentre.Depth, "\nВес предмета:", musicCentre.Weight, "\nЦвет предмета:", musicCentre.Colour, "\nГарантия:", musicCentre.Guarantee, "\nСтрана производитель:", musicCentre.Country)
	fmt.Println("\tНазвание предмета:", laptop.Name, "\nВысота предмета:", laptop.Height, "\nШирина предмета:", laptop.Width, "\nГлубина предмета:", laptop.Depth, "\nВес предмета:", laptop.Weight, "\nЦвет предмета:", laptop.Colour, "\nГарантия:", laptop.Guarantee, "\nСтрана производитель:", laptop.Country)
	fmt.Println("\tНазвание предмета:", airConditioner.Name, "\nВысота предмета:", airConditioner.Height, "\nШирина предмета:", airConditioner.Width, "\nГлубина предмета:", airConditioner.Depth, "\nВес предмета:", airConditioner.Weight, "\nЦвет предмета:", airConditioner.Colour, "\nГарантия:", airConditioner.Guarantee, "\nСтрана производитель:", airConditioner.Country)
	fmt.Println("\tНазвание предмета:", tv.Name, "\nВысота предмета:", tv.Height, "\nШирина предмета:", tv.Width, "\nГлубина предмета:", tv.Depth, "\nВес предмета:", tv.Weight, "\nЦвет предмета:", tv.Colour, "\nГарантия:", tv.Guarantee, "\nСтрана производитель:", tv.Country)
	fmt.Println("\tНазвание предмета:", ps5.Name, "\nВысота предмета:", ps5.Height, "\nШирина предмета:", ps5.Width, "\nГлубина предмета:", ps5.Depth, "\nВес предмета:", ps5.Weight, "\nЦвет предмета:", ps5.Colour, "\nГарантия:", ps5.Guarantee, "\nСтрана производитель:", ps5.Country)
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
	fmt.Println("\t\t\tЭлектроприборы в кухне")
	fmt.Println("\tНазвание предмета:", electricStove.Name, "\nВысота предмета:", electricStove.Height, "\nШирина предмета:", electricStove.Width, "\nГлубина предмета:", electricStove.Depth, "\nВес предмета:", electricStove.Weight, "\nЦвет предмета:", electricStove.Colour, "\nГарантия:", electricStove.Guarantee, "\nСтрана производитель:", electricStove.Country)
	fmt.Println("\tНазвание предмета:", mixer.Name, "\nВысота предмета:", mixer.Height, "\nШирина предмета:", mixer.Width, "\nГлубина предмета:", mixer.Depth, "\nВес предмета:", mixer.Weight, "\nЦвет предмета:", mixer.Colour, "\nГарантия:", mixer.Guarantee, "\nСтрана производитель:", mixer.Country)
	fmt.Println("\tНазвание предмета:", multicooker.Name, "\nВысота предмета:", multicooker.Height, "\nШирина предмета:", multicooker.Width, "\nГлубина предмета:", multicooker.Depth, "\nВес предмета:", multicooker.Weight, "\nЦвет предмета:", multicooker.Colour, "\nГарантия:", multicooker.Guarantee, "\nСтрана производитель:", multicooker.Country)
	fmt.Println("\tНазвание предмета:", coffemachine.Name, "\nВысота предмета:", coffemachine.Height, "\nШирина предмета:", coffemachine.Width, "\nГлубина предмета:", coffemachine.Depth, "\nВес предмета:", coffemachine.Weight, "\nЦвет предмета:", coffemachine.Colour, "\nГарантия:", coffemachine.Guarantee, "\nСтрана производитель:", coffemachine.Country)
	fmt.Println("\tНазвание предмета:", dishwasher.Name, "\nВысота предмета:", dishwasher.Height, "\nШирина предмета:", dishwasher.Width, "\nГлубина предмета:", dishwasher.Depth, "\nВес предмета:", dishwasher.Weight, "\nЦвет предмета:", dishwasher.Colour, "\nГарантия:", dishwasher.Guarantee, "\nСтрана производитель:", dishwasher.Country)
	fmt.Println("\tНазвание предмета:", kettle.Name, "\nВысота предмета:", kettle.Height, "\nШирина предмета:", kettle.Width, "\nГлубина предмета:", kettle.Depth, "\nВес предмета:", kettle.Weight, "\nЦвет предмета:", kettle.Colour, "\nГарантия:", kettle.Guarantee, "\nСтрана производитель:", kettle.Country)
	fmt.Println("\tНазвание предмета:", fridge.Name, "\nВысота предмета:", fridge.Height, "\nШирина предмета:", fridge.Width, "\nГлубина предмета:", fridge.Depth, "\nВес предмета:", fridge.Weight, "\nЦвет предмета:", fridge.Colour, "\nГарантия:", fridge.Guarantee, "\nСтрана производитель:", fridge.Country)
	return Appliances{}
}
