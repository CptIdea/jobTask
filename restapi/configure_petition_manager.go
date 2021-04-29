// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"jobTask/models"
	"jobTask/restapi/operations"
	"jobTask/restapi/operations/petitions"
)

//go:generate swagger generate server --target ..\..\jobTask --name PetitionManager --spec ..\openAPI\swagger.yaml --template-dir .\templates --principal models.Principal

func configureFlags(api *operations.PetitionManagerAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

var headerToken = "123"

func configureAPI(api *operations.PetitionManagerAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "token" header is set

	api.HeaderTokenAuth = func(token string) (*models.Principal, error) {
		if token == headerToken {
			p := models.Principal(token)
			return &p, nil
		}
		return nil, fmt.Errorf("authorization failed")
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	if api.PetitionsPetitionsDeleteHandler == nil {
		api.PetitionsPetitionsDeleteHandler = petitions.PetitionsDeleteHandlerFunc(func(params petitions.PetitionsDeleteParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation petitions.PetitionsDelete has not yet been implemented")
		})
	}
	if api.PetitionsPetitionsGetHandler == nil {
		api.PetitionsPetitionsGetHandler = petitions.PetitionsGetHandlerFunc(func(params petitions.PetitionsGetParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation petitions.PetitionsGet has not yet been implemented")
		})
	}
	if api.PetitionsPetitionsGetByIDHandler == nil {
		api.PetitionsPetitionsGetByIDHandler = petitions.PetitionsGetByIDHandlerFunc(func(params petitions.PetitionsGetByIDParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation petitions.PetitionsGetByID has not yet been implemented")
		})
	}
	if api.PetitionsPetitionsPostHandler == nil {
		api.PetitionsPetitionsPostHandler = petitions.PetitionsPostHandlerFunc(func(params petitions.PetitionsPostParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation petitions.PetitionsPost has not yet been implemented")
		})
	}
	if api.PetitionsPetitionsUpdateHandler == nil {
		api.PetitionsPetitionsUpdateHandler = petitions.PetitionsUpdateHandlerFunc(func(params petitions.PetitionsUpdateParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation petitions.PetitionsUpdate has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
