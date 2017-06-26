package universe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/pascaldmier/esi/models"
)

// GetUniverseStargatesStargateIDOKCode is the HTTP code returned for type GetUniverseStargatesStargateIDOK
const GetUniverseStargatesStargateIDOKCode int = 200

/*GetUniverseStargatesStargateIDOK Information about a stargate

swagger:response getUniverseStargatesStargateIdOK
*/
type GetUniverseStargatesStargateIDOK struct {
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
	Payload *models.GetUniverseStargatesStargateIDOKBody `json:"body,omitempty"`
}

// NewGetUniverseStargatesStargateIDOK creates GetUniverseStargatesStargateIDOK with default headers values
func NewGetUniverseStargatesStargateIDOK() *GetUniverseStargatesStargateIDOK {
	return &GetUniverseStargatesStargateIDOK{}
}

// WithCacheControl adds the cacheControl to the get universe stargates stargate Id o k response
func (o *GetUniverseStargatesStargateIDOK) WithCacheControl(cacheControl string) *GetUniverseStargatesStargateIDOK {
	o.CacheControl = cacheControl
	return o
}

// SetCacheControl sets the cacheControl to the get universe stargates stargate Id o k response
func (o *GetUniverseStargatesStargateIDOK) SetCacheControl(cacheControl string) {
	o.CacheControl = cacheControl
}

// WithExpires adds the expires to the get universe stargates stargate Id o k response
func (o *GetUniverseStargatesStargateIDOK) WithExpires(expires string) *GetUniverseStargatesStargateIDOK {
	o.Expires = expires
	return o
}

// SetExpires sets the expires to the get universe stargates stargate Id o k response
func (o *GetUniverseStargatesStargateIDOK) SetExpires(expires string) {
	o.Expires = expires
}

// WithLastModified adds the lastModified to the get universe stargates stargate Id o k response
func (o *GetUniverseStargatesStargateIDOK) WithLastModified(lastModified string) *GetUniverseStargatesStargateIDOK {
	o.LastModified = lastModified
	return o
}

// SetLastModified sets the lastModified to the get universe stargates stargate Id o k response
func (o *GetUniverseStargatesStargateIDOK) SetLastModified(lastModified string) {
	o.LastModified = lastModified
}

// WithPayload adds the payload to the get universe stargates stargate Id o k response
func (o *GetUniverseStargatesStargateIDOK) WithPayload(payload *models.GetUniverseStargatesStargateIDOKBody) *GetUniverseStargatesStargateIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get universe stargates stargate Id o k response
func (o *GetUniverseStargatesStargateIDOK) SetPayload(payload *models.GetUniverseStargatesStargateIDOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUniverseStargatesStargateIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetUniverseStargatesStargateIDNotFoundCode is the HTTP code returned for type GetUniverseStargatesStargateIDNotFound
const GetUniverseStargatesStargateIDNotFoundCode int = 404

/*GetUniverseStargatesStargateIDNotFound Stargate not found

swagger:response getUniverseStargatesStargateIdNotFound
*/
type GetUniverseStargatesStargateIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.GetUniverseStargatesStargateIDNotFoundBody `json:"body,omitempty"`
}

// NewGetUniverseStargatesStargateIDNotFound creates GetUniverseStargatesStargateIDNotFound with default headers values
func NewGetUniverseStargatesStargateIDNotFound() *GetUniverseStargatesStargateIDNotFound {
	return &GetUniverseStargatesStargateIDNotFound{}
}

// WithPayload adds the payload to the get universe stargates stargate Id not found response
func (o *GetUniverseStargatesStargateIDNotFound) WithPayload(payload *models.GetUniverseStargatesStargateIDNotFoundBody) *GetUniverseStargatesStargateIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get universe stargates stargate Id not found response
func (o *GetUniverseStargatesStargateIDNotFound) SetPayload(payload *models.GetUniverseStargatesStargateIDNotFoundBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUniverseStargatesStargateIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUniverseStargatesStargateIDInternalServerErrorCode is the HTTP code returned for type GetUniverseStargatesStargateIDInternalServerError
const GetUniverseStargatesStargateIDInternalServerErrorCode int = 500

/*GetUniverseStargatesStargateIDInternalServerError Internal server error

swagger:response getUniverseStargatesStargateIdInternalServerError
*/
type GetUniverseStargatesStargateIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.GetUniverseStargatesStargateIDInternalServerErrorBody `json:"body,omitempty"`
}

// NewGetUniverseStargatesStargateIDInternalServerError creates GetUniverseStargatesStargateIDInternalServerError with default headers values
func NewGetUniverseStargatesStargateIDInternalServerError() *GetUniverseStargatesStargateIDInternalServerError {
	return &GetUniverseStargatesStargateIDInternalServerError{}
}

// WithPayload adds the payload to the get universe stargates stargate Id internal server error response
func (o *GetUniverseStargatesStargateIDInternalServerError) WithPayload(payload *models.GetUniverseStargatesStargateIDInternalServerErrorBody) *GetUniverseStargatesStargateIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get universe stargates stargate Id internal server error response
func (o *GetUniverseStargatesStargateIDInternalServerError) SetPayload(payload *models.GetUniverseStargatesStargateIDInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUniverseStargatesStargateIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}