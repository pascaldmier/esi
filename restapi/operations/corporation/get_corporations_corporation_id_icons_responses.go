package corporation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/pascaldmier/esi/models"
)

// GetCorporationsCorporationIDIconsOKCode is the HTTP code returned for type GetCorporationsCorporationIDIconsOK
const GetCorporationsCorporationIDIconsOKCode int = 200

/*GetCorporationsCorporationIDIconsOK Urls for icons for the given corporation id and server

swagger:response getCorporationsCorporationIdIconsOK
*/
type GetCorporationsCorporationIDIconsOK struct {
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
	Payload *models.GetCorporationsCorporationIDIconsOKBody `json:"body,omitempty"`
}

// NewGetCorporationsCorporationIDIconsOK creates GetCorporationsCorporationIDIconsOK with default headers values
func NewGetCorporationsCorporationIDIconsOK() *GetCorporationsCorporationIDIconsOK {
	return &GetCorporationsCorporationIDIconsOK{}
}

// WithCacheControl adds the cacheControl to the get corporations corporation Id icons o k response
func (o *GetCorporationsCorporationIDIconsOK) WithCacheControl(cacheControl string) *GetCorporationsCorporationIDIconsOK {
	o.CacheControl = cacheControl
	return o
}

// SetCacheControl sets the cacheControl to the get corporations corporation Id icons o k response
func (o *GetCorporationsCorporationIDIconsOK) SetCacheControl(cacheControl string) {
	o.CacheControl = cacheControl
}

// WithExpires adds the expires to the get corporations corporation Id icons o k response
func (o *GetCorporationsCorporationIDIconsOK) WithExpires(expires string) *GetCorporationsCorporationIDIconsOK {
	o.Expires = expires
	return o
}

// SetExpires sets the expires to the get corporations corporation Id icons o k response
func (o *GetCorporationsCorporationIDIconsOK) SetExpires(expires string) {
	o.Expires = expires
}

// WithLastModified adds the lastModified to the get corporations corporation Id icons o k response
func (o *GetCorporationsCorporationIDIconsOK) WithLastModified(lastModified string) *GetCorporationsCorporationIDIconsOK {
	o.LastModified = lastModified
	return o
}

// SetLastModified sets the lastModified to the get corporations corporation Id icons o k response
func (o *GetCorporationsCorporationIDIconsOK) SetLastModified(lastModified string) {
	o.LastModified = lastModified
}

// WithPayload adds the payload to the get corporations corporation Id icons o k response
func (o *GetCorporationsCorporationIDIconsOK) WithPayload(payload *models.GetCorporationsCorporationIDIconsOKBody) *GetCorporationsCorporationIDIconsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get corporations corporation Id icons o k response
func (o *GetCorporationsCorporationIDIconsOK) SetPayload(payload *models.GetCorporationsCorporationIDIconsOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCorporationsCorporationIDIconsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetCorporationsCorporationIDIconsNotFoundCode is the HTTP code returned for type GetCorporationsCorporationIDIconsNotFound
const GetCorporationsCorporationIDIconsNotFoundCode int = 404

/*GetCorporationsCorporationIDIconsNotFound No image server for this datasource

swagger:response getCorporationsCorporationIdIconsNotFound
*/
type GetCorporationsCorporationIDIconsNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.GetCorporationsCorporationIDIconsNotFoundBody `json:"body,omitempty"`
}

// NewGetCorporationsCorporationIDIconsNotFound creates GetCorporationsCorporationIDIconsNotFound with default headers values
func NewGetCorporationsCorporationIDIconsNotFound() *GetCorporationsCorporationIDIconsNotFound {
	return &GetCorporationsCorporationIDIconsNotFound{}
}

// WithPayload adds the payload to the get corporations corporation Id icons not found response
func (o *GetCorporationsCorporationIDIconsNotFound) WithPayload(payload *models.GetCorporationsCorporationIDIconsNotFoundBody) *GetCorporationsCorporationIDIconsNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get corporations corporation Id icons not found response
func (o *GetCorporationsCorporationIDIconsNotFound) SetPayload(payload *models.GetCorporationsCorporationIDIconsNotFoundBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCorporationsCorporationIDIconsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetCorporationsCorporationIDIconsInternalServerErrorCode is the HTTP code returned for type GetCorporationsCorporationIDIconsInternalServerError
const GetCorporationsCorporationIDIconsInternalServerErrorCode int = 500

/*GetCorporationsCorporationIDIconsInternalServerError Internal server error

swagger:response getCorporationsCorporationIdIconsInternalServerError
*/
type GetCorporationsCorporationIDIconsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.GetCorporationsCorporationIDIconsInternalServerErrorBody `json:"body,omitempty"`
}

// NewGetCorporationsCorporationIDIconsInternalServerError creates GetCorporationsCorporationIDIconsInternalServerError with default headers values
func NewGetCorporationsCorporationIDIconsInternalServerError() *GetCorporationsCorporationIDIconsInternalServerError {
	return &GetCorporationsCorporationIDIconsInternalServerError{}
}

// WithPayload adds the payload to the get corporations corporation Id icons internal server error response
func (o *GetCorporationsCorporationIDIconsInternalServerError) WithPayload(payload *models.GetCorporationsCorporationIDIconsInternalServerErrorBody) *GetCorporationsCorporationIDIconsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get corporations corporation Id icons internal server error response
func (o *GetCorporationsCorporationIDIconsInternalServerError) SetPayload(payload *models.GetCorporationsCorporationIDIconsInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCorporationsCorporationIDIconsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
