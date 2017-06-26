package calendar

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/pascaldmier/esi/models"
)

// GetCharactersCharacterIDCalendarOKCode is the HTTP code returned for type GetCharactersCharacterIDCalendarOK
const GetCharactersCharacterIDCalendarOKCode int = 200

/*GetCharactersCharacterIDCalendarOK A collection of event summaries

swagger:response getCharactersCharacterIdCalendarOK
*/
type GetCharactersCharacterIDCalendarOK struct {
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
	Payload models.GetCharactersCharacterIDCalendarOKBody `json:"body,omitempty"`
}

// NewGetCharactersCharacterIDCalendarOK creates GetCharactersCharacterIDCalendarOK with default headers values
func NewGetCharactersCharacterIDCalendarOK() *GetCharactersCharacterIDCalendarOK {
	return &GetCharactersCharacterIDCalendarOK{}
}

// WithCacheControl adds the cacheControl to the get characters character Id calendar o k response
func (o *GetCharactersCharacterIDCalendarOK) WithCacheControl(cacheControl string) *GetCharactersCharacterIDCalendarOK {
	o.CacheControl = cacheControl
	return o
}

// SetCacheControl sets the cacheControl to the get characters character Id calendar o k response
func (o *GetCharactersCharacterIDCalendarOK) SetCacheControl(cacheControl string) {
	o.CacheControl = cacheControl
}

// WithExpires adds the expires to the get characters character Id calendar o k response
func (o *GetCharactersCharacterIDCalendarOK) WithExpires(expires string) *GetCharactersCharacterIDCalendarOK {
	o.Expires = expires
	return o
}

// SetExpires sets the expires to the get characters character Id calendar o k response
func (o *GetCharactersCharacterIDCalendarOK) SetExpires(expires string) {
	o.Expires = expires
}

// WithLastModified adds the lastModified to the get characters character Id calendar o k response
func (o *GetCharactersCharacterIDCalendarOK) WithLastModified(lastModified string) *GetCharactersCharacterIDCalendarOK {
	o.LastModified = lastModified
	return o
}

// SetLastModified sets the lastModified to the get characters character Id calendar o k response
func (o *GetCharactersCharacterIDCalendarOK) SetLastModified(lastModified string) {
	o.LastModified = lastModified
}

// WithPayload adds the payload to the get characters character Id calendar o k response
func (o *GetCharactersCharacterIDCalendarOK) WithPayload(payload models.GetCharactersCharacterIDCalendarOKBody) *GetCharactersCharacterIDCalendarOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get characters character Id calendar o k response
func (o *GetCharactersCharacterIDCalendarOK) SetPayload(payload models.GetCharactersCharacterIDCalendarOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCharactersCharacterIDCalendarOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
		payload = make(models.GetCharactersCharacterIDCalendarOKBody, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetCharactersCharacterIDCalendarForbiddenCode is the HTTP code returned for type GetCharactersCharacterIDCalendarForbidden
const GetCharactersCharacterIDCalendarForbiddenCode int = 403

/*GetCharactersCharacterIDCalendarForbidden Forbidden

swagger:response getCharactersCharacterIdCalendarForbidden
*/
type GetCharactersCharacterIDCalendarForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GetCharactersCharacterIDCalendarForbiddenBody `json:"body,omitempty"`
}

// NewGetCharactersCharacterIDCalendarForbidden creates GetCharactersCharacterIDCalendarForbidden with default headers values
func NewGetCharactersCharacterIDCalendarForbidden() *GetCharactersCharacterIDCalendarForbidden {
	return &GetCharactersCharacterIDCalendarForbidden{}
}

// WithPayload adds the payload to the get characters character Id calendar forbidden response
func (o *GetCharactersCharacterIDCalendarForbidden) WithPayload(payload *models.GetCharactersCharacterIDCalendarForbiddenBody) *GetCharactersCharacterIDCalendarForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get characters character Id calendar forbidden response
func (o *GetCharactersCharacterIDCalendarForbidden) SetPayload(payload *models.GetCharactersCharacterIDCalendarForbiddenBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCharactersCharacterIDCalendarForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetCharactersCharacterIDCalendarInternalServerErrorCode is the HTTP code returned for type GetCharactersCharacterIDCalendarInternalServerError
const GetCharactersCharacterIDCalendarInternalServerErrorCode int = 500

/*GetCharactersCharacterIDCalendarInternalServerError Internal server error

swagger:response getCharactersCharacterIdCalendarInternalServerError
*/
type GetCharactersCharacterIDCalendarInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.GetCharactersCharacterIDCalendarInternalServerErrorBody `json:"body,omitempty"`
}

// NewGetCharactersCharacterIDCalendarInternalServerError creates GetCharactersCharacterIDCalendarInternalServerError with default headers values
func NewGetCharactersCharacterIDCalendarInternalServerError() *GetCharactersCharacterIDCalendarInternalServerError {
	return &GetCharactersCharacterIDCalendarInternalServerError{}
}

// WithPayload adds the payload to the get characters character Id calendar internal server error response
func (o *GetCharactersCharacterIDCalendarInternalServerError) WithPayload(payload *models.GetCharactersCharacterIDCalendarInternalServerErrorBody) *GetCharactersCharacterIDCalendarInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get characters character Id calendar internal server error response
func (o *GetCharactersCharacterIDCalendarInternalServerError) SetPayload(payload *models.GetCharactersCharacterIDCalendarInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCharactersCharacterIDCalendarInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
