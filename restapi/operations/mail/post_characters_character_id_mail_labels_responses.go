package mail

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/pascaldmier/esi/models"
)

// PostCharactersCharacterIDMailLabelsCreatedCode is the HTTP code returned for type PostCharactersCharacterIDMailLabelsCreated
const PostCharactersCharacterIDMailLabelsCreatedCode int = 201

/*PostCharactersCharacterIDMailLabelsCreated Label created

swagger:response postCharactersCharacterIdMailLabelsCreated
*/
type PostCharactersCharacterIDMailLabelsCreated struct {

	/*Label ID
	  In: Body
	*/
	Payload int64 `json:"body,omitempty"`
}

// NewPostCharactersCharacterIDMailLabelsCreated creates PostCharactersCharacterIDMailLabelsCreated with default headers values
func NewPostCharactersCharacterIDMailLabelsCreated() *PostCharactersCharacterIDMailLabelsCreated {
	return &PostCharactersCharacterIDMailLabelsCreated{}
}

// WithPayload adds the payload to the post characters character Id mail labels created response
func (o *PostCharactersCharacterIDMailLabelsCreated) WithPayload(payload int64) *PostCharactersCharacterIDMailLabelsCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post characters character Id mail labels created response
func (o *PostCharactersCharacterIDMailLabelsCreated) SetPayload(payload int64) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostCharactersCharacterIDMailLabelsCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// PostCharactersCharacterIDMailLabelsForbiddenCode is the HTTP code returned for type PostCharactersCharacterIDMailLabelsForbidden
const PostCharactersCharacterIDMailLabelsForbiddenCode int = 403

/*PostCharactersCharacterIDMailLabelsForbidden Forbidden

swagger:response postCharactersCharacterIdMailLabelsForbidden
*/
type PostCharactersCharacterIDMailLabelsForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.PostCharactersCharacterIDMailLabelsForbiddenBody `json:"body,omitempty"`
}

// NewPostCharactersCharacterIDMailLabelsForbidden creates PostCharactersCharacterIDMailLabelsForbidden with default headers values
func NewPostCharactersCharacterIDMailLabelsForbidden() *PostCharactersCharacterIDMailLabelsForbidden {
	return &PostCharactersCharacterIDMailLabelsForbidden{}
}

// WithPayload adds the payload to the post characters character Id mail labels forbidden response
func (o *PostCharactersCharacterIDMailLabelsForbidden) WithPayload(payload *models.PostCharactersCharacterIDMailLabelsForbiddenBody) *PostCharactersCharacterIDMailLabelsForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post characters character Id mail labels forbidden response
func (o *PostCharactersCharacterIDMailLabelsForbidden) SetPayload(payload *models.PostCharactersCharacterIDMailLabelsForbiddenBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostCharactersCharacterIDMailLabelsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostCharactersCharacterIDMailLabelsInternalServerErrorCode is the HTTP code returned for type PostCharactersCharacterIDMailLabelsInternalServerError
const PostCharactersCharacterIDMailLabelsInternalServerErrorCode int = 500

/*PostCharactersCharacterIDMailLabelsInternalServerError Internal server error

swagger:response postCharactersCharacterIdMailLabelsInternalServerError
*/
type PostCharactersCharacterIDMailLabelsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.PostCharactersCharacterIDMailLabelsInternalServerErrorBody `json:"body,omitempty"`
}

// NewPostCharactersCharacterIDMailLabelsInternalServerError creates PostCharactersCharacterIDMailLabelsInternalServerError with default headers values
func NewPostCharactersCharacterIDMailLabelsInternalServerError() *PostCharactersCharacterIDMailLabelsInternalServerError {
	return &PostCharactersCharacterIDMailLabelsInternalServerError{}
}

// WithPayload adds the payload to the post characters character Id mail labels internal server error response
func (o *PostCharactersCharacterIDMailLabelsInternalServerError) WithPayload(payload *models.PostCharactersCharacterIDMailLabelsInternalServerErrorBody) *PostCharactersCharacterIDMailLabelsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post characters character Id mail labels internal server error response
func (o *PostCharactersCharacterIDMailLabelsInternalServerError) SetPayload(payload *models.PostCharactersCharacterIDMailLabelsInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostCharactersCharacterIDMailLabelsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}