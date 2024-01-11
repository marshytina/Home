package rooms

import (
	"fmt"
	"house/appliances"
	"house/dishes"
)

func PrintKitchen() Rooms {
	kitchen := Rooms{
		Name:    "Kitchen",
		Length:  7,
		Width:   10,
		Windows: 2,
	}

	fmt.Println("\tНазвание комнаты:", kitchen.Name)
	fmt.Println("Длина комнаты:", kitchen.Length)
	fmt.Println("Ширина комнаты:", kitchen.Width)
	fmt.Println("Kоличество окон:", kitchen.Windows)
	fmt.Println("Площадь комнаты:", kitchen.Length*kitchen.Width)
	fmt.Println("\tПредметы имеющиеся в кухне:")
	fmt.Println(appliances.PrintKitchenAppliances())
	fmt.Println(dishes.PrintKitchenDishes())
	return Rooms{}
}
