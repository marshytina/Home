package appliances

import (
	"fmt"
)

type Appliances struct { //бытовая техника
	Length       int
	Height       int
	Width        int
	Depth        int
	Name         string
	Weight       float32
	Colour       string
	Guarantee    int
	Country      string
	SmartHome    bool
	Exploitation int
}

func (a Appliances) IntelligentHome() {
	if a.SmartHome {
		fmt.Println("Устройство", a.Name, "подключено к программе Умный дом")
	} else {
		fmt.Println("Устройство", a.Name, "не подключено к программе Умный дом")
	}
}
func (a Appliances) GuaranteePeriod() {
	if a.Exploitation >= a.Guarantee {
		fmt.Println("Гарантийный срок устройства", a.Name, "истек")
	} else {
		fmt.Println("Гарантийный срок устройства", a.Name, "не истек")

	}
}
func PrintBathroomAppliances() Appliances {
	washingMachine := Appliances{
		Height:       85,
		Width:        59,
		Depth:        40,
		Name:         "WashingMachine",
		Weight:       58,
		Colour:       "white",
		Guarantee:    2,
		Country:      "China",
		SmartHome:    false,
		Exploitation: 3,
	}
	electricToothBrush := Appliances{
		Length:       22,
		Name:         "ElectricToothbrush",
		Weight:       0.208,
		Colour:       "black",
		Guarantee:    1,
		Country:      "Germany",
		SmartHome:    false,
		Exploitation: 1,
	}
	hairDryer := Appliances{
		Length:       35,
		Name:         "Hairdryer",
		Weight:       0.84,
		Colour:       "black",
		Guarantee:    2,
		Country:      "China",
		SmartHome:    false,
		Exploitation: 1}
	electricShaver := Appliances{
		Length:       17,
		Name:         "ElectricShaver",
		Weight:       0.330,
		Colour:       "black",
		Guarantee:    2,
		Country:      "China",
		Exploitation: 3,
	}
	curlingIron := Appliances{
		Length:       32,
		Name:         "CurlingIron",
		Weight:       0.5,
		Colour:       "black",
		Guarantee:    5,
		Country:      "China",
		SmartHome:    false,
		Exploitation: 2}
	fmt.Println("\t\t\tЭлектроприборы в ванной комнате")
	fmt.Println("\tНазвание предмета:", washingMachine.Name, "\nВысота предмета:", washingMachine.Height, "\nШирина предмета:", washingMachine.Width, "\nГлубина предмета:", washingMachine.Depth, "\nВес предмета:", washingMachine.Weight, "\nЦвет предмета:", washingMachine.Colour, "\nГарантия:", washingMachine.Guarantee, "\nСтрана производитель:", washingMachine.Country)
	Appliances.IntelligentHome(washingMachine)
	fmt.Println("\tНазвание предмета:", electricShaver.Name, "\nВысота предмета:", electricShaver.Height, "\nШирина предмета:", electricShaver.Width, "\nГлубина предмета:", electricShaver.Depth, "\nВес предмета:", electricShaver.Weight, "\nЦвет предмета:", electricShaver.Colour, "\nГарантия:", electricShaver.Guarantee, "\nСтрана производитель:", electricShaver.Country)
	fmt.Println("\tНазвание предмета:", electricToothBrush.Name, "\nВысота предмета:", electricToothBrush.Height, "\nШирина предмета:", electricToothBrush.Width, "\nГлубина предмета:", electricToothBrush.Depth, "\nВес предмета:", electricToothBrush.Weight, "\nЦвет предмета:", electricToothBrush.Colour, "\nГарантия:", electricToothBrush.Guarantee, "\nСтрана производитель:", electricToothBrush.Country)
	Appliances.GuaranteePeriod(electricToothBrush)
	fmt.Println("\tНазвание предмета:", hairDryer.Name, "\nВысота предмета:", hairDryer.Height, "\nШирина предмета:", hairDryer.Width, "\nГлубина предмета:", hairDryer.Depth, "\nВес предмета:", hairDryer.Weight, "\nЦвет предмета:", hairDryer.Colour, "\nГарантия:", hairDryer.Guarantee, "\nСтрана производитель:", hairDryer.Country)
	Appliances.IntelligentHome(hairDryer)
	fmt.Println("\tНазвание предмета:", curlingIron.Name, "\nВысота предмета:", curlingIron.Height, "\nШирина предмета:", curlingIron.Width, "\nГлубина предмета:", curlingIron.Depth, "\nВес предмета:", curlingIron.Weight, "\nЦвет предмета:", curlingIron.Colour, "\nГарантия:", curlingIron.Guarantee, "\nСтрана производитель:", curlingIron.Country)
	return Appliances{}
}

func PrintHallAppliances() Appliances {
	tv := Appliances{
		Height:       71,
		Width:        122,
		Depth:        8,
		Name:         "TV",
		Weight:       14.68,
		Colour:       "black",
		Guarantee:    3,
		Country:      "China",
		SmartHome:    true,
		Exploitation: 5}
	ps5 := Appliances{
		Height:       39,
		Width:        16,
		Depth:        9,
		Name:         "PS5",
		Weight:       3.9,
		Colour:       "white",
		Guarantee:    2,
		Country:      "Japan",
		SmartHome:    true,
		Exploitation: 1}
	airConditioner := Appliances{
		Height:       77,
		Width:        49,
		Depth:        29,
		Name:         "AirConditioner",
		Weight:       26,
		Colour:       "white",
		Guarantee:    1,
		Country:      "China",
		SmartHome:    true,
		Exploitation: 2}
	musicCentre := Appliances{
		Height:       72,
		Width:        56,
		Depth:        37,
		Name:         "MusicCentre",
		Weight:       7,
		Colour:       "black",
		Guarantee:    3,
		Country:      "China",
		SmartHome:    true,
		Exploitation: 3,
	}
	laptop := Appliances{
		Height:       38,
		Width:        27,
		Depth:        15,
		Name:         "Laptop",
		Weight:       1.9,
		Colour:       "black",
		Guarantee:    5,
		Country:      "China",
		SmartHome:    false,
		Exploitation: 4,
	}
	fmt.Println("\t\t\tЭлектроприборы в основной комнате")
	fmt.Println("\tНазвание предмета:", musicCentre.Name, "\nВысота предмета:", musicCentre.Height, "\nШирина предмета:", musicCentre.Width, "\nГлубина предмета:", musicCentre.Depth, "\nВес предмета:", musicCentre.Weight, "\nЦвет предмета:", musicCentre.Colour, "\nГарантия:", musicCentre.Guarantee, "\nСтрана производитель:", musicCentre.Country)
	fmt.Println("\tНазвание предмета:", laptop.Name, "\nВысота предмета:", laptop.Height, "\nШирина предмета:", laptop.Width, "\nГлубина предмета:", laptop.Depth, "\nВес предмета:", laptop.Weight, "\nЦвет предмета:", laptop.Colour, "\nГарантия:", laptop.Guarantee, "\nСтрана производитель:", laptop.Country)
	Appliances.IntelligentHome(laptop)
	fmt.Println("\tНазвание предмета:", airConditioner.Name, "\nВысота предмета:", airConditioner.Height, "\nШирина предмета:", airConditioner.Width, "\nГлубина предмета:", airConditioner.Depth, "\nВес предмета:", airConditioner.Weight, "\nЦвет предмета:", airConditioner.Colour, "\nГарантия:", airConditioner.Guarantee, "\nСтрана производитель:", airConditioner.Country)
	fmt.Println("\tНазвание предмета:", tv.Name, "\nВысота предмета:", tv.Height, "\nШирина предмета:", tv.Width, "\nГлубина предмета:", tv.Depth, "\nВес предмета:", tv.Weight, "\nЦвет предмета:", tv.Colour, "\nГарантия:", tv.Guarantee, "\nСтрана производитель:", tv.Country)
	fmt.Println("\tНазвание предмета:", ps5.Name, "\nВысота предмета:", ps5.Height, "\nШирина предмета:", ps5.Width, "\nГлубина предмета:", ps5.Depth, "\nВес предмета:", ps5.Weight, "\nЦвет предмета:", ps5.Colour, "\nГарантия:", ps5.Guarantee, "\nСтрана производитель:", ps5.Country)
	Appliances.GuaranteePeriod(ps5)
	return Appliances{}
}
func PrintKitchenAppliances() Appliances {
	dishwasher := Appliances{
		Height:       89,
		Width:        64,
		Depth:        65,
		Name:         "Dishwasher",
		Weight:       40.58,
		Colour:       "silver",
		Guarantee:    3,
		Country:      "Italy",
		SmartHome:    false,
		Exploitation: 1,
	}
	kettle := Appliances{
		Height:       25,
		Width:        21,
		Depth:        18,
		Name:         "Kettle",
		Weight:       1.25,
		Colour:       "black",
		Guarantee:    2,
		Country:      "Italy",
		SmartHome:    true,
		Exploitation: 1,
	}
	coffemachine := Appliances{
		Height:       37,
		Width:        25,
		Depth:        43,
		Name:         "Coffeemachine",
		Weight:       7.5,
		Colour:       "black",
		Guarantee:    1,
		Country:      "Netherlands",
		SmartHome:    false,
		Exploitation: 1,
	}
	electricStove := Appliances{
		Height:       85,
		Width:        50,
		Depth:        59,
		Name:         "ElectricStove",
		Weight:       38.8,
		Colour:       "black",
		Guarantee:    4,
		Country:      "Czech",
		SmartHome:    false,
		Exploitation: 2,
	}
	fridge := Appliances{
		Height:       185,
		Width:        60,
		Depth:        64,
		Name:         "Fridge",
		Weight:       65.4,
		Colour:       "black",
		Guarantee:    5,
		Country:      "Italy",
		SmartHome:    false,
		Exploitation: 2,
	}
	multicooker := Appliances{
		Height:       28,
		Width:        24,
		Depth:        37,
		Name:         "Multicooker",
		Weight:       3.4,
		Colour:       "black",
		Guarantee:    1,
		Country:      "China",
		SmartHome:    false,
		Exploitation: 2,
	}
	fmt.Println("\t\t\tЭлектроприборы в кухне")
	fmt.Println("\tНазвание предмета:", electricStove.Name, "\nВысота предмета:", electricStove.Height, "\nШирина предмета:", electricStove.Width, "\nГлубина предмета:", electricStove.Depth, "\nВес предмета:", electricStove.Weight, "\nЦвет предмета:", electricStove.Colour, "\nГарантия:", electricStove.Guarantee, "\nСтрана производитель:", electricStove.Country)
	Appliances.IntelligentHome(electricStove)
	fmt.Println("\tНазвание предмета:", multicooker.Name, "\nВысота предмета:", multicooker.Height, "\nШирина предмета:", multicooker.Width, "\nГлубина предмета:", multicooker.Depth, "\nВес предмета:", multicooker.Weight, "\nЦвет предмета:", multicooker.Colour, "\nГарантия:", multicooker.Guarantee, "\nСтрана производитель:", multicooker.Country)
	fmt.Println("\tНазвание предмета:", coffemachine.Name, "\nВысота предмета:", coffemachine.Height, "\nШирина предмета:", coffemachine.Width, "\nГлубина предмета:", coffemachine.Depth, "\nВес предмета:", coffemachine.Weight, "\nЦвет предмета:", coffemachine.Colour, "\nГарантия:", coffemachine.Guarantee, "\nСтрана производитель:", coffemachine.Country)
	Appliances.GuaranteePeriod(coffemachine)
	fmt.Println("\tНазвание предмета:", dishwasher.Name, "\nВысота предмета:", dishwasher.Height, "\nШирина предмета:", dishwasher.Width, "\nГлубина предмета:", dishwasher.Depth, "\nВес предмета:", dishwasher.Weight, "\nЦвет предмета:", dishwasher.Colour, "\nГарантия:", dishwasher.Guarantee, "\nСтрана производитель:", dishwasher.Country)
	Appliances.IntelligentHome(dishwasher)
	fmt.Println("\tНазвание предмета:", kettle.Name, "\nВысота предмета:", kettle.Height, "\nШирина предмета:", kettle.Width, "\nГлубина предмета:", kettle.Depth, "\nВес предмета:", kettle.Weight, "\nЦвет предмета:", kettle.Colour, "\nГарантия:", kettle.Guarantee, "\nСтрана производитель:", kettle.Country)
	fmt.Println("\tНазвание предмета:", fridge.Name, "\nВысота предмета:", fridge.Height, "\nШирина предмета:", fridge.Width, "\nГлубина предмета:", fridge.Depth, "\nВес предмета:", fridge.Weight, "\nЦвет предмета:", fridge.Colour, "\nГарантия:", fridge.Guarantee, "\nСтрана производитель:", fridge.Country)
	Appliances.GuaranteePeriod(fridge)
	return Appliances{}
}
func PrintBedroomAppliances() Appliances {
	computer := Appliances{
		Height:       60,
		Width:        75,
		Depth:        0,
		Name:         "Computer",
		Weight:       5.58,
		Colour:       "silver",
		Guarantee:    3,
		Country:      "China",
		SmartHome:    false,
		Exploitation: 1,
	}
	roomba := Appliances{
		Height:       10,
		Width:        20,
		Depth:        0,
		Name:         "Roomba",
		Weight:       2.5,
		Colour:       "white",
		Guarantee:    2,
		Country:      "Italy",
		SmartHome:    true,
		Exploitation: 1,
	}
	tv := Appliances{
		Height:       50,
		Width:        80,
		Depth:        0,
		Name:         "TV",
		Weight:       6,
		Colour:       "black",
		Guarantee:    2,
		Country:      "China",
		SmartHome:    false,
		Exploitation: 1,
	}
	fmt.Println("\t\t\tЭлектроприборы в спальне")
	fmt.Println("\tНазвание предмета:", roomba.Name, "\nВысота предмета:", roomba.Height, "\nШирина предмета:", roomba.Width, "\nГлубина предмета:", roomba.Depth, "\nВес предмета:", roomba.Weight, "\nЦвет предмета:", roomba.Colour, "\nГарантия:", roomba.Guarantee, "\nСтрана производитель:", roomba.Country)
	fmt.Println("\tНазвание предмета:", computer.Name, "\nВысота предмета:", computer.Height, "\nШирина предмета:", computer.Width, "\nГлубина предмета:", computer.Depth, "\nВес предмета:", computer.Weight, "\nЦвет предмета:", computer.Colour, "\nГарантия:", computer.Guarantee, "\nСтрана производитель:", computer.Country)
	Appliances.IntelligentHome(computer)
	fmt.Println("\tНазвание предмета:", tv.Name, "\nВысота предмета:", tv.Height, "\nШирина предмета:", tv.Width, "\nГлубина предмета:", tv.Depth, "\nВес предмета:", tv.Weight, "\nЦвет предмета:", tv.Colour, "\nГарантия:", tv.Guarantee, "\nСтрана производитель:", tv.Country)
	return Appliances{}
}
