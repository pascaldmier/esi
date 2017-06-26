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

// GetCharactersCharacterIDOnlineReader is a Reader for the GetCharactersCharacterIDOnline structure.
type GetCharactersCharacterIDOnlineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDOnlineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCharactersCharacterIDOnlineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetCharactersCharacterIDOnlineForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCharactersCharacterIDOnlineInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCharactersCharacterIDOnlineOK creates a GetCharactersCharacterIDOnlineOK with default headers values
func NewGetCharactersCharacterIDOnlineOK() *GetCharactersCharacterIDOnlineOK {
	return &GetCharactersCharacterIDOnlineOK{}
}

/*GetCharactersCharacterIDOnlineOK handles this case with default header values.

Boolean of if the character is currently online
*/
type GetCharactersCharacterIDOnlineOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload bool
}

func (o *GetCharactersCharacterIDOnlineOK) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/online/][%d] getCharactersCharacterIdOnlineOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDOnlineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDOnlineForbidden creates a GetCharactersCharacterIDOnlineForbidden with default headers values
func NewGetCharactersCharacterIDOnlineForbidden() *GetCharactersCharacterIDOnlineForbidden {
	return &GetCharactersCharacterIDOnlineForbidden{}
}

/*GetCharactersCharacterIDOnlineForbidden handles this case with default header values.

Forbidden
*/
type GetCharactersCharacterIDOnlineForbidden struct {
	Payload *models.GetCharactersCharacterIDOnlineForbiddenBody
}

func (o *GetCharactersCharacterIDOnlineForbidden) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/online/][%d] getCharactersCharacterIdOnlineForbidden  %+v", 403, o.Payload)
}

func (o *GetCharactersCharacterIDOnlineForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCharactersCharacterIDOnlineForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDOnlineInternalServerError creates a GetCharactersCharacterIDOnlineInternalServerError with default headers values
func NewGetCharactersCharacterIDOnlineInternalServerError() *GetCharactersCharacterIDOnlineInternalServerError {
	return &GetCharactersCharacterIDOnlineInternalServerError{}
}

/*GetCharactersCharacterIDOnlineInternalServerError handles this case with default header values.

Internal server error
*/
type GetCharactersCharacterIDOnlineInternalServerError struct {
	Payload *models.GetCharactersCharacterIDOnlineInternalServerErrorBody
}

func (o *GetCharactersCharacterIDOnlineInternalServerError) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/online/][%d] getCharactersCharacterIdOnlineInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDOnlineInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCharactersCharacterIDOnlineInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}