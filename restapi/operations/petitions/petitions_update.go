// Code generated by go-swagger; DO NOT EDIT.

package petitions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"jobTask/models"
)

// PetitionsUpdateHandlerFunc turns a function with the right signature into a petitions update handler
type PetitionsUpdateHandlerFunc func(PetitionsUpdateParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn PetitionsUpdateHandlerFunc) Handle(params PetitionsUpdateParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// PetitionsUpdateHandler interface for that can handle valid petitions update params
type PetitionsUpdateHandler interface {
	Handle(PetitionsUpdateParams, *models.Principal) middleware.Responder
}

// NewPetitionsUpdate creates a new http.Handler for the petitions update operation
func NewPetitionsUpdate(ctx *middleware.Context, handler PetitionsUpdateHandler) *PetitionsUpdate {
	return &PetitionsUpdate{Context: ctx, Handler: handler}
}

/* PetitionsUpdate swagger:route PUT /petitions/{petitionId} petitions petitionsUpdate

Update a petition

*/
type PetitionsUpdate struct {
	Context *middleware.Context
	Handler PetitionsUpdateHandler
}

func (o *PetitionsUpdate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPetitionsUpdateParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
