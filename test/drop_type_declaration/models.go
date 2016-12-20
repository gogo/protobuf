package droptypedeclaration

type Dropped struct {
	Name string
	Age  int32
}

func (d *Dropped) Drop() bool {
	return true
}

func (d *Dropped) GetName() string {
	return d.Name
}

func (d *Dropped) GetAge() int32 {
	return d.Age
}
