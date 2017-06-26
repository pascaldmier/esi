package market

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/pascaldmier/esi/models"
)

// GetMarketsRegionIDHistoryOKCode is the HTTP code returned for type GetMarketsRegionIDHistoryOK
const GetMarketsRegionIDHistoryOKCode int = 200

/*GetMarketsRegionIDHistoryOK A list of historical market statistics

swagger:response getMarketsRegionIdHistoryOK
*/
type GetMarketsRegionIDHistoryOK struct {
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
	Payload models.GetMarketsRegionIDHistoryOKBody `json:"body,omitempty"`
}

// NewGetMarketsRegionIDHistoryOK creates GetMarketsRegionIDHistoryOK with default headers values
func NewGetMarketsRegionIDHistoryOK() *GetMarketsRegionIDHistoryOK {
	return &GetMarketsRegionIDHistoryOK{}
}

// WithCacheControl adds the cacheControl to the get markets region Id history o k response
func (o *GetMarketsRegionIDHistoryOK) WithCacheControl(cacheControl string) *GetMarketsRegionIDHistoryOK {
	o.CacheControl = cacheControl
	return o
}

// SetCacheControl sets the cacheControl to the get markets region Id history o k response
func (o *GetMarketsRegionIDHistoryOK) SetCacheControl(cacheControl string) {
	o.CacheControl = cacheControl
}

// WithExpires adds the expires to the get markets region Id history o k response
func (o *GetMarketsRegionIDHistoryOK) WithExpires(expires string) *GetMarketsRegionIDHistoryOK {
	o.Expires = expires
	return o
}

// SetExpires sets the expires to the get markets region Id history o k response
func (o *GetMarketsRegionIDHistoryOK) SetExpires(expires string) {
	o.Expires = expires
}

// WithLastModified adds the lastModified to the get markets region Id history o k response
func (o *GetMarketsRegionIDHistoryOK) WithLastModified(lastModified string) *GetMarketsRegionIDHistoryOK {
	o.LastModified = lastModified
	return o
}

// SetLastModified sets the lastModified to the get markets region Id history o k response
func (o *GetMarketsRegionIDHistoryOK) SetLastModified(lastModified string) {
	o.LastModified = lastModified
}

// WithPayload adds the payload to the get markets region Id history o k response
func (o *GetMarketsRegionIDHistoryOK) WithPayload(payload models.GetMarketsRegionIDHistoryOKBody) *GetMarketsRegionIDHistoryOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get markets region Id history o k response
func (o *GetMarketsRegionIDHistoryOK) SetPayload(payload models.GetMarketsRegionIDHistoryOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMarketsRegionIDHistoryOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
		payload = make(models.GetMarketsRegionIDHistoryOKBody, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetMarketsRegionIDHistoryUnprocessableEntityCode is the HTTP code returned for type GetMarketsRegionIDHistoryUnprocessableEntity
const GetMarketsRegionIDHistoryUnprocessableEntityCode int = 422

/*GetMarketsRegionIDHistoryUnprocessableEntity Not found

swagger:response getMarketsRegionIdHistoryUnprocessableEntity
*/
type GetMarketsRegionIDHistoryUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.GetMarketsRegionIDHistoryUnprocessableEntityBody `json:"body,omitempty"`
}

// NewGetMarketsRegionIDHistoryUnprocessableEntity creates GetMarketsRegionIDHistoryUnprocessableEntity with default headers values
func NewGetMarketsRegionIDHistoryUnprocessableEntity() *GetMarketsRegionIDHistoryUnprocessableEntity {
	return &GetMarketsRegionIDHistoryUnprocessableEntity{}
}

// WithPayload adds the payload to the get markets region Id history unprocessable entity response
func (o *GetMarketsRegionIDHistoryUnprocessableEntity) WithPayload(payload *models.GetMarketsRegionIDHistoryUnprocessableEntityBody) *GetMarketsRegionIDHistoryUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get markets region Id history unprocessable entity response
func (o *GetMarketsRegionIDHistoryUnprocessableEntity) SetPayload(payload *models.GetMarketsRegionIDHistoryUnprocessableEntityBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMarketsRegionIDHistoryUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetMarketsRegionIDHistoryInternalServerErrorCode is the HTTP code returned for type GetMarketsRegionIDHistoryInternalServerError
const GetMarketsRegionIDHistoryInternalServerErrorCode int = 500

/*GetMarketsRegionIDHistoryInternalServerError Internal server error

swagger:response getMarketsRegionIdHistoryInternalServerError
*/
type GetMarketsRegionIDHistoryInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.GetMarketsRegionIDHistoryInternalServerErrorBody `json:"body,omitempty"`
}

// NewGetMarketsRegionIDHistoryInternalServerError creates GetMarketsRegionIDHistoryInternalServerError with default headers values
func NewGetMarketsRegionIDHistoryInternalServerError() *GetMarketsRegionIDHistoryInternalServerError {
	return &GetMarketsRegionIDHistoryInternalServerError{}
}

// WithPayload adds the payload to the get markets region Id history internal server error response
func (o *GetMarketsRegionIDHistoryInternalServerError) WithPayload(payload *models.GetMarketsRegionIDHistoryInternalServerErrorBody) *GetMarketsRegionIDHistoryInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get markets region Id history internal server error response
func (o *GetMarketsRegionIDHistoryInternalServerError) SetPayload(payload *models.GetMarketsRegionIDHistoryInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMarketsRegionIDHistoryInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
