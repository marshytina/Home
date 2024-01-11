package rooms

import (
	"fmt"
	"My house/Accessories"
)

type Rooms struct {
	Name    string
	Length  float32
	Width   float32
	Windows int
}

func PrintRoom() Rooms {
	bathroom := Rooms{
		Name:    "Bathroom",
		Length:  5,
		Width:   8,
		Windows: 1,
	}
	hall := Rooms{
		Name:    "Hall",
		Length:  9,
		Width:   12,
		Windows: 3,
	}
	kitchen := Rooms{
		Name:    "Kitchen",
		Length:  7,
		Width:   10,
		Windows: 2,
	}
	fmt.Println("Название комнаты:", bathroom.Name, "Длина комнаты:", bathroom.Length, "Ширина комнаты:", bathroom.Width, "Количество окон:", bathroom.Windows, "Площадь комнаты:", bathroom.Length*bathroom.Width)
	fmt.Println("Предметы имеющиеся в ванной:", PrintBathroomAcc
	fmt.Println("Название комнаты:", hall.Name, "Длина комнаты:", hall.Length, "Ширина комнаты:", hall.Width, "Количество окон:", hall.Windows, "Вес предмета:", "Площадь комнаты:", hall.Length*hall.Width)
	fmt.Println("Название комнаты:", kitchen.Name, "Длина комнаты:", kitchen.Length, "Ширина комнаты:", kitchen.Width, "Количество окон:", kitchen.Windows, "Площадь комнаты:", kitchen.Length*kitchen.Width)
	return Rooms{}
}
