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

// GetCharactersCharacterIDContactsLabelsReader is a Reader for the GetCharactersCharacterIDContactsLabels structure.
type GetCharactersCharacterIDContactsLabelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDContactsLabelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCharactersCharacterIDContactsLabelsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetCharactersCharacterIDContactsLabelsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCharactersCharacterIDContactsLabelsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCharactersCharacterIDContactsLabelsOK creates a GetCharactersCharacterIDContactsLabelsOK with default headers values
func NewGetCharactersCharacterIDContactsLabelsOK() *GetCharactersCharacterIDContactsLabelsOK {
	return &GetCharactersCharacterIDContactsLabelsOK{}
}

/*GetCharactersCharacterIDContactsLabelsOK handles this case with default header values.

A list of contact labels
*/
type GetCharactersCharacterIDContactsLabelsOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload models.GetCharactersCharacterIDContactsLabelsOKBody
}

func (o *GetCharactersCharacterIDContactsLabelsOK) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/contacts/labels/][%d] getCharactersCharacterIdContactsLabelsOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDContactsLabelsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDContactsLabelsForbidden creates a GetCharactersCharacterIDContactsLabelsForbidden with default headers values
func NewGetCharactersCharacterIDContactsLabelsForbidden() *GetCharactersCharacterIDContactsLabelsForbidden {
	return &GetCharactersCharacterIDContactsLabelsForbidden{}
}

/*GetCharactersCharacterIDContactsLabelsForbidden handles this case with default header values.

Forbidden
*/
type GetCharactersCharacterIDContactsLabelsForbidden struct {
	Payload *models.GetCharactersCharacterIDContactsLabelsForbiddenBody
}

func (o *GetCharactersCharacterIDContactsLabelsForbidden) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/contacts/labels/][%d] getCharactersCharacterIdContactsLabelsForbidden  %+v", 403, o.Payload)
}

func (o *GetCharactersCharacterIDContactsLabelsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCharactersCharacterIDContactsLabelsForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDContactsLabelsInternalServerError creates a GetCharactersCharacterIDContactsLabelsInternalServerError with default headers values
func NewGetCharactersCharacterIDContactsLabelsInternalServerError() *GetCharactersCharacterIDContactsLabelsInternalServerError {
	return &GetCharactersCharacterIDContactsLabelsInternalServerError{}
}

/*GetCharactersCharacterIDContactsLabelsInternalServerError handles this case with default header values.

Internal server error
*/
type GetCharactersCharacterIDContactsLabelsInternalServerError struct {
	Payload *models.GetCharactersCharacterIDContactsLabelsInternalServerErrorBody
}

func (o *GetCharactersCharacterIDContactsLabelsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/contacts/labels/][%d] getCharactersCharacterIdContactsLabelsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDContactsLabelsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCharactersCharacterIDContactsLabelsInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
