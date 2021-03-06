package alliance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/pascaldmier/esi/models"
)

// GetAlliancesAllianceIDCorporationsOKCode is the HTTP code returned for type GetAlliancesAllianceIDCorporationsOK
const GetAlliancesAllianceIDCorporationsOKCode int = 200

/*GetAlliancesAllianceIDCorporationsOK List of corporation IDs

swagger:response getAlliancesAllianceIdCorporationsOK
*/
type GetAlliancesAllianceIDCorporationsOK struct {
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

	/*200 ok array
	  Max Items: 1000
	  In: Body
	*/
	Payload []*int32 `json:"body,omitempty"`
}

// NewGetAlliancesAllianceIDCorporationsOK creates GetAlliancesAllianceIDCorporationsOK with default headers values
func NewGetAlliancesAllianceIDCorporationsOK() *GetAlliancesAllianceIDCorporationsOK {
	return &GetAlliancesAllianceIDCorporationsOK{}
}

// WithCacheControl adds the cacheControl to the get alliances alliance Id corporations o k response
func (o *GetAlliancesAllianceIDCorporationsOK) WithCacheControl(cacheControl string) *GetAlliancesAllianceIDCorporationsOK {
	o.CacheControl = cacheControl
	return o
}

// SetCacheControl sets the cacheControl to the get alliances alliance Id corporations o k response
func (o *GetAlliancesAllianceIDCorporationsOK) SetCacheControl(cacheControl string) {
	o.CacheControl = cacheControl
}

// WithExpires adds the expires to the get alliances alliance Id corporations o k response
func (o *GetAlliancesAllianceIDCorporationsOK) WithExpires(expires string) *GetAlliancesAllianceIDCorporationsOK {
	o.Expires = expires
	return o
}

// SetExpires sets the expires to the get alliances alliance Id corporations o k response
func (o *GetAlliancesAllianceIDCorporationsOK) SetExpires(expires string) {
	o.Expires = expires
}

// WithLastModified adds the lastModified to the get alliances alliance Id corporations o k response
func (o *GetAlliancesAllianceIDCorporationsOK) WithLastModified(lastModified string) *GetAlliancesAllianceIDCorporationsOK {
	o.LastModified = lastModified
	return o
}

// SetLastModified sets the lastModified to the get alliances alliance Id corporations o k response
func (o *GetAlliancesAllianceIDCorporationsOK) SetLastModified(lastModified string) {
	o.LastModified = lastModified
}

// WithPayload adds the payload to the get alliances alliance Id corporations o k response
func (o *GetAlliancesAllianceIDCorporationsOK) WithPayload(payload []*int32) *GetAlliancesAllianceIDCorporationsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get alliances alliance Id corporations o k response
func (o *GetAlliancesAllianceIDCorporationsOK) SetPayload(payload []*int32) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAlliancesAllianceIDCorporationsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
		payload = make([]*int32, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetAlliancesAllianceIDCorporationsInternalServerErrorCode is the HTTP code returned for type GetAlliancesAllianceIDCorporationsInternalServerError
const GetAlliancesAllianceIDCorporationsInternalServerErrorCode int = 500

/*GetAlliancesAllianceIDCorporationsInternalServerError Internal server error

swagger:response getAlliancesAllianceIdCorporationsInternalServerError
*/
type GetAlliancesAllianceIDCorporationsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.GetAlliancesAllianceIDCorporationsInternalServerErrorBody `json:"body,omitempty"`
}

// NewGetAlliancesAllianceIDCorporationsInternalServerError creates GetAlliancesAllianceIDCorporationsInternalServerError with default headers values
func NewGetAlliancesAllianceIDCorporationsInternalServerError() *GetAlliancesAllianceIDCorporationsInternalServerError {
	return &GetAlliancesAllianceIDCorporationsInternalServerError{}
}

// WithPayload adds the payload to the get alliances alliance Id corporations internal server error response
func (o *GetAlliancesAllianceIDCorporationsInternalServerError) WithPayload(payload *models.GetAlliancesAllianceIDCorporationsInternalServerErrorBody) *GetAlliancesAllianceIDCorporationsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get alliances alliance Id corporations internal server error response
func (o *GetAlliancesAllianceIDCorporationsInternalServerError) SetPayload(payload *models.GetAlliancesAllianceIDCorporationsInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAlliancesAllianceIDCorporationsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
