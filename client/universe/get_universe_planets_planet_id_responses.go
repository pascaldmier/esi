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

// GetUniversePlanetsPlanetIDReader is a Reader for the GetUniversePlanetsPlanetID structure.
type GetUniversePlanetsPlanetIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUniversePlanetsPlanetIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUniversePlanetsPlanetIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetUniversePlanetsPlanetIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetUniversePlanetsPlanetIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUniversePlanetsPlanetIDOK creates a GetUniversePlanetsPlanetIDOK with default headers values
func NewGetUniversePlanetsPlanetIDOK() *GetUniversePlanetsPlanetIDOK {
	return &GetUniversePlanetsPlanetIDOK{}
}

/*GetUniversePlanetsPlanetIDOK handles this case with default header values.

Information about a planet
*/
type GetUniversePlanetsPlanetIDOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload *models.GetUniversePlanetsPlanetIDOKBody
}

func (o *GetUniversePlanetsPlanetIDOK) Error() string {
	return fmt.Sprintf("[GET /universe/planets/{planet_id}/][%d] getUniversePlanetsPlanetIdOK  %+v", 200, o.Payload)
}

func (o *GetUniversePlanetsPlanetIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Cache-Control
	o.CacheControl = response.GetHeader("Cache-Control")

	// response header Expires
	o.Expires = response.GetHeader("Expires")

	// response header Last-Modified
	o.LastModified = response.GetHeader("Last-Modified")

	o.Payload = new(models.GetUniversePlanetsPlanetIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniversePlanetsPlanetIDNotFound creates a GetUniversePlanetsPlanetIDNotFound with default headers values
func NewGetUniversePlanetsPlanetIDNotFound() *GetUniversePlanetsPlanetIDNotFound {
	return &GetUniversePlanetsPlanetIDNotFound{}
}

/*GetUniversePlanetsPlanetIDNotFound handles this case with default header values.

Planet not found
*/
type GetUniversePlanetsPlanetIDNotFound struct {
	Payload *models.GetUniversePlanetsPlanetIDNotFoundBody
}

func (o *GetUniversePlanetsPlanetIDNotFound) Error() string {
	return fmt.Sprintf("[GET /universe/planets/{planet_id}/][%d] getUniversePlanetsPlanetIdNotFound  %+v", 404, o.Payload)
}

func (o *GetUniversePlanetsPlanetIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetUniversePlanetsPlanetIDNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniversePlanetsPlanetIDInternalServerError creates a GetUniversePlanetsPlanetIDInternalServerError with default headers values
func NewGetUniversePlanetsPlanetIDInternalServerError() *GetUniversePlanetsPlanetIDInternalServerError {
	return &GetUniversePlanetsPlanetIDInternalServerError{}
}

/*GetUniversePlanetsPlanetIDInternalServerError handles this case with default header values.

Internal server error
*/
type GetUniversePlanetsPlanetIDInternalServerError struct {
	Payload *models.GetUniversePlanetsPlanetIDInternalServerErrorBody
}

func (o *GetUniversePlanetsPlanetIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /universe/planets/{planet_id}/][%d] getUniversePlanetsPlanetIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUniversePlanetsPlanetIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetUniversePlanetsPlanetIDInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}