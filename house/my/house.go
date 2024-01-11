package my

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
	fmt.Println("Дом находится по адресу:", house.Address)
	fmt.Println("Количество этажей:", house.Floors)
	fmt.Println(family.PrintFamily())
	fmt.Println("Общая характеристика комнат:")
	fmt.Println(rooms.PrintRoom())
	return house
}
