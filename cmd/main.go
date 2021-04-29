package main

import (
	"log"
	"os"

	"jobTask/restapi"
	"jobTask/restapi/operations"

	apiPetitions "jobTask/restapi/operations/petitions"

	"github.com/go-openapi/loads"
	flags "github.com/jessevdk/go-flags"

	"jobTask/app"
)

func main() {
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	srv, err := app.New()
	if err != nil {
		log.Fatalln(err)
	}
	api := operations.NewPetitionManagerAPI(swaggerSpec)

	api.PetitionsPetitionsDeleteHandler = apiPetitions.PetitionsDeleteHandlerFunc(srv.PetitionsDeleteHandler)
	api.PetitionsPetitionsGetHandler = apiPetitions.PetitionsGetHandlerFunc(srv.PetitionsGetHandler)
	api.PetitionsPetitionsGetByIDHandler = apiPetitions.PetitionsGetByIDHandlerFunc(srv.PetitionsGetByIDHandler)
	api.PetitionsPetitionsPostHandler = apiPetitions.PetitionsPostHandlerFunc(srv.PetitionsPostHandler)
	api.PetitionsPetitionsUpdateHandler = apiPetitions.PetitionsUpdateHandlerFunc(srv.PetitionsUpdateHandler)
	api.ServerShutdown = srv.OnShutdown

	server := restapi.NewServer(api)
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "Petition manager"
	parser.LongDescription = "Petition manager"

	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
