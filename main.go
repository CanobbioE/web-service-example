package main

import (
	"net/http"

	"github.com/CanobbioE/web-service-example/infrastructure"
	"github.com/CanobbioE/web-service-example/interfaces"
	"github.com/CanobbioE/web-service-example/interfaces/repositories"
	"github.com/CanobbioE/web-service-example/usecases"
)

func main() {

	dbHandler, err := infrastructure.NewSqliteHandler("")
	if err != nil {
		panic(err)
	}

	dbHandlers := make(map[string]repositories.DbHandler)
	for _, repo := range repositories.All {
		dbHandlers[repo] = dbHandler
	}

	adoptionInteractor := usecases.AdoptionInteractor{
		UserRepository:     repositories.NewDbUserRepo(dbHandlers),
		AdoptionRepository: repositories.NewDbAdoptionRepo(dbHandlers),
		AnimalRepository:   repositories.NewDbAnimalRepo(dbHandlers),
		Logger:             nil,
	}

	webserviceHandler := interfaces.WebserviceHandler{
		AdoptionInteractor: adoptionInteractor,
	}

	http.HandleFunc("/adopt", webserviceHandler.AdoptAnimal)
	http.HandleFunc("/adoptions", webserviceHandler.ShowAdoptions)
	http.HandleFunc("/adoptable", webserviceHandler.ShowAnimals)

	http.ListenAndServe(":8080", nil)

}
