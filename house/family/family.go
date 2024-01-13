package family

import (
	"fmt"
)

type Family struct {
	Name          string
	Surname       string
	Age           int
	MaritalStatus bool
	Profession    string
	Hobby         string
	FavouriteToy  string
	Who           string
	Partner       string
	Period        int
}
type Members struct {
	Members []Family
}

func (f Family) Marriage() {
	if f.MaritalStatus {
		fmt.Println(f.Name, "находится в браке с", f.Partner, "\nСрок отношений:", f.Period)
	} else {
		fmt.Println(f.Name, "одинок(a)")
	}
}
func PrintFamily() Family {
	fmt.Println("Члены семьи:")
	mom := Family{
		Who:           "MOM",
		Name:          "Eva",
		Surname:       "Koshkina",
		Age:           45,
		MaritalStatus: true,
		Profession:    "Teacher",
		Period:        16,
		Partner:       "Ivan Koshkin "}
	dad := Family{
		Who:           "DAD",
		Name:          "Ivan",
		Surname:       "Koshkin",
		Age:           49,
		MaritalStatus: true,
		Profession:    "Doctor",
		Period:        16,
		Partner:       "Eva Koshkina"}
	daughter := Family{
		Who:           "DAUGHTER",
		Name:          "Veronika",
		Surname:       "Koshkina",
		Age:           20,
		MaritalStatus: false,
		Profession:    "Singer",
		Partner:       "No",
		Period:        0}
	grandmother := Family{
		Who:           "GRANDMOTHER",
		Name:          "Olga",
		Surname:       "Petrova",
		Age:           86,
		MaritalStatus: true,
		Hobby:         "Knitting",
		Period:        59,
		Partner:       "Egor Petrov "}
	grandfather := Family{
		Who:           "GRANDFATHER",
		Name:          "Egor",
		Surname:       "Petrov",
		Age:           87,
		MaritalStatus: true,
		Hobby:         "Fishing",
		Period:        59,
		Partner:       "Olga Petrova"}
	cat := Family{
		Who:   "CAT",
		Name:  "Mem",
		Age:   4,
		Hobby: "Sleeping"}
	fmt.Println("\t\t\tЧлены семьи")
	fmt.Println("\t\t\t", grandmother.Who, "\nИмя:", grandfather.Name, "\nФамилия:", grandfather.Surname, "\nВозраст:", grandfather.Age, "\nХобби:", grandfather.Hobby)
	Family.Marriage(grandmother)
	fmt.Println("\t\t\t", grandfather.Who, "\nИмя:", grandmother.Name, "\nФамилия:", grandmother.Surname, "\nВозраст:", grandmother.Age, "\nХобби:", grandmother.Hobby)
	Family.Marriage(grandfather)
	fmt.Println("\t\t\t", mom.Who, "\nИмя:", mom.Name, "\nФамилия:", mom.Surname, "\nВозраст:", mom.Age, "\nПрофессия:", mom.Profession)
	Family.Marriage(mom)
	fmt.Println("\t\t\t", dad.Who, "\nИмя:", dad.Name, "\nФамилия:", dad.Surname, "\nВозраст:", dad.Age, "\nПрофессия:", dad.Profession)
	Family.Marriage(dad)
	fmt.Println("\t\t\t", daughter.Who, "\nИмя:", daughter.Name, "\nФамилия:", daughter.Surname, "\nВозраст:", daughter.Age, "\nПрофессия:", daughter.Profession)
	Family.Marriage(daughter)
	fmt.Println("\t\t\t", cat.Who, "\nИмя:", cat.Name, "\nВозраст:", cat.Age, "\nХобби:", cat.Hobby)
	Family.Marriage(cat)
	return Family{}
}
