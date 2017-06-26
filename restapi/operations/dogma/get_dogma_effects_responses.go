package dogma

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/pascaldmier/esi/models"
)

// GetDogmaEffectsOKCode is the HTTP code returned for type GetDogmaEffectsOK
const GetDogmaEffectsOKCode int = 200

/*GetDogmaEffectsOK A list of dogma effect ids

swagger:response getDogmaEffectsOK
*/
type GetDogmaEffectsOK struct {
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

	/*200 ok array
	  Max Items: 10000
	  In: Body
	*/
	Payload []int32 `json:"body,omitempty"`
}

// NewGetDogmaEffectsOK creates GetDogmaEffectsOK with default headers values
func NewGetDogmaEffectsOK() *GetDogmaEffectsOK {
	return &GetDogmaEffectsOK{}
}

// WithCacheControl adds the cacheControl to the get dogma effects o k response
func (o *GetDogmaEffectsOK) WithCacheControl(cacheControl string) *GetDogmaEffectsOK {
	o.CacheControl = cacheControl
	return o
}

// SetCacheControl sets the cacheControl to the get dogma effects o k response
func (o *GetDogmaEffectsOK) SetCacheControl(cacheControl string) {
	o.CacheControl = cacheControl
}

// WithExpires adds the expires to the get dogma effects o k response
func (o *GetDogmaEffectsOK) WithExpires(expires string) *GetDogmaEffectsOK {
	o.Expires = expires
	return o
}

// SetExpires sets the expires to the get dogma effects o k response
func (o *GetDogmaEffectsOK) SetExpires(expires string) {
	o.Expires = expires
}

// WithLastModified adds the lastModified to the get dogma effects o k response
func (o *GetDogmaEffectsOK) WithLastModified(lastModified string) *GetDogmaEffectsOK {
	o.LastModified = lastModified
	return o
}

// SetLastModified sets the lastModified to the get dogma effects o k response
func (o *GetDogmaEffectsOK) SetLastModified(lastModified string) {
	o.LastModified = lastModified
}

// WithPayload adds the payload to the get dogma effects o k response
func (o *GetDogmaEffectsOK) WithPayload(payload []int32) *GetDogmaEffectsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get dogma effects o k response
func (o *GetDogmaEffectsOK) SetPayload(payload []int32) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDogmaEffectsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
		payload = make([]int32, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetDogmaEffectsInternalServerErrorCode is the HTTP code returned for type GetDogmaEffectsInternalServerError
const GetDogmaEffectsInternalServerErrorCode int = 500

/*GetDogmaEffectsInternalServerError Internal server error

swagger:response getDogmaEffectsInternalServerError
*/
type GetDogmaEffectsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.GetDogmaEffectsInternalServerErrorBody `json:"body,omitempty"`
}

// NewGetDogmaEffectsInternalServerError creates GetDogmaEffectsInternalServerError with default headers values
func NewGetDogmaEffectsInternalServerError() *GetDogmaEffectsInternalServerError {
	return &GetDogmaEffectsInternalServerError{}
}

// WithPayload adds the payload to the get dogma effects internal server error response
func (o *GetDogmaEffectsInternalServerError) WithPayload(payload *models.GetDogmaEffectsInternalServerErrorBody) *GetDogmaEffectsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get dogma effects internal server error response
func (o *GetDogmaEffectsInternalServerError) SetPayload(payload *models.GetDogmaEffectsInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDogmaEffectsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}