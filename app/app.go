package app

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"jobTask/models"
	"os"
)

type Service struct {
	db *gorm.DB

	emailValidator emailValidator
}

func New() (*Service, error) {
	s := Service{}

	var err error

	dbPath, has := os.LookupEnv("DATABASE")
	if !has {
		return nil, fmt.Errorf("DATABASE string not found in env")
	}

	s.db, err = gorm.Open(postgres.Open(dbPath))
	if err != nil {
		return nil, err
	}

	s.emailValidator.Token, has = os.LookupEnv("EMAIL_VALIDATOR_TOKEN")
	if !has {
		return nil, fmt.Errorf("EMAIL_VALIDATOR_TOKEN string not found in env")
	}
	s.emailValidator.Secret, has = os.LookupEnv("EMAIL_VALIDATOR_SECRET")
	if !has {
		return nil, fmt.Errorf("EMAIL_VALIDATOR_TOKEN string not found in env")
	}

	// Попытка миграции
	migrate, has := os.LookupEnv("TRY_MIGRATE")
	if migrate == "TRUE" {
		err = s.AutoMigrate()
		if err != nil {
			return nil, err
		}
	}

	return &s, nil
}

func (srv *Service) OnShutdown() {}

func (srv *Service) AutoMigrate() error {
	return srv.db.AutoMigrate(&models.Petition{})
}
