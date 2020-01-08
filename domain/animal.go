package domain

// Animal represents the entity that can be adopted.
type Animal struct {
	ID     int
	Specie string
}

// AnimalRepository provides an interface for
// the animals persistency container.
type AnimalRepository interface {
	Store(animal Animal) error
	FindById(id int) (Animal, error)
	FindAll() ([]Animal, error)
}
