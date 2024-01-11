package rooms

type Rooms struct {
	Name    string
	Length  float32
	Width   float32
	Windows int
	Square  int
}

func PrintRoom() Rooms {
	PrintBathroom()
	PrintHall()
	PrintKitchen()
	return Rooms{}
}
