package app

import (
	"jobTask/models"
	apiPetitions "jobTask/restapi/operations/petitions"
)

func (srv *Service) petitionGetOne(ID int64) (*models.Petition, error) {
	p := models.Petition{ID: ID}
	err := srv.db.Find(&p).Error
	return &p, err
}

func (srv *Service) petitionDelete(ID int64) error {
	p := models.Petition{ID: ID}
	err := srv.db.Delete(&p).Error
	return err
}

func (srv *Service) petitionGetAll() ([]*models.Petition, error) {
	var p []*models.Petition
	err := srv.db.Find(&p).Error
	return p, err
}

func (srv *Service) petitionCreate(petition apiPetitions.PetitionsPostParams) (*models.Petition, error) {
	p := models.Petition{
		AuthorEmail: petition.Body.AuthorEmail,
		Body:        petition.Body.Body,
		Name:        petition.Body.Name,
	}
	err := srv.db.Create(&p).Error
	return &p, err
}

func (srv *Service) petitionUpdate(petition apiPetitions.PetitionsUpdateParams) (*models.Petition, error) {
	p := models.Petition{
		AuthorEmail: petition.Body.AuthorEmail,
		Body:        petition.Body.Body,
		Name:        petition.Body.Name,
	}
	err := srv.db.Create(&p).Error
	return &p, err
}
