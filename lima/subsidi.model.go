package lima

import "fmt"

type Subsidi struct {
	Model string
	Floor uint
	Color string
}

func (s Subsidi) GetHomeFullName() string {
	homeFullName := fmt.Sprintf("%s %d", s.Model, s.Floor)
	return homeFullName
}

func (s Subsidi) Jenis() {
	fmt.Println("Subsidi")
}


