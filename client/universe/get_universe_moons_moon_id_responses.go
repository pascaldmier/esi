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

// GetUniverseMoonsMoonIDReader is a Reader for the GetUniverseMoonsMoonID structure.
type GetUniverseMoonsMoonIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUniverseMoonsMoonIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUniverseMoonsMoonIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetUniverseMoonsMoonIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetUniverseMoonsMoonIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUniverseMoonsMoonIDOK creates a GetUniverseMoonsMoonIDOK with default headers values
func NewGetUniverseMoonsMoonIDOK() *GetUniverseMoonsMoonIDOK {
	return &GetUniverseMoonsMoonIDOK{}
}

/*GetUniverseMoonsMoonIDOK handles this case with default header values.

Information about a moon
*/
type GetUniverseMoonsMoonIDOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload *models.GetUniverseMoonsMoonIDOKBody
}

func (o *GetUniverseMoonsMoonIDOK) Error() string {
	return fmt.Sprintf("[GET /universe/moons/{moon_id}/][%d] getUniverseMoonsMoonIdOK  %+v", 200, o.Payload)
}

func (o *GetUniverseMoonsMoonIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Cache-Control
	o.CacheControl = response.GetHeader("Cache-Control")

	// response header Expires
	o.Expires = response.GetHeader("Expires")

	// response header Last-Modified
	o.LastModified = response.GetHeader("Last-Modified")

	o.Payload = new(models.GetUniverseMoonsMoonIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseMoonsMoonIDNotFound creates a GetUniverseMoonsMoonIDNotFound with default headers values
func NewGetUniverseMoonsMoonIDNotFound() *GetUniverseMoonsMoonIDNotFound {
	return &GetUniverseMoonsMoonIDNotFound{}
}

/*GetUniverseMoonsMoonIDNotFound handles this case with default header values.

Moon not found
*/
type GetUniverseMoonsMoonIDNotFound struct {
	Payload *models.GetUniverseMoonsMoonIDNotFoundBody
}

func (o *GetUniverseMoonsMoonIDNotFound) Error() string {
	return fmt.Sprintf("[GET /universe/moons/{moon_id}/][%d] getUniverseMoonsMoonIdNotFound  %+v", 404, o.Payload)
}

func (o *GetUniverseMoonsMoonIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetUniverseMoonsMoonIDNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseMoonsMoonIDInternalServerError creates a GetUniverseMoonsMoonIDInternalServerError with default headers values
func NewGetUniverseMoonsMoonIDInternalServerError() *GetUniverseMoonsMoonIDInternalServerError {
	return &GetUniverseMoonsMoonIDInternalServerError{}
}

/*GetUniverseMoonsMoonIDInternalServerError handles this case with default header values.

Internal server error
*/
type GetUniverseMoonsMoonIDInternalServerError struct {
	Payload *models.GetUniverseMoonsMoonIDInternalServerErrorBody
}

func (o *GetUniverseMoonsMoonIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /universe/moons/{moon_id}/][%d] getUniverseMoonsMoonIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUniverseMoonsMoonIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetUniverseMoonsMoonIDInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}