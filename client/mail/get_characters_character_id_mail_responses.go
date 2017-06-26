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

// GetCharactersCharacterIDMailReader is a Reader for the GetCharactersCharacterIDMail structure.
type GetCharactersCharacterIDMailReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDMailReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCharactersCharacterIDMailOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetCharactersCharacterIDMailForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCharactersCharacterIDMailInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCharactersCharacterIDMailOK creates a GetCharactersCharacterIDMailOK with default headers values
func NewGetCharactersCharacterIDMailOK() *GetCharactersCharacterIDMailOK {
	return &GetCharactersCharacterIDMailOK{}
}

/*GetCharactersCharacterIDMailOK handles this case with default header values.

The requested mail
*/
type GetCharactersCharacterIDMailOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload models.GetCharactersCharacterIDMailOKBody
}

func (o *GetCharactersCharacterIDMailOK) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/mail/][%d] getCharactersCharacterIdMailOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDMailOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDMailForbidden creates a GetCharactersCharacterIDMailForbidden with default headers values
func NewGetCharactersCharacterIDMailForbidden() *GetCharactersCharacterIDMailForbidden {
	return &GetCharactersCharacterIDMailForbidden{}
}

/*GetCharactersCharacterIDMailForbidden handles this case with default header values.

Forbidden
*/
type GetCharactersCharacterIDMailForbidden struct {
	Payload *models.GetCharactersCharacterIDMailForbiddenBody
}

func (o *GetCharactersCharacterIDMailForbidden) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/mail/][%d] getCharactersCharacterIdMailForbidden  %+v", 403, o.Payload)
}

func (o *GetCharactersCharacterIDMailForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCharactersCharacterIDMailForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDMailInternalServerError creates a GetCharactersCharacterIDMailInternalServerError with default headers values
func NewGetCharactersCharacterIDMailInternalServerError() *GetCharactersCharacterIDMailInternalServerError {
	return &GetCharactersCharacterIDMailInternalServerError{}
}

/*GetCharactersCharacterIDMailInternalServerError handles this case with default header values.

Internal server error
*/
type GetCharactersCharacterIDMailInternalServerError struct {
	Payload *models.GetCharactersCharacterIDMailInternalServerErrorBody
}

func (o *GetCharactersCharacterIDMailInternalServerError) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/mail/][%d] getCharactersCharacterIdMailInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDMailInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCharactersCharacterIDMailInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
