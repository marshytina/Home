package family

import "fmt"

type Family struct {
	Name          string
	Surname       string
	Age           int
	MaritalStatus string
	Profession    string
	Hobby         string
	FavouriteToy  string
	Who           string
}

func PrintFamily() Family {
	fmt.Println("Члены семьи:")
	mom := Family{
		Who:           "MOM",
		Name:          "Eva",
		Surname:       "Koshkina",
		Age:           45,
		MaritalStatus: "Married",
		Profession:    "Teacher",
	}
	dad := Family{
		Who:           "DAD",
		Name:          "Ivan",
		Surname:       "Koshkin",
		Age:           49,
		MaritalStatus: "Married",
		Profession:    "Doctor",
	}
	son := Family{
		Who:           "SON",
		Name:          "Oleg",
		Surname:       "Koshkin",
		Age:           23,
		MaritalStatus: "Married",
		Profession:    "Designer",
	}
	daughter := Family{
		Who:           "DAUGHTER",
		Name:          "Veronika",
		Surname:       "Koshkina",
		Age:           20,
		MaritalStatus: "Single",
		Profession:    "Singer",
	}
	grandmother := Family{
		Who:           "GRANDMOTHER",
		Name:          "Olga",
		Surname:       "Petrova",
		Age:           86,
		MaritalStatus: "Married",
		Hobby:         "Knitting",
	}
	grandfather := Family{
		Who:           "GRANDFATHER",
		Name:          "Egor",
		Surname:       "Petrov",
		Age:           87,
		MaritalStatus: "Married",
		Hobby:         "Fishing",
	}
	cat := Family{
		Who:   "CAT",
		Name:  "Mem",
		Age:   4,
		Hobby: "Sleeping",
	}
	fmt.Println("\t\t\tЧлены семьи")
	fmt.Println("\t\t\t", grandmother.Who, "\nИмя:", grandfather.Name, "\nФамилия:", grandfather.Surname, "\nВозраст:", grandfather.Age, "\nЛичный статус:", grandfather.MaritalStatus, "\nХобби:", grandfather.Hobby)
	fmt.Println("\t\t\t", grandfather.Who, "\nИмя:", grandmother.Name, "\nФамилия:", grandmother.Surname, "\nВозраст:", grandmother.Age, "\nЛичный статус:", grandmother.MaritalStatus, "\nХобби:", grandmother.Hobby)
	fmt.Println("\t\t\t", mom.Who, "\nИмя:", mom.Name, "\nФамилия:", mom.Surname, "\nВозраст:", mom.Age, "\nЛичный статус:", mom.MaritalStatus, "\nПрофессия:", mom.Profession)
	fmt.Println("\t\t\t", dad.Who, "\nИмя:", dad.Name, "\nФамилия:", dad.Surname, "\nВозраст:", dad.Age, "\nЛичный статус:", dad.MaritalStatus, "\nПрофессия:", dad.Profession)
	fmt.Println("\t\t\t", son.Who, "\nИмя:", son.Name, "\nФамилия:", son.Surname, "\nВозраст:", son.Age, "\nЛичный статус:", son.MaritalStatus, "\nПрофессия:", son.Profession)
	fmt.Println("\t\t\t", daughter.Who, "\nИмя:", daughter.Name, "\nФамилия:", daughter.Surname, "\nВозраст:", daughter.Age, "\nЛичный статус:", daughter.MaritalStatus, "\nПрофессия:", daughter.Profession)
	fmt.Println("\t\t\t", cat.Who, "\nИмя:", cat.Name, "\nВозраст:", cat.Age, "\nХобби:", cat.Hobby)
	return Family{}
}
