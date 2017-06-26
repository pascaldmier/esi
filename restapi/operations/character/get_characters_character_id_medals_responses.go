package character

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/pascaldmier/esi/models"
)

// GetCharactersCharacterIDMedalsOKCode is the HTTP code returned for type GetCharactersCharacterIDMedalsOK
const GetCharactersCharacterIDMedalsOKCode int = 200

/*GetCharactersCharacterIDMedalsOK A list of medals

swagger:response getCharactersCharacterIdMedalsOK
*/
type GetCharactersCharacterIDMedalsOK struct {
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
	Payload models.GetCharactersCharacterIDMedalsOKBody `json:"body,omitempty"`
}

// NewGetCharactersCharacterIDMedalsOK creates GetCharactersCharacterIDMedalsOK with default headers values
func NewGetCharactersCharacterIDMedalsOK() *GetCharactersCharacterIDMedalsOK {
	return &GetCharactersCharacterIDMedalsOK{}
}

// WithCacheControl adds the cacheControl to the get characters character Id medals o k response
func (o *GetCharactersCharacterIDMedalsOK) WithCacheControl(cacheControl string) *GetCharactersCharacterIDMedalsOK {
	o.CacheControl = cacheControl
	return o
}

// SetCacheControl sets the cacheControl to the get characters character Id medals o k response
func (o *GetCharactersCharacterIDMedalsOK) SetCacheControl(cacheControl string) {
	o.CacheControl = cacheControl
}

// WithExpires adds the expires to the get characters character Id medals o k response
func (o *GetCharactersCharacterIDMedalsOK) WithExpires(expires string) *GetCharactersCharacterIDMedalsOK {
	o.Expires = expires
	return o
}

// SetExpires sets the expires to the get characters character Id medals o k response
func (o *GetCharactersCharacterIDMedalsOK) SetExpires(expires string) {
	o.Expires = expires
}

// WithLastModified adds the lastModified to the get characters character Id medals o k response
func (o *GetCharactersCharacterIDMedalsOK) WithLastModified(lastModified string) *GetCharactersCharacterIDMedalsOK {
	o.LastModified = lastModified
	return o
}

// SetLastModified sets the lastModified to the get characters character Id medals o k response
func (o *GetCharactersCharacterIDMedalsOK) SetLastModified(lastModified string) {
	o.LastModified = lastModified
}

// WithPayload adds the payload to the get characters character Id medals o k response
func (o *GetCharactersCharacterIDMedalsOK) WithPayload(payload models.GetCharactersCharacterIDMedalsOKBody) *GetCharactersCharacterIDMedalsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get characters character Id medals o k response
func (o *GetCharactersCharacterIDMedalsOK) SetPayload(payload models.GetCharactersCharacterIDMedalsOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCharactersCharacterIDMedalsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
		payload = make(models.GetCharactersCharacterIDMedalsOKBody, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetCharactersCharacterIDMedalsForbiddenCode is the HTTP code returned for type GetCharactersCharacterIDMedalsForbidden
const GetCharactersCharacterIDMedalsForbiddenCode int = 403

/*GetCharactersCharacterIDMedalsForbidden Forbidden

swagger:response getCharactersCharacterIdMedalsForbidden
*/
type GetCharactersCharacterIDMedalsForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GetCharactersCharacterIDMedalsForbiddenBody `json:"body,omitempty"`
}

// NewGetCharactersCharacterIDMedalsForbidden creates GetCharactersCharacterIDMedalsForbidden with default headers values
func NewGetCharactersCharacterIDMedalsForbidden() *GetCharactersCharacterIDMedalsForbidden {
	return &GetCharactersCharacterIDMedalsForbidden{}
}

// WithPayload adds the payload to the get characters character Id medals forbidden response
func (o *GetCharactersCharacterIDMedalsForbidden) WithPayload(payload *models.GetCharactersCharacterIDMedalsForbiddenBody) *GetCharactersCharacterIDMedalsForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get characters character Id medals forbidden response
func (o *GetCharactersCharacterIDMedalsForbidden) SetPayload(payload *models.GetCharactersCharacterIDMedalsForbiddenBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCharactersCharacterIDMedalsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetCharactersCharacterIDMedalsInternalServerErrorCode is the HTTP code returned for type GetCharactersCharacterIDMedalsInternalServerError
const GetCharactersCharacterIDMedalsInternalServerErrorCode int = 500

/*GetCharactersCharacterIDMedalsInternalServerError Internal server error

swagger:response getCharactersCharacterIdMedalsInternalServerError
*/
type GetCharactersCharacterIDMedalsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.GetCharactersCharacterIDMedalsInternalServerErrorBody `json:"body,omitempty"`
}

// NewGetCharactersCharacterIDMedalsInternalServerError creates GetCharactersCharacterIDMedalsInternalServerError with default headers values
func NewGetCharactersCharacterIDMedalsInternalServerError() *GetCharactersCharacterIDMedalsInternalServerError {
	return &GetCharactersCharacterIDMedalsInternalServerError{}
}

// WithPayload adds the payload to the get characters character Id medals internal server error response
func (o *GetCharactersCharacterIDMedalsInternalServerError) WithPayload(payload *models.GetCharactersCharacterIDMedalsInternalServerErrorBody) *GetCharactersCharacterIDMedalsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get characters character Id medals internal server error response
func (o *GetCharactersCharacterIDMedalsInternalServerError) SetPayload(payload *models.GetCharactersCharacterIDMedalsInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCharactersCharacterIDMedalsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}