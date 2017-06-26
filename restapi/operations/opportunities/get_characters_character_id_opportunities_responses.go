package opportunities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/pascaldmier/esi/models"
)

// GetCharactersCharacterIDOpportunitiesOKCode is the HTTP code returned for type GetCharactersCharacterIDOpportunitiesOK
const GetCharactersCharacterIDOpportunitiesOKCode int = 200

/*GetCharactersCharacterIDOpportunitiesOK A list of opportunities task ids

swagger:response getCharactersCharacterIdOpportunitiesOK
*/
type GetCharactersCharacterIDOpportunitiesOK struct {
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
	Payload models.GetCharactersCharacterIDOpportunitiesOKBody `json:"body,omitempty"`
}

// NewGetCharactersCharacterIDOpportunitiesOK creates GetCharactersCharacterIDOpportunitiesOK with default headers values
func NewGetCharactersCharacterIDOpportunitiesOK() *GetCharactersCharacterIDOpportunitiesOK {
	return &GetCharactersCharacterIDOpportunitiesOK{}
}

// WithCacheControl adds the cacheControl to the get characters character Id opportunities o k response
func (o *GetCharactersCharacterIDOpportunitiesOK) WithCacheControl(cacheControl string) *GetCharactersCharacterIDOpportunitiesOK {
	o.CacheControl = cacheControl
	return o
}

// SetCacheControl sets the cacheControl to the get characters character Id opportunities o k response
func (o *GetCharactersCharacterIDOpportunitiesOK) SetCacheControl(cacheControl string) {
	o.CacheControl = cacheControl
}

// WithExpires adds the expires to the get characters character Id opportunities o k response
func (o *GetCharactersCharacterIDOpportunitiesOK) WithExpires(expires string) *GetCharactersCharacterIDOpportunitiesOK {
	o.Expires = expires
	return o
}

// SetExpires sets the expires to the get characters character Id opportunities o k response
func (o *GetCharactersCharacterIDOpportunitiesOK) SetExpires(expires string) {
	o.Expires = expires
}

// WithLastModified adds the lastModified to the get characters character Id opportunities o k response
func (o *GetCharactersCharacterIDOpportunitiesOK) WithLastModified(lastModified string) *GetCharactersCharacterIDOpportunitiesOK {
	o.LastModified = lastModified
	return o
}

// SetLastModified sets the lastModified to the get characters character Id opportunities o k response
func (o *GetCharactersCharacterIDOpportunitiesOK) SetLastModified(lastModified string) {
	o.LastModified = lastModified
}

// WithPayload adds the payload to the get characters character Id opportunities o k response
func (o *GetCharactersCharacterIDOpportunitiesOK) WithPayload(payload models.GetCharactersCharacterIDOpportunitiesOKBody) *GetCharactersCharacterIDOpportunitiesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get characters character Id opportunities o k response
func (o *GetCharactersCharacterIDOpportunitiesOK) SetPayload(payload models.GetCharactersCharacterIDOpportunitiesOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCharactersCharacterIDOpportunitiesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
		payload = make(models.GetCharactersCharacterIDOpportunitiesOKBody, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetCharactersCharacterIDOpportunitiesForbiddenCode is the HTTP code returned for type GetCharactersCharacterIDOpportunitiesForbidden
const GetCharactersCharacterIDOpportunitiesForbiddenCode int = 403

/*GetCharactersCharacterIDOpportunitiesForbidden Forbidden

swagger:response getCharactersCharacterIdOpportunitiesForbidden
*/
type GetCharactersCharacterIDOpportunitiesForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GetCharactersCharacterIDOpportunitiesForbiddenBody `json:"body,omitempty"`
}

// NewGetCharactersCharacterIDOpportunitiesForbidden creates GetCharactersCharacterIDOpportunitiesForbidden with default headers values
func NewGetCharactersCharacterIDOpportunitiesForbidden() *GetCharactersCharacterIDOpportunitiesForbidden {
	return &GetCharactersCharacterIDOpportunitiesForbidden{}
}

// WithPayload adds the payload to the get characters character Id opportunities forbidden response
func (o *GetCharactersCharacterIDOpportunitiesForbidden) WithPayload(payload *models.GetCharactersCharacterIDOpportunitiesForbiddenBody) *GetCharactersCharacterIDOpportunitiesForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get characters character Id opportunities forbidden response
func (o *GetCharactersCharacterIDOpportunitiesForbidden) SetPayload(payload *models.GetCharactersCharacterIDOpportunitiesForbiddenBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCharactersCharacterIDOpportunitiesForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetCharactersCharacterIDOpportunitiesInternalServerErrorCode is the HTTP code returned for type GetCharactersCharacterIDOpportunitiesInternalServerError
const GetCharactersCharacterIDOpportunitiesInternalServerErrorCode int = 500

/*GetCharactersCharacterIDOpportunitiesInternalServerError Internal server error

swagger:response getCharactersCharacterIdOpportunitiesInternalServerError
*/
type GetCharactersCharacterIDOpportunitiesInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.GetCharactersCharacterIDOpportunitiesInternalServerErrorBody `json:"body,omitempty"`
}

// NewGetCharactersCharacterIDOpportunitiesInternalServerError creates GetCharactersCharacterIDOpportunitiesInternalServerError with default headers values
func NewGetCharactersCharacterIDOpportunitiesInternalServerError() *GetCharactersCharacterIDOpportunitiesInternalServerError {
	return &GetCharactersCharacterIDOpportunitiesInternalServerError{}
}

// WithPayload adds the payload to the get characters character Id opportunities internal server error response
func (o *GetCharactersCharacterIDOpportunitiesInternalServerError) WithPayload(payload *models.GetCharactersCharacterIDOpportunitiesInternalServerErrorBody) *GetCharactersCharacterIDOpportunitiesInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get characters character Id opportunities internal server error response
func (o *GetCharactersCharacterIDOpportunitiesInternalServerError) SetPayload(payload *models.GetCharactersCharacterIDOpportunitiesInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCharactersCharacterIDOpportunitiesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
