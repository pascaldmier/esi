package character

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/pascaldmier/esi/models"
)

// GetCharactersCharacterIDPortraitOKCode is the HTTP code returned for type GetCharactersCharacterIDPortraitOK
const GetCharactersCharacterIDPortraitOKCode int = 200

/*GetCharactersCharacterIDPortraitOK Public data for the given character

swagger:response getCharactersCharacterIdPortraitOK
*/
type GetCharactersCharacterIDPortraitOK struct {
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
	Payload *models.GetCharactersCharacterIDPortraitOKBody `json:"body,omitempty"`
}

// NewGetCharactersCharacterIDPortraitOK creates GetCharactersCharacterIDPortraitOK with default headers values
func NewGetCharactersCharacterIDPortraitOK() *GetCharactersCharacterIDPortraitOK {
	return &GetCharactersCharacterIDPortraitOK{}
}

// WithCacheControl adds the cacheControl to the get characters character Id portrait o k response
func (o *GetCharactersCharacterIDPortraitOK) WithCacheControl(cacheControl string) *GetCharactersCharacterIDPortraitOK {
	o.CacheControl = cacheControl
	return o
}

// SetCacheControl sets the cacheControl to the get characters character Id portrait o k response
func (o *GetCharactersCharacterIDPortraitOK) SetCacheControl(cacheControl string) {
	o.CacheControl = cacheControl
}

// WithExpires adds the expires to the get characters character Id portrait o k response
func (o *GetCharactersCharacterIDPortraitOK) WithExpires(expires string) *GetCharactersCharacterIDPortraitOK {
	o.Expires = expires
	return o
}

// SetExpires sets the expires to the get characters character Id portrait o k response
func (o *GetCharactersCharacterIDPortraitOK) SetExpires(expires string) {
	o.Expires = expires
}

// WithLastModified adds the lastModified to the get characters character Id portrait o k response
func (o *GetCharactersCharacterIDPortraitOK) WithLastModified(lastModified string) *GetCharactersCharacterIDPortraitOK {
	o.LastModified = lastModified
	return o
}

// SetLastModified sets the lastModified to the get characters character Id portrait o k response
func (o *GetCharactersCharacterIDPortraitOK) SetLastModified(lastModified string) {
	o.LastModified = lastModified
}

// WithPayload adds the payload to the get characters character Id portrait o k response
func (o *GetCharactersCharacterIDPortraitOK) WithPayload(payload *models.GetCharactersCharacterIDPortraitOKBody) *GetCharactersCharacterIDPortraitOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get characters character Id portrait o k response
func (o *GetCharactersCharacterIDPortraitOK) SetPayload(payload *models.GetCharactersCharacterIDPortraitOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCharactersCharacterIDPortraitOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetCharactersCharacterIDPortraitNotFoundCode is the HTTP code returned for type GetCharactersCharacterIDPortraitNotFound
const GetCharactersCharacterIDPortraitNotFoundCode int = 404

/*GetCharactersCharacterIDPortraitNotFound No image server for this datasource

swagger:response getCharactersCharacterIdPortraitNotFound
*/
type GetCharactersCharacterIDPortraitNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.GetCharactersCharacterIDPortraitNotFoundBody `json:"body,omitempty"`
}

// NewGetCharactersCharacterIDPortraitNotFound creates GetCharactersCharacterIDPortraitNotFound with default headers values
func NewGetCharactersCharacterIDPortraitNotFound() *GetCharactersCharacterIDPortraitNotFound {
	return &GetCharactersCharacterIDPortraitNotFound{}
}

// WithPayload adds the payload to the get characters character Id portrait not found response
func (o *GetCharactersCharacterIDPortraitNotFound) WithPayload(payload *models.GetCharactersCharacterIDPortraitNotFoundBody) *GetCharactersCharacterIDPortraitNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get characters character Id portrait not found response
func (o *GetCharactersCharacterIDPortraitNotFound) SetPayload(payload *models.GetCharactersCharacterIDPortraitNotFoundBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCharactersCharacterIDPortraitNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetCharactersCharacterIDPortraitInternalServerErrorCode is the HTTP code returned for type GetCharactersCharacterIDPortraitInternalServerError
const GetCharactersCharacterIDPortraitInternalServerErrorCode int = 500

/*GetCharactersCharacterIDPortraitInternalServerError Internal server error

swagger:response getCharactersCharacterIdPortraitInternalServerError
*/
type GetCharactersCharacterIDPortraitInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.GetCharactersCharacterIDPortraitInternalServerErrorBody `json:"body,omitempty"`
}

// NewGetCharactersCharacterIDPortraitInternalServerError creates GetCharactersCharacterIDPortraitInternalServerError with default headers values
func NewGetCharactersCharacterIDPortraitInternalServerError() *GetCharactersCharacterIDPortraitInternalServerError {
	return &GetCharactersCharacterIDPortraitInternalServerError{}
}

// WithPayload adds the payload to the get characters character Id portrait internal server error response
func (o *GetCharactersCharacterIDPortraitInternalServerError) WithPayload(payload *models.GetCharactersCharacterIDPortraitInternalServerErrorBody) *GetCharactersCharacterIDPortraitInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get characters character Id portrait internal server error response
func (o *GetCharactersCharacterIDPortraitInternalServerError) SetPayload(payload *models.GetCharactersCharacterIDPortraitInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCharactersCharacterIDPortraitInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
