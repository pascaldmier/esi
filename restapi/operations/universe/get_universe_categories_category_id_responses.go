package universe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/pascaldmier/esi/models"
)

// GetUniverseCategoriesCategoryIDOKCode is the HTTP code returned for type GetUniverseCategoriesCategoryIDOK
const GetUniverseCategoriesCategoryIDOKCode int = 200

/*GetUniverseCategoriesCategoryIDOK Information about an item category

swagger:response getUniverseCategoriesCategoryIdOK
*/
type GetUniverseCategoriesCategoryIDOK struct {
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
	Payload *models.GetUniverseCategoriesCategoryIDOKBody `json:"body,omitempty"`
}

// NewGetUniverseCategoriesCategoryIDOK creates GetUniverseCategoriesCategoryIDOK with default headers values
func NewGetUniverseCategoriesCategoryIDOK() *GetUniverseCategoriesCategoryIDOK {
	return &GetUniverseCategoriesCategoryIDOK{}
}

// WithCacheControl adds the cacheControl to the get universe categories category Id o k response
func (o *GetUniverseCategoriesCategoryIDOK) WithCacheControl(cacheControl string) *GetUniverseCategoriesCategoryIDOK {
	o.CacheControl = cacheControl
	return o
}

// SetCacheControl sets the cacheControl to the get universe categories category Id o k response
func (o *GetUniverseCategoriesCategoryIDOK) SetCacheControl(cacheControl string) {
	o.CacheControl = cacheControl
}

// WithContentLanguage adds the contentLanguage to the get universe categories category Id o k response
func (o *GetUniverseCategoriesCategoryIDOK) WithContentLanguage(contentLanguage string) *GetUniverseCategoriesCategoryIDOK {
	o.ContentLanguage = contentLanguage
	return o
}

// SetContentLanguage sets the contentLanguage to the get universe categories category Id o k response
func (o *GetUniverseCategoriesCategoryIDOK) SetContentLanguage(contentLanguage string) {
	o.ContentLanguage = contentLanguage
}

// WithExpires adds the expires to the get universe categories category Id o k response
func (o *GetUniverseCategoriesCategoryIDOK) WithExpires(expires string) *GetUniverseCategoriesCategoryIDOK {
	o.Expires = expires
	return o
}

// SetExpires sets the expires to the get universe categories category Id o k response
func (o *GetUniverseCategoriesCategoryIDOK) SetExpires(expires string) {
	o.Expires = expires
}

// WithLastModified adds the lastModified to the get universe categories category Id o k response
func (o *GetUniverseCategoriesCategoryIDOK) WithLastModified(lastModified string) *GetUniverseCategoriesCategoryIDOK {
	o.LastModified = lastModified
	return o
}

// SetLastModified sets the lastModified to the get universe categories category Id o k response
func (o *GetUniverseCategoriesCategoryIDOK) SetLastModified(lastModified string) {
	o.LastModified = lastModified
}

// WithPayload adds the payload to the get universe categories category Id o k response
func (o *GetUniverseCategoriesCategoryIDOK) WithPayload(payload *models.GetUniverseCategoriesCategoryIDOKBody) *GetUniverseCategoriesCategoryIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get universe categories category Id o k response
func (o *GetUniverseCategoriesCategoryIDOK) SetPayload(payload *models.GetUniverseCategoriesCategoryIDOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUniverseCategoriesCategoryIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetUniverseCategoriesCategoryIDNotFoundCode is the HTTP code returned for type GetUniverseCategoriesCategoryIDNotFound
const GetUniverseCategoriesCategoryIDNotFoundCode int = 404

/*GetUniverseCategoriesCategoryIDNotFound Category not found

swagger:response getUniverseCategoriesCategoryIdNotFound
*/
type GetUniverseCategoriesCategoryIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.GetUniverseCategoriesCategoryIDNotFoundBody `json:"body,omitempty"`
}

// NewGetUniverseCategoriesCategoryIDNotFound creates GetUniverseCategoriesCategoryIDNotFound with default headers values
func NewGetUniverseCategoriesCategoryIDNotFound() *GetUniverseCategoriesCategoryIDNotFound {
	return &GetUniverseCategoriesCategoryIDNotFound{}
}

// WithPayload adds the payload to the get universe categories category Id not found response
func (o *GetUniverseCategoriesCategoryIDNotFound) WithPayload(payload *models.GetUniverseCategoriesCategoryIDNotFoundBody) *GetUniverseCategoriesCategoryIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get universe categories category Id not found response
func (o *GetUniverseCategoriesCategoryIDNotFound) SetPayload(payload *models.GetUniverseCategoriesCategoryIDNotFoundBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUniverseCategoriesCategoryIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUniverseCategoriesCategoryIDInternalServerErrorCode is the HTTP code returned for type GetUniverseCategoriesCategoryIDInternalServerError
const GetUniverseCategoriesCategoryIDInternalServerErrorCode int = 500

/*GetUniverseCategoriesCategoryIDInternalServerError Internal server error

swagger:response getUniverseCategoriesCategoryIdInternalServerError
*/
type GetUniverseCategoriesCategoryIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.GetUniverseCategoriesCategoryIDInternalServerErrorBody `json:"body,omitempty"`
}

// NewGetUniverseCategoriesCategoryIDInternalServerError creates GetUniverseCategoriesCategoryIDInternalServerError with default headers values
func NewGetUniverseCategoriesCategoryIDInternalServerError() *GetUniverseCategoriesCategoryIDInternalServerError {
	return &GetUniverseCategoriesCategoryIDInternalServerError{}
}

// WithPayload adds the payload to the get universe categories category Id internal server error response
func (o *GetUniverseCategoriesCategoryIDInternalServerError) WithPayload(payload *models.GetUniverseCategoriesCategoryIDInternalServerErrorBody) *GetUniverseCategoriesCategoryIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get universe categories category Id internal server error response
func (o *GetUniverseCategoriesCategoryIDInternalServerError) SetPayload(payload *models.GetUniverseCategoriesCategoryIDInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUniverseCategoriesCategoryIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
