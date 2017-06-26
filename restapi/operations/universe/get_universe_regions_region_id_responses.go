package universe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/pascaldmier/esi/models"
)

// GetUniverseRegionsRegionIDOKCode is the HTTP code returned for type GetUniverseRegionsRegionIDOK
const GetUniverseRegionsRegionIDOKCode int = 200

/*GetUniverseRegionsRegionIDOK Information about a region

swagger:response getUniverseRegionsRegionIdOK
*/
type GetUniverseRegionsRegionIDOK struct {
	/*The caching mechanism used
	  Required: true
	*/
	CacheControl string `json:"Cache-Control"`
	/*The language used in the response
	  Required: true
	*/
	ContentLanguage string `json:"Content-Language"`
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
	Payload *models.GetUniverseRegionsRegionIDOKBody `json:"body,omitempty"`
}

// NewGetUniverseRegionsRegionIDOK creates GetUniverseRegionsRegionIDOK with default headers values
func NewGetUniverseRegionsRegionIDOK() *GetUniverseRegionsRegionIDOK {
	return &GetUniverseRegionsRegionIDOK{}
}

// WithCacheControl adds the cacheControl to the get universe regions region Id o k response
func (o *GetUniverseRegionsRegionIDOK) WithCacheControl(cacheControl string) *GetUniverseRegionsRegionIDOK {
	o.CacheControl = cacheControl
	return o
}

// SetCacheControl sets the cacheControl to the get universe regions region Id o k response
func (o *GetUniverseRegionsRegionIDOK) SetCacheControl(cacheControl string) {
	o.CacheControl = cacheControl
}

// WithContentLanguage adds the contentLanguage to the get universe regions region Id o k response
func (o *GetUniverseRegionsRegionIDOK) WithContentLanguage(contentLanguage string) *GetUniverseRegionsRegionIDOK {
	o.ContentLanguage = contentLanguage
	return o
}

// SetContentLanguage sets the contentLanguage to the get universe regions region Id o k response
func (o *GetUniverseRegionsRegionIDOK) SetContentLanguage(contentLanguage string) {
	o.ContentLanguage = contentLanguage
}

// WithExpires adds the expires to the get universe regions region Id o k response
func (o *GetUniverseRegionsRegionIDOK) WithExpires(expires string) *GetUniverseRegionsRegionIDOK {
	o.Expires = expires
	return o
}

// SetExpires sets the expires to the get universe regions region Id o k response
func (o *GetUniverseRegionsRegionIDOK) SetExpires(expires string) {
	o.Expires = expires
}

// WithLastModified adds the lastModified to the get universe regions region Id o k response
func (o *GetUniverseRegionsRegionIDOK) WithLastModified(lastModified string) *GetUniverseRegionsRegionIDOK {
	o.LastModified = lastModified
	return o
}

// SetLastModified sets the lastModified to the get universe regions region Id o k response
func (o *GetUniverseRegionsRegionIDOK) SetLastModified(lastModified string) {
	o.LastModified = lastModified
}

// WithPayload adds the payload to the get universe regions region Id o k response
func (o *GetUniverseRegionsRegionIDOK) WithPayload(payload *models.GetUniverseRegionsRegionIDOKBody) *GetUniverseRegionsRegionIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get universe regions region Id o k response
func (o *GetUniverseRegionsRegionIDOK) SetPayload(payload *models.GetUniverseRegionsRegionIDOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUniverseRegionsRegionIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Cache-Control

	cacheControl := o.CacheControl
	if cacheControl != "" {
		rw.Header().Set("Cache-Control", cacheControl)
	}

	// response header Content-Language

	contentLanguage := o.ContentLanguage
	if contentLanguage != "" {
		rw.Header().Set("Content-Language", contentLanguage)
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

// GetUniverseRegionsRegionIDNotFoundCode is the HTTP code returned for type GetUniverseRegionsRegionIDNotFound
const GetUniverseRegionsRegionIDNotFoundCode int = 404

/*GetUniverseRegionsRegionIDNotFound Region not found

swagger:response getUniverseRegionsRegionIdNotFound
*/
type GetUniverseRegionsRegionIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.GetUniverseRegionsRegionIDNotFoundBody `json:"body,omitempty"`
}

// NewGetUniverseRegionsRegionIDNotFound creates GetUniverseRegionsRegionIDNotFound with default headers values
func NewGetUniverseRegionsRegionIDNotFound() *GetUniverseRegionsRegionIDNotFound {
	return &GetUniverseRegionsRegionIDNotFound{}
}

// WithPayload adds the payload to the get universe regions region Id not found response
func (o *GetUniverseRegionsRegionIDNotFound) WithPayload(payload *models.GetUniverseRegionsRegionIDNotFoundBody) *GetUniverseRegionsRegionIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get universe regions region Id not found response
func (o *GetUniverseRegionsRegionIDNotFound) SetPayload(payload *models.GetUniverseRegionsRegionIDNotFoundBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUniverseRegionsRegionIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUniverseRegionsRegionIDInternalServerErrorCode is the HTTP code returned for type GetUniverseRegionsRegionIDInternalServerError
const GetUniverseRegionsRegionIDInternalServerErrorCode int = 500

/*GetUniverseRegionsRegionIDInternalServerError Internal server error

swagger:response getUniverseRegionsRegionIdInternalServerError
*/
type GetUniverseRegionsRegionIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.GetUniverseRegionsRegionIDInternalServerErrorBody `json:"body,omitempty"`
}

// NewGetUniverseRegionsRegionIDInternalServerError creates GetUniverseRegionsRegionIDInternalServerError with default headers values
func NewGetUniverseRegionsRegionIDInternalServerError() *GetUniverseRegionsRegionIDInternalServerError {
	return &GetUniverseRegionsRegionIDInternalServerError{}
}

// WithPayload adds the payload to the get universe regions region Id internal server error response
func (o *GetUniverseRegionsRegionIDInternalServerError) WithPayload(payload *models.GetUniverseRegionsRegionIDInternalServerErrorBody) *GetUniverseRegionsRegionIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get universe regions region Id internal server error response
func (o *GetUniverseRegionsRegionIDInternalServerError) SetPayload(payload *models.GetUniverseRegionsRegionIDInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUniverseRegionsRegionIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}