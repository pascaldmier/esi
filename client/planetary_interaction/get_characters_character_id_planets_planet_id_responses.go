package planetary_interaction

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/pascaldmier/esi/models"
)

// GetCharactersCharacterIDPlanetsPlanetIDReader is a Reader for the GetCharactersCharacterIDPlanetsPlanetID structure.
type GetCharactersCharacterIDPlanetsPlanetIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDPlanetsPlanetIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCharactersCharacterIDPlanetsPlanetIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetCharactersCharacterIDPlanetsPlanetIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetCharactersCharacterIDPlanetsPlanetIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCharactersCharacterIDPlanetsPlanetIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCharactersCharacterIDPlanetsPlanetIDOK creates a GetCharactersCharacterIDPlanetsPlanetIDOK with default headers values
func NewGetCharactersCharacterIDPlanetsPlanetIDOK() *GetCharactersCharacterIDPlanetsPlanetIDOK {
	return &GetCharactersCharacterIDPlanetsPlanetIDOK{}
}

/*GetCharactersCharacterIDPlanetsPlanetIDOK handles this case with default header values.

Colony layout
*/
type GetCharactersCharacterIDPlanetsPlanetIDOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload *models.GetCharactersCharacterIDPlanetsPlanetIDOKBody
}

func (o *GetCharactersCharacterIDPlanetsPlanetIDOK) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/planets/{planet_id}/][%d] getCharactersCharacterIdPlanetsPlanetIdOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDPlanetsPlanetIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Cache-Control
	o.CacheControl = response.GetHeader("Cache-Control")

	// response header Expires
	o.Expires = response.GetHeader("Expires")

	// response header Last-Modified
	o.LastModified = response.GetHeader("Last-Modified")

	o.Payload = new(models.GetCharactersCharacterIDPlanetsPlanetIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDPlanetsPlanetIDForbidden creates a GetCharactersCharacterIDPlanetsPlanetIDForbidden with default headers values
func NewGetCharactersCharacterIDPlanetsPlanetIDForbidden() *GetCharactersCharacterIDPlanetsPlanetIDForbidden {
	return &GetCharactersCharacterIDPlanetsPlanetIDForbidden{}
}

/*GetCharactersCharacterIDPlanetsPlanetIDForbidden handles this case with default header values.

Forbidden
*/
type GetCharactersCharacterIDPlanetsPlanetIDForbidden struct {
	Payload *models.GetCharactersCharacterIDPlanetsPlanetIDForbiddenBody
}

func (o *GetCharactersCharacterIDPlanetsPlanetIDForbidden) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/planets/{planet_id}/][%d] getCharactersCharacterIdPlanetsPlanetIdForbidden  %+v", 403, o.Payload)
}

func (o *GetCharactersCharacterIDPlanetsPlanetIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCharactersCharacterIDPlanetsPlanetIDForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDPlanetsPlanetIDNotFound creates a GetCharactersCharacterIDPlanetsPlanetIDNotFound with default headers values
func NewGetCharactersCharacterIDPlanetsPlanetIDNotFound() *GetCharactersCharacterIDPlanetsPlanetIDNotFound {
	return &GetCharactersCharacterIDPlanetsPlanetIDNotFound{}
}

/*GetCharactersCharacterIDPlanetsPlanetIDNotFound handles this case with default header values.

Colony not found
*/
type GetCharactersCharacterIDPlanetsPlanetIDNotFound struct {
	Payload *models.GetCharactersCharacterIDPlanetsPlanetIDNotFoundBody
}

func (o *GetCharactersCharacterIDPlanetsPlanetIDNotFound) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/planets/{planet_id}/][%d] getCharactersCharacterIdPlanetsPlanetIdNotFound  %+v", 404, o.Payload)
}

func (o *GetCharactersCharacterIDPlanetsPlanetIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCharactersCharacterIDPlanetsPlanetIDNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDPlanetsPlanetIDInternalServerError creates a GetCharactersCharacterIDPlanetsPlanetIDInternalServerError with default headers values
func NewGetCharactersCharacterIDPlanetsPlanetIDInternalServerError() *GetCharactersCharacterIDPlanetsPlanetIDInternalServerError {
	return &GetCharactersCharacterIDPlanetsPlanetIDInternalServerError{}
}

/*GetCharactersCharacterIDPlanetsPlanetIDInternalServerError handles this case with default header values.

Internal server error
*/
type GetCharactersCharacterIDPlanetsPlanetIDInternalServerError struct {
	Payload *models.GetCharactersCharacterIDPlanetsPlanetIDInternalServerErrorBody
}

func (o *GetCharactersCharacterIDPlanetsPlanetIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/planets/{planet_id}/][%d] getCharactersCharacterIdPlanetsPlanetIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDPlanetsPlanetIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCharactersCharacterIDPlanetsPlanetIDInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
