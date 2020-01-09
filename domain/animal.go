package domain

// Animal represents the entity that can be adopted.
type Animal struct {
	ID          int
	Specie      string
	IsAdoptable bool
}

// AnimalRepository provides an interface for
// the animals persistency container.
type AnimalRepository interface {
	Store(animal Animal) error
	FindByID(id int) (Animal, error)
	FindAllAdoptable() ([]Animal, error)
}
