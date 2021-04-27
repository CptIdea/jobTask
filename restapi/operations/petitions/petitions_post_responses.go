// Code generated by go-swagger; DO NOT EDIT.

package petitions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"jobTask/models"
)

// PetitionsPostOKCode is the HTTP code returned for type PetitionsPostOK
const PetitionsPostOKCode int = 200

/*PetitionsPostOK OK

swagger:response petitionsPostOK
*/
type PetitionsPostOK struct {

	/*
	  In: Body
	*/
	Payload *models.Petition `json:"body,omitempty"`
}

// NewPetitionsPostOK creates PetitionsPostOK with default headers values
func NewPetitionsPostOK() *PetitionsPostOK {

	return &PetitionsPostOK{}
}

// WithPayload adds the payload to the petitions post o k response
func (o *PetitionsPostOK) WithPayload(payload *models.Petition) *PetitionsPostOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the petitions post o k response
func (o *PetitionsPostOK) SetPayload(payload *models.Petition) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PetitionsPostOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
