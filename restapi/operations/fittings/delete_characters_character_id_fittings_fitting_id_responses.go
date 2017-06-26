package fittings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/pascaldmier/esi/models"
)

// DeleteCharactersCharacterIDFittingsFittingIDNoContentCode is the HTTP code returned for type DeleteCharactersCharacterIDFittingsFittingIDNoContent
const DeleteCharactersCharacterIDFittingsFittingIDNoContentCode int = 204

/*DeleteCharactersCharacterIDFittingsFittingIDNoContent Fitting deleted

swagger:response deleteCharactersCharacterIdFittingsFittingIdNoContent
*/
type DeleteCharactersCharacterIDFittingsFittingIDNoContent struct {
}

// NewDeleteCharactersCharacterIDFittingsFittingIDNoContent creates DeleteCharactersCharacterIDFittingsFittingIDNoContent with default headers values
func NewDeleteCharactersCharacterIDFittingsFittingIDNoContent() *DeleteCharactersCharacterIDFittingsFittingIDNoContent {
	return &DeleteCharactersCharacterIDFittingsFittingIDNoContent{}
}

// WriteResponse to the client
func (o *DeleteCharactersCharacterIDFittingsFittingIDNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(204)
}

// DeleteCharactersCharacterIDFittingsFittingIDForbiddenCode is the HTTP code returned for type DeleteCharactersCharacterIDFittingsFittingIDForbidden
const DeleteCharactersCharacterIDFittingsFittingIDForbiddenCode int = 403

/*DeleteCharactersCharacterIDFittingsFittingIDForbidden Forbidden

swagger:response deleteCharactersCharacterIdFittingsFittingIdForbidden
*/
type DeleteCharactersCharacterIDFittingsFittingIDForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.DeleteCharactersCharacterIDFittingsFittingIDForbiddenBody `json:"body,omitempty"`
}

// NewDeleteCharactersCharacterIDFittingsFittingIDForbidden creates DeleteCharactersCharacterIDFittingsFittingIDForbidden with default headers values
func NewDeleteCharactersCharacterIDFittingsFittingIDForbidden() *DeleteCharactersCharacterIDFittingsFittingIDForbidden {
	return &DeleteCharactersCharacterIDFittingsFittingIDForbidden{}
}

// WithPayload adds the payload to the delete characters character Id fittings fitting Id forbidden response
func (o *DeleteCharactersCharacterIDFittingsFittingIDForbidden) WithPayload(payload *models.DeleteCharactersCharacterIDFittingsFittingIDForbiddenBody) *DeleteCharactersCharacterIDFittingsFittingIDForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete characters character Id fittings fitting Id forbidden response
func (o *DeleteCharactersCharacterIDFittingsFittingIDForbidden) SetPayload(payload *models.DeleteCharactersCharacterIDFittingsFittingIDForbiddenBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteCharactersCharacterIDFittingsFittingIDForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteCharactersCharacterIDFittingsFittingIDInternalServerErrorCode is the HTTP code returned for type DeleteCharactersCharacterIDFittingsFittingIDInternalServerError
const DeleteCharactersCharacterIDFittingsFittingIDInternalServerErrorCode int = 500

/*DeleteCharactersCharacterIDFittingsFittingIDInternalServerError Internal server error

swagger:response deleteCharactersCharacterIdFittingsFittingIdInternalServerError
*/
type DeleteCharactersCharacterIDFittingsFittingIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.DeleteCharactersCharacterIDFittingsFittingIDInternalServerErrorBody `json:"body,omitempty"`
}

// NewDeleteCharactersCharacterIDFittingsFittingIDInternalServerError creates DeleteCharactersCharacterIDFittingsFittingIDInternalServerError with default headers values
func NewDeleteCharactersCharacterIDFittingsFittingIDInternalServerError() *DeleteCharactersCharacterIDFittingsFittingIDInternalServerError {
	return &DeleteCharactersCharacterIDFittingsFittingIDInternalServerError{}
}

// WithPayload adds the payload to the delete characters character Id fittings fitting Id internal server error response
func (o *DeleteCharactersCharacterIDFittingsFittingIDInternalServerError) WithPayload(payload *models.DeleteCharactersCharacterIDFittingsFittingIDInternalServerErrorBody) *DeleteCharactersCharacterIDFittingsFittingIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete characters character Id fittings fitting Id internal server error response
func (o *DeleteCharactersCharacterIDFittingsFittingIDInternalServerError) SetPayload(payload *models.DeleteCharactersCharacterIDFittingsFittingIDInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteCharactersCharacterIDFittingsFittingIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}