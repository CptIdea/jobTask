package app

import (
	"fmt"
	models "jobTask/models"
	apiPetitions "jobTask/restapi/operations/petitions"

	"github.com/go-openapi/runtime/middleware"
)

func (srv *Service) PetitionsDeleteHandler(params apiPetitions.PetitionsDeleteParams, principal *models.Principal) middleware.Responder {
	err := srv.petitionDelete(params.PetitionID)
	if err != nil {
		return middleware.Error(500, fmt.Sprintf("database error: %s", err.Error()))
	}

	return apiPetitions.NewPetitionsDeleteOK().WithPayload(&models.Petition{})
}
