package domain

// Adopter represents the entity that can adopt an Animal.
type Adopter struct {
	ID   int
	Name string
}

// AdopterRepository provides an interface for
// the adopters persistency container
type AdopterRepository interface {
	Store(user Adopter) error
	FindById(id int) (Adopter, error)
}
