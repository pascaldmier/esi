package fleets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/pascaldmier/esi/models"
)

// PutFleetsFleetIDMembersMemberIDNoContentCode is the HTTP code returned for type PutFleetsFleetIDMembersMemberIDNoContent
const PutFleetsFleetIDMembersMemberIDNoContentCode int = 204

/*PutFleetsFleetIDMembersMemberIDNoContent Fleet invitation sent

swagger:response putFleetsFleetIdMembersMemberIdNoContent
*/
type PutFleetsFleetIDMembersMemberIDNoContent struct {
}

// NewPutFleetsFleetIDMembersMemberIDNoContent creates PutFleetsFleetIDMembersMemberIDNoContent with default headers values
func NewPutFleetsFleetIDMembersMemberIDNoContent() *PutFleetsFleetIDMembersMemberIDNoContent {
	return &PutFleetsFleetIDMembersMemberIDNoContent{}
}

// WriteResponse to the client
func (o *PutFleetsFleetIDMembersMemberIDNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(204)
}

// PutFleetsFleetIDMembersMemberIDForbiddenCode is the HTTP code returned for type PutFleetsFleetIDMembersMemberIDForbidden
const PutFleetsFleetIDMembersMemberIDForbiddenCode int = 403

/*PutFleetsFleetIDMembersMemberIDForbidden Forbidden

swagger:response putFleetsFleetIdMembersMemberIdForbidden
*/
type PutFleetsFleetIDMembersMemberIDForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.PutFleetsFleetIDMembersMemberIDForbiddenBody `json:"body,omitempty"`
}

// NewPutFleetsFleetIDMembersMemberIDForbidden creates PutFleetsFleetIDMembersMemberIDForbidden with default headers values
func NewPutFleetsFleetIDMembersMemberIDForbidden() *PutFleetsFleetIDMembersMemberIDForbidden {
	return &PutFleetsFleetIDMembersMemberIDForbidden{}
}

// WithPayload adds the payload to the put fleets fleet Id members member Id forbidden response
func (o *PutFleetsFleetIDMembersMemberIDForbidden) WithPayload(payload *models.PutFleetsFleetIDMembersMemberIDForbiddenBody) *PutFleetsFleetIDMembersMemberIDForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put fleets fleet Id members member Id forbidden response
func (o *PutFleetsFleetIDMembersMemberIDForbidden) SetPayload(payload *models.PutFleetsFleetIDMembersMemberIDForbiddenBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutFleetsFleetIDMembersMemberIDForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutFleetsFleetIDMembersMemberIDNotFoundCode is the HTTP code returned for type PutFleetsFleetIDMembersMemberIDNotFound
const PutFleetsFleetIDMembersMemberIDNotFoundCode int = 404

/*PutFleetsFleetIDMembersMemberIDNotFound The fleet does not exist or you don't have access to it

swagger:response putFleetsFleetIdMembersMemberIdNotFound
*/
type PutFleetsFleetIDMembersMemberIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.PutFleetsFleetIDMembersMemberIDNotFoundBody `json:"body,omitempty"`
}

// NewPutFleetsFleetIDMembersMemberIDNotFound creates PutFleetsFleetIDMembersMemberIDNotFound with default headers values
func NewPutFleetsFleetIDMembersMemberIDNotFound() *PutFleetsFleetIDMembersMemberIDNotFound {
	return &PutFleetsFleetIDMembersMemberIDNotFound{}
}

// WithPayload adds the payload to the put fleets fleet Id members member Id not found response
func (o *PutFleetsFleetIDMembersMemberIDNotFound) WithPayload(payload *models.PutFleetsFleetIDMembersMemberIDNotFoundBody) *PutFleetsFleetIDMembersMemberIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put fleets fleet Id members member Id not found response
func (o *PutFleetsFleetIDMembersMemberIDNotFound) SetPayload(payload *models.PutFleetsFleetIDMembersMemberIDNotFoundBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutFleetsFleetIDMembersMemberIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutFleetsFleetIDMembersMemberIDUnprocessableEntityCode is the HTTP code returned for type PutFleetsFleetIDMembersMemberIDUnprocessableEntity
const PutFleetsFleetIDMembersMemberIDUnprocessableEntityCode int = 422

/*PutFleetsFleetIDMembersMemberIDUnprocessableEntity Errors in invitation

swagger:response putFleetsFleetIdMembersMemberIdUnprocessableEntity
*/
type PutFleetsFleetIDMembersMemberIDUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.PutFleetsFleetIDMembersMemberIDUnprocessableEntityBody `json:"body,omitempty"`
}

// NewPutFleetsFleetIDMembersMemberIDUnprocessableEntity creates PutFleetsFleetIDMembersMemberIDUnprocessableEntity with default headers values
func NewPutFleetsFleetIDMembersMemberIDUnprocessableEntity() *PutFleetsFleetIDMembersMemberIDUnprocessableEntity {
	return &PutFleetsFleetIDMembersMemberIDUnprocessableEntity{}
}

// WithPayload adds the payload to the put fleets fleet Id members member Id unprocessable entity response
func (o *PutFleetsFleetIDMembersMemberIDUnprocessableEntity) WithPayload(payload *models.PutFleetsFleetIDMembersMemberIDUnprocessableEntityBody) *PutFleetsFleetIDMembersMemberIDUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put fleets fleet Id members member Id unprocessable entity response
func (o *PutFleetsFleetIDMembersMemberIDUnprocessableEntity) SetPayload(payload *models.PutFleetsFleetIDMembersMemberIDUnprocessableEntityBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutFleetsFleetIDMembersMemberIDUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutFleetsFleetIDMembersMemberIDInternalServerErrorCode is the HTTP code returned for type PutFleetsFleetIDMembersMemberIDInternalServerError
const PutFleetsFleetIDMembersMemberIDInternalServerErrorCode int = 500

/*PutFleetsFleetIDMembersMemberIDInternalServerError Internal server error

swagger:response putFleetsFleetIdMembersMemberIdInternalServerError
*/
type PutFleetsFleetIDMembersMemberIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.PutFleetsFleetIDMembersMemberIDInternalServerErrorBody `json:"body,omitempty"`
}

// NewPutFleetsFleetIDMembersMemberIDInternalServerError creates PutFleetsFleetIDMembersMemberIDInternalServerError with default headers values
func NewPutFleetsFleetIDMembersMemberIDInternalServerError() *PutFleetsFleetIDMembersMemberIDInternalServerError {
	return &PutFleetsFleetIDMembersMemberIDInternalServerError{}
}

// WithPayload adds the payload to the put fleets fleet Id members member Id internal server error response
func (o *PutFleetsFleetIDMembersMemberIDInternalServerError) WithPayload(payload *models.PutFleetsFleetIDMembersMemberIDInternalServerErrorBody) *PutFleetsFleetIDMembersMemberIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put fleets fleet Id members member Id internal server error response
func (o *PutFleetsFleetIDMembersMemberIDInternalServerError) SetPayload(payload *models.PutFleetsFleetIDMembersMemberIDInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutFleetsFleetIDMembersMemberIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}