package main

import "fmt"

type Family struct {
	Name          string
	Surname       string
	Age           int
	MaritalStatus string
	Profession    string
	Hobby         string
	FavouriteToy  string
}

func PrintFamily() Family {
	fmt.Println("Члены семьи:")
	mom := Family{
		Name:          "Eva",
		Surname:       "Koshkina",
		Age:           45,
		MaritalStatus: "Married",
		Profession:    "Teacher",
	}
	dad := Family{
		Name:          "Ivan",
		Surname:       "Koshkin",
		Age:           49,
		MaritalStatus: "Married",
		Profession:    "Doctor",
	}
	son := Family{
		Name:          "Oleg",
		Surname:       "Koshkin",
		Age:           23,
		MaritalStatus: "Married",
		Profession:    "Designer",
	}
	daughter := Family{
		Name:          "Veronika",
		Surname:       "Koshkina",
		Age:           20,
		MaritalStatus: "Single",
		Profession:    "Singer",
	}
	grandmother := Family{
		Name:          "Olga",
		Surname:       "Petrova",
		Age:           86,
		MaritalStatus: "Married",
		Hobby:         "Knitting",
	}
	grandfather := Family{
		Name:          "Egor",
		Surname:       "Petrov",
		Age:           87,
		MaritalStatus: "Married",
		Hobby:         "Fishing",
	}
	cat := Family{
		Name:  "Mem",
		Age:   4,
		Hobby: "Sleeping",
	}
	fmt.Println("Имя:", grandfather.Name, "Фамилия:", grandfather.Surname, "Возраст:", grandfather.Age, "Личный статус:", grandfather.MaritalStatus, "Хобби:", grandfather.Hobby)
	fmt.Println("Имя:", grandmother.Name, "Фамилия:", grandmother.Surname, "Возраст:", grandmother.Age, "Личный статус:", grandmother.MaritalStatus, "Хобби:", grandmother.Hobby)
	fmt.Println("Имя:", mom.Name, "Фамилия:", mom.Surname, "Возраст:", mom.Age, "Личный статус:", mom.MaritalStatus, "Профессия:", mom.Profession)
	fmt.Println("Имя:", dad.Name, "Фамилия:", dad.Surname, "Возраст:", dad.Age, "Личный статус:", dad.MaritalStatus, "Профессия:", dad.Profession)
	fmt.Println("Имя:", son.Name, "Фамилия:", son.Surname, "Возраст:", son.Age, "Личный статус:", son.MaritalStatus, "Профессия:", son.Profession)
	fmt.Println("Имя:", daughter.Name, "Фамилия:", daughter.Surname, "Возраст:", daughter.Age, "Личный статус:", daughter.MaritalStatus, "Профессия:", daughter.Profession)
	fmt.Println("Имя:", cat.Name, "Возраст:", cat.Age, "Хобби:", cat.Hobby)
	return Family{}
}
