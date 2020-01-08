package domain

// Adoption represents the relationship between
// an Animal and its Adopter.
type Adoption struct {
	ID      int
	Adopter Adopter
	Animal  Animal
}

// AdoptionRepository provides an interface for
// the adoptions persistency container.
type AdoptionRepository interface {
	Store(adoption Adoption) error
	FindById(id int) (Adoption, error)
	FindAllByAdopterID(id int) ([]Adoption, error)
}
