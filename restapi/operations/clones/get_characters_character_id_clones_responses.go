package clones

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/pascaldmier/esi/models"
)

// GetCharactersCharacterIDClonesOKCode is the HTTP code returned for type GetCharactersCharacterIDClonesOK
const GetCharactersCharacterIDClonesOKCode int = 200

/*GetCharactersCharacterIDClonesOK Clone information for the given character

swagger:response getCharactersCharacterIdClonesOK
*/
type GetCharactersCharacterIDClonesOK struct {
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
	Payload *models.GetCharactersCharacterIDClonesOKBody `json:"body,omitempty"`
}

// NewGetCharactersCharacterIDClonesOK creates GetCharactersCharacterIDClonesOK with default headers values
func NewGetCharactersCharacterIDClonesOK() *GetCharactersCharacterIDClonesOK {
	return &GetCharactersCharacterIDClonesOK{}
}

// WithCacheControl adds the cacheControl to the get characters character Id clones o k response
func (o *GetCharactersCharacterIDClonesOK) WithCacheControl(cacheControl string) *GetCharactersCharacterIDClonesOK {
	o.CacheControl = cacheControl
	return o
}

// SetCacheControl sets the cacheControl to the get characters character Id clones o k response
func (o *GetCharactersCharacterIDClonesOK) SetCacheControl(cacheControl string) {
	o.CacheControl = cacheControl
}

// WithExpires adds the expires to the get characters character Id clones o k response
func (o *GetCharactersCharacterIDClonesOK) WithExpires(expires string) *GetCharactersCharacterIDClonesOK {
	o.Expires = expires
	return o
}

// SetExpires sets the expires to the get characters character Id clones o k response
func (o *GetCharactersCharacterIDClonesOK) SetExpires(expires string) {
	o.Expires = expires
}

// WithLastModified adds the lastModified to the get characters character Id clones o k response
func (o *GetCharactersCharacterIDClonesOK) WithLastModified(lastModified string) *GetCharactersCharacterIDClonesOK {
	o.LastModified = lastModified
	return o
}

// SetLastModified sets the lastModified to the get characters character Id clones o k response
func (o *GetCharactersCharacterIDClonesOK) SetLastModified(lastModified string) {
	o.LastModified = lastModified
}

// WithPayload adds the payload to the get characters character Id clones o k response
func (o *GetCharactersCharacterIDClonesOK) WithPayload(payload *models.GetCharactersCharacterIDClonesOKBody) *GetCharactersCharacterIDClonesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get characters character Id clones o k response
func (o *GetCharactersCharacterIDClonesOK) SetPayload(payload *models.GetCharactersCharacterIDClonesOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCharactersCharacterIDClonesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetCharactersCharacterIDClonesForbiddenCode is the HTTP code returned for type GetCharactersCharacterIDClonesForbidden
const GetCharactersCharacterIDClonesForbiddenCode int = 403

/*GetCharactersCharacterIDClonesForbidden Forbidden

swagger:response getCharactersCharacterIdClonesForbidden
*/
type GetCharactersCharacterIDClonesForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GetCharactersCharacterIDClonesForbiddenBody `json:"body,omitempty"`
}

// NewGetCharactersCharacterIDClonesForbidden creates GetCharactersCharacterIDClonesForbidden with default headers values
func NewGetCharactersCharacterIDClonesForbidden() *GetCharactersCharacterIDClonesForbidden {
	return &GetCharactersCharacterIDClonesForbidden{}
}

// WithPayload adds the payload to the get characters character Id clones forbidden response
func (o *GetCharactersCharacterIDClonesForbidden) WithPayload(payload *models.GetCharactersCharacterIDClonesForbiddenBody) *GetCharactersCharacterIDClonesForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get characters character Id clones forbidden response
func (o *GetCharactersCharacterIDClonesForbidden) SetPayload(payload *models.GetCharactersCharacterIDClonesForbiddenBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCharactersCharacterIDClonesForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetCharactersCharacterIDClonesInternalServerErrorCode is the HTTP code returned for type GetCharactersCharacterIDClonesInternalServerError
const GetCharactersCharacterIDClonesInternalServerErrorCode int = 500

/*GetCharactersCharacterIDClonesInternalServerError Internal server error

swagger:response getCharactersCharacterIdClonesInternalServerError
*/
type GetCharactersCharacterIDClonesInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.GetCharactersCharacterIDClonesInternalServerErrorBody `json:"body,omitempty"`
}

// NewGetCharactersCharacterIDClonesInternalServerError creates GetCharactersCharacterIDClonesInternalServerError with default headers values
func NewGetCharactersCharacterIDClonesInternalServerError() *GetCharactersCharacterIDClonesInternalServerError {
	return &GetCharactersCharacterIDClonesInternalServerError{}
}

// WithPayload adds the payload to the get characters character Id clones internal server error response
func (o *GetCharactersCharacterIDClonesInternalServerError) WithPayload(payload *models.GetCharactersCharacterIDClonesInternalServerErrorBody) *GetCharactersCharacterIDClonesInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get characters character Id clones internal server error response
func (o *GetCharactersCharacterIDClonesInternalServerError) SetPayload(payload *models.GetCharactersCharacterIDClonesInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCharactersCharacterIDClonesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
