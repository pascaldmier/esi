package loyalty

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/pascaldmier/esi/models"
)

// GetLoyaltyStoresCorporationIDOffersReader is a Reader for the GetLoyaltyStoresCorporationIDOffers structure.
type GetLoyaltyStoresCorporationIDOffersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLoyaltyStoresCorporationIDOffersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetLoyaltyStoresCorporationIDOffersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetLoyaltyStoresCorporationIDOffersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLoyaltyStoresCorporationIDOffersOK creates a GetLoyaltyStoresCorporationIDOffersOK with default headers values
func NewGetLoyaltyStoresCorporationIDOffersOK() *GetLoyaltyStoresCorporationIDOffersOK {
	return &GetLoyaltyStoresCorporationIDOffersOK{}
}

/*GetLoyaltyStoresCorporationIDOffersOK handles this case with default header values.

A list of offers
*/
type GetLoyaltyStoresCorporationIDOffersOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload models.GetLoyaltyStoresCorporationIDOffersOKBody
}

func (o *GetLoyaltyStoresCorporationIDOffersOK) Error() string {
	return fmt.Sprintf("[GET /loyalty/stores/{corporation_id}/offers/][%d] getLoyaltyStoresCorporationIdOffersOK  %+v", 200, o.Payload)
}

func (o *GetLoyaltyStoresCorporationIDOffersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Cache-Control
	o.CacheControl = response.GetHeader("Cache-Control")

	// response header Expires
	o.Expires = response.GetHeader("Expires")

	// response header Last-Modified
	o.LastModified = response.GetHeader("Last-Modified")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLoyaltyStoresCorporationIDOffersInternalServerError creates a GetLoyaltyStoresCorporationIDOffersInternalServerError with default headers values
func NewGetLoyaltyStoresCorporationIDOffersInternalServerError() *GetLoyaltyStoresCorporationIDOffersInternalServerError {
	return &GetLoyaltyStoresCorporationIDOffersInternalServerError{}
}

/*GetLoyaltyStoresCorporationIDOffersInternalServerError handles this case with default header values.

Internal server error
*/
type GetLoyaltyStoresCorporationIDOffersInternalServerError struct {
	Payload *models.GetLoyaltyStoresCorporationIDOffersInternalServerErrorBody
}

func (o *GetLoyaltyStoresCorporationIDOffersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /loyalty/stores/{corporation_id}/offers/][%d] getLoyaltyStoresCorporationIdOffersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLoyaltyStoresCorporationIDOffersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetLoyaltyStoresCorporationIDOffersInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
