package lima

import "fmt"

type Home struct {
	Model string
	Floor uint
	Color string
}

func (h Home) GetHomeFullName() string {
	homeFullName := fmt.Sprintf("%s %d", h.Model, h.Floor)
	return homeFullName
}

func (h Home) Jenis() {
	fmt.Println("Cluster")
}


