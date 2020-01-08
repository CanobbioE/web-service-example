package repositories

import (
	"fmt"

	"github.com/CanobbioE/web-service-example/domain"
)

// DbAnimalRepo is the repository for animals.
type DbAnimalRepo DbRepo

// NewDbAnimalRepo istanciates and returns a user repository.
func NewDbAnimalRepo(dbHandlers map[string]DbHandler) *DbAnimalRepo {
	return &DbAnimalRepo{
		dbHandlers: dbHandlers,
		dbHandler:  dbHandlers["DbAnimalRepo"],
	}
}

// Store saves the new animal into the repository.
func (db *DbAnimalRepo) Store(animal domain.Animal) error {
	s := fmt.Sprintf("INSERT INTO animals VALUES (%d, \"%s\")", animal.ID, animal.Specie)
	return db.dbHandler.Execute(s)
}

// FindByID retrieves an animals given its ID.
func (db *DbAnimalRepo) FindByID(id int) (domain.Animal, error) {
	var animal domain.Animal
	q := fmt.Sprintf("SELECT specie FROM animals WHERE id = %d", id)

	row, err := db.dbHandler.Query(q)
	if err != nil {
		return animal, fmt.Errorf("can't find animal with id %d:\n\t%v", id, err)
	}

	var specie string
	err = row.Scan(&specie)
	if err != nil {
		return animal, fmt.Errorf("can't find animal with id %d:\n\t%v", id, err)
	}

	animal.ID = id
	animal.Specie = specie
	return animal, nil
}

// FindAll returns a list of all the animals in the repository.
func (db *DbAnimalRepo) FindAll() ([]domain.Animal, error) {
	var animals []domain.Animal
	q := `SELECT id, specie FROM animals`

	row, err := db.dbHandler.Query(q)
	if err != nil {
		return animals, fmt.Errorf("can't find all animals: %v", err)
	}
	for row.Next() {
		var id int
		var specie string
		row.Scan(&id, &specie)
		animals = append(animals, domain.Animal{id, specie})
	}
	return animals, nil
}
