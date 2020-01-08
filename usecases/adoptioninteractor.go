package usecases

import "github.com/CanobbioE/web-service-example/domain"

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
		return err
	}

	user, err := ai.UserRepository.FindByID(userID)
	if err != nil {
		return err
	}

	adoption := domain.Adoption{
		Adopter: user.Adopter,
		Animal:  animal,
	}

	err = ai.AdoptionRepository.Store(adoption)
	if err != nil {
		return err
	}

	return nil
}

// AdoptedAnimals lists all the animals that have been adopted
// by the specified user.
func (ai AdoptionInteractor) AdoptedAnimals(userID int) ([]domain.Animal, error) {
	var animals []domain.Animal

	user, err := ai.UserRepository.FindByID(userID)
	if err != nil {
		return animals, err
	}

	adoptions, err := ai.AdoptionRepository.FindAllByAdopterID(user.Adopter.ID)
	if err != nil {
		return animals, err
	}

	for _, adoption := range adoptions {
		animals = append(animals, adoption.Animal)
	}

	return animals, nil
}

// AdoptableAnimals lists all the animals that can be adopted.
func (ai AdoptionInteractor) AdoptableAnimals() (animals []domain.Animal, err error) {
	animals, err = ai.AnimalRepository.FindAll()
	if err != nil {
		return
	}
	return animals, nil
}
