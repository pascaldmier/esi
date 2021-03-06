package universe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/pascaldmier/esi/models"
)

// GetUniverseFactionsReader is a Reader for the GetUniverseFactions structure.
type GetUniverseFactionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUniverseFactionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUniverseFactionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetUniverseFactionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUniverseFactionsOK creates a GetUniverseFactionsOK with default headers values
func NewGetUniverseFactionsOK() *GetUniverseFactionsOK {
	return &GetUniverseFactionsOK{}
}

/*GetUniverseFactionsOK handles this case with default header values.

A list of factions
*/
type GetUniverseFactionsOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*The language used in the response
	 */
	ContentLanguage string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload models.GetUniverseFactionsOKBody
}

func (o *GetUniverseFactionsOK) Error() string {
	return fmt.Sprintf("[GET /universe/factions/][%d] getUniverseFactionsOK  %+v", 200, o.Payload)
}

func (o *GetUniverseFactionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Cache-Control
	o.CacheControl = response.GetHeader("Cache-Control")

	// response header Content-Language
	o.ContentLanguage = response.GetHeader("Content-Language")

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

// NewGetUniverseFactionsInternalServerError creates a GetUniverseFactionsInternalServerError with default headers values
func NewGetUniverseFactionsInternalServerError() *GetUniverseFactionsInternalServerError {
	return &GetUniverseFactionsInternalServerError{}
}

/*GetUniverseFactionsInternalServerError handles this case with default header values.

Internal server error
*/
type GetUniverseFactionsInternalServerError struct {
	Payload *models.GetUniverseFactionsInternalServerErrorBody
}

func (o *GetUniverseFactionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /universe/factions/][%d] getUniverseFactionsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUniverseFactionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetUniverseFactionsInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
