package app

import (
	"fmt"
	models "jobTask/models"
	apiPetitions "jobTask/restapi/operations/petitions"

	"github.com/go-openapi/runtime/middleware"
)

func (srv *Service) PetitionsGetHandler(params apiPetitions.PetitionsGetParams, principal *models.Principal) middleware.Responder {
	petitions, err := srv.petitionGetAll()
	if err != nil {
		return middleware.Error(500, fmt.Sprintf("database error: %s", err.Error()))
	}

	return apiPetitions.NewPetitionsGetOK().WithPayload(petitions)
}
