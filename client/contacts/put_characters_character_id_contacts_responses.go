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

// PutCharactersCharacterIDContactsReader is a Reader for the PutCharactersCharacterIDContacts structure.
type PutCharactersCharacterIDContactsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCharactersCharacterIDContactsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewPutCharactersCharacterIDContactsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewPutCharactersCharacterIDContactsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPutCharactersCharacterIDContactsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutCharactersCharacterIDContactsNoContent creates a PutCharactersCharacterIDContactsNoContent with default headers values
func NewPutCharactersCharacterIDContactsNoContent() *PutCharactersCharacterIDContactsNoContent {
	return &PutCharactersCharacterIDContactsNoContent{}
}

/*PutCharactersCharacterIDContactsNoContent handles this case with default header values.

Contacts updated
*/
type PutCharactersCharacterIDContactsNoContent struct {
}

func (o *PutCharactersCharacterIDContactsNoContent) Error() string {
	return fmt.Sprintf("[PUT /characters/{character_id}/contacts/][%d] putCharactersCharacterIdContactsNoContent ", 204)
}

func (o *PutCharactersCharacterIDContactsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCharactersCharacterIDContactsForbidden creates a PutCharactersCharacterIDContactsForbidden with default headers values
func NewPutCharactersCharacterIDContactsForbidden() *PutCharactersCharacterIDContactsForbidden {
	return &PutCharactersCharacterIDContactsForbidden{}
}

/*PutCharactersCharacterIDContactsForbidden handles this case with default header values.

Forbidden
*/
type PutCharactersCharacterIDContactsForbidden struct {
	Payload *models.PutCharactersCharacterIDContactsForbiddenBody
}

func (o *PutCharactersCharacterIDContactsForbidden) Error() string {
	return fmt.Sprintf("[PUT /characters/{character_id}/contacts/][%d] putCharactersCharacterIdContactsForbidden  %+v", 403, o.Payload)
}

func (o *PutCharactersCharacterIDContactsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PutCharactersCharacterIDContactsForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutCharactersCharacterIDContactsInternalServerError creates a PutCharactersCharacterIDContactsInternalServerError with default headers values
func NewPutCharactersCharacterIDContactsInternalServerError() *PutCharactersCharacterIDContactsInternalServerError {
	return &PutCharactersCharacterIDContactsInternalServerError{}
}

/*PutCharactersCharacterIDContactsInternalServerError handles this case with default header values.

Internal server error
*/
type PutCharactersCharacterIDContactsInternalServerError struct {
	Payload *models.PutCharactersCharacterIDContactsInternalServerErrorBody
}

func (o *PutCharactersCharacterIDContactsInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /characters/{character_id}/contacts/][%d] putCharactersCharacterIdContactsInternalServerError  %+v", 500, o.Payload)
}

func (o *PutCharactersCharacterIDContactsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PutCharactersCharacterIDContactsInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
