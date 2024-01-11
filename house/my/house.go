package My_house

import (
	"fmt"
	"house/family"
	"house/rooms"
)

type House struct {
	Address string
	Square  int
	Floors  int
	Rooms   rooms.Rooms
	Family  family.Family
}

func CreateHouse() House {
	house := House{
		Address: "ул. Ленинградская д.157",
		Floors:  2,
	}
	fmt.Println("Дом находится по адресу:", house.Address, "Количество этажей:", house.Floors, "Общая характеристика комнат:\n", rooms.PrintRoom(), "Жильцы дома:\n", family.PrintFamily())
	return house
}
