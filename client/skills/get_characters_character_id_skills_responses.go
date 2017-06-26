package skills

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/pascaldmier/esi/models"
)

// GetCharactersCharacterIDSkillsReader is a Reader for the GetCharactersCharacterIDSkills structure.
type GetCharactersCharacterIDSkillsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDSkillsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCharactersCharacterIDSkillsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetCharactersCharacterIDSkillsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCharactersCharacterIDSkillsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCharactersCharacterIDSkillsOK creates a GetCharactersCharacterIDSkillsOK with default headers values
func NewGetCharactersCharacterIDSkillsOK() *GetCharactersCharacterIDSkillsOK {
	return &GetCharactersCharacterIDSkillsOK{}
}

/*GetCharactersCharacterIDSkillsOK handles this case with default header values.

Known skills for the character
*/
type GetCharactersCharacterIDSkillsOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload *models.GetCharactersCharacterIDSkillsOKBody
}

func (o *GetCharactersCharacterIDSkillsOK) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/skills/][%d] getCharactersCharacterIdSkillsOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDSkillsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Cache-Control
	o.CacheControl = response.GetHeader("Cache-Control")

	// response header Expires
	o.Expires = response.GetHeader("Expires")

	// response header Last-Modified
	o.LastModified = response.GetHeader("Last-Modified")

	o.Payload = new(models.GetCharactersCharacterIDSkillsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDSkillsForbidden creates a GetCharactersCharacterIDSkillsForbidden with default headers values
func NewGetCharactersCharacterIDSkillsForbidden() *GetCharactersCharacterIDSkillsForbidden {
	return &GetCharactersCharacterIDSkillsForbidden{}
}

/*GetCharactersCharacterIDSkillsForbidden handles this case with default header values.

Forbidden
*/
type GetCharactersCharacterIDSkillsForbidden struct {
	Payload *models.GetCharactersCharacterIDSkillsForbiddenBody
}

func (o *GetCharactersCharacterIDSkillsForbidden) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/skills/][%d] getCharactersCharacterIdSkillsForbidden  %+v", 403, o.Payload)
}

func (o *GetCharactersCharacterIDSkillsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCharactersCharacterIDSkillsForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDSkillsInternalServerError creates a GetCharactersCharacterIDSkillsInternalServerError with default headers values
func NewGetCharactersCharacterIDSkillsInternalServerError() *GetCharactersCharacterIDSkillsInternalServerError {
	return &GetCharactersCharacterIDSkillsInternalServerError{}
}

/*GetCharactersCharacterIDSkillsInternalServerError handles this case with default header values.

Internal server error
*/
type GetCharactersCharacterIDSkillsInternalServerError struct {
	Payload *models.GetCharactersCharacterIDSkillsInternalServerErrorBody
}

func (o *GetCharactersCharacterIDSkillsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/skills/][%d] getCharactersCharacterIdSkillsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDSkillsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCharactersCharacterIDSkillsInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
