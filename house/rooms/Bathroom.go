package rooms

import (
	"fmt"
	"house/house/accessories"
	"house/house/appliances"
	"house/house/furniture"
)

func PrintBathroom() Rooms {
	bathroom := Rooms{
		Name:    "Bathroom",
		Length:  5,
		Width:   3,
		Windows: 1,
	}
	fmt.Println("\tНазвание комнаты:", bathroom.Name)
	fmt.Println("Длина комнаты:", bathroom.Length)
	fmt.Println("Ширина комнаты:", bathroom.Width)
	fmt.Println("Kоличество окон:", bathroom.Windows)
	fmt.Println("Площадь комнаты:", bathroom.Length*bathroom.Width)
	fmt.Println("\tПредметы имеющиеся в ванной:")
	fmt.Println(accessories.PrintBathroomAccessories())
	fmt.Println(appliances.PrintBathroomAppliances())
	fmt.Println(furniture.PrintBathroomFurniture())
	return Rooms{}
}
