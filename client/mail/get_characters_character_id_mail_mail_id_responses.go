package mail

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/pascaldmier/esi/models"
)

// GetCharactersCharacterIDMailMailIDReader is a Reader for the GetCharactersCharacterIDMailMailID structure.
type GetCharactersCharacterIDMailMailIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDMailMailIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCharactersCharacterIDMailMailIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetCharactersCharacterIDMailMailIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetCharactersCharacterIDMailMailIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCharactersCharacterIDMailMailIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCharactersCharacterIDMailMailIDOK creates a GetCharactersCharacterIDMailMailIDOK with default headers values
func NewGetCharactersCharacterIDMailMailIDOK() *GetCharactersCharacterIDMailMailIDOK {
	return &GetCharactersCharacterIDMailMailIDOK{}
}

/*GetCharactersCharacterIDMailMailIDOK handles this case with default header values.

Contents of a mail
*/
type GetCharactersCharacterIDMailMailIDOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload *models.GetCharactersCharacterIDMailMailIDOKBody
}

func (o *GetCharactersCharacterIDMailMailIDOK) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/mail/{mail_id}/][%d] getCharactersCharacterIdMailMailIdOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDMailMailIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Cache-Control
	o.CacheControl = response.GetHeader("Cache-Control")

	// response header Expires
	o.Expires = response.GetHeader("Expires")

	// response header Last-Modified
	o.LastModified = response.GetHeader("Last-Modified")

	o.Payload = new(models.GetCharactersCharacterIDMailMailIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDMailMailIDForbidden creates a GetCharactersCharacterIDMailMailIDForbidden with default headers values
func NewGetCharactersCharacterIDMailMailIDForbidden() *GetCharactersCharacterIDMailMailIDForbidden {
	return &GetCharactersCharacterIDMailMailIDForbidden{}
}

/*GetCharactersCharacterIDMailMailIDForbidden handles this case with default header values.

Forbidden
*/
type GetCharactersCharacterIDMailMailIDForbidden struct {
	Payload *models.GetCharactersCharacterIDMailMailIDForbiddenBody
}

func (o *GetCharactersCharacterIDMailMailIDForbidden) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/mail/{mail_id}/][%d] getCharactersCharacterIdMailMailIdForbidden  %+v", 403, o.Payload)
}

func (o *GetCharactersCharacterIDMailMailIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCharactersCharacterIDMailMailIDForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDMailMailIDNotFound creates a GetCharactersCharacterIDMailMailIDNotFound with default headers values
func NewGetCharactersCharacterIDMailMailIDNotFound() *GetCharactersCharacterIDMailMailIDNotFound {
	return &GetCharactersCharacterIDMailMailIDNotFound{}
}

/*GetCharactersCharacterIDMailMailIDNotFound handles this case with default header values.

Mail not found
*/
type GetCharactersCharacterIDMailMailIDNotFound struct {
	Payload *models.GetCharactersCharacterIDMailMailIDNotFoundBody
}

func (o *GetCharactersCharacterIDMailMailIDNotFound) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/mail/{mail_id}/][%d] getCharactersCharacterIdMailMailIdNotFound  %+v", 404, o.Payload)
}

func (o *GetCharactersCharacterIDMailMailIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCharactersCharacterIDMailMailIDNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDMailMailIDInternalServerError creates a GetCharactersCharacterIDMailMailIDInternalServerError with default headers values
func NewGetCharactersCharacterIDMailMailIDInternalServerError() *GetCharactersCharacterIDMailMailIDInternalServerError {
	return &GetCharactersCharacterIDMailMailIDInternalServerError{}
}

/*GetCharactersCharacterIDMailMailIDInternalServerError handles this case with default header values.

Internal server error
*/
type GetCharactersCharacterIDMailMailIDInternalServerError struct {
	Payload *models.GetCharactersCharacterIDMailMailIDInternalServerErrorBody
}

func (o *GetCharactersCharacterIDMailMailIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/mail/{mail_id}/][%d] getCharactersCharacterIdMailMailIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDMailMailIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCharactersCharacterIDMailMailIDInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
