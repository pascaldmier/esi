package wallet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/pascaldmier/esi/models"
)

// GetCharactersCharacterIDWalletsOKCode is the HTTP code returned for type GetCharactersCharacterIDWalletsOK
const GetCharactersCharacterIDWalletsOKCode int = 200

/*GetCharactersCharacterIDWalletsOK Wallet data for selected user

swagger:response getCharactersCharacterIdWalletsOK
*/
type GetCharactersCharacterIDWalletsOK struct {
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
	Payload models.GetCharactersCharacterIDWalletsOKBody `json:"body,omitempty"`
}

// NewGetCharactersCharacterIDWalletsOK creates GetCharactersCharacterIDWalletsOK with default headers values
func NewGetCharactersCharacterIDWalletsOK() *GetCharactersCharacterIDWalletsOK {
	return &GetCharactersCharacterIDWalletsOK{}
}

// WithCacheControl adds the cacheControl to the get characters character Id wallets o k response
func (o *GetCharactersCharacterIDWalletsOK) WithCacheControl(cacheControl string) *GetCharactersCharacterIDWalletsOK {
	o.CacheControl = cacheControl
	return o
}

// SetCacheControl sets the cacheControl to the get characters character Id wallets o k response
func (o *GetCharactersCharacterIDWalletsOK) SetCacheControl(cacheControl string) {
	o.CacheControl = cacheControl
}

// WithExpires adds the expires to the get characters character Id wallets o k response
func (o *GetCharactersCharacterIDWalletsOK) WithExpires(expires string) *GetCharactersCharacterIDWalletsOK {
	o.Expires = expires
	return o
}

// SetExpires sets the expires to the get characters character Id wallets o k response
func (o *GetCharactersCharacterIDWalletsOK) SetExpires(expires string) {
	o.Expires = expires
}

// WithLastModified adds the lastModified to the get characters character Id wallets o k response
func (o *GetCharactersCharacterIDWalletsOK) WithLastModified(lastModified string) *GetCharactersCharacterIDWalletsOK {
	o.LastModified = lastModified
	return o
}

// SetLastModified sets the lastModified to the get characters character Id wallets o k response
func (o *GetCharactersCharacterIDWalletsOK) SetLastModified(lastModified string) {
	o.LastModified = lastModified
}

// WithPayload adds the payload to the get characters character Id wallets o k response
func (o *GetCharactersCharacterIDWalletsOK) WithPayload(payload models.GetCharactersCharacterIDWalletsOKBody) *GetCharactersCharacterIDWalletsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get characters character Id wallets o k response
func (o *GetCharactersCharacterIDWalletsOK) SetPayload(payload models.GetCharactersCharacterIDWalletsOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCharactersCharacterIDWalletsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
		payload = make(models.GetCharactersCharacterIDWalletsOKBody, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetCharactersCharacterIDWalletsForbiddenCode is the HTTP code returned for type GetCharactersCharacterIDWalletsForbidden
const GetCharactersCharacterIDWalletsForbiddenCode int = 403

/*GetCharactersCharacterIDWalletsForbidden Forbidden

swagger:response getCharactersCharacterIdWalletsForbidden
*/
type GetCharactersCharacterIDWalletsForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GetCharactersCharacterIDWalletsForbiddenBody `json:"body,omitempty"`
}

// NewGetCharactersCharacterIDWalletsForbidden creates GetCharactersCharacterIDWalletsForbidden with default headers values
func NewGetCharactersCharacterIDWalletsForbidden() *GetCharactersCharacterIDWalletsForbidden {
	return &GetCharactersCharacterIDWalletsForbidden{}
}

// WithPayload adds the payload to the get characters character Id wallets forbidden response
func (o *GetCharactersCharacterIDWalletsForbidden) WithPayload(payload *models.GetCharactersCharacterIDWalletsForbiddenBody) *GetCharactersCharacterIDWalletsForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get characters character Id wallets forbidden response
func (o *GetCharactersCharacterIDWalletsForbidden) SetPayload(payload *models.GetCharactersCharacterIDWalletsForbiddenBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCharactersCharacterIDWalletsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetCharactersCharacterIDWalletsInternalServerErrorCode is the HTTP code returned for type GetCharactersCharacterIDWalletsInternalServerError
const GetCharactersCharacterIDWalletsInternalServerErrorCode int = 500

/*GetCharactersCharacterIDWalletsInternalServerError Internal server error

swagger:response getCharactersCharacterIdWalletsInternalServerError
*/
type GetCharactersCharacterIDWalletsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.GetCharactersCharacterIDWalletsInternalServerErrorBody `json:"body,omitempty"`
}

// NewGetCharactersCharacterIDWalletsInternalServerError creates GetCharactersCharacterIDWalletsInternalServerError with default headers values
func NewGetCharactersCharacterIDWalletsInternalServerError() *GetCharactersCharacterIDWalletsInternalServerError {
	return &GetCharactersCharacterIDWalletsInternalServerError{}
}

// WithPayload adds the payload to the get characters character Id wallets internal server error response
func (o *GetCharactersCharacterIDWalletsInternalServerError) WithPayload(payload *models.GetCharactersCharacterIDWalletsInternalServerErrorBody) *GetCharactersCharacterIDWalletsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get characters character Id wallets internal server error response
func (o *GetCharactersCharacterIDWalletsInternalServerError) SetPayload(payload *models.GetCharactersCharacterIDWalletsInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCharactersCharacterIDWalletsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
