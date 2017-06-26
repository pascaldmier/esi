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

// GetUniverseRegionsReader is a Reader for the GetUniverseRegions structure.
type GetUniverseRegionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUniverseRegionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUniverseRegionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetUniverseRegionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUniverseRegionsOK creates a GetUniverseRegionsOK with default headers values
func NewGetUniverseRegionsOK() *GetUniverseRegionsOK {
	return &GetUniverseRegionsOK{}
}

/*GetUniverseRegionsOK handles this case with default header values.

A list of region ids
*/
type GetUniverseRegionsOK struct {
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

func (o *GetUniverseRegionsOK) Error() string {
	return fmt.Sprintf("[GET /universe/regions/][%d] getUniverseRegionsOK  %+v", 200, o.Payload)
}

func (o *GetUniverseRegionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetUniverseRegionsInternalServerError creates a GetUniverseRegionsInternalServerError with default headers values
func NewGetUniverseRegionsInternalServerError() *GetUniverseRegionsInternalServerError {
	return &GetUniverseRegionsInternalServerError{}
}

/*GetUniverseRegionsInternalServerError handles this case with default header values.

Internal server error
*/
type GetUniverseRegionsInternalServerError struct {
	Payload *models.GetUniverseRegionsInternalServerErrorBody
}

func (o *GetUniverseRegionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /universe/regions/][%d] getUniverseRegionsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUniverseRegionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetUniverseRegionsInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}