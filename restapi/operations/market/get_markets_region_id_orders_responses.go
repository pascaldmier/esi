package market

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/pascaldmier/esi/models"
)

// GetMarketsRegionIDOrdersOKCode is the HTTP code returned for type GetMarketsRegionIDOrdersOK
const GetMarketsRegionIDOrdersOKCode int = 200

/*GetMarketsRegionIDOrdersOK A list of orders

swagger:response getMarketsRegionIdOrdersOK
*/
type GetMarketsRegionIDOrdersOK struct {
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
	Payload models.GetMarketsRegionIDOrdersOKBody `json:"body,omitempty"`
}

// NewGetMarketsRegionIDOrdersOK creates GetMarketsRegionIDOrdersOK with default headers values
func NewGetMarketsRegionIDOrdersOK() *GetMarketsRegionIDOrdersOK {
	return &GetMarketsRegionIDOrdersOK{}
}

// WithCacheControl adds the cacheControl to the get markets region Id orders o k response
func (o *GetMarketsRegionIDOrdersOK) WithCacheControl(cacheControl string) *GetMarketsRegionIDOrdersOK {
	o.CacheControl = cacheControl
	return o
}

// SetCacheControl sets the cacheControl to the get markets region Id orders o k response
func (o *GetMarketsRegionIDOrdersOK) SetCacheControl(cacheControl string) {
	o.CacheControl = cacheControl
}

// WithExpires adds the expires to the get markets region Id orders o k response
func (o *GetMarketsRegionIDOrdersOK) WithExpires(expires string) *GetMarketsRegionIDOrdersOK {
	o.Expires = expires
	return o
}

// SetExpires sets the expires to the get markets region Id orders o k response
func (o *GetMarketsRegionIDOrdersOK) SetExpires(expires string) {
	o.Expires = expires
}

// WithLastModified adds the lastModified to the get markets region Id orders o k response
func (o *GetMarketsRegionIDOrdersOK) WithLastModified(lastModified string) *GetMarketsRegionIDOrdersOK {
	o.LastModified = lastModified
	return o
}

// SetLastModified sets the lastModified to the get markets region Id orders o k response
func (o *GetMarketsRegionIDOrdersOK) SetLastModified(lastModified string) {
	o.LastModified = lastModified
}

// WithPayload adds the payload to the get markets region Id orders o k response
func (o *GetMarketsRegionIDOrdersOK) WithPayload(payload models.GetMarketsRegionIDOrdersOKBody) *GetMarketsRegionIDOrdersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get markets region Id orders o k response
func (o *GetMarketsRegionIDOrdersOK) SetPayload(payload models.GetMarketsRegionIDOrdersOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMarketsRegionIDOrdersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
		payload = make(models.GetMarketsRegionIDOrdersOKBody, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetMarketsRegionIDOrdersUnprocessableEntityCode is the HTTP code returned for type GetMarketsRegionIDOrdersUnprocessableEntity
const GetMarketsRegionIDOrdersUnprocessableEntityCode int = 422

/*GetMarketsRegionIDOrdersUnprocessableEntity Not found

swagger:response getMarketsRegionIdOrdersUnprocessableEntity
*/
type GetMarketsRegionIDOrdersUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.GetMarketsRegionIDOrdersUnprocessableEntityBody `json:"body,omitempty"`
}

// NewGetMarketsRegionIDOrdersUnprocessableEntity creates GetMarketsRegionIDOrdersUnprocessableEntity with default headers values
func NewGetMarketsRegionIDOrdersUnprocessableEntity() *GetMarketsRegionIDOrdersUnprocessableEntity {
	return &GetMarketsRegionIDOrdersUnprocessableEntity{}
}

// WithPayload adds the payload to the get markets region Id orders unprocessable entity response
func (o *GetMarketsRegionIDOrdersUnprocessableEntity) WithPayload(payload *models.GetMarketsRegionIDOrdersUnprocessableEntityBody) *GetMarketsRegionIDOrdersUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get markets region Id orders unprocessable entity response
func (o *GetMarketsRegionIDOrdersUnprocessableEntity) SetPayload(payload *models.GetMarketsRegionIDOrdersUnprocessableEntityBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMarketsRegionIDOrdersUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetMarketsRegionIDOrdersInternalServerErrorCode is the HTTP code returned for type GetMarketsRegionIDOrdersInternalServerError
const GetMarketsRegionIDOrdersInternalServerErrorCode int = 500

/*GetMarketsRegionIDOrdersInternalServerError Internal server error

swagger:response getMarketsRegionIdOrdersInternalServerError
*/
type GetMarketsRegionIDOrdersInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.GetMarketsRegionIDOrdersInternalServerErrorBody `json:"body,omitempty"`
}

// NewGetMarketsRegionIDOrdersInternalServerError creates GetMarketsRegionIDOrdersInternalServerError with default headers values
func NewGetMarketsRegionIDOrdersInternalServerError() *GetMarketsRegionIDOrdersInternalServerError {
	return &GetMarketsRegionIDOrdersInternalServerError{}
}

// WithPayload adds the payload to the get markets region Id orders internal server error response
func (o *GetMarketsRegionIDOrdersInternalServerError) WithPayload(payload *models.GetMarketsRegionIDOrdersInternalServerErrorBody) *GetMarketsRegionIDOrdersInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get markets region Id orders internal server error response
func (o *GetMarketsRegionIDOrdersInternalServerError) SetPayload(payload *models.GetMarketsRegionIDOrdersInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMarketsRegionIDOrdersInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
