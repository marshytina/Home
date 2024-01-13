package rooms

import (
	"fmt"
	"house/house/appliances"
	"house/house/decor"
	"house/house/furniture"
)

func PrintBedroom() Rooms {
	bedroom := Rooms{
		Name:    "Bedroom",
		Length:  11,
		Width:   10,
		Windows: 2,
	}
	fmt.Println("\tНазвание комнаты:", bedroom.Name)
	fmt.Println("Длина комнаты:", bedroom.Length)
	fmt.Println("Ширина комнаты:", bedroom.Width)
	fmt.Println("Количество окон:", bedroom.Windows)
	fmt.Println("Площадь комнаты:", bedroom.Length*bedroom.Width)
	fmt.Println("\tПредметы имеющиеся в спальне:")
	fmt.Println(appliances.PrintBedroomAppliances())
	fmt.Println(furniture.PrintBedroomFurniture())
	fmt.Println(decor.PrintBedroomDecor())
	return Rooms{}
}
