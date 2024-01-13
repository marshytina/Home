package rooms

import (
	"fmt"
	"house/house/appliances"
	"house/house/decor"
	"house/house/furniture"
)

func PrintHall() Rooms {
	hall := Rooms{
		Name:    "Hall",
		Length:  6,
		Width:   10,
		Windows: 3,
	}
	fmt.Println("\tНазвание комнаты:", hall.Name)
	fmt.Println("Длина комнаты:", hall.Length)
	fmt.Println("Ширина комнаты:", hall.Width)
	fmt.Println("Количество окон:", hall.Windows)
	fmt.Println("Площадь комнаты:", hall.Length*hall.Width)
	fmt.Println("\tПредметы имеющиеся в основной комнате:")
	fmt.Println(appliances.PrintHallAppliances())
	fmt.Println(furniture.PrintHallFirnuture())
	fmt.Println(decor.PrintHallDecor())
	return Rooms{}
}
