package universe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/pascaldmier/esi/models"
)

// GetUniverseStationsStationIDOKCode is the HTTP code returned for type GetUniverseStationsStationIDOK
const GetUniverseStationsStationIDOKCode int = 200

/*GetUniverseStationsStationIDOK Information about a station

swagger:response getUniverseStationsStationIdOK
*/
type GetUniverseStationsStationIDOK struct {
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
	Payload *models.GetUniverseStationsStationIDOKBody `json:"body,omitempty"`
}

// NewGetUniverseStationsStationIDOK creates GetUniverseStationsStationIDOK with default headers values
func NewGetUniverseStationsStationIDOK() *GetUniverseStationsStationIDOK {
	return &GetUniverseStationsStationIDOK{}
}

// WithCacheControl adds the cacheControl to the get universe stations station Id o k response
func (o *GetUniverseStationsStationIDOK) WithCacheControl(cacheControl string) *GetUniverseStationsStationIDOK {
	o.CacheControl = cacheControl
	return o
}

// SetCacheControl sets the cacheControl to the get universe stations station Id o k response
func (o *GetUniverseStationsStationIDOK) SetCacheControl(cacheControl string) {
	o.CacheControl = cacheControl
}

// WithExpires adds the expires to the get universe stations station Id o k response
func (o *GetUniverseStationsStationIDOK) WithExpires(expires string) *GetUniverseStationsStationIDOK {
	o.Expires = expires
	return o
}

// SetExpires sets the expires to the get universe stations station Id o k response
func (o *GetUniverseStationsStationIDOK) SetExpires(expires string) {
	o.Expires = expires
}

// WithLastModified adds the lastModified to the get universe stations station Id o k response
func (o *GetUniverseStationsStationIDOK) WithLastModified(lastModified string) *GetUniverseStationsStationIDOK {
	o.LastModified = lastModified
	return o
}

// SetLastModified sets the lastModified to the get universe stations station Id o k response
func (o *GetUniverseStationsStationIDOK) SetLastModified(lastModified string) {
	o.LastModified = lastModified
}

// WithPayload adds the payload to the get universe stations station Id o k response
func (o *GetUniverseStationsStationIDOK) WithPayload(payload *models.GetUniverseStationsStationIDOKBody) *GetUniverseStationsStationIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get universe stations station Id o k response
func (o *GetUniverseStationsStationIDOK) SetPayload(payload *models.GetUniverseStationsStationIDOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUniverseStationsStationIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetUniverseStationsStationIDNotFoundCode is the HTTP code returned for type GetUniverseStationsStationIDNotFound
const GetUniverseStationsStationIDNotFoundCode int = 404

/*GetUniverseStationsStationIDNotFound Station not found

swagger:response getUniverseStationsStationIdNotFound
*/
type GetUniverseStationsStationIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.GetUniverseStationsStationIDNotFoundBody `json:"body,omitempty"`
}

// NewGetUniverseStationsStationIDNotFound creates GetUniverseStationsStationIDNotFound with default headers values
func NewGetUniverseStationsStationIDNotFound() *GetUniverseStationsStationIDNotFound {
	return &GetUniverseStationsStationIDNotFound{}
}

// WithPayload adds the payload to the get universe stations station Id not found response
func (o *GetUniverseStationsStationIDNotFound) WithPayload(payload *models.GetUniverseStationsStationIDNotFoundBody) *GetUniverseStationsStationIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get universe stations station Id not found response
func (o *GetUniverseStationsStationIDNotFound) SetPayload(payload *models.GetUniverseStationsStationIDNotFoundBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUniverseStationsStationIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUniverseStationsStationIDInternalServerErrorCode is the HTTP code returned for type GetUniverseStationsStationIDInternalServerError
const GetUniverseStationsStationIDInternalServerErrorCode int = 500

/*GetUniverseStationsStationIDInternalServerError Internal server error

swagger:response getUniverseStationsStationIdInternalServerError
*/
type GetUniverseStationsStationIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.GetUniverseStationsStationIDInternalServerErrorBody `json:"body,omitempty"`
}

// NewGetUniverseStationsStationIDInternalServerError creates GetUniverseStationsStationIDInternalServerError with default headers values
func NewGetUniverseStationsStationIDInternalServerError() *GetUniverseStationsStationIDInternalServerError {
	return &GetUniverseStationsStationIDInternalServerError{}
}

// WithPayload adds the payload to the get universe stations station Id internal server error response
func (o *GetUniverseStationsStationIDInternalServerError) WithPayload(payload *models.GetUniverseStationsStationIDInternalServerErrorBody) *GetUniverseStationsStationIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get universe stations station Id internal server error response
func (o *GetUniverseStationsStationIDInternalServerError) SetPayload(payload *models.GetUniverseStationsStationIDInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUniverseStationsStationIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}