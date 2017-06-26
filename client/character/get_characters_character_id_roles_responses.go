package character

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/pascaldmier/esi/models"
)

// GetCharactersCharacterIDRolesReader is a Reader for the GetCharactersCharacterIDRoles structure.
type GetCharactersCharacterIDRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCharactersCharacterIDRolesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetCharactersCharacterIDRolesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCharactersCharacterIDRolesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCharactersCharacterIDRolesOK creates a GetCharactersCharacterIDRolesOK with default headers values
func NewGetCharactersCharacterIDRolesOK() *GetCharactersCharacterIDRolesOK {
	return &GetCharactersCharacterIDRolesOK{}
}

/*GetCharactersCharacterIDRolesOK handles this case with default header values.

The character's roles in thier corporation
*/
type GetCharactersCharacterIDRolesOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload []string
}

func (o *GetCharactersCharacterIDRolesOK) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/roles/][%d] getCharactersCharacterIdRolesOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDRolesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDRolesForbidden creates a GetCharactersCharacterIDRolesForbidden with default headers values
func NewGetCharactersCharacterIDRolesForbidden() *GetCharactersCharacterIDRolesForbidden {
	return &GetCharactersCharacterIDRolesForbidden{}
}

/*GetCharactersCharacterIDRolesForbidden handles this case with default header values.

Forbidden
*/
type GetCharactersCharacterIDRolesForbidden struct {
	Payload *models.GetCharactersCharacterIDRolesForbiddenBody
}

func (o *GetCharactersCharacterIDRolesForbidden) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/roles/][%d] getCharactersCharacterIdRolesForbidden  %+v", 403, o.Payload)
}

func (o *GetCharactersCharacterIDRolesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCharactersCharacterIDRolesForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDRolesInternalServerError creates a GetCharactersCharacterIDRolesInternalServerError with default headers values
func NewGetCharactersCharacterIDRolesInternalServerError() *GetCharactersCharacterIDRolesInternalServerError {
	return &GetCharactersCharacterIDRolesInternalServerError{}
}

/*GetCharactersCharacterIDRolesInternalServerError handles this case with default header values.

Internal server error
*/
type GetCharactersCharacterIDRolesInternalServerError struct {
	Payload *models.GetCharactersCharacterIDRolesInternalServerErrorBody
}

func (o *GetCharactersCharacterIDRolesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/roles/][%d] getCharactersCharacterIdRolesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDRolesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCharactersCharacterIDRolesInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}