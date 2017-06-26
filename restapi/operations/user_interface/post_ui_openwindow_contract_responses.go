package user_interface

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/pascaldmier/esi/models"
)

// PostUIOpenwindowContractNoContentCode is the HTTP code returned for type PostUIOpenwindowContractNoContent
const PostUIOpenwindowContractNoContentCode int = 204

/*PostUIOpenwindowContractNoContent Open window request received

swagger:response postUiOpenwindowContractNoContent
*/
type PostUIOpenwindowContractNoContent struct {
}

// NewPostUIOpenwindowContractNoContent creates PostUIOpenwindowContractNoContent with default headers values
func NewPostUIOpenwindowContractNoContent() *PostUIOpenwindowContractNoContent {
	return &PostUIOpenwindowContractNoContent{}
}

// WriteResponse to the client
func (o *PostUIOpenwindowContractNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(204)
}

// PostUIOpenwindowContractForbiddenCode is the HTTP code returned for type PostUIOpenwindowContractForbidden
const PostUIOpenwindowContractForbiddenCode int = 403

/*PostUIOpenwindowContractForbidden Forbidden

swagger:response postUiOpenwindowContractForbidden
*/
type PostUIOpenwindowContractForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.PostUIOpenwindowContractForbiddenBody `json:"body,omitempty"`
}

// NewPostUIOpenwindowContractForbidden creates PostUIOpenwindowContractForbidden with default headers values
func NewPostUIOpenwindowContractForbidden() *PostUIOpenwindowContractForbidden {
	return &PostUIOpenwindowContractForbidden{}
}

// WithPayload adds the payload to the post Ui openwindow contract forbidden response
func (o *PostUIOpenwindowContractForbidden) WithPayload(payload *models.PostUIOpenwindowContractForbiddenBody) *PostUIOpenwindowContractForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post Ui openwindow contract forbidden response
func (o *PostUIOpenwindowContractForbidden) SetPayload(payload *models.PostUIOpenwindowContractForbiddenBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostUIOpenwindowContractForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostUIOpenwindowContractInternalServerErrorCode is the HTTP code returned for type PostUIOpenwindowContractInternalServerError
const PostUIOpenwindowContractInternalServerErrorCode int = 500

/*PostUIOpenwindowContractInternalServerError Internal server error

swagger:response postUiOpenwindowContractInternalServerError
*/
type PostUIOpenwindowContractInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.PostUIOpenwindowContractInternalServerErrorBody `json:"body,omitempty"`
}

// NewPostUIOpenwindowContractInternalServerError creates PostUIOpenwindowContractInternalServerError with default headers values
func NewPostUIOpenwindowContractInternalServerError() *PostUIOpenwindowContractInternalServerError {
	return &PostUIOpenwindowContractInternalServerError{}
}

// WithPayload adds the payload to the post Ui openwindow contract internal server error response
func (o *PostUIOpenwindowContractInternalServerError) WithPayload(payload *models.PostUIOpenwindowContractInternalServerErrorBody) *PostUIOpenwindowContractInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post Ui openwindow contract internal server error response
func (o *PostUIOpenwindowContractInternalServerError) SetPayload(payload *models.PostUIOpenwindowContractInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostUIOpenwindowContractInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}