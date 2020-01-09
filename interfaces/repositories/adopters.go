package repositories

import (
	"fmt"

	"github.com/CanobbioE/web-service-example/domain"
)

// DbAdopterRepo is the repository for adopters.
type DbAdopterRepo DbRepo

// NewDbAdopterRepo istanciates and returns a user repository.
func NewDbAdopterRepo(dbHandlers map[string]DbHandler) *DbAdopterRepo {
	return &DbAdopterRepo{
		dbHandlers: dbHandlers,
		dbHandler:  dbHandlers["DbAdopterRepo"],
	}
}

// Store saves the new adopter into the repository
func (db *DbAdopterRepo) Store(adopter domain.Adopter) error {
	s := fmt.Sprintf("INSERT INTO adopters VALUES (%d, \"%s\")", adopter.ID, adopter.Name)
	return db.dbHandler.Execute(s)
}

// FindByID retrieves a user by the specified id.
func (db *DbAdopterRepo) FindByID(id int) (domain.Adopter, error) {
	var adopter domain.Adopter

	q := fmt.Sprintf("SELECT name FROM adopters WHERE id = %d", id)

	row, err := db.dbHandler.Query(q)
	if err != nil {
		return adopter, fmt.Errorf("can't find adopter with id %d:\n\t%v", id, err)
	}
	defer row.Close()

	var name string
	err = row.Scan(&name)
	if err != nil {
		return adopter, fmt.Errorf("can't find adopter with id %d:\n\t%v", id, err)
	}

	adopter.ID = id
	adopter.Name = name

	return adopter, nil
}
