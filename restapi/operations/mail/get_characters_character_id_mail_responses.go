package mail

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/pascaldmier/esi/models"
)

// GetCharactersCharacterIDMailOKCode is the HTTP code returned for type GetCharactersCharacterIDMailOK
const GetCharactersCharacterIDMailOKCode int = 200

/*GetCharactersCharacterIDMailOK The requested mail

swagger:response getCharactersCharacterIdMailOK
*/
type GetCharactersCharacterIDMailOK struct {
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
	Payload models.GetCharactersCharacterIDMailOKBody `json:"body,omitempty"`
}

// NewGetCharactersCharacterIDMailOK creates GetCharactersCharacterIDMailOK with default headers values
func NewGetCharactersCharacterIDMailOK() *GetCharactersCharacterIDMailOK {
	return &GetCharactersCharacterIDMailOK{}
}

// WithCacheControl adds the cacheControl to the get characters character Id mail o k response
func (o *GetCharactersCharacterIDMailOK) WithCacheControl(cacheControl string) *GetCharactersCharacterIDMailOK {
	o.CacheControl = cacheControl
	return o
}

// SetCacheControl sets the cacheControl to the get characters character Id mail o k response
func (o *GetCharactersCharacterIDMailOK) SetCacheControl(cacheControl string) {
	o.CacheControl = cacheControl
}

// WithExpires adds the expires to the get characters character Id mail o k response
func (o *GetCharactersCharacterIDMailOK) WithExpires(expires string) *GetCharactersCharacterIDMailOK {
	o.Expires = expires
	return o
}

// SetExpires sets the expires to the get characters character Id mail o k response
func (o *GetCharactersCharacterIDMailOK) SetExpires(expires string) {
	o.Expires = expires
}

// WithLastModified adds the lastModified to the get characters character Id mail o k response
func (o *GetCharactersCharacterIDMailOK) WithLastModified(lastModified string) *GetCharactersCharacterIDMailOK {
	o.LastModified = lastModified
	return o
}

// SetLastModified sets the lastModified to the get characters character Id mail o k response
func (o *GetCharactersCharacterIDMailOK) SetLastModified(lastModified string) {
	o.LastModified = lastModified
}

// WithPayload adds the payload to the get characters character Id mail o k response
func (o *GetCharactersCharacterIDMailOK) WithPayload(payload models.GetCharactersCharacterIDMailOKBody) *GetCharactersCharacterIDMailOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get characters character Id mail o k response
func (o *GetCharactersCharacterIDMailOK) SetPayload(payload models.GetCharactersCharacterIDMailOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCharactersCharacterIDMailOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
		payload = make(models.GetCharactersCharacterIDMailOKBody, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetCharactersCharacterIDMailForbiddenCode is the HTTP code returned for type GetCharactersCharacterIDMailForbidden
const GetCharactersCharacterIDMailForbiddenCode int = 403

/*GetCharactersCharacterIDMailForbidden Forbidden

swagger:response getCharactersCharacterIdMailForbidden
*/
type GetCharactersCharacterIDMailForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GetCharactersCharacterIDMailForbiddenBody `json:"body,omitempty"`
}

// NewGetCharactersCharacterIDMailForbidden creates GetCharactersCharacterIDMailForbidden with default headers values
func NewGetCharactersCharacterIDMailForbidden() *GetCharactersCharacterIDMailForbidden {
	return &GetCharactersCharacterIDMailForbidden{}
}

// WithPayload adds the payload to the get characters character Id mail forbidden response
func (o *GetCharactersCharacterIDMailForbidden) WithPayload(payload *models.GetCharactersCharacterIDMailForbiddenBody) *GetCharactersCharacterIDMailForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get characters character Id mail forbidden response
func (o *GetCharactersCharacterIDMailForbidden) SetPayload(payload *models.GetCharactersCharacterIDMailForbiddenBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCharactersCharacterIDMailForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetCharactersCharacterIDMailInternalServerErrorCode is the HTTP code returned for type GetCharactersCharacterIDMailInternalServerError
const GetCharactersCharacterIDMailInternalServerErrorCode int = 500

/*GetCharactersCharacterIDMailInternalServerError Internal server error

swagger:response getCharactersCharacterIdMailInternalServerError
*/
type GetCharactersCharacterIDMailInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.GetCharactersCharacterIDMailInternalServerErrorBody `json:"body,omitempty"`
}

// NewGetCharactersCharacterIDMailInternalServerError creates GetCharactersCharacterIDMailInternalServerError with default headers values
func NewGetCharactersCharacterIDMailInternalServerError() *GetCharactersCharacterIDMailInternalServerError {
	return &GetCharactersCharacterIDMailInternalServerError{}
}

// WithPayload adds the payload to the get characters character Id mail internal server error response
func (o *GetCharactersCharacterIDMailInternalServerError) WithPayload(payload *models.GetCharactersCharacterIDMailInternalServerErrorBody) *GetCharactersCharacterIDMailInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get characters character Id mail internal server error response
func (o *GetCharactersCharacterIDMailInternalServerError) SetPayload(payload *models.GetCharactersCharacterIDMailInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCharactersCharacterIDMailInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
