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

// GetCharactersNamesReader is a Reader for the GetCharactersNames structure.
type GetCharactersNamesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersNamesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCharactersNamesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetCharactersNamesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCharactersNamesOK creates a GetCharactersNamesOK with default headers values
func NewGetCharactersNamesOK() *GetCharactersNamesOK {
	return &GetCharactersNamesOK{}
}

/*GetCharactersNamesOK handles this case with default header values.

List of id/name associations
*/
type GetCharactersNamesOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload models.GetCharactersNamesOKBody
}

func (o *GetCharactersNamesOK) Error() string {
	return fmt.Sprintf("[GET /characters/names/][%d] getCharactersNamesOK  %+v", 200, o.Payload)
}

func (o *GetCharactersNamesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersNamesInternalServerError creates a GetCharactersNamesInternalServerError with default headers values
func NewGetCharactersNamesInternalServerError() *GetCharactersNamesInternalServerError {
	return &GetCharactersNamesInternalServerError{}
}

/*GetCharactersNamesInternalServerError handles this case with default header values.

Internal server error
*/
type GetCharactersNamesInternalServerError struct {
	Payload *models.GetCharactersNamesInternalServerErrorBody
}

func (o *GetCharactersNamesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /characters/names/][%d] getCharactersNamesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersNamesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCharactersNamesInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}