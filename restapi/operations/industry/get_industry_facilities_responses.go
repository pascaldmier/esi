package industry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/pascaldmier/esi/models"
)

// GetIndustryFacilitiesOKCode is the HTTP code returned for type GetIndustryFacilitiesOK
const GetIndustryFacilitiesOKCode int = 200

/*GetIndustryFacilitiesOK A list of prices

swagger:response getIndustryFacilitiesOK
*/
type GetIndustryFacilitiesOK struct {
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
	Payload models.GetIndustryFacilitiesOKBody `json:"body,omitempty"`
}

// NewGetIndustryFacilitiesOK creates GetIndustryFacilitiesOK with default headers values
func NewGetIndustryFacilitiesOK() *GetIndustryFacilitiesOK {
	return &GetIndustryFacilitiesOK{}
}

// WithCacheControl adds the cacheControl to the get industry facilities o k response
func (o *GetIndustryFacilitiesOK) WithCacheControl(cacheControl string) *GetIndustryFacilitiesOK {
	o.CacheControl = cacheControl
	return o
}

// SetCacheControl sets the cacheControl to the get industry facilities o k response
func (o *GetIndustryFacilitiesOK) SetCacheControl(cacheControl string) {
	o.CacheControl = cacheControl
}

// WithExpires adds the expires to the get industry facilities o k response
func (o *GetIndustryFacilitiesOK) WithExpires(expires string) *GetIndustryFacilitiesOK {
	o.Expires = expires
	return o
}

// SetExpires sets the expires to the get industry facilities o k response
func (o *GetIndustryFacilitiesOK) SetExpires(expires string) {
	o.Expires = expires
}

// WithLastModified adds the lastModified to the get industry facilities o k response
func (o *GetIndustryFacilitiesOK) WithLastModified(lastModified string) *GetIndustryFacilitiesOK {
	o.LastModified = lastModified
	return o
}

// SetLastModified sets the lastModified to the get industry facilities o k response
func (o *GetIndustryFacilitiesOK) SetLastModified(lastModified string) {
	o.LastModified = lastModified
}

// WithPayload adds the payload to the get industry facilities o k response
func (o *GetIndustryFacilitiesOK) WithPayload(payload models.GetIndustryFacilitiesOKBody) *GetIndustryFacilitiesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get industry facilities o k response
func (o *GetIndustryFacilitiesOK) SetPayload(payload models.GetIndustryFacilitiesOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetIndustryFacilitiesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
		payload = make(models.GetIndustryFacilitiesOKBody, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetIndustryFacilitiesInternalServerErrorCode is the HTTP code returned for type GetIndustryFacilitiesInternalServerError
const GetIndustryFacilitiesInternalServerErrorCode int = 500

/*GetIndustryFacilitiesInternalServerError Internal server error

swagger:response getIndustryFacilitiesInternalServerError
*/
type GetIndustryFacilitiesInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.GetIndustryFacilitiesInternalServerErrorBody `json:"body,omitempty"`
}

// NewGetIndustryFacilitiesInternalServerError creates GetIndustryFacilitiesInternalServerError with default headers values
func NewGetIndustryFacilitiesInternalServerError() *GetIndustryFacilitiesInternalServerError {
	return &GetIndustryFacilitiesInternalServerError{}
}

// WithPayload adds the payload to the get industry facilities internal server error response
func (o *GetIndustryFacilitiesInternalServerError) WithPayload(payload *models.GetIndustryFacilitiesInternalServerErrorBody) *GetIndustryFacilitiesInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get industry facilities internal server error response
func (o *GetIndustryFacilitiesInternalServerError) SetPayload(payload *models.GetIndustryFacilitiesInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetIndustryFacilitiesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
