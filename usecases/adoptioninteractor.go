package usecases

import (
	"fmt"

	"github.com/CanobbioE/web-service-example/domain"
)

// Logger is used for logging. Using an interface allows
// us to hook up anything as a logger.
type Logger interface {
	Log(message string)
}

// AdoptionInteractor collects all the externally injected repositories.
// The idea is that this package doesn't depend on anything outside
// the "domain" or the "usecases" packages.
type AdoptionInteractor struct {
	UserRepository     UserRepository
	AdoptionRepository domain.AdoptionRepository
	AnimalRepository   domain.AnimalRepository
	Logger             Logger
}

// Adopt registers the specified user as the adopter for the specified animal.
func (ai AdoptionInteractor) Adopt(userID, animalID int) error {
	animal, err := ai.AnimalRepository.FindByID(animalID)
	if err != nil {
		ai.Logger.Log(err.Error())
		return err
	}

	user, err := ai.UserRepository.FindByID(userID)
	if err != nil {
		ai.Logger.Log(err.Error())
		return err
	}

	if !animal.IsAdoptable {
		err := fmt.Errorf("tried to adopt not available animal, id: %d", animalID)
		ai.Logger.Log(err.Error())
		return err
	}

	maxID, err := ai.AdoptionRepository.FindMaxID()
	if err != nil {
		ai.Logger.Log(err.Error())
		return err
	}
	adoption := domain.Adoption{
		ID:      maxID + 1,
		Adopter: user.Adopter,
		Animal:  animal,
	}

	err = ai.AdoptionRepository.Store(adoption)
	if err != nil {
		ai.Logger.Log(err.Error())
		return err
	}

	ai.Logger.Log(fmt.Sprintf("User %d has adopted animal %d", userID, animalID))

	return nil
}

// AdoptedAnimals lists all the animals that have been adopted
// by the specified user.
func (ai AdoptionInteractor) AdoptedAnimals(userID int) ([]domain.Animal, error) {
	var animals []domain.Animal

	user, err := ai.UserRepository.FindByID(userID)
	if err != nil {
		ai.Logger.Log(err.Error())
		return animals, err
	}

	adoptions, err := ai.AdoptionRepository.FindAllByAdopterID(user.Adopter.ID)
	if err != nil {
		ai.Logger.Log(err.Error())
		return animals, err
	}

	for _, adoption := range adoptions {
		animals = append(animals, adoption.Animal)
	}

	ai.Logger.Log(fmt.Sprintf("Listed all adopted animals for user %d", userID))

	return animals, nil
}

// AdoptableAnimals lists all the animals that can be adopted.
func (ai AdoptionInteractor) AdoptableAnimals() (animals []domain.Animal, err error) {
	animals, err = ai.AnimalRepository.FindAllAdoptable()
	if err != nil {
		ai.Logger.Log(err.Error())
		return
	}

	ai.Logger.Log(`Listed all adoptable animals`)
	return animals, nil
}
