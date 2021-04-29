package app

import (
	"fmt"
	models "jobTask/models"
	apiPetitions "jobTask/restapi/operations/petitions"

	"github.com/go-openapi/runtime/middleware"
)

func (srv *Service) PetitionsGetByIDHandler(params apiPetitions.PetitionsGetByIDParams, principal *models.Principal) middleware.Responder {
	petition, err := srv.petitionGetOne(params.PetitionID)
	if err != nil {
		return middleware.Error(500, fmt.Sprintf("database error: %s", err.Error()))
	}

	return apiPetitions.NewPetitionsGetByIDOK().WithPayload(petition)
}
