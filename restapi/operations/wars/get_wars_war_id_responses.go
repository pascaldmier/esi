package wars

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/pascaldmier/esi/models"
)

// GetWarsWarIDOKCode is the HTTP code returned for type GetWarsWarIDOK
const GetWarsWarIDOKCode int = 200

/*GetWarsWarIDOK Details about a war

swagger:response getWarsWarIdOK
*/
type GetWarsWarIDOK struct {
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
	Payload *models.GetWarsWarIDOKBody `json:"body,omitempty"`
}

// NewGetWarsWarIDOK creates GetWarsWarIDOK with default headers values
func NewGetWarsWarIDOK() *GetWarsWarIDOK {
	return &GetWarsWarIDOK{}
}

// WithCacheControl adds the cacheControl to the get wars war Id o k response
func (o *GetWarsWarIDOK) WithCacheControl(cacheControl string) *GetWarsWarIDOK {
	o.CacheControl = cacheControl
	return o
}

// SetCacheControl sets the cacheControl to the get wars war Id o k response
func (o *GetWarsWarIDOK) SetCacheControl(cacheControl string) {
	o.CacheControl = cacheControl
}

// WithExpires adds the expires to the get wars war Id o k response
func (o *GetWarsWarIDOK) WithExpires(expires string) *GetWarsWarIDOK {
	o.Expires = expires
	return o
}

// SetExpires sets the expires to the get wars war Id o k response
func (o *GetWarsWarIDOK) SetExpires(expires string) {
	o.Expires = expires
}

// WithLastModified adds the lastModified to the get wars war Id o k response
func (o *GetWarsWarIDOK) WithLastModified(lastModified string) *GetWarsWarIDOK {
	o.LastModified = lastModified
	return o
}

// SetLastModified sets the lastModified to the get wars war Id o k response
func (o *GetWarsWarIDOK) SetLastModified(lastModified string) {
	o.LastModified = lastModified
}

// WithPayload adds the payload to the get wars war Id o k response
func (o *GetWarsWarIDOK) WithPayload(payload *models.GetWarsWarIDOKBody) *GetWarsWarIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get wars war Id o k response
func (o *GetWarsWarIDOK) SetPayload(payload *models.GetWarsWarIDOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetWarsWarIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetWarsWarIDUnprocessableEntityCode is the HTTP code returned for type GetWarsWarIDUnprocessableEntity
const GetWarsWarIDUnprocessableEntityCode int = 422

/*GetWarsWarIDUnprocessableEntity War not found

swagger:response getWarsWarIdUnprocessableEntity
*/
type GetWarsWarIDUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.GetWarsWarIDUnprocessableEntityBody `json:"body,omitempty"`
}

// NewGetWarsWarIDUnprocessableEntity creates GetWarsWarIDUnprocessableEntity with default headers values
func NewGetWarsWarIDUnprocessableEntity() *GetWarsWarIDUnprocessableEntity {
	return &GetWarsWarIDUnprocessableEntity{}
}

// WithPayload adds the payload to the get wars war Id unprocessable entity response
func (o *GetWarsWarIDUnprocessableEntity) WithPayload(payload *models.GetWarsWarIDUnprocessableEntityBody) *GetWarsWarIDUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get wars war Id unprocessable entity response
func (o *GetWarsWarIDUnprocessableEntity) SetPayload(payload *models.GetWarsWarIDUnprocessableEntityBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetWarsWarIDUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetWarsWarIDInternalServerErrorCode is the HTTP code returned for type GetWarsWarIDInternalServerError
const GetWarsWarIDInternalServerErrorCode int = 500

/*GetWarsWarIDInternalServerError Internal server error

swagger:response getWarsWarIdInternalServerError
*/
type GetWarsWarIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.GetWarsWarIDInternalServerErrorBody `json:"body,omitempty"`
}

// NewGetWarsWarIDInternalServerError creates GetWarsWarIDInternalServerError with default headers values
func NewGetWarsWarIDInternalServerError() *GetWarsWarIDInternalServerError {
	return &GetWarsWarIDInternalServerError{}
}

// WithPayload adds the payload to the get wars war Id internal server error response
func (o *GetWarsWarIDInternalServerError) WithPayload(payload *models.GetWarsWarIDInternalServerErrorBody) *GetWarsWarIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get wars war Id internal server error response
func (o *GetWarsWarIDInternalServerError) SetPayload(payload *models.GetWarsWarIDInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetWarsWarIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
