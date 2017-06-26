package location

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/pascaldmier/esi/models"
)

// GetCharactersCharacterIDLocationOKCode is the HTTP code returned for type GetCharactersCharacterIDLocationOK
const GetCharactersCharacterIDLocationOKCode int = 200

/*GetCharactersCharacterIDLocationOK Information about the characters current location. Returns the current solar system id, and also the current station or structure ID if applicable.

swagger:response getCharactersCharacterIdLocationOK
*/
type GetCharactersCharacterIDLocationOK struct {
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
	Payload *models.GetCharactersCharacterIDLocationOKBody `json:"body,omitempty"`
}

// NewGetCharactersCharacterIDLocationOK creates GetCharactersCharacterIDLocationOK with default headers values
func NewGetCharactersCharacterIDLocationOK() *GetCharactersCharacterIDLocationOK {
	return &GetCharactersCharacterIDLocationOK{}
}

// WithCacheControl adds the cacheControl to the get characters character Id location o k response
func (o *GetCharactersCharacterIDLocationOK) WithCacheControl(cacheControl string) *GetCharactersCharacterIDLocationOK {
	o.CacheControl = cacheControl
	return o
}

// SetCacheControl sets the cacheControl to the get characters character Id location o k response
func (o *GetCharactersCharacterIDLocationOK) SetCacheControl(cacheControl string) {
	o.CacheControl = cacheControl
}

// WithExpires adds the expires to the get characters character Id location o k response
func (o *GetCharactersCharacterIDLocationOK) WithExpires(expires string) *GetCharactersCharacterIDLocationOK {
	o.Expires = expires
	return o
}

// SetExpires sets the expires to the get characters character Id location o k response
func (o *GetCharactersCharacterIDLocationOK) SetExpires(expires string) {
	o.Expires = expires
}

// WithLastModified adds the lastModified to the get characters character Id location o k response
func (o *GetCharactersCharacterIDLocationOK) WithLastModified(lastModified string) *GetCharactersCharacterIDLocationOK {
	o.LastModified = lastModified
	return o
}

// SetLastModified sets the lastModified to the get characters character Id location o k response
func (o *GetCharactersCharacterIDLocationOK) SetLastModified(lastModified string) {
	o.LastModified = lastModified
}

// WithPayload adds the payload to the get characters character Id location o k response
func (o *GetCharactersCharacterIDLocationOK) WithPayload(payload *models.GetCharactersCharacterIDLocationOKBody) *GetCharactersCharacterIDLocationOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get characters character Id location o k response
func (o *GetCharactersCharacterIDLocationOK) SetPayload(payload *models.GetCharactersCharacterIDLocationOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCharactersCharacterIDLocationOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetCharactersCharacterIDLocationForbiddenCode is the HTTP code returned for type GetCharactersCharacterIDLocationForbidden
const GetCharactersCharacterIDLocationForbiddenCode int = 403

/*GetCharactersCharacterIDLocationForbidden Forbidden

swagger:response getCharactersCharacterIdLocationForbidden
*/
type GetCharactersCharacterIDLocationForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GetCharactersCharacterIDLocationForbiddenBody `json:"body,omitempty"`
}

// NewGetCharactersCharacterIDLocationForbidden creates GetCharactersCharacterIDLocationForbidden with default headers values
func NewGetCharactersCharacterIDLocationForbidden() *GetCharactersCharacterIDLocationForbidden {
	return &GetCharactersCharacterIDLocationForbidden{}
}

// WithPayload adds the payload to the get characters character Id location forbidden response
func (o *GetCharactersCharacterIDLocationForbidden) WithPayload(payload *models.GetCharactersCharacterIDLocationForbiddenBody) *GetCharactersCharacterIDLocationForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get characters character Id location forbidden response
func (o *GetCharactersCharacterIDLocationForbidden) SetPayload(payload *models.GetCharactersCharacterIDLocationForbiddenBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCharactersCharacterIDLocationForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetCharactersCharacterIDLocationInternalServerErrorCode is the HTTP code returned for type GetCharactersCharacterIDLocationInternalServerError
const GetCharactersCharacterIDLocationInternalServerErrorCode int = 500

/*GetCharactersCharacterIDLocationInternalServerError Internal server error

swagger:response getCharactersCharacterIdLocationInternalServerError
*/
type GetCharactersCharacterIDLocationInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.GetCharactersCharacterIDLocationInternalServerErrorBody `json:"body,omitempty"`
}

// NewGetCharactersCharacterIDLocationInternalServerError creates GetCharactersCharacterIDLocationInternalServerError with default headers values
func NewGetCharactersCharacterIDLocationInternalServerError() *GetCharactersCharacterIDLocationInternalServerError {
	return &GetCharactersCharacterIDLocationInternalServerError{}
}

// WithPayload adds the payload to the get characters character Id location internal server error response
func (o *GetCharactersCharacterIDLocationInternalServerError) WithPayload(payload *models.GetCharactersCharacterIDLocationInternalServerErrorBody) *GetCharactersCharacterIDLocationInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get characters character Id location internal server error response
func (o *GetCharactersCharacterIDLocationInternalServerError) SetPayload(payload *models.GetCharactersCharacterIDLocationInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCharactersCharacterIDLocationInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}