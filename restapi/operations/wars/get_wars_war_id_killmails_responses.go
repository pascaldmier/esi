package wars

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/pascaldmier/esi/models"
)

// GetWarsWarIDKillmailsOKCode is the HTTP code returned for type GetWarsWarIDKillmailsOK
const GetWarsWarIDKillmailsOKCode int = 200

/*GetWarsWarIDKillmailsOK A list of killmail IDs and hashes

swagger:response getWarsWarIdKillmailsOK
*/
type GetWarsWarIDKillmailsOK struct {
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
	Payload models.GetWarsWarIDKillmailsOKBody `json:"body,omitempty"`
}

// NewGetWarsWarIDKillmailsOK creates GetWarsWarIDKillmailsOK with default headers values
func NewGetWarsWarIDKillmailsOK() *GetWarsWarIDKillmailsOK {
	return &GetWarsWarIDKillmailsOK{}
}

// WithCacheControl adds the cacheControl to the get wars war Id killmails o k response
func (o *GetWarsWarIDKillmailsOK) WithCacheControl(cacheControl string) *GetWarsWarIDKillmailsOK {
	o.CacheControl = cacheControl
	return o
}

// SetCacheControl sets the cacheControl to the get wars war Id killmails o k response
func (o *GetWarsWarIDKillmailsOK) SetCacheControl(cacheControl string) {
	o.CacheControl = cacheControl
}

// WithExpires adds the expires to the get wars war Id killmails o k response
func (o *GetWarsWarIDKillmailsOK) WithExpires(expires string) *GetWarsWarIDKillmailsOK {
	o.Expires = expires
	return o
}

// SetExpires sets the expires to the get wars war Id killmails o k response
func (o *GetWarsWarIDKillmailsOK) SetExpires(expires string) {
	o.Expires = expires
}

// WithLastModified adds the lastModified to the get wars war Id killmails o k response
func (o *GetWarsWarIDKillmailsOK) WithLastModified(lastModified string) *GetWarsWarIDKillmailsOK {
	o.LastModified = lastModified
	return o
}

// SetLastModified sets the lastModified to the get wars war Id killmails o k response
func (o *GetWarsWarIDKillmailsOK) SetLastModified(lastModified string) {
	o.LastModified = lastModified
}

// WithPayload adds the payload to the get wars war Id killmails o k response
func (o *GetWarsWarIDKillmailsOK) WithPayload(payload models.GetWarsWarIDKillmailsOKBody) *GetWarsWarIDKillmailsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get wars war Id killmails o k response
func (o *GetWarsWarIDKillmailsOK) SetPayload(payload models.GetWarsWarIDKillmailsOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetWarsWarIDKillmailsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
		payload = make(models.GetWarsWarIDKillmailsOKBody, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetWarsWarIDKillmailsUnprocessableEntityCode is the HTTP code returned for type GetWarsWarIDKillmailsUnprocessableEntity
const GetWarsWarIDKillmailsUnprocessableEntityCode int = 422

/*GetWarsWarIDKillmailsUnprocessableEntity War not found

swagger:response getWarsWarIdKillmailsUnprocessableEntity
*/
type GetWarsWarIDKillmailsUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.GetWarsWarIDKillmailsUnprocessableEntityBody `json:"body,omitempty"`
}

// NewGetWarsWarIDKillmailsUnprocessableEntity creates GetWarsWarIDKillmailsUnprocessableEntity with default headers values
func NewGetWarsWarIDKillmailsUnprocessableEntity() *GetWarsWarIDKillmailsUnprocessableEntity {
	return &GetWarsWarIDKillmailsUnprocessableEntity{}
}

// WithPayload adds the payload to the get wars war Id killmails unprocessable entity response
func (o *GetWarsWarIDKillmailsUnprocessableEntity) WithPayload(payload *models.GetWarsWarIDKillmailsUnprocessableEntityBody) *GetWarsWarIDKillmailsUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get wars war Id killmails unprocessable entity response
func (o *GetWarsWarIDKillmailsUnprocessableEntity) SetPayload(payload *models.GetWarsWarIDKillmailsUnprocessableEntityBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetWarsWarIDKillmailsUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetWarsWarIDKillmailsInternalServerErrorCode is the HTTP code returned for type GetWarsWarIDKillmailsInternalServerError
const GetWarsWarIDKillmailsInternalServerErrorCode int = 500

/*GetWarsWarIDKillmailsInternalServerError Internal server error

swagger:response getWarsWarIdKillmailsInternalServerError
*/
type GetWarsWarIDKillmailsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.GetWarsWarIDKillmailsInternalServerErrorBody `json:"body,omitempty"`
}

// NewGetWarsWarIDKillmailsInternalServerError creates GetWarsWarIDKillmailsInternalServerError with default headers values
func NewGetWarsWarIDKillmailsInternalServerError() *GetWarsWarIDKillmailsInternalServerError {
	return &GetWarsWarIDKillmailsInternalServerError{}
}

// WithPayload adds the payload to the get wars war Id killmails internal server error response
func (o *GetWarsWarIDKillmailsInternalServerError) WithPayload(payload *models.GetWarsWarIDKillmailsInternalServerErrorBody) *GetWarsWarIDKillmailsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get wars war Id killmails internal server error response
func (o *GetWarsWarIDKillmailsInternalServerError) SetPayload(payload *models.GetWarsWarIDKillmailsInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetWarsWarIDKillmailsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
