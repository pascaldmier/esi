package mail

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/pascaldmier/esi/models"
)

// DeleteCharactersCharacterIDMailMailIDNoContentCode is the HTTP code returned for type DeleteCharactersCharacterIDMailMailIDNoContent
const DeleteCharactersCharacterIDMailMailIDNoContentCode int = 204

/*DeleteCharactersCharacterIDMailMailIDNoContent Mail deleted

swagger:response deleteCharactersCharacterIdMailMailIdNoContent
*/
type DeleteCharactersCharacterIDMailMailIDNoContent struct {
}

// NewDeleteCharactersCharacterIDMailMailIDNoContent creates DeleteCharactersCharacterIDMailMailIDNoContent with default headers values
func NewDeleteCharactersCharacterIDMailMailIDNoContent() *DeleteCharactersCharacterIDMailMailIDNoContent {
	return &DeleteCharactersCharacterIDMailMailIDNoContent{}
}

// WriteResponse to the client
func (o *DeleteCharactersCharacterIDMailMailIDNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(204)
}

// DeleteCharactersCharacterIDMailMailIDForbiddenCode is the HTTP code returned for type DeleteCharactersCharacterIDMailMailIDForbidden
const DeleteCharactersCharacterIDMailMailIDForbiddenCode int = 403

/*DeleteCharactersCharacterIDMailMailIDForbidden Forbidden

swagger:response deleteCharactersCharacterIdMailMailIdForbidden
*/
type DeleteCharactersCharacterIDMailMailIDForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.DeleteCharactersCharacterIDMailMailIDForbiddenBody `json:"body,omitempty"`
}

// NewDeleteCharactersCharacterIDMailMailIDForbidden creates DeleteCharactersCharacterIDMailMailIDForbidden with default headers values
func NewDeleteCharactersCharacterIDMailMailIDForbidden() *DeleteCharactersCharacterIDMailMailIDForbidden {
	return &DeleteCharactersCharacterIDMailMailIDForbidden{}
}

// WithPayload adds the payload to the delete characters character Id mail mail Id forbidden response
func (o *DeleteCharactersCharacterIDMailMailIDForbidden) WithPayload(payload *models.DeleteCharactersCharacterIDMailMailIDForbiddenBody) *DeleteCharactersCharacterIDMailMailIDForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete characters character Id mail mail Id forbidden response
func (o *DeleteCharactersCharacterIDMailMailIDForbidden) SetPayload(payload *models.DeleteCharactersCharacterIDMailMailIDForbiddenBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteCharactersCharacterIDMailMailIDForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteCharactersCharacterIDMailMailIDInternalServerErrorCode is the HTTP code returned for type DeleteCharactersCharacterIDMailMailIDInternalServerError
const DeleteCharactersCharacterIDMailMailIDInternalServerErrorCode int = 500

/*DeleteCharactersCharacterIDMailMailIDInternalServerError Internal server error

swagger:response deleteCharactersCharacterIdMailMailIdInternalServerError
*/
type DeleteCharactersCharacterIDMailMailIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.DeleteCharactersCharacterIDMailMailIDInternalServerErrorBody `json:"body,omitempty"`
}

// NewDeleteCharactersCharacterIDMailMailIDInternalServerError creates DeleteCharactersCharacterIDMailMailIDInternalServerError with default headers values
func NewDeleteCharactersCharacterIDMailMailIDInternalServerError() *DeleteCharactersCharacterIDMailMailIDInternalServerError {
	return &DeleteCharactersCharacterIDMailMailIDInternalServerError{}
}

// WithPayload adds the payload to the delete characters character Id mail mail Id internal server error response
func (o *DeleteCharactersCharacterIDMailMailIDInternalServerError) WithPayload(payload *models.DeleteCharactersCharacterIDMailMailIDInternalServerErrorBody) *DeleteCharactersCharacterIDMailMailIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete characters character Id mail mail Id internal server error response
func (o *DeleteCharactersCharacterIDMailMailIDInternalServerError) SetPayload(payload *models.DeleteCharactersCharacterIDMailMailIDInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteCharactersCharacterIDMailMailIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
