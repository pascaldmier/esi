package corporation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/pascaldmier/esi/models"
)

// PutCorporationsCorporationIDStructuresStructureIDNoContentCode is the HTTP code returned for type PutCorporationsCorporationIDStructuresStructureIDNoContent
const PutCorporationsCorporationIDStructuresStructureIDNoContentCode int = 204

/*PutCorporationsCorporationIDStructuresStructureIDNoContent Structure vulnerability window updated

swagger:response putCorporationsCorporationIdStructuresStructureIdNoContent
*/
type PutCorporationsCorporationIDStructuresStructureIDNoContent struct {
}

// NewPutCorporationsCorporationIDStructuresStructureIDNoContent creates PutCorporationsCorporationIDStructuresStructureIDNoContent with default headers values
func NewPutCorporationsCorporationIDStructuresStructureIDNoContent() *PutCorporationsCorporationIDStructuresStructureIDNoContent {
	return &PutCorporationsCorporationIDStructuresStructureIDNoContent{}
}

// WriteResponse to the client
func (o *PutCorporationsCorporationIDStructuresStructureIDNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(204)
}

// PutCorporationsCorporationIDStructuresStructureIDForbiddenCode is the HTTP code returned for type PutCorporationsCorporationIDStructuresStructureIDForbidden
const PutCorporationsCorporationIDStructuresStructureIDForbiddenCode int = 403

/*PutCorporationsCorporationIDStructuresStructureIDForbidden Forbidden

swagger:response putCorporationsCorporationIdStructuresStructureIdForbidden
*/
type PutCorporationsCorporationIDStructuresStructureIDForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.PutCorporationsCorporationIDStructuresStructureIDForbiddenBody `json:"body,omitempty"`
}

// NewPutCorporationsCorporationIDStructuresStructureIDForbidden creates PutCorporationsCorporationIDStructuresStructureIDForbidden with default headers values
func NewPutCorporationsCorporationIDStructuresStructureIDForbidden() *PutCorporationsCorporationIDStructuresStructureIDForbidden {
	return &PutCorporationsCorporationIDStructuresStructureIDForbidden{}
}

// WithPayload adds the payload to the put corporations corporation Id structures structure Id forbidden response
func (o *PutCorporationsCorporationIDStructuresStructureIDForbidden) WithPayload(payload *models.PutCorporationsCorporationIDStructuresStructureIDForbiddenBody) *PutCorporationsCorporationIDStructuresStructureIDForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put corporations corporation Id structures structure Id forbidden response
func (o *PutCorporationsCorporationIDStructuresStructureIDForbidden) SetPayload(payload *models.PutCorporationsCorporationIDStructuresStructureIDForbiddenBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutCorporationsCorporationIDStructuresStructureIDForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutCorporationsCorporationIDStructuresStructureIDInternalServerErrorCode is the HTTP code returned for type PutCorporationsCorporationIDStructuresStructureIDInternalServerError
const PutCorporationsCorporationIDStructuresStructureIDInternalServerErrorCode int = 500

/*PutCorporationsCorporationIDStructuresStructureIDInternalServerError Internal server error

swagger:response putCorporationsCorporationIdStructuresStructureIdInternalServerError
*/
type PutCorporationsCorporationIDStructuresStructureIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.PutCorporationsCorporationIDStructuresStructureIDInternalServerErrorBody `json:"body,omitempty"`
}

// NewPutCorporationsCorporationIDStructuresStructureIDInternalServerError creates PutCorporationsCorporationIDStructuresStructureIDInternalServerError with default headers values
func NewPutCorporationsCorporationIDStructuresStructureIDInternalServerError() *PutCorporationsCorporationIDStructuresStructureIDInternalServerError {
	return &PutCorporationsCorporationIDStructuresStructureIDInternalServerError{}
}

// WithPayload adds the payload to the put corporations corporation Id structures structure Id internal server error response
func (o *PutCorporationsCorporationIDStructuresStructureIDInternalServerError) WithPayload(payload *models.PutCorporationsCorporationIDStructuresStructureIDInternalServerErrorBody) *PutCorporationsCorporationIDStructuresStructureIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put corporations corporation Id structures structure Id internal server error response
func (o *PutCorporationsCorporationIDStructuresStructureIDInternalServerError) SetPayload(payload *models.PutCorporationsCorporationIDStructuresStructureIDInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutCorporationsCorporationIDStructuresStructureIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}