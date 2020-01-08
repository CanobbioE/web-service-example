package usecases

import "github.com/CanobbioE/web-service-example/domain"

// UserRepository provides an interface for the users persistency container.
type UserRepository interface {
	Store(user User) error
	FindById(id int) (User, error)
}

// User represents a use case's level entity.
// A user and an Adopter are basically the same entity but
// it is good practice to separate the two concepts.
type User struct {
	ID      int
	Adopter domain.Adopter
}
