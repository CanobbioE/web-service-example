package repositories

import (
	"fmt"

	"github.com/CanobbioE/web-service-example/usecases"
)

// DbUserRepo is the repository for the users.
type DbUserRepo DbRepo

// NewDbUserRepo istanciates and returns a user repository.
func NewDbUserRepo(dbHandlers map[string]DbHandler) *DbUserRepo {
	return &DbUserRepo{
		dbHandlers: dbHandlers,
		dbHandler:  dbHandlers["DbUserRepo"],
	}
}

// Store saves the new user into the repository
func (db *DbUserRepo) Store(user usecases.User) error {
	s := fmt.Sprintf("INSERT INTO users VALUES (%d, %d)", user.ID, user.Adopter.ID)
	err := db.dbHandler.Execute(s)
	if err != nil {
		return err
	}

	return NewDbAdopterRepo(db.dbHandlers).Store(user.Adopter)
}

// FindByID retrieves a user given its ID.
func (db *DbUserRepo) FindByID(id int) (usecases.User, error) {
	var user usecases.User

	q := fmt.Sprintf("SELECT adopter_id FROM users WHERE id = %d", id)
	row, err := db.dbHandler.Query(q)
	if err != nil {
		return user, fmt.Errorf("can't find user with id %d:\n\t%v", id, err)
	}

	var adopterID int
	row.Next()
	err = row.Scan(&adopterID)
	if err != nil {
		return user, fmt.Errorf("can't find user with id %d:\n\t%v", id, err)
	}

	adopter, err := NewDbAdopterRepo(db.dbHandlers).FindByID(adopterID)
	if err != nil {
		return user, fmt.Errorf("can't find user with id %d:\n\t%v", id, err)
	}

	user.Adopter = adopter
	user.ID = id

	return user, nil

}
