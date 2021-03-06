package fleets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/pascaldmier/esi/models"
)

// DeleteFleetsFleetIDSquadsSquadIDNoContentCode is the HTTP code returned for type DeleteFleetsFleetIDSquadsSquadIDNoContent
const DeleteFleetsFleetIDSquadsSquadIDNoContentCode int = 204

/*DeleteFleetsFleetIDSquadsSquadIDNoContent Squad deleted

swagger:response deleteFleetsFleetIdSquadsSquadIdNoContent
*/
type DeleteFleetsFleetIDSquadsSquadIDNoContent struct {
}

// NewDeleteFleetsFleetIDSquadsSquadIDNoContent creates DeleteFleetsFleetIDSquadsSquadIDNoContent with default headers values
func NewDeleteFleetsFleetIDSquadsSquadIDNoContent() *DeleteFleetsFleetIDSquadsSquadIDNoContent {
	return &DeleteFleetsFleetIDSquadsSquadIDNoContent{}
}

// WriteResponse to the client
func (o *DeleteFleetsFleetIDSquadsSquadIDNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(204)
}

// DeleteFleetsFleetIDSquadsSquadIDForbiddenCode is the HTTP code returned for type DeleteFleetsFleetIDSquadsSquadIDForbidden
const DeleteFleetsFleetIDSquadsSquadIDForbiddenCode int = 403

/*DeleteFleetsFleetIDSquadsSquadIDForbidden Forbidden

swagger:response deleteFleetsFleetIdSquadsSquadIdForbidden
*/
type DeleteFleetsFleetIDSquadsSquadIDForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.DeleteFleetsFleetIDSquadsSquadIDForbiddenBody `json:"body,omitempty"`
}

// NewDeleteFleetsFleetIDSquadsSquadIDForbidden creates DeleteFleetsFleetIDSquadsSquadIDForbidden with default headers values
func NewDeleteFleetsFleetIDSquadsSquadIDForbidden() *DeleteFleetsFleetIDSquadsSquadIDForbidden {
	return &DeleteFleetsFleetIDSquadsSquadIDForbidden{}
}

// WithPayload adds the payload to the delete fleets fleet Id squads squad Id forbidden response
func (o *DeleteFleetsFleetIDSquadsSquadIDForbidden) WithPayload(payload *models.DeleteFleetsFleetIDSquadsSquadIDForbiddenBody) *DeleteFleetsFleetIDSquadsSquadIDForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete fleets fleet Id squads squad Id forbidden response
func (o *DeleteFleetsFleetIDSquadsSquadIDForbidden) SetPayload(payload *models.DeleteFleetsFleetIDSquadsSquadIDForbiddenBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteFleetsFleetIDSquadsSquadIDForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteFleetsFleetIDSquadsSquadIDNotFoundCode is the HTTP code returned for type DeleteFleetsFleetIDSquadsSquadIDNotFound
const DeleteFleetsFleetIDSquadsSquadIDNotFoundCode int = 404

/*DeleteFleetsFleetIDSquadsSquadIDNotFound The fleet does not exist or you don't have access to it

swagger:response deleteFleetsFleetIdSquadsSquadIdNotFound
*/
type DeleteFleetsFleetIDSquadsSquadIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.DeleteFleetsFleetIDSquadsSquadIDNotFoundBody `json:"body,omitempty"`
}

// NewDeleteFleetsFleetIDSquadsSquadIDNotFound creates DeleteFleetsFleetIDSquadsSquadIDNotFound with default headers values
func NewDeleteFleetsFleetIDSquadsSquadIDNotFound() *DeleteFleetsFleetIDSquadsSquadIDNotFound {
	return &DeleteFleetsFleetIDSquadsSquadIDNotFound{}
}

// WithPayload adds the payload to the delete fleets fleet Id squads squad Id not found response
func (o *DeleteFleetsFleetIDSquadsSquadIDNotFound) WithPayload(payload *models.DeleteFleetsFleetIDSquadsSquadIDNotFoundBody) *DeleteFleetsFleetIDSquadsSquadIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete fleets fleet Id squads squad Id not found response
func (o *DeleteFleetsFleetIDSquadsSquadIDNotFound) SetPayload(payload *models.DeleteFleetsFleetIDSquadsSquadIDNotFoundBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteFleetsFleetIDSquadsSquadIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteFleetsFleetIDSquadsSquadIDInternalServerErrorCode is the HTTP code returned for type DeleteFleetsFleetIDSquadsSquadIDInternalServerError
const DeleteFleetsFleetIDSquadsSquadIDInternalServerErrorCode int = 500

/*DeleteFleetsFleetIDSquadsSquadIDInternalServerError Internal server error

swagger:response deleteFleetsFleetIdSquadsSquadIdInternalServerError
*/
type DeleteFleetsFleetIDSquadsSquadIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.DeleteFleetsFleetIDSquadsSquadIDInternalServerErrorBody `json:"body,omitempty"`
}

// NewDeleteFleetsFleetIDSquadsSquadIDInternalServerError creates DeleteFleetsFleetIDSquadsSquadIDInternalServerError with default headers values
func NewDeleteFleetsFleetIDSquadsSquadIDInternalServerError() *DeleteFleetsFleetIDSquadsSquadIDInternalServerError {
	return &DeleteFleetsFleetIDSquadsSquadIDInternalServerError{}
}

// WithPayload adds the payload to the delete fleets fleet Id squads squad Id internal server error response
func (o *DeleteFleetsFleetIDSquadsSquadIDInternalServerError) WithPayload(payload *models.DeleteFleetsFleetIDSquadsSquadIDInternalServerErrorBody) *DeleteFleetsFleetIDSquadsSquadIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete fleets fleet Id squads squad Id internal server error response
func (o *DeleteFleetsFleetIDSquadsSquadIDInternalServerError) SetPayload(payload *models.DeleteFleetsFleetIDSquadsSquadIDInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteFleetsFleetIDSquadsSquadIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
