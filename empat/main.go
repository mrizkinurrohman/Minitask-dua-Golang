package empat

type Person struct {
	name    string
	address string
	phone   string
}

func NewPerson(name, address, phone string) Person {
	return Person{
		name:    name,
		address: address,
		phone:   phone,
	}
}

type Persons struct {
	Data     []Person
	IsSucces bool
	IsFailed bool
}

func (p *Person) GetPersonName() string {
	return p.name
}
func (p *Person) GetPersonAddress() string {
	return p.address
}
func (p *Person) GetPersonPhone() string {
	return p.phone
}

func (p *Person) GreetName(newName string) {
	p.name = newName
}

// func (p *Person) ChangeAddress(newAddress string) {
// 	p.Address = newAddress
// }
