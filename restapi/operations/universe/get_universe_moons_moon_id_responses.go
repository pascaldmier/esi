package universe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/pascaldmier/esi/models"
)

// GetUniverseMoonsMoonIDOKCode is the HTTP code returned for type GetUniverseMoonsMoonIDOK
const GetUniverseMoonsMoonIDOKCode int = 200

/*GetUniverseMoonsMoonIDOK Information about a moon

swagger:response getUniverseMoonsMoonIdOK
*/
type GetUniverseMoonsMoonIDOK struct {
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
	Payload *models.GetUniverseMoonsMoonIDOKBody `json:"body,omitempty"`
}

// NewGetUniverseMoonsMoonIDOK creates GetUniverseMoonsMoonIDOK with default headers values
func NewGetUniverseMoonsMoonIDOK() *GetUniverseMoonsMoonIDOK {
	return &GetUniverseMoonsMoonIDOK{}
}

// WithCacheControl adds the cacheControl to the get universe moons moon Id o k response
func (o *GetUniverseMoonsMoonIDOK) WithCacheControl(cacheControl string) *GetUniverseMoonsMoonIDOK {
	o.CacheControl = cacheControl
	return o
}

// SetCacheControl sets the cacheControl to the get universe moons moon Id o k response
func (o *GetUniverseMoonsMoonIDOK) SetCacheControl(cacheControl string) {
	o.CacheControl = cacheControl
}

// WithExpires adds the expires to the get universe moons moon Id o k response
func (o *GetUniverseMoonsMoonIDOK) WithExpires(expires string) *GetUniverseMoonsMoonIDOK {
	o.Expires = expires
	return o
}

// SetExpires sets the expires to the get universe moons moon Id o k response
func (o *GetUniverseMoonsMoonIDOK) SetExpires(expires string) {
	o.Expires = expires
}

// WithLastModified adds the lastModified to the get universe moons moon Id o k response
func (o *GetUniverseMoonsMoonIDOK) WithLastModified(lastModified string) *GetUniverseMoonsMoonIDOK {
	o.LastModified = lastModified
	return o
}

// SetLastModified sets the lastModified to the get universe moons moon Id o k response
func (o *GetUniverseMoonsMoonIDOK) SetLastModified(lastModified string) {
	o.LastModified = lastModified
}

// WithPayload adds the payload to the get universe moons moon Id o k response
func (o *GetUniverseMoonsMoonIDOK) WithPayload(payload *models.GetUniverseMoonsMoonIDOKBody) *GetUniverseMoonsMoonIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get universe moons moon Id o k response
func (o *GetUniverseMoonsMoonIDOK) SetPayload(payload *models.GetUniverseMoonsMoonIDOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUniverseMoonsMoonIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetUniverseMoonsMoonIDNotFoundCode is the HTTP code returned for type GetUniverseMoonsMoonIDNotFound
const GetUniverseMoonsMoonIDNotFoundCode int = 404

/*GetUniverseMoonsMoonIDNotFound Moon not found

swagger:response getUniverseMoonsMoonIdNotFound
*/
type GetUniverseMoonsMoonIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.GetUniverseMoonsMoonIDNotFoundBody `json:"body,omitempty"`
}

// NewGetUniverseMoonsMoonIDNotFound creates GetUniverseMoonsMoonIDNotFound with default headers values
func NewGetUniverseMoonsMoonIDNotFound() *GetUniverseMoonsMoonIDNotFound {
	return &GetUniverseMoonsMoonIDNotFound{}
}

// WithPayload adds the payload to the get universe moons moon Id not found response
func (o *GetUniverseMoonsMoonIDNotFound) WithPayload(payload *models.GetUniverseMoonsMoonIDNotFoundBody) *GetUniverseMoonsMoonIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get universe moons moon Id not found response
func (o *GetUniverseMoonsMoonIDNotFound) SetPayload(payload *models.GetUniverseMoonsMoonIDNotFoundBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUniverseMoonsMoonIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUniverseMoonsMoonIDInternalServerErrorCode is the HTTP code returned for type GetUniverseMoonsMoonIDInternalServerError
const GetUniverseMoonsMoonIDInternalServerErrorCode int = 500

/*GetUniverseMoonsMoonIDInternalServerError Internal server error

swagger:response getUniverseMoonsMoonIdInternalServerError
*/
type GetUniverseMoonsMoonIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.GetUniverseMoonsMoonIDInternalServerErrorBody `json:"body,omitempty"`
}

// NewGetUniverseMoonsMoonIDInternalServerError creates GetUniverseMoonsMoonIDInternalServerError with default headers values
func NewGetUniverseMoonsMoonIDInternalServerError() *GetUniverseMoonsMoonIDInternalServerError {
	return &GetUniverseMoonsMoonIDInternalServerError{}
}

// WithPayload adds the payload to the get universe moons moon Id internal server error response
func (o *GetUniverseMoonsMoonIDInternalServerError) WithPayload(payload *models.GetUniverseMoonsMoonIDInternalServerErrorBody) *GetUniverseMoonsMoonIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get universe moons moon Id internal server error response
func (o *GetUniverseMoonsMoonIDInternalServerError) SetPayload(payload *models.GetUniverseMoonsMoonIDInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUniverseMoonsMoonIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
