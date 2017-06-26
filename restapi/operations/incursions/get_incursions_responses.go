package incursions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/pascaldmier/esi/models"
)

// GetIncursionsOKCode is the HTTP code returned for type GetIncursionsOK
const GetIncursionsOKCode int = 200

/*GetIncursionsOK A list of incursions

swagger:response getIncursionsOK
*/
type GetIncursionsOK struct {
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
	Payload models.GetIncursionsOKBody `json:"body,omitempty"`
}

// NewGetIncursionsOK creates GetIncursionsOK with default headers values
func NewGetIncursionsOK() *GetIncursionsOK {
	return &GetIncursionsOK{}
}

// WithCacheControl adds the cacheControl to the get incursions o k response
func (o *GetIncursionsOK) WithCacheControl(cacheControl string) *GetIncursionsOK {
	o.CacheControl = cacheControl
	return o
}

// SetCacheControl sets the cacheControl to the get incursions o k response
func (o *GetIncursionsOK) SetCacheControl(cacheControl string) {
	o.CacheControl = cacheControl
}

// WithExpires adds the expires to the get incursions o k response
func (o *GetIncursionsOK) WithExpires(expires string) *GetIncursionsOK {
	o.Expires = expires
	return o
}

// SetExpires sets the expires to the get incursions o k response
func (o *GetIncursionsOK) SetExpires(expires string) {
	o.Expires = expires
}

// WithLastModified adds the lastModified to the get incursions o k response
func (o *GetIncursionsOK) WithLastModified(lastModified string) *GetIncursionsOK {
	o.LastModified = lastModified
	return o
}

// SetLastModified sets the lastModified to the get incursions o k response
func (o *GetIncursionsOK) SetLastModified(lastModified string) {
	o.LastModified = lastModified
}

// WithPayload adds the payload to the get incursions o k response
func (o *GetIncursionsOK) WithPayload(payload models.GetIncursionsOKBody) *GetIncursionsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get incursions o k response
func (o *GetIncursionsOK) SetPayload(payload models.GetIncursionsOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetIncursionsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
		payload = make(models.GetIncursionsOKBody, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetIncursionsInternalServerErrorCode is the HTTP code returned for type GetIncursionsInternalServerError
const GetIncursionsInternalServerErrorCode int = 500

/*GetIncursionsInternalServerError Internal server error

swagger:response getIncursionsInternalServerError
*/
type GetIncursionsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.GetIncursionsInternalServerErrorBody `json:"body,omitempty"`
}

// NewGetIncursionsInternalServerError creates GetIncursionsInternalServerError with default headers values
func NewGetIncursionsInternalServerError() *GetIncursionsInternalServerError {
	return &GetIncursionsInternalServerError{}
}

// WithPayload adds the payload to the get incursions internal server error response
func (o *GetIncursionsInternalServerError) WithPayload(payload *models.GetIncursionsInternalServerErrorBody) *GetIncursionsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get incursions internal server error response
func (o *GetIncursionsInternalServerError) SetPayload(payload *models.GetIncursionsInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetIncursionsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
