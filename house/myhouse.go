package house

import (
	"fmt"
	"house/house/appliances"
	"house/house/dishes"
	"house/house/family"
	"house/house/rooms"
)

type House struct {
	Address    string
	Square     int
	Floors     int
	Rooms      rooms.Rooms
	Family     family.Family
	Dishes     dishes.Dishes
	Appliances appliances.Appliances
}

func CreateHouse() House {
	house := House{
		Address: "ул. Ленинградская д.157",
		Floors:  2,
		Square:  250,
	}
	fmt.Println("Дом находится по адресу:", house.Address)
	fmt.Println("Количество этажей:", house.Floors)
	fmt.Println(family.PrintFamily())
	fmt.Println("Общая характеристика комнат:")
	fmt.Println(rooms.PrintRoom())
	return house
}
