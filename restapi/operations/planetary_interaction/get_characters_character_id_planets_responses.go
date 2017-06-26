package planetary_interaction

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/pascaldmier/esi/models"
)

// GetCharactersCharacterIDPlanetsOKCode is the HTTP code returned for type GetCharactersCharacterIDPlanetsOK
const GetCharactersCharacterIDPlanetsOKCode int = 200

/*GetCharactersCharacterIDPlanetsOK List of colonies

swagger:response getCharactersCharacterIdPlanetsOK
*/
type GetCharactersCharacterIDPlanetsOK struct {
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
	Payload models.GetCharactersCharacterIDPlanetsOKBody `json:"body,omitempty"`
}

// NewGetCharactersCharacterIDPlanetsOK creates GetCharactersCharacterIDPlanetsOK with default headers values
func NewGetCharactersCharacterIDPlanetsOK() *GetCharactersCharacterIDPlanetsOK {
	return &GetCharactersCharacterIDPlanetsOK{}
}

// WithCacheControl adds the cacheControl to the get characters character Id planets o k response
func (o *GetCharactersCharacterIDPlanetsOK) WithCacheControl(cacheControl string) *GetCharactersCharacterIDPlanetsOK {
	o.CacheControl = cacheControl
	return o
}

// SetCacheControl sets the cacheControl to the get characters character Id planets o k response
func (o *GetCharactersCharacterIDPlanetsOK) SetCacheControl(cacheControl string) {
	o.CacheControl = cacheControl
}

// WithExpires adds the expires to the get characters character Id planets o k response
func (o *GetCharactersCharacterIDPlanetsOK) WithExpires(expires string) *GetCharactersCharacterIDPlanetsOK {
	o.Expires = expires
	return o
}

// SetExpires sets the expires to the get characters character Id planets o k response
func (o *GetCharactersCharacterIDPlanetsOK) SetExpires(expires string) {
	o.Expires = expires
}

// WithLastModified adds the lastModified to the get characters character Id planets o k response
func (o *GetCharactersCharacterIDPlanetsOK) WithLastModified(lastModified string) *GetCharactersCharacterIDPlanetsOK {
	o.LastModified = lastModified
	return o
}

// SetLastModified sets the lastModified to the get characters character Id planets o k response
func (o *GetCharactersCharacterIDPlanetsOK) SetLastModified(lastModified string) {
	o.LastModified = lastModified
}

// WithPayload adds the payload to the get characters character Id planets o k response
func (o *GetCharactersCharacterIDPlanetsOK) WithPayload(payload models.GetCharactersCharacterIDPlanetsOKBody) *GetCharactersCharacterIDPlanetsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get characters character Id planets o k response
func (o *GetCharactersCharacterIDPlanetsOK) SetPayload(payload models.GetCharactersCharacterIDPlanetsOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCharactersCharacterIDPlanetsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
		payload = make(models.GetCharactersCharacterIDPlanetsOKBody, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetCharactersCharacterIDPlanetsForbiddenCode is the HTTP code returned for type GetCharactersCharacterIDPlanetsForbidden
const GetCharactersCharacterIDPlanetsForbiddenCode int = 403

/*GetCharactersCharacterIDPlanetsForbidden Forbidden

swagger:response getCharactersCharacterIdPlanetsForbidden
*/
type GetCharactersCharacterIDPlanetsForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GetCharactersCharacterIDPlanetsForbiddenBody `json:"body,omitempty"`
}

// NewGetCharactersCharacterIDPlanetsForbidden creates GetCharactersCharacterIDPlanetsForbidden with default headers values
func NewGetCharactersCharacterIDPlanetsForbidden() *GetCharactersCharacterIDPlanetsForbidden {
	return &GetCharactersCharacterIDPlanetsForbidden{}
}

// WithPayload adds the payload to the get characters character Id planets forbidden response
func (o *GetCharactersCharacterIDPlanetsForbidden) WithPayload(payload *models.GetCharactersCharacterIDPlanetsForbiddenBody) *GetCharactersCharacterIDPlanetsForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get characters character Id planets forbidden response
func (o *GetCharactersCharacterIDPlanetsForbidden) SetPayload(payload *models.GetCharactersCharacterIDPlanetsForbiddenBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCharactersCharacterIDPlanetsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetCharactersCharacterIDPlanetsInternalServerErrorCode is the HTTP code returned for type GetCharactersCharacterIDPlanetsInternalServerError
const GetCharactersCharacterIDPlanetsInternalServerErrorCode int = 500

/*GetCharactersCharacterIDPlanetsInternalServerError Internal server error

swagger:response getCharactersCharacterIdPlanetsInternalServerError
*/
type GetCharactersCharacterIDPlanetsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.GetCharactersCharacterIDPlanetsInternalServerErrorBody `json:"body,omitempty"`
}

// NewGetCharactersCharacterIDPlanetsInternalServerError creates GetCharactersCharacterIDPlanetsInternalServerError with default headers values
func NewGetCharactersCharacterIDPlanetsInternalServerError() *GetCharactersCharacterIDPlanetsInternalServerError {
	return &GetCharactersCharacterIDPlanetsInternalServerError{}
}

// WithPayload adds the payload to the get characters character Id planets internal server error response
func (o *GetCharactersCharacterIDPlanetsInternalServerError) WithPayload(payload *models.GetCharactersCharacterIDPlanetsInternalServerErrorBody) *GetCharactersCharacterIDPlanetsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get characters character Id planets internal server error response
func (o *GetCharactersCharacterIDPlanetsInternalServerError) SetPayload(payload *models.GetCharactersCharacterIDPlanetsInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCharactersCharacterIDPlanetsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}