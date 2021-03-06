package universe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/pascaldmier/esi/models"
)

// GetUniverseStructuresStructureIDOKCode is the HTTP code returned for type GetUniverseStructuresStructureIDOK
const GetUniverseStructuresStructureIDOKCode int = 200

/*GetUniverseStructuresStructureIDOK Data about a structure

swagger:response getUniverseStructuresStructureIdOK
*/
type GetUniverseStructuresStructureIDOK struct {
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
	Payload *models.GetUniverseStructuresStructureIDOKBody `json:"body,omitempty"`
}

// NewGetUniverseStructuresStructureIDOK creates GetUniverseStructuresStructureIDOK with default headers values
func NewGetUniverseStructuresStructureIDOK() *GetUniverseStructuresStructureIDOK {
	return &GetUniverseStructuresStructureIDOK{}
}

// WithCacheControl adds the cacheControl to the get universe structures structure Id o k response
func (o *GetUniverseStructuresStructureIDOK) WithCacheControl(cacheControl string) *GetUniverseStructuresStructureIDOK {
	o.CacheControl = cacheControl
	return o
}

// SetCacheControl sets the cacheControl to the get universe structures structure Id o k response
func (o *GetUniverseStructuresStructureIDOK) SetCacheControl(cacheControl string) {
	o.CacheControl = cacheControl
}

// WithExpires adds the expires to the get universe structures structure Id o k response
func (o *GetUniverseStructuresStructureIDOK) WithExpires(expires string) *GetUniverseStructuresStructureIDOK {
	o.Expires = expires
	return o
}

// SetExpires sets the expires to the get universe structures structure Id o k response
func (o *GetUniverseStructuresStructureIDOK) SetExpires(expires string) {
	o.Expires = expires
}

// WithLastModified adds the lastModified to the get universe structures structure Id o k response
func (o *GetUniverseStructuresStructureIDOK) WithLastModified(lastModified string) *GetUniverseStructuresStructureIDOK {
	o.LastModified = lastModified
	return o
}

// SetLastModified sets the lastModified to the get universe structures structure Id o k response
func (o *GetUniverseStructuresStructureIDOK) SetLastModified(lastModified string) {
	o.LastModified = lastModified
}

// WithPayload adds the payload to the get universe structures structure Id o k response
func (o *GetUniverseStructuresStructureIDOK) WithPayload(payload *models.GetUniverseStructuresStructureIDOKBody) *GetUniverseStructuresStructureIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get universe structures structure Id o k response
func (o *GetUniverseStructuresStructureIDOK) SetPayload(payload *models.GetUniverseStructuresStructureIDOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUniverseStructuresStructureIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetUniverseStructuresStructureIDForbiddenCode is the HTTP code returned for type GetUniverseStructuresStructureIDForbidden
const GetUniverseStructuresStructureIDForbiddenCode int = 403

/*GetUniverseStructuresStructureIDForbidden Forbidden

swagger:response getUniverseStructuresStructureIdForbidden
*/
type GetUniverseStructuresStructureIDForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GetUniverseStructuresStructureIDForbiddenBody `json:"body,omitempty"`
}

// NewGetUniverseStructuresStructureIDForbidden creates GetUniverseStructuresStructureIDForbidden with default headers values
func NewGetUniverseStructuresStructureIDForbidden() *GetUniverseStructuresStructureIDForbidden {
	return &GetUniverseStructuresStructureIDForbidden{}
}

// WithPayload adds the payload to the get universe structures structure Id forbidden response
func (o *GetUniverseStructuresStructureIDForbidden) WithPayload(payload *models.GetUniverseStructuresStructureIDForbiddenBody) *GetUniverseStructuresStructureIDForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get universe structures structure Id forbidden response
func (o *GetUniverseStructuresStructureIDForbidden) SetPayload(payload *models.GetUniverseStructuresStructureIDForbiddenBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUniverseStructuresStructureIDForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUniverseStructuresStructureIDNotFoundCode is the HTTP code returned for type GetUniverseStructuresStructureIDNotFound
const GetUniverseStructuresStructureIDNotFoundCode int = 404

/*GetUniverseStructuresStructureIDNotFound Structure not found

swagger:response getUniverseStructuresStructureIdNotFound
*/
type GetUniverseStructuresStructureIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.GetUniverseStructuresStructureIDNotFoundBody `json:"body,omitempty"`
}

// NewGetUniverseStructuresStructureIDNotFound creates GetUniverseStructuresStructureIDNotFound with default headers values
func NewGetUniverseStructuresStructureIDNotFound() *GetUniverseStructuresStructureIDNotFound {
	return &GetUniverseStructuresStructureIDNotFound{}
}

// WithPayload adds the payload to the get universe structures structure Id not found response
func (o *GetUniverseStructuresStructureIDNotFound) WithPayload(payload *models.GetUniverseStructuresStructureIDNotFoundBody) *GetUniverseStructuresStructureIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get universe structures structure Id not found response
func (o *GetUniverseStructuresStructureIDNotFound) SetPayload(payload *models.GetUniverseStructuresStructureIDNotFoundBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUniverseStructuresStructureIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUniverseStructuresStructureIDInternalServerErrorCode is the HTTP code returned for type GetUniverseStructuresStructureIDInternalServerError
const GetUniverseStructuresStructureIDInternalServerErrorCode int = 500

/*GetUniverseStructuresStructureIDInternalServerError Internal server error

swagger:response getUniverseStructuresStructureIdInternalServerError
*/
type GetUniverseStructuresStructureIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.GetUniverseStructuresStructureIDInternalServerErrorBody `json:"body,omitempty"`
}

// NewGetUniverseStructuresStructureIDInternalServerError creates GetUniverseStructuresStructureIDInternalServerError with default headers values
func NewGetUniverseStructuresStructureIDInternalServerError() *GetUniverseStructuresStructureIDInternalServerError {
	return &GetUniverseStructuresStructureIDInternalServerError{}
}

// WithPayload adds the payload to the get universe structures structure Id internal server error response
func (o *GetUniverseStructuresStructureIDInternalServerError) WithPayload(payload *models.GetUniverseStructuresStructureIDInternalServerErrorBody) *GetUniverseStructuresStructureIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get universe structures structure Id internal server error response
func (o *GetUniverseStructuresStructureIDInternalServerError) SetPayload(payload *models.GetUniverseStructuresStructureIDInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUniverseStructuresStructureIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
