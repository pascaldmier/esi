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

// GetUniverseTypesReader is a Reader for the GetUniverseTypes structure.
type GetUniverseTypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUniverseTypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUniverseTypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetUniverseTypesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUniverseTypesOK creates a GetUniverseTypesOK with default headers values
func NewGetUniverseTypesOK() *GetUniverseTypesOK {
	return &GetUniverseTypesOK{}
}

/*GetUniverseTypesOK handles this case with default header values.

A list of type ids
*/
type GetUniverseTypesOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload []int32
}

func (o *GetUniverseTypesOK) Error() string {
	return fmt.Sprintf("[GET /universe/types/][%d] getUniverseTypesOK  %+v", 200, o.Payload)
}

func (o *GetUniverseTypesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetUniverseTypesInternalServerError creates a GetUniverseTypesInternalServerError with default headers values
func NewGetUniverseTypesInternalServerError() *GetUniverseTypesInternalServerError {
	return &GetUniverseTypesInternalServerError{}
}

/*GetUniverseTypesInternalServerError handles this case with default header values.

Internal server error
*/
type GetUniverseTypesInternalServerError struct {
	Payload *models.GetUniverseTypesInternalServerErrorBody
}

func (o *GetUniverseTypesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /universe/types/][%d] getUniverseTypesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUniverseTypesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetUniverseTypesInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
