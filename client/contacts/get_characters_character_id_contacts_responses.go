package contacts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/pascaldmier/esi/models"
)

// GetCharactersCharacterIDContactsReader is a Reader for the GetCharactersCharacterIDContacts structure.
type GetCharactersCharacterIDContactsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDContactsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCharactersCharacterIDContactsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetCharactersCharacterIDContactsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCharactersCharacterIDContactsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCharactersCharacterIDContactsOK creates a GetCharactersCharacterIDContactsOK with default headers values
func NewGetCharactersCharacterIDContactsOK() *GetCharactersCharacterIDContactsOK {
	return &GetCharactersCharacterIDContactsOK{}
}

/*GetCharactersCharacterIDContactsOK handles this case with default header values.

A list of contacts
*/
type GetCharactersCharacterIDContactsOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload models.GetCharactersCharacterIDContactsOKBody
}

func (o *GetCharactersCharacterIDContactsOK) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/contacts/][%d] getCharactersCharacterIdContactsOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDContactsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDContactsForbidden creates a GetCharactersCharacterIDContactsForbidden with default headers values
func NewGetCharactersCharacterIDContactsForbidden() *GetCharactersCharacterIDContactsForbidden {
	return &GetCharactersCharacterIDContactsForbidden{}
}

/*GetCharactersCharacterIDContactsForbidden handles this case with default header values.

Forbidden
*/
type GetCharactersCharacterIDContactsForbidden struct {
	Payload *models.GetCharactersCharacterIDContactsForbiddenBody
}

func (o *GetCharactersCharacterIDContactsForbidden) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/contacts/][%d] getCharactersCharacterIdContactsForbidden  %+v", 403, o.Payload)
}

func (o *GetCharactersCharacterIDContactsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCharactersCharacterIDContactsForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDContactsInternalServerError creates a GetCharactersCharacterIDContactsInternalServerError with default headers values
func NewGetCharactersCharacterIDContactsInternalServerError() *GetCharactersCharacterIDContactsInternalServerError {
	return &GetCharactersCharacterIDContactsInternalServerError{}
}

/*GetCharactersCharacterIDContactsInternalServerError handles this case with default header values.

Internal server error
*/
type GetCharactersCharacterIDContactsInternalServerError struct {
	Payload *models.GetCharactersCharacterIDContactsInternalServerErrorBody
}

func (o *GetCharactersCharacterIDContactsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/contacts/][%d] getCharactersCharacterIdContactsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDContactsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCharactersCharacterIDContactsInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
