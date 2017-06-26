package character

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/pascaldmier/esi/models"
)

// GetCharactersCharacterIDStandingsOKCode is the HTTP code returned for type GetCharactersCharacterIDStandingsOK
const GetCharactersCharacterIDStandingsOKCode int = 200

/*GetCharactersCharacterIDStandingsOK A list of standings

swagger:response getCharactersCharacterIdStandingsOK
*/
type GetCharactersCharacterIDStandingsOK struct {
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
	Payload models.GetCharactersCharacterIDStandingsOKBody `json:"body,omitempty"`
}

// NewGetCharactersCharacterIDStandingsOK creates GetCharactersCharacterIDStandingsOK with default headers values
func NewGetCharactersCharacterIDStandingsOK() *GetCharactersCharacterIDStandingsOK {
	return &GetCharactersCharacterIDStandingsOK{}
}

// WithCacheControl adds the cacheControl to the get characters character Id standings o k response
func (o *GetCharactersCharacterIDStandingsOK) WithCacheControl(cacheControl string) *GetCharactersCharacterIDStandingsOK {
	o.CacheControl = cacheControl
	return o
}

// SetCacheControl sets the cacheControl to the get characters character Id standings o k response
func (o *GetCharactersCharacterIDStandingsOK) SetCacheControl(cacheControl string) {
	o.CacheControl = cacheControl
}

// WithExpires adds the expires to the get characters character Id standings o k response
func (o *GetCharactersCharacterIDStandingsOK) WithExpires(expires string) *GetCharactersCharacterIDStandingsOK {
	o.Expires = expires
	return o
}

// SetExpires sets the expires to the get characters character Id standings o k response
func (o *GetCharactersCharacterIDStandingsOK) SetExpires(expires string) {
	o.Expires = expires
}

// WithLastModified adds the lastModified to the get characters character Id standings o k response
func (o *GetCharactersCharacterIDStandingsOK) WithLastModified(lastModified string) *GetCharactersCharacterIDStandingsOK {
	o.LastModified = lastModified
	return o
}

// SetLastModified sets the lastModified to the get characters character Id standings o k response
func (o *GetCharactersCharacterIDStandingsOK) SetLastModified(lastModified string) {
	o.LastModified = lastModified
}

// WithPayload adds the payload to the get characters character Id standings o k response
func (o *GetCharactersCharacterIDStandingsOK) WithPayload(payload models.GetCharactersCharacterIDStandingsOKBody) *GetCharactersCharacterIDStandingsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get characters character Id standings o k response
func (o *GetCharactersCharacterIDStandingsOK) SetPayload(payload models.GetCharactersCharacterIDStandingsOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCharactersCharacterIDStandingsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
		payload = make(models.GetCharactersCharacterIDStandingsOKBody, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetCharactersCharacterIDStandingsForbiddenCode is the HTTP code returned for type GetCharactersCharacterIDStandingsForbidden
const GetCharactersCharacterIDStandingsForbiddenCode int = 403

/*GetCharactersCharacterIDStandingsForbidden Forbidden

swagger:response getCharactersCharacterIdStandingsForbidden
*/
type GetCharactersCharacterIDStandingsForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GetCharactersCharacterIDStandingsForbiddenBody `json:"body,omitempty"`
}

// NewGetCharactersCharacterIDStandingsForbidden creates GetCharactersCharacterIDStandingsForbidden with default headers values
func NewGetCharactersCharacterIDStandingsForbidden() *GetCharactersCharacterIDStandingsForbidden {
	return &GetCharactersCharacterIDStandingsForbidden{}
}

// WithPayload adds the payload to the get characters character Id standings forbidden response
func (o *GetCharactersCharacterIDStandingsForbidden) WithPayload(payload *models.GetCharactersCharacterIDStandingsForbiddenBody) *GetCharactersCharacterIDStandingsForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get characters character Id standings forbidden response
func (o *GetCharactersCharacterIDStandingsForbidden) SetPayload(payload *models.GetCharactersCharacterIDStandingsForbiddenBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCharactersCharacterIDStandingsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetCharactersCharacterIDStandingsInternalServerErrorCode is the HTTP code returned for type GetCharactersCharacterIDStandingsInternalServerError
const GetCharactersCharacterIDStandingsInternalServerErrorCode int = 500

/*GetCharactersCharacterIDStandingsInternalServerError Internal server error

swagger:response getCharactersCharacterIdStandingsInternalServerError
*/
type GetCharactersCharacterIDStandingsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.GetCharactersCharacterIDStandingsInternalServerErrorBody `json:"body,omitempty"`
}

// NewGetCharactersCharacterIDStandingsInternalServerError creates GetCharactersCharacterIDStandingsInternalServerError with default headers values
func NewGetCharactersCharacterIDStandingsInternalServerError() *GetCharactersCharacterIDStandingsInternalServerError {
	return &GetCharactersCharacterIDStandingsInternalServerError{}
}

// WithPayload adds the payload to the get characters character Id standings internal server error response
func (o *GetCharactersCharacterIDStandingsInternalServerError) WithPayload(payload *models.GetCharactersCharacterIDStandingsInternalServerErrorBody) *GetCharactersCharacterIDStandingsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get characters character Id standings internal server error response
func (o *GetCharactersCharacterIDStandingsInternalServerError) SetPayload(payload *models.GetCharactersCharacterIDStandingsInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCharactersCharacterIDStandingsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}