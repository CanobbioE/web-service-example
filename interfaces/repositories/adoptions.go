package repositories

import (
	"fmt"

	"github.com/CanobbioE/web-service-example/domain"
)

// DbAdoptionRepo is the repository for adoptions.
type DbAdoptionRepo DbRepo

// NewDbAdoptionRepo istanciates and returns a user repository.
func NewDbAdoptionRepo(dbHandlers map[string]DbHandler) *DbAdoptionRepo {
	return &DbAdoptionRepo{
		dbHandlers: dbHandlers,
		dbHandler:  dbHandlers["DbAdoptionRepo"],
	}
}

// Store saves the new adoption into the repository.
func (db *DbAdoptionRepo) Store(adoption domain.Adoption) error {
	s := fmt.Sprintf("INSERT INTO adoptions VALUES (%d, %d, %d)",
		adoption.ID, adoption.Adopter.ID, adoption.Animal.ID)
	return db.dbHandler.Execute(s)
}

// FindByID retrieves an adoptions given its ID
func (db *DbAdoptionRepo) FindByID(id int) (domain.Adoption, error) {
	var adoption domain.Adoption

	q := fmt.Sprintf("SELECT adopter_id, animal_id FROM adoptions WHERE id = %d", id)

	row, err := db.dbHandler.Query(q)
	if err != nil {
		return adoption, fmt.Errorf("can't find adoption with id %d:\n\t%v", id, err)
	}
	defer row.Close()

	row.Next()
	var adopterID, animalID int
	err = row.Scan(&adopterID, &animalID)
	if err != nil {
		return adoption, fmt.Errorf("can't find adoption with id %d:\n\t%v", id, err)
	}

	animal, err := NewDbAnimalRepo(db.dbHandlers).FindByID(animalID)
	if err != nil {
		return adoption, fmt.Errorf("can't find adoption with id %d:\n\t%v", id, err)
	}

	adopter, err := NewDbAdopterRepo(db.dbHandlers).FindByID(adopterID)
	if err != nil {
		return adoption, fmt.Errorf("can't find adoption with id %d:\n\t%v", id, err)
	}
	adoption = domain.Adoption{
		Adopter: adopter,
		Animal:  animal,
		ID:      id,
	}

	return adoption, nil
}

// FindAllByAdopterID returns a list of all the adoptions related
// to the specified user ID.
func (db *DbAdoptionRepo) FindAllByAdopterID(id int) ([]domain.Adoption, error) {
	var adoptions []domain.Adoption

	q := fmt.Sprintf("SELECT id FROM adoptions WHERE adopter_id = %d", id)

	row, err := db.dbHandler.Query(q)
	if err != nil {
		return adoptions, fmt.Errorf("can't create list of adoptions for adopters %d:\n\t%v", id, err)
	}
	defer row.Close()

	for row.Next() {
		var adoptionID int
		row.Scan(&adoptionID)
		adoption, err := db.FindByID(adoptionID)
		if err != nil {
			return adoptions, fmt.Errorf("can't create list of adoptions for adopters %d:\n\t%v", id, err)
		}

		adoptions = append(adoptions, adoption)
	}

	return adoptions, nil
}

// FindMaxID retrieves the last inserted ID.
// This is a bad practice, we are using it here just to enable the project
// to run as a demo.
func (db *DbAdoptionRepo) FindMaxID() (id int, err error) {
	q := `SELECT MAX(id) FROM adoptions`

	row, err := db.dbHandler.Query(q)
	if err != nil {
		return
	}
	defer row.Close()

	row.Next()
	err = row.Scan(&id)
	if err != nil {
		return
	}

	return id, nil
}
