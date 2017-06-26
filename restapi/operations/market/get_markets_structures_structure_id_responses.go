package market

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/pascaldmier/esi/models"
)

// GetMarketsStructuresStructureIDOKCode is the HTTP code returned for type GetMarketsStructuresStructureIDOK
const GetMarketsStructuresStructureIDOKCode int = 200

/*GetMarketsStructuresStructureIDOK A list of orders

swagger:response getMarketsStructuresStructureIdOK
*/
type GetMarketsStructuresStructureIDOK struct {
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
	Payload models.GetMarketsStructuresStructureIDOKBody `json:"body,omitempty"`
}

// NewGetMarketsStructuresStructureIDOK creates GetMarketsStructuresStructureIDOK with default headers values
func NewGetMarketsStructuresStructureIDOK() *GetMarketsStructuresStructureIDOK {
	return &GetMarketsStructuresStructureIDOK{}
}

// WithCacheControl adds the cacheControl to the get markets structures structure Id o k response
func (o *GetMarketsStructuresStructureIDOK) WithCacheControl(cacheControl string) *GetMarketsStructuresStructureIDOK {
	o.CacheControl = cacheControl
	return o
}

// SetCacheControl sets the cacheControl to the get markets structures structure Id o k response
func (o *GetMarketsStructuresStructureIDOK) SetCacheControl(cacheControl string) {
	o.CacheControl = cacheControl
}

// WithExpires adds the expires to the get markets structures structure Id o k response
func (o *GetMarketsStructuresStructureIDOK) WithExpires(expires string) *GetMarketsStructuresStructureIDOK {
	o.Expires = expires
	return o
}

// SetExpires sets the expires to the get markets structures structure Id o k response
func (o *GetMarketsStructuresStructureIDOK) SetExpires(expires string) {
	o.Expires = expires
}

// WithLastModified adds the lastModified to the get markets structures structure Id o k response
func (o *GetMarketsStructuresStructureIDOK) WithLastModified(lastModified string) *GetMarketsStructuresStructureIDOK {
	o.LastModified = lastModified
	return o
}

// SetLastModified sets the lastModified to the get markets structures structure Id o k response
func (o *GetMarketsStructuresStructureIDOK) SetLastModified(lastModified string) {
	o.LastModified = lastModified
}

// WithPayload adds the payload to the get markets structures structure Id o k response
func (o *GetMarketsStructuresStructureIDOK) WithPayload(payload models.GetMarketsStructuresStructureIDOKBody) *GetMarketsStructuresStructureIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get markets structures structure Id o k response
func (o *GetMarketsStructuresStructureIDOK) SetPayload(payload models.GetMarketsStructuresStructureIDOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMarketsStructuresStructureIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
	payload := o.Payload
	if payload == nil {
		payload = make(models.GetMarketsStructuresStructureIDOKBody, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetMarketsStructuresStructureIDForbiddenCode is the HTTP code returned for type GetMarketsStructuresStructureIDForbidden
const GetMarketsStructuresStructureIDForbiddenCode int = 403

/*GetMarketsStructuresStructureIDForbidden Forbidden

swagger:response getMarketsStructuresStructureIdForbidden
*/
type GetMarketsStructuresStructureIDForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GetMarketsStructuresStructureIDForbiddenBody `json:"body,omitempty"`
}

// NewGetMarketsStructuresStructureIDForbidden creates GetMarketsStructuresStructureIDForbidden with default headers values
func NewGetMarketsStructuresStructureIDForbidden() *GetMarketsStructuresStructureIDForbidden {
	return &GetMarketsStructuresStructureIDForbidden{}
}

// WithPayload adds the payload to the get markets structures structure Id forbidden response
func (o *GetMarketsStructuresStructureIDForbidden) WithPayload(payload *models.GetMarketsStructuresStructureIDForbiddenBody) *GetMarketsStructuresStructureIDForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get markets structures structure Id forbidden response
func (o *GetMarketsStructuresStructureIDForbidden) SetPayload(payload *models.GetMarketsStructuresStructureIDForbiddenBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMarketsStructuresStructureIDForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetMarketsStructuresStructureIDInternalServerErrorCode is the HTTP code returned for type GetMarketsStructuresStructureIDInternalServerError
const GetMarketsStructuresStructureIDInternalServerErrorCode int = 500

/*GetMarketsStructuresStructureIDInternalServerError Internal server error

swagger:response getMarketsStructuresStructureIdInternalServerError
*/
type GetMarketsStructuresStructureIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.GetMarketsStructuresStructureIDInternalServerErrorBody `json:"body,omitempty"`
}

// NewGetMarketsStructuresStructureIDInternalServerError creates GetMarketsStructuresStructureIDInternalServerError with default headers values
func NewGetMarketsStructuresStructureIDInternalServerError() *GetMarketsStructuresStructureIDInternalServerError {
	return &GetMarketsStructuresStructureIDInternalServerError{}
}

// WithPayload adds the payload to the get markets structures structure Id internal server error response
func (o *GetMarketsStructuresStructureIDInternalServerError) WithPayload(payload *models.GetMarketsStructuresStructureIDInternalServerErrorBody) *GetMarketsStructuresStructureIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get markets structures structure Id internal server error response
func (o *GetMarketsStructuresStructureIDInternalServerError) SetPayload(payload *models.GetMarketsStructuresStructureIDInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMarketsStructuresStructureIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}