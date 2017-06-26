package character

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/pascaldmier/esi/models"
)

// GetCharactersCharacterIDBlueprintsOKCode is the HTTP code returned for type GetCharactersCharacterIDBlueprintsOK
const GetCharactersCharacterIDBlueprintsOKCode int = 200

/*GetCharactersCharacterIDBlueprintsOK A list of blueprints

swagger:response getCharactersCharacterIdBlueprintsOK
*/
type GetCharactersCharacterIDBlueprintsOK struct {
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
	Payload models.GetCharactersCharacterIDBlueprintsOKBody `json:"body,omitempty"`
}

// NewGetCharactersCharacterIDBlueprintsOK creates GetCharactersCharacterIDBlueprintsOK with default headers values
func NewGetCharactersCharacterIDBlueprintsOK() *GetCharactersCharacterIDBlueprintsOK {
	return &GetCharactersCharacterIDBlueprintsOK{}
}

// WithCacheControl adds the cacheControl to the get characters character Id blueprints o k response
func (o *GetCharactersCharacterIDBlueprintsOK) WithCacheControl(cacheControl string) *GetCharactersCharacterIDBlueprintsOK {
	o.CacheControl = cacheControl
	return o
}

// SetCacheControl sets the cacheControl to the get characters character Id blueprints o k response
func (o *GetCharactersCharacterIDBlueprintsOK) SetCacheControl(cacheControl string) {
	o.CacheControl = cacheControl
}

// WithExpires adds the expires to the get characters character Id blueprints o k response
func (o *GetCharactersCharacterIDBlueprintsOK) WithExpires(expires string) *GetCharactersCharacterIDBlueprintsOK {
	o.Expires = expires
	return o
}

// SetExpires sets the expires to the get characters character Id blueprints o k response
func (o *GetCharactersCharacterIDBlueprintsOK) SetExpires(expires string) {
	o.Expires = expires
}

// WithLastModified adds the lastModified to the get characters character Id blueprints o k response
func (o *GetCharactersCharacterIDBlueprintsOK) WithLastModified(lastModified string) *GetCharactersCharacterIDBlueprintsOK {
	o.LastModified = lastModified
	return o
}

// SetLastModified sets the lastModified to the get characters character Id blueprints o k response
func (o *GetCharactersCharacterIDBlueprintsOK) SetLastModified(lastModified string) {
	o.LastModified = lastModified
}

// WithPayload adds the payload to the get characters character Id blueprints o k response
func (o *GetCharactersCharacterIDBlueprintsOK) WithPayload(payload models.GetCharactersCharacterIDBlueprintsOKBody) *GetCharactersCharacterIDBlueprintsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get characters character Id blueprints o k response
func (o *GetCharactersCharacterIDBlueprintsOK) SetPayload(payload models.GetCharactersCharacterIDBlueprintsOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCharactersCharacterIDBlueprintsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
		payload = make(models.GetCharactersCharacterIDBlueprintsOKBody, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetCharactersCharacterIDBlueprintsForbiddenCode is the HTTP code returned for type GetCharactersCharacterIDBlueprintsForbidden
const GetCharactersCharacterIDBlueprintsForbiddenCode int = 403

/*GetCharactersCharacterIDBlueprintsForbidden Forbidden

swagger:response getCharactersCharacterIdBlueprintsForbidden
*/
type GetCharactersCharacterIDBlueprintsForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GetCharactersCharacterIDBlueprintsForbiddenBody `json:"body,omitempty"`
}

// NewGetCharactersCharacterIDBlueprintsForbidden creates GetCharactersCharacterIDBlueprintsForbidden with default headers values
func NewGetCharactersCharacterIDBlueprintsForbidden() *GetCharactersCharacterIDBlueprintsForbidden {
	return &GetCharactersCharacterIDBlueprintsForbidden{}
}

// WithPayload adds the payload to the get characters character Id blueprints forbidden response
func (o *GetCharactersCharacterIDBlueprintsForbidden) WithPayload(payload *models.GetCharactersCharacterIDBlueprintsForbiddenBody) *GetCharactersCharacterIDBlueprintsForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get characters character Id blueprints forbidden response
func (o *GetCharactersCharacterIDBlueprintsForbidden) SetPayload(payload *models.GetCharactersCharacterIDBlueprintsForbiddenBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCharactersCharacterIDBlueprintsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetCharactersCharacterIDBlueprintsInternalServerErrorCode is the HTTP code returned for type GetCharactersCharacterIDBlueprintsInternalServerError
const GetCharactersCharacterIDBlueprintsInternalServerErrorCode int = 500

/*GetCharactersCharacterIDBlueprintsInternalServerError Internal server error

swagger:response getCharactersCharacterIdBlueprintsInternalServerError
*/
type GetCharactersCharacterIDBlueprintsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.GetCharactersCharacterIDBlueprintsInternalServerErrorBody `json:"body,omitempty"`
}

// NewGetCharactersCharacterIDBlueprintsInternalServerError creates GetCharactersCharacterIDBlueprintsInternalServerError with default headers values
func NewGetCharactersCharacterIDBlueprintsInternalServerError() *GetCharactersCharacterIDBlueprintsInternalServerError {
	return &GetCharactersCharacterIDBlueprintsInternalServerError{}
}

// WithPayload adds the payload to the get characters character Id blueprints internal server error response
func (o *GetCharactersCharacterIDBlueprintsInternalServerError) WithPayload(payload *models.GetCharactersCharacterIDBlueprintsInternalServerErrorBody) *GetCharactersCharacterIDBlueprintsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get characters character Id blueprints internal server error response
func (o *GetCharactersCharacterIDBlueprintsInternalServerError) SetPayload(payload *models.GetCharactersCharacterIDBlueprintsInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCharactersCharacterIDBlueprintsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}