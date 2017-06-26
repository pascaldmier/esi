package user_interface

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/pascaldmier/esi/models"
)

// PostUIOpenwindowNewmailNoContentCode is the HTTP code returned for type PostUIOpenwindowNewmailNoContent
const PostUIOpenwindowNewmailNoContentCode int = 204

/*PostUIOpenwindowNewmailNoContent Open window request received

swagger:response postUiOpenwindowNewmailNoContent
*/
type PostUIOpenwindowNewmailNoContent struct {
}

// NewPostUIOpenwindowNewmailNoContent creates PostUIOpenwindowNewmailNoContent with default headers values
func NewPostUIOpenwindowNewmailNoContent() *PostUIOpenwindowNewmailNoContent {
	return &PostUIOpenwindowNewmailNoContent{}
}

// WriteResponse to the client
func (o *PostUIOpenwindowNewmailNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(204)
}

// PostUIOpenwindowNewmailForbiddenCode is the HTTP code returned for type PostUIOpenwindowNewmailForbidden
const PostUIOpenwindowNewmailForbiddenCode int = 403

/*PostUIOpenwindowNewmailForbidden Forbidden

swagger:response postUiOpenwindowNewmailForbidden
*/
type PostUIOpenwindowNewmailForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.PostUIOpenwindowNewmailForbiddenBody `json:"body,omitempty"`
}

// NewPostUIOpenwindowNewmailForbidden creates PostUIOpenwindowNewmailForbidden with default headers values
func NewPostUIOpenwindowNewmailForbidden() *PostUIOpenwindowNewmailForbidden {
	return &PostUIOpenwindowNewmailForbidden{}
}

// WithPayload adds the payload to the post Ui openwindow newmail forbidden response
func (o *PostUIOpenwindowNewmailForbidden) WithPayload(payload *models.PostUIOpenwindowNewmailForbiddenBody) *PostUIOpenwindowNewmailForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post Ui openwindow newmail forbidden response
func (o *PostUIOpenwindowNewmailForbidden) SetPayload(payload *models.PostUIOpenwindowNewmailForbiddenBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostUIOpenwindowNewmailForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostUIOpenwindowNewmailUnprocessableEntityCode is the HTTP code returned for type PostUIOpenwindowNewmailUnprocessableEntity
const PostUIOpenwindowNewmailUnprocessableEntityCode int = 422

/*PostUIOpenwindowNewmailUnprocessableEntity Invalid request

swagger:response postUiOpenwindowNewmailUnprocessableEntity
*/
type PostUIOpenwindowNewmailUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.PostUIOpenwindowNewmailUnprocessableEntityBody `json:"body,omitempty"`
}

// NewPostUIOpenwindowNewmailUnprocessableEntity creates PostUIOpenwindowNewmailUnprocessableEntity with default headers values
func NewPostUIOpenwindowNewmailUnprocessableEntity() *PostUIOpenwindowNewmailUnprocessableEntity {
	return &PostUIOpenwindowNewmailUnprocessableEntity{}
}

// WithPayload adds the payload to the post Ui openwindow newmail unprocessable entity response
func (o *PostUIOpenwindowNewmailUnprocessableEntity) WithPayload(payload *models.PostUIOpenwindowNewmailUnprocessableEntityBody) *PostUIOpenwindowNewmailUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post Ui openwindow newmail unprocessable entity response
func (o *PostUIOpenwindowNewmailUnprocessableEntity) SetPayload(payload *models.PostUIOpenwindowNewmailUnprocessableEntityBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostUIOpenwindowNewmailUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostUIOpenwindowNewmailInternalServerErrorCode is the HTTP code returned for type PostUIOpenwindowNewmailInternalServerError
const PostUIOpenwindowNewmailInternalServerErrorCode int = 500

/*PostUIOpenwindowNewmailInternalServerError Internal server error

swagger:response postUiOpenwindowNewmailInternalServerError
*/
type PostUIOpenwindowNewmailInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.PostUIOpenwindowNewmailInternalServerErrorBody `json:"body,omitempty"`
}

// NewPostUIOpenwindowNewmailInternalServerError creates PostUIOpenwindowNewmailInternalServerError with default headers values
func NewPostUIOpenwindowNewmailInternalServerError() *PostUIOpenwindowNewmailInternalServerError {
	return &PostUIOpenwindowNewmailInternalServerError{}
}

// WithPayload adds the payload to the post Ui openwindow newmail internal server error response
func (o *PostUIOpenwindowNewmailInternalServerError) WithPayload(payload *models.PostUIOpenwindowNewmailInternalServerErrorBody) *PostUIOpenwindowNewmailInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post Ui openwindow newmail internal server error response
func (o *PostUIOpenwindowNewmailInternalServerError) SetPayload(payload *models.PostUIOpenwindowNewmailInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostUIOpenwindowNewmailInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
