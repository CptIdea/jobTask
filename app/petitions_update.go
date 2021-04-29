package app

import (
	"fmt"
	models "jobTask/models"
	apiPetitions "jobTask/restapi/operations/petitions"

	"github.com/go-openapi/runtime/middleware"
)

func (srv *Service) PetitionsUpdateHandler(params apiPetitions.PetitionsUpdateParams, principal *models.Principal) middleware.Responder {
	//Email validation
	validEmail, err := srv.emailValidator.validEmail(params.Body.AuthorEmail)
	if err != nil {
		return middleware.Error(500, fmt.Sprintf("email validation error: %s", err.Error()))
	}
	if !validEmail {
		return middleware.Error(400, "email validation failed")
	}

	petition, err := srv.petitionUpdate(params)
	if err != nil {
		return middleware.Error(500, fmt.Sprintf("database error: %s", err.Error()))
	}

	return apiPetitions.NewPetitionsPostOK().WithPayload(petition)
}
