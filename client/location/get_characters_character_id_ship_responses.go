package location

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/pascaldmier/esi/models"
)

// GetCharactersCharacterIDShipReader is a Reader for the GetCharactersCharacterIDShip structure.
type GetCharactersCharacterIDShipReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDShipReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCharactersCharacterIDShipOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetCharactersCharacterIDShipForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCharactersCharacterIDShipInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCharactersCharacterIDShipOK creates a GetCharactersCharacterIDShipOK with default headers values
func NewGetCharactersCharacterIDShipOK() *GetCharactersCharacterIDShipOK {
	return &GetCharactersCharacterIDShipOK{}
}

/*GetCharactersCharacterIDShipOK handles this case with default header values.

Get the current ship type, name and id
*/
type GetCharactersCharacterIDShipOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload *models.GetCharactersCharacterIDShipOKBody
}

func (o *GetCharactersCharacterIDShipOK) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/ship/][%d] getCharactersCharacterIdShipOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDShipOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Cache-Control
	o.CacheControl = response.GetHeader("Cache-Control")

	// response header Expires
	o.Expires = response.GetHeader("Expires")

	// response header Last-Modified
	o.LastModified = response.GetHeader("Last-Modified")

	o.Payload = new(models.GetCharactersCharacterIDShipOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDShipForbidden creates a GetCharactersCharacterIDShipForbidden with default headers values
func NewGetCharactersCharacterIDShipForbidden() *GetCharactersCharacterIDShipForbidden {
	return &GetCharactersCharacterIDShipForbidden{}
}

/*GetCharactersCharacterIDShipForbidden handles this case with default header values.

Forbidden
*/
type GetCharactersCharacterIDShipForbidden struct {
	Payload *models.GetCharactersCharacterIDShipForbiddenBody
}

func (o *GetCharactersCharacterIDShipForbidden) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/ship/][%d] getCharactersCharacterIdShipForbidden  %+v", 403, o.Payload)
}

func (o *GetCharactersCharacterIDShipForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCharactersCharacterIDShipForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDShipInternalServerError creates a GetCharactersCharacterIDShipInternalServerError with default headers values
func NewGetCharactersCharacterIDShipInternalServerError() *GetCharactersCharacterIDShipInternalServerError {
	return &GetCharactersCharacterIDShipInternalServerError{}
}

/*GetCharactersCharacterIDShipInternalServerError handles this case with default header values.

Internal server error
*/
type GetCharactersCharacterIDShipInternalServerError struct {
	Payload *models.GetCharactersCharacterIDShipInternalServerErrorBody
}

func (o *GetCharactersCharacterIDShipInternalServerError) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/ship/][%d] getCharactersCharacterIdShipInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDShipInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCharactersCharacterIDShipInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}