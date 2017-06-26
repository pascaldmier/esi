package market

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/pascaldmier/esi/models"
)

// GetMarketsPricesOKCode is the HTTP code returned for type GetMarketsPricesOK
const GetMarketsPricesOKCode int = 200

/*GetMarketsPricesOK A list of prices

swagger:response getMarketsPricesOK
*/
type GetMarketsPricesOK struct {
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
	Payload models.GetMarketsPricesOKBody `json:"body,omitempty"`
}

// NewGetMarketsPricesOK creates GetMarketsPricesOK with default headers values
func NewGetMarketsPricesOK() *GetMarketsPricesOK {
	return &GetMarketsPricesOK{}
}

// WithCacheControl adds the cacheControl to the get markets prices o k response
func (o *GetMarketsPricesOK) WithCacheControl(cacheControl string) *GetMarketsPricesOK {
	o.CacheControl = cacheControl
	return o
}

// SetCacheControl sets the cacheControl to the get markets prices o k response
func (o *GetMarketsPricesOK) SetCacheControl(cacheControl string) {
	o.CacheControl = cacheControl
}

// WithExpires adds the expires to the get markets prices o k response
func (o *GetMarketsPricesOK) WithExpires(expires string) *GetMarketsPricesOK {
	o.Expires = expires
	return o
}

// SetExpires sets the expires to the get markets prices o k response
func (o *GetMarketsPricesOK) SetExpires(expires string) {
	o.Expires = expires
}

// WithLastModified adds the lastModified to the get markets prices o k response
func (o *GetMarketsPricesOK) WithLastModified(lastModified string) *GetMarketsPricesOK {
	o.LastModified = lastModified
	return o
}

// SetLastModified sets the lastModified to the get markets prices o k response
func (o *GetMarketsPricesOK) SetLastModified(lastModified string) {
	o.LastModified = lastModified
}

// WithPayload adds the payload to the get markets prices o k response
func (o *GetMarketsPricesOK) WithPayload(payload models.GetMarketsPricesOKBody) *GetMarketsPricesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get markets prices o k response
func (o *GetMarketsPricesOK) SetPayload(payload models.GetMarketsPricesOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMarketsPricesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
		payload = make(models.GetMarketsPricesOKBody, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetMarketsPricesInternalServerErrorCode is the HTTP code returned for type GetMarketsPricesInternalServerError
const GetMarketsPricesInternalServerErrorCode int = 500

/*GetMarketsPricesInternalServerError Internal server error

swagger:response getMarketsPricesInternalServerError
*/
type GetMarketsPricesInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.GetMarketsPricesInternalServerErrorBody `json:"body,omitempty"`
}

// NewGetMarketsPricesInternalServerError creates GetMarketsPricesInternalServerError with default headers values
func NewGetMarketsPricesInternalServerError() *GetMarketsPricesInternalServerError {
	return &GetMarketsPricesInternalServerError{}
}

// WithPayload adds the payload to the get markets prices internal server error response
func (o *GetMarketsPricesInternalServerError) WithPayload(payload *models.GetMarketsPricesInternalServerErrorBody) *GetMarketsPricesInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get markets prices internal server error response
func (o *GetMarketsPricesInternalServerError) SetPayload(payload *models.GetMarketsPricesInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMarketsPricesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}