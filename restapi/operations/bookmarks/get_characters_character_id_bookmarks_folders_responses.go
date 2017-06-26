package bookmarks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/pascaldmier/esi/models"
)

// GetCharactersCharacterIDBookmarksFoldersOKCode is the HTTP code returned for type GetCharactersCharacterIDBookmarksFoldersOK
const GetCharactersCharacterIDBookmarksFoldersOKCode int = 200

/*GetCharactersCharacterIDBookmarksFoldersOK List of bookmark folders

swagger:response getCharactersCharacterIdBookmarksFoldersOK
*/
type GetCharactersCharacterIDBookmarksFoldersOK struct {
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
	Payload models.GetCharactersCharacterIDBookmarksFoldersOKBody `json:"body,omitempty"`
}

// NewGetCharactersCharacterIDBookmarksFoldersOK creates GetCharactersCharacterIDBookmarksFoldersOK with default headers values
func NewGetCharactersCharacterIDBookmarksFoldersOK() *GetCharactersCharacterIDBookmarksFoldersOK {
	return &GetCharactersCharacterIDBookmarksFoldersOK{}
}

// WithCacheControl adds the cacheControl to the get characters character Id bookmarks folders o k response
func (o *GetCharactersCharacterIDBookmarksFoldersOK) WithCacheControl(cacheControl string) *GetCharactersCharacterIDBookmarksFoldersOK {
	o.CacheControl = cacheControl
	return o
}

// SetCacheControl sets the cacheControl to the get characters character Id bookmarks folders o k response
func (o *GetCharactersCharacterIDBookmarksFoldersOK) SetCacheControl(cacheControl string) {
	o.CacheControl = cacheControl
}

// WithExpires adds the expires to the get characters character Id bookmarks folders o k response
func (o *GetCharactersCharacterIDBookmarksFoldersOK) WithExpires(expires string) *GetCharactersCharacterIDBookmarksFoldersOK {
	o.Expires = expires
	return o
}

// SetExpires sets the expires to the get characters character Id bookmarks folders o k response
func (o *GetCharactersCharacterIDBookmarksFoldersOK) SetExpires(expires string) {
	o.Expires = expires
}

// WithLastModified adds the lastModified to the get characters character Id bookmarks folders o k response
func (o *GetCharactersCharacterIDBookmarksFoldersOK) WithLastModified(lastModified string) *GetCharactersCharacterIDBookmarksFoldersOK {
	o.LastModified = lastModified
	return o
}

// SetLastModified sets the lastModified to the get characters character Id bookmarks folders o k response
func (o *GetCharactersCharacterIDBookmarksFoldersOK) SetLastModified(lastModified string) {
	o.LastModified = lastModified
}

// WithPayload adds the payload to the get characters character Id bookmarks folders o k response
func (o *GetCharactersCharacterIDBookmarksFoldersOK) WithPayload(payload models.GetCharactersCharacterIDBookmarksFoldersOKBody) *GetCharactersCharacterIDBookmarksFoldersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get characters character Id bookmarks folders o k response
func (o *GetCharactersCharacterIDBookmarksFoldersOK) SetPayload(payload models.GetCharactersCharacterIDBookmarksFoldersOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCharactersCharacterIDBookmarksFoldersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
		payload = make(models.GetCharactersCharacterIDBookmarksFoldersOKBody, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetCharactersCharacterIDBookmarksFoldersForbiddenCode is the HTTP code returned for type GetCharactersCharacterIDBookmarksFoldersForbidden
const GetCharactersCharacterIDBookmarksFoldersForbiddenCode int = 403

/*GetCharactersCharacterIDBookmarksFoldersForbidden Forbidden

swagger:response getCharactersCharacterIdBookmarksFoldersForbidden
*/
type GetCharactersCharacterIDBookmarksFoldersForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GetCharactersCharacterIDBookmarksFoldersForbiddenBody `json:"body,omitempty"`
}

// NewGetCharactersCharacterIDBookmarksFoldersForbidden creates GetCharactersCharacterIDBookmarksFoldersForbidden with default headers values
func NewGetCharactersCharacterIDBookmarksFoldersForbidden() *GetCharactersCharacterIDBookmarksFoldersForbidden {
	return &GetCharactersCharacterIDBookmarksFoldersForbidden{}
}

// WithPayload adds the payload to the get characters character Id bookmarks folders forbidden response
func (o *GetCharactersCharacterIDBookmarksFoldersForbidden) WithPayload(payload *models.GetCharactersCharacterIDBookmarksFoldersForbiddenBody) *GetCharactersCharacterIDBookmarksFoldersForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get characters character Id bookmarks folders forbidden response
func (o *GetCharactersCharacterIDBookmarksFoldersForbidden) SetPayload(payload *models.GetCharactersCharacterIDBookmarksFoldersForbiddenBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCharactersCharacterIDBookmarksFoldersForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetCharactersCharacterIDBookmarksFoldersInternalServerErrorCode is the HTTP code returned for type GetCharactersCharacterIDBookmarksFoldersInternalServerError
const GetCharactersCharacterIDBookmarksFoldersInternalServerErrorCode int = 500

/*GetCharactersCharacterIDBookmarksFoldersInternalServerError Internal server error

swagger:response getCharactersCharacterIdBookmarksFoldersInternalServerError
*/
type GetCharactersCharacterIDBookmarksFoldersInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.GetCharactersCharacterIDBookmarksFoldersInternalServerErrorBody `json:"body,omitempty"`
}

// NewGetCharactersCharacterIDBookmarksFoldersInternalServerError creates GetCharactersCharacterIDBookmarksFoldersInternalServerError with default headers values
func NewGetCharactersCharacterIDBookmarksFoldersInternalServerError() *GetCharactersCharacterIDBookmarksFoldersInternalServerError {
	return &GetCharactersCharacterIDBookmarksFoldersInternalServerError{}
}

// WithPayload adds the payload to the get characters character Id bookmarks folders internal server error response
func (o *GetCharactersCharacterIDBookmarksFoldersInternalServerError) WithPayload(payload *models.GetCharactersCharacterIDBookmarksFoldersInternalServerErrorBody) *GetCharactersCharacterIDBookmarksFoldersInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get characters character Id bookmarks folders internal server error response
func (o *GetCharactersCharacterIDBookmarksFoldersInternalServerError) SetPayload(payload *models.GetCharactersCharacterIDBookmarksFoldersInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCharactersCharacterIDBookmarksFoldersInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
