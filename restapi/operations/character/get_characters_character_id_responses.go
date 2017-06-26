package character

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/pascaldmier/esi/models"
)

// GetCharactersCharacterIDOKCode is the HTTP code returned for type GetCharactersCharacterIDOK
const GetCharactersCharacterIDOKCode int = 200

/*GetCharactersCharacterIDOK Public data for the given character

swagger:response getCharactersCharacterIdOK
*/
type GetCharactersCharacterIDOK struct {
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
	Payload *models.GetCharactersCharacterIDOKBody `json:"body,omitempty"`
}

// NewGetCharactersCharacterIDOK creates GetCharactersCharacterIDOK with default headers values
func NewGetCharactersCharacterIDOK() *GetCharactersCharacterIDOK {
	return &GetCharactersCharacterIDOK{}
}

// WithCacheControl adds the cacheControl to the get characters character Id o k response
func (o *GetCharactersCharacterIDOK) WithCacheControl(cacheControl string) *GetCharactersCharacterIDOK {
	o.CacheControl = cacheControl
	return o
}

// SetCacheControl sets the cacheControl to the get characters character Id o k response
func (o *GetCharactersCharacterIDOK) SetCacheControl(cacheControl string) {
	o.CacheControl = cacheControl
}

// WithExpires adds the expires to the get characters character Id o k response
func (o *GetCharactersCharacterIDOK) WithExpires(expires string) *GetCharactersCharacterIDOK {
	o.Expires = expires
	return o
}

// SetExpires sets the expires to the get characters character Id o k response
func (o *GetCharactersCharacterIDOK) SetExpires(expires string) {
	o.Expires = expires
}

// WithLastModified adds the lastModified to the get characters character Id o k response
func (o *GetCharactersCharacterIDOK) WithLastModified(lastModified string) *GetCharactersCharacterIDOK {
	o.LastModified = lastModified
	return o
}

// SetLastModified sets the lastModified to the get characters character Id o k response
func (o *GetCharactersCharacterIDOK) SetLastModified(lastModified string) {
	o.LastModified = lastModified
}

// WithPayload adds the payload to the get characters character Id o k response
func (o *GetCharactersCharacterIDOK) WithPayload(payload *models.GetCharactersCharacterIDOKBody) *GetCharactersCharacterIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get characters character Id o k response
func (o *GetCharactersCharacterIDOK) SetPayload(payload *models.GetCharactersCharacterIDOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCharactersCharacterIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetCharactersCharacterIDNotFoundCode is the HTTP code returned for type GetCharactersCharacterIDNotFound
const GetCharactersCharacterIDNotFoundCode int = 404

/*GetCharactersCharacterIDNotFound Character not found

swagger:response getCharactersCharacterIdNotFound
*/
type GetCharactersCharacterIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.GetCharactersCharacterIDNotFoundBody `json:"body,omitempty"`
}

// NewGetCharactersCharacterIDNotFound creates GetCharactersCharacterIDNotFound with default headers values
func NewGetCharactersCharacterIDNotFound() *GetCharactersCharacterIDNotFound {
	return &GetCharactersCharacterIDNotFound{}
}

// WithPayload adds the payload to the get characters character Id not found response
func (o *GetCharactersCharacterIDNotFound) WithPayload(payload *models.GetCharactersCharacterIDNotFoundBody) *GetCharactersCharacterIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get characters character Id not found response
func (o *GetCharactersCharacterIDNotFound) SetPayload(payload *models.GetCharactersCharacterIDNotFoundBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCharactersCharacterIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetCharactersCharacterIDInternalServerErrorCode is the HTTP code returned for type GetCharactersCharacterIDInternalServerError
const GetCharactersCharacterIDInternalServerErrorCode int = 500

/*GetCharactersCharacterIDInternalServerError Internal server error

swagger:response getCharactersCharacterIdInternalServerError
*/
type GetCharactersCharacterIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.GetCharactersCharacterIDInternalServerErrorBody `json:"body,omitempty"`
}

// NewGetCharactersCharacterIDInternalServerError creates GetCharactersCharacterIDInternalServerError with default headers values
func NewGetCharactersCharacterIDInternalServerError() *GetCharactersCharacterIDInternalServerError {
	return &GetCharactersCharacterIDInternalServerError{}
}

// WithPayload adds the payload to the get characters character Id internal server error response
func (o *GetCharactersCharacterIDInternalServerError) WithPayload(payload *models.GetCharactersCharacterIDInternalServerErrorBody) *GetCharactersCharacterIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get characters character Id internal server error response
func (o *GetCharactersCharacterIDInternalServerError) SetPayload(payload *models.GetCharactersCharacterIDInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCharactersCharacterIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}