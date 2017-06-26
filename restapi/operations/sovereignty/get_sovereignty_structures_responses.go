package sovereignty

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/pascaldmier/esi/models"
)

// GetSovereigntyStructuresOKCode is the HTTP code returned for type GetSovereigntyStructuresOK
const GetSovereigntyStructuresOKCode int = 200

/*GetSovereigntyStructuresOK A list of sovereignty structures

swagger:response getSovereigntyStructuresOK
*/
type GetSovereigntyStructuresOK struct {
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
	Payload models.GetSovereigntyStructuresOKBody `json:"body,omitempty"`
}

// NewGetSovereigntyStructuresOK creates GetSovereigntyStructuresOK with default headers values
func NewGetSovereigntyStructuresOK() *GetSovereigntyStructuresOK {
	return &GetSovereigntyStructuresOK{}
}

// WithCacheControl adds the cacheControl to the get sovereignty structures o k response
func (o *GetSovereigntyStructuresOK) WithCacheControl(cacheControl string) *GetSovereigntyStructuresOK {
	o.CacheControl = cacheControl
	return o
}

// SetCacheControl sets the cacheControl to the get sovereignty structures o k response
func (o *GetSovereigntyStructuresOK) SetCacheControl(cacheControl string) {
	o.CacheControl = cacheControl
}

// WithExpires adds the expires to the get sovereignty structures o k response
func (o *GetSovereigntyStructuresOK) WithExpires(expires string) *GetSovereigntyStructuresOK {
	o.Expires = expires
	return o
}

// SetExpires sets the expires to the get sovereignty structures o k response
func (o *GetSovereigntyStructuresOK) SetExpires(expires string) {
	o.Expires = expires
}

// WithLastModified adds the lastModified to the get sovereignty structures o k response
func (o *GetSovereigntyStructuresOK) WithLastModified(lastModified string) *GetSovereigntyStructuresOK {
	o.LastModified = lastModified
	return o
}

// SetLastModified sets the lastModified to the get sovereignty structures o k response
func (o *GetSovereigntyStructuresOK) SetLastModified(lastModified string) {
	o.LastModified = lastModified
}

// WithPayload adds the payload to the get sovereignty structures o k response
func (o *GetSovereigntyStructuresOK) WithPayload(payload models.GetSovereigntyStructuresOKBody) *GetSovereigntyStructuresOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get sovereignty structures o k response
func (o *GetSovereigntyStructuresOK) SetPayload(payload models.GetSovereigntyStructuresOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSovereigntyStructuresOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
		payload = make(models.GetSovereigntyStructuresOKBody, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetSovereigntyStructuresInternalServerErrorCode is the HTTP code returned for type GetSovereigntyStructuresInternalServerError
const GetSovereigntyStructuresInternalServerErrorCode int = 500

/*GetSovereigntyStructuresInternalServerError Internal server error

swagger:response getSovereigntyStructuresInternalServerError
*/
type GetSovereigntyStructuresInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.GetSovereigntyStructuresInternalServerErrorBody `json:"body,omitempty"`
}

// NewGetSovereigntyStructuresInternalServerError creates GetSovereigntyStructuresInternalServerError with default headers values
func NewGetSovereigntyStructuresInternalServerError() *GetSovereigntyStructuresInternalServerError {
	return &GetSovereigntyStructuresInternalServerError{}
}

// WithPayload adds the payload to the get sovereignty structures internal server error response
func (o *GetSovereigntyStructuresInternalServerError) WithPayload(payload *models.GetSovereigntyStructuresInternalServerErrorBody) *GetSovereigntyStructuresInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get sovereignty structures internal server error response
func (o *GetSovereigntyStructuresInternalServerError) SetPayload(payload *models.GetSovereigntyStructuresInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSovereigntyStructuresInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
