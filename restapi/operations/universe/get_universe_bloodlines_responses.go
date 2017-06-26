package universe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/pascaldmier/esi/models"
)

// GetUniverseBloodlinesOKCode is the HTTP code returned for type GetUniverseBloodlinesOK
const GetUniverseBloodlinesOKCode int = 200

/*GetUniverseBloodlinesOK A list of bloodlines

swagger:response getUniverseBloodlinesOK
*/
type GetUniverseBloodlinesOK struct {
	/*The caching mechanism used
	  Required: true
	*/
	CacheControl string `json:"Cache-Control"`
	/*The language used in the response
	  Required: true
	*/
	ContentLanguage string `json:"Content-Language"`
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
	Payload models.GetUniverseBloodlinesOKBody `json:"body,omitempty"`
}

// NewGetUniverseBloodlinesOK creates GetUniverseBloodlinesOK with default headers values
func NewGetUniverseBloodlinesOK() *GetUniverseBloodlinesOK {
	return &GetUniverseBloodlinesOK{}
}

// WithCacheControl adds the cacheControl to the get universe bloodlines o k response
func (o *GetUniverseBloodlinesOK) WithCacheControl(cacheControl string) *GetUniverseBloodlinesOK {
	o.CacheControl = cacheControl
	return o
}

// SetCacheControl sets the cacheControl to the get universe bloodlines o k response
func (o *GetUniverseBloodlinesOK) SetCacheControl(cacheControl string) {
	o.CacheControl = cacheControl
}

// WithContentLanguage adds the contentLanguage to the get universe bloodlines o k response
func (o *GetUniverseBloodlinesOK) WithContentLanguage(contentLanguage string) *GetUniverseBloodlinesOK {
	o.ContentLanguage = contentLanguage
	return o
}

// SetContentLanguage sets the contentLanguage to the get universe bloodlines o k response
func (o *GetUniverseBloodlinesOK) SetContentLanguage(contentLanguage string) {
	o.ContentLanguage = contentLanguage
}

// WithExpires adds the expires to the get universe bloodlines o k response
func (o *GetUniverseBloodlinesOK) WithExpires(expires string) *GetUniverseBloodlinesOK {
	o.Expires = expires
	return o
}

// SetExpires sets the expires to the get universe bloodlines o k response
func (o *GetUniverseBloodlinesOK) SetExpires(expires string) {
	o.Expires = expires
}

// WithLastModified adds the lastModified to the get universe bloodlines o k response
func (o *GetUniverseBloodlinesOK) WithLastModified(lastModified string) *GetUniverseBloodlinesOK {
	o.LastModified = lastModified
	return o
}

// SetLastModified sets the lastModified to the get universe bloodlines o k response
func (o *GetUniverseBloodlinesOK) SetLastModified(lastModified string) {
	o.LastModified = lastModified
}

// WithPayload adds the payload to the get universe bloodlines o k response
func (o *GetUniverseBloodlinesOK) WithPayload(payload models.GetUniverseBloodlinesOKBody) *GetUniverseBloodlinesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get universe bloodlines o k response
func (o *GetUniverseBloodlinesOK) SetPayload(payload models.GetUniverseBloodlinesOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUniverseBloodlinesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Cache-Control

	cacheControl := o.CacheControl
	if cacheControl != "" {
		rw.Header().Set("Cache-Control", cacheControl)
	}

	// response header Content-Language

	contentLanguage := o.ContentLanguage
	if contentLanguage != "" {
		rw.Header().Set("Content-Language", contentLanguage)
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
		payload = make(models.GetUniverseBloodlinesOKBody, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetUniverseBloodlinesInternalServerErrorCode is the HTTP code returned for type GetUniverseBloodlinesInternalServerError
const GetUniverseBloodlinesInternalServerErrorCode int = 500

/*GetUniverseBloodlinesInternalServerError Internal server error

swagger:response getUniverseBloodlinesInternalServerError
*/
type GetUniverseBloodlinesInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.GetUniverseBloodlinesInternalServerErrorBody `json:"body,omitempty"`
}

// NewGetUniverseBloodlinesInternalServerError creates GetUniverseBloodlinesInternalServerError with default headers values
func NewGetUniverseBloodlinesInternalServerError() *GetUniverseBloodlinesInternalServerError {
	return &GetUniverseBloodlinesInternalServerError{}
}

// WithPayload adds the payload to the get universe bloodlines internal server error response
func (o *GetUniverseBloodlinesInternalServerError) WithPayload(payload *models.GetUniverseBloodlinesInternalServerErrorBody) *GetUniverseBloodlinesInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get universe bloodlines internal server error response
func (o *GetUniverseBloodlinesInternalServerError) SetPayload(payload *models.GetUniverseBloodlinesInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUniverseBloodlinesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}