package interfaces

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/CanobbioE/web-service-example/domain"
)

// AdoptionInteractor define the usecases behavior.
// Using an interface here allows to change implementation
// as needed, for example mocking it during testing.
type AdoptionInteractor interface {
	AdoptedAnimals(userID int) ([]domain.Animal, error)
	Adopt(userID, animalID int) error
	AdoptableAnimals() ([]domain.Animal, error)
}

// WebserviceHandler represents the mechanism that transform
// HTTP requests to data that the usecases layer can comprehend.
type WebserviceHandler struct {
	AdoptionInteractor AdoptionInteractor
}

// AdoptAnimal receives a POST request to adopt an animal
// on behalf of the requesting client.
func (wh WebserviceHandler) AdoptAnimal(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.URL.Query().Get("user"))
	if err != nil {
		http.Error(w, fmt.Sprintf("\n%v, %v", err, r.URL.Query()), http.StatusBadRequest)
		return
	}

	animalID, err := strconv.Atoi(r.URL.Query().Get("animal"))
	if err != nil {
		http.Error(w, fmt.Sprintf("\n%v, %v", err, r.URL.Query()), http.StatusBadRequest)
		return
	}

	if err := wh.AdoptionInteractor.Adopt(userID, animalID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}

// ShowAdoptions receives a GET request to list all the animals
// adopted by the requesting user.
func (wh WebserviceHandler) ShowAdoptions(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.URL.Query().Get("user"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	adoptions, err := wh.AdoptionInteractor.AdoptedAnimals(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ret, err := json.Marshal(adoptions)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(ret)

	return
}

// ShowAnimals receives a GET request to list all the available animals,
func (wh WebserviceHandler) ShowAnimals(w http.ResponseWriter, r *http.Request) {
	animals, err := wh.AdoptionInteractor.AdoptableAnimals()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ret, err := json.Marshal(animals)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(ret)
}
