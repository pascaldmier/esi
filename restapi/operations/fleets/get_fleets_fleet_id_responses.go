package fleets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/pascaldmier/esi/models"
)

// GetFleetsFleetIDOKCode is the HTTP code returned for type GetFleetsFleetIDOK
const GetFleetsFleetIDOKCode int = 200

/*GetFleetsFleetIDOK Details about a fleet

swagger:response getFleetsFleetIdOK
*/
type GetFleetsFleetIDOK struct {
	/*The caching mechanism used
	  Required: true
	*/
	CacheControl string `json:"Cache-Control"`
	/*RFC7231 formatted datetime string
	  Required: true
	*/
	Expires string `json:"Expires"`
	/*RFC7231 formatted datetime string
	  Required: true
	*/
	LastModified string `json:"Last-Modified"`

	/*
	  In: Body
	*/
	Payload *models.GetFleetsFleetIDOKBody `json:"body,omitempty"`
}

// NewGetFleetsFleetIDOK creates GetFleetsFleetIDOK with default headers values
func NewGetFleetsFleetIDOK() *GetFleetsFleetIDOK {
	return &GetFleetsFleetIDOK{}
}

// WithCacheControl adds the cacheControl to the get fleets fleet Id o k response
func (o *GetFleetsFleetIDOK) WithCacheControl(cacheControl string) *GetFleetsFleetIDOK {
	o.CacheControl = cacheControl
	return o
}

// SetCacheControl sets the cacheControl to the get fleets fleet Id o k response
func (o *GetFleetsFleetIDOK) SetCacheControl(cacheControl string) {
	o.CacheControl = cacheControl
}

// WithExpires adds the expires to the get fleets fleet Id o k response
func (o *GetFleetsFleetIDOK) WithExpires(expires string) *GetFleetsFleetIDOK {
	o.Expires = expires
	return o
}

// SetExpires sets the expires to the get fleets fleet Id o k response
func (o *GetFleetsFleetIDOK) SetExpires(expires string) {
	o.Expires = expires
}

// WithLastModified adds the lastModified to the get fleets fleet Id o k response
func (o *GetFleetsFleetIDOK) WithLastModified(lastModified string) *GetFleetsFleetIDOK {
	o.LastModified = lastModified
	return o
}

// SetLastModified sets the lastModified to the get fleets fleet Id o k response
func (o *GetFleetsFleetIDOK) SetLastModified(lastModified string) {
	o.LastModified = lastModified
}

// WithPayload adds the payload to the get fleets fleet Id o k response
func (o *GetFleetsFleetIDOK) WithPayload(payload *models.GetFleetsFleetIDOKBody) *GetFleetsFleetIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get fleets fleet Id o k response
func (o *GetFleetsFleetIDOK) SetPayload(payload *models.GetFleetsFleetIDOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFleetsFleetIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Cache-Control

	cacheControl := o.CacheControl
	if cacheControl != "" {
		rw.Header().Set("Cache-Control", cacheControl)
	}

	// response header Expires

	expires := o.Expires
	if expires != "" {
		rw.Header().Set("Expires", expires)
	}

	// response header Last-Modified

	lastModified := o.LastModified
	if lastModified != "" {
		rw.Header().Set("Last-Modified", lastModified)
	}

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetFleetsFleetIDForbiddenCode is the HTTP code returned for type GetFleetsFleetIDForbidden
const GetFleetsFleetIDForbiddenCode int = 403

/*GetFleetsFleetIDForbidden Forbidden

swagger:response getFleetsFleetIdForbidden
*/
type GetFleetsFleetIDForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GetFleetsFleetIDForbiddenBody `json:"body,omitempty"`
}

// NewGetFleetsFleetIDForbidden creates GetFleetsFleetIDForbidden with default headers values
func NewGetFleetsFleetIDForbidden() *GetFleetsFleetIDForbidden {
	return &GetFleetsFleetIDForbidden{}
}

// WithPayload adds the payload to the get fleets fleet Id forbidden response
func (o *GetFleetsFleetIDForbidden) WithPayload(payload *models.GetFleetsFleetIDForbiddenBody) *GetFleetsFleetIDForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get fleets fleet Id forbidden response
func (o *GetFleetsFleetIDForbidden) SetPayload(payload *models.GetFleetsFleetIDForbiddenBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFleetsFleetIDForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetFleetsFleetIDNotFoundCode is the HTTP code returned for type GetFleetsFleetIDNotFound
const GetFleetsFleetIDNotFoundCode int = 404

/*GetFleetsFleetIDNotFound The fleet does not exist or you don't have access to it

swagger:response getFleetsFleetIdNotFound
*/
type GetFleetsFleetIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.GetFleetsFleetIDNotFoundBody `json:"body,omitempty"`
}

// NewGetFleetsFleetIDNotFound creates GetFleetsFleetIDNotFound with default headers values
func NewGetFleetsFleetIDNotFound() *GetFleetsFleetIDNotFound {
	return &GetFleetsFleetIDNotFound{}
}

// WithPayload adds the payload to the get fleets fleet Id not found response
func (o *GetFleetsFleetIDNotFound) WithPayload(payload *models.GetFleetsFleetIDNotFoundBody) *GetFleetsFleetIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get fleets fleet Id not found response
func (o *GetFleetsFleetIDNotFound) SetPayload(payload *models.GetFleetsFleetIDNotFoundBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFleetsFleetIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetFleetsFleetIDInternalServerErrorCode is the HTTP code returned for type GetFleetsFleetIDInternalServerError
const GetFleetsFleetIDInternalServerErrorCode int = 500

/*GetFleetsFleetIDInternalServerError Internal server error

swagger:response getFleetsFleetIdInternalServerError
*/
type GetFleetsFleetIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.GetFleetsFleetIDInternalServerErrorBody `json:"body,omitempty"`
}

// NewGetFleetsFleetIDInternalServerError creates GetFleetsFleetIDInternalServerError with default headers values
func NewGetFleetsFleetIDInternalServerError() *GetFleetsFleetIDInternalServerError {
	return &GetFleetsFleetIDInternalServerError{}
}

// WithPayload adds the payload to the get fleets fleet Id internal server error response
func (o *GetFleetsFleetIDInternalServerError) WithPayload(payload *models.GetFleetsFleetIDInternalServerErrorBody) *GetFleetsFleetIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get fleets fleet Id internal server error response
func (o *GetFleetsFleetIDInternalServerError) SetPayload(payload *models.GetFleetsFleetIDInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFleetsFleetIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
