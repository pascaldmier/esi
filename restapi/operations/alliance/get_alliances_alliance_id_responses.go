package alliance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/pascaldmier/esi/models"
)

// GetAlliancesAllianceIDOKCode is the HTTP code returned for type GetAlliancesAllianceIDOK
const GetAlliancesAllianceIDOKCode int = 200

/*GetAlliancesAllianceIDOK Public data about an alliance

swagger:response getAlliancesAllianceIdOK
*/
type GetAlliancesAllianceIDOK struct {
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
	Payload *models.GetAlliancesAllianceIDOKBody `json:"body,omitempty"`
}

// NewGetAlliancesAllianceIDOK creates GetAlliancesAllianceIDOK with default headers values
func NewGetAlliancesAllianceIDOK() *GetAlliancesAllianceIDOK {
	return &GetAlliancesAllianceIDOK{}
}

// WithCacheControl adds the cacheControl to the get alliances alliance Id o k response
func (o *GetAlliancesAllianceIDOK) WithCacheControl(cacheControl string) *GetAlliancesAllianceIDOK {
	o.CacheControl = cacheControl
	return o
}

// SetCacheControl sets the cacheControl to the get alliances alliance Id o k response
func (o *GetAlliancesAllianceIDOK) SetCacheControl(cacheControl string) {
	o.CacheControl = cacheControl
}

// WithExpires adds the expires to the get alliances alliance Id o k response
func (o *GetAlliancesAllianceIDOK) WithExpires(expires string) *GetAlliancesAllianceIDOK {
	o.Expires = expires
	return o
}

// SetExpires sets the expires to the get alliances alliance Id o k response
func (o *GetAlliancesAllianceIDOK) SetExpires(expires string) {
	o.Expires = expires
}

// WithLastModified adds the lastModified to the get alliances alliance Id o k response
func (o *GetAlliancesAllianceIDOK) WithLastModified(lastModified string) *GetAlliancesAllianceIDOK {
	o.LastModified = lastModified
	return o
}

// SetLastModified sets the lastModified to the get alliances alliance Id o k response
func (o *GetAlliancesAllianceIDOK) SetLastModified(lastModified string) {
	o.LastModified = lastModified
}

// WithPayload adds the payload to the get alliances alliance Id o k response
func (o *GetAlliancesAllianceIDOK) WithPayload(payload *models.GetAlliancesAllianceIDOKBody) *GetAlliancesAllianceIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get alliances alliance Id o k response
func (o *GetAlliancesAllianceIDOK) SetPayload(payload *models.GetAlliancesAllianceIDOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAlliancesAllianceIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetAlliancesAllianceIDNotFoundCode is the HTTP code returned for type GetAlliancesAllianceIDNotFound
const GetAlliancesAllianceIDNotFoundCode int = 404

/*GetAlliancesAllianceIDNotFound Alliance not found

swagger:response getAlliancesAllianceIdNotFound
*/
type GetAlliancesAllianceIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.GetAlliancesAllianceIDNotFoundBody `json:"body,omitempty"`
}

// NewGetAlliancesAllianceIDNotFound creates GetAlliancesAllianceIDNotFound with default headers values
func NewGetAlliancesAllianceIDNotFound() *GetAlliancesAllianceIDNotFound {
	return &GetAlliancesAllianceIDNotFound{}
}

// WithPayload adds the payload to the get alliances alliance Id not found response
func (o *GetAlliancesAllianceIDNotFound) WithPayload(payload *models.GetAlliancesAllianceIDNotFoundBody) *GetAlliancesAllianceIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get alliances alliance Id not found response
func (o *GetAlliancesAllianceIDNotFound) SetPayload(payload *models.GetAlliancesAllianceIDNotFoundBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAlliancesAllianceIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAlliancesAllianceIDInternalServerErrorCode is the HTTP code returned for type GetAlliancesAllianceIDInternalServerError
const GetAlliancesAllianceIDInternalServerErrorCode int = 500

/*GetAlliancesAllianceIDInternalServerError Internal server error

swagger:response getAlliancesAllianceIdInternalServerError
*/
type GetAlliancesAllianceIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.GetAlliancesAllianceIDInternalServerErrorBody `json:"body,omitempty"`
}

// NewGetAlliancesAllianceIDInternalServerError creates GetAlliancesAllianceIDInternalServerError with default headers values
func NewGetAlliancesAllianceIDInternalServerError() *GetAlliancesAllianceIDInternalServerError {
	return &GetAlliancesAllianceIDInternalServerError{}
}

// WithPayload adds the payload to the get alliances alliance Id internal server error response
func (o *GetAlliancesAllianceIDInternalServerError) WithPayload(payload *models.GetAlliancesAllianceIDInternalServerErrorBody) *GetAlliancesAllianceIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get alliances alliance Id internal server error response
func (o *GetAlliancesAllianceIDInternalServerError) SetPayload(payload *models.GetAlliancesAllianceIDInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAlliancesAllianceIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
