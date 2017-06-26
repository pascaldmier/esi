package fleets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/pascaldmier/esi/models"
)

// GetFleetsFleetIDReader is a Reader for the GetFleetsFleetID structure.
type GetFleetsFleetIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFleetsFleetIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetFleetsFleetIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetFleetsFleetIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetFleetsFleetIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetFleetsFleetIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFleetsFleetIDOK creates a GetFleetsFleetIDOK with default headers values
func NewGetFleetsFleetIDOK() *GetFleetsFleetIDOK {
	return &GetFleetsFleetIDOK{}
}

/*GetFleetsFleetIDOK handles this case with default header values.

Details about a fleet
*/
type GetFleetsFleetIDOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload *models.GetFleetsFleetIDOKBody
}

func (o *GetFleetsFleetIDOK) Error() string {
	return fmt.Sprintf("[GET /fleets/{fleet_id}/][%d] getFleetsFleetIdOK  %+v", 200, o.Payload)
}

func (o *GetFleetsFleetIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Cache-Control
	o.CacheControl = response.GetHeader("Cache-Control")

	// response header Expires
	o.Expires = response.GetHeader("Expires")

	// response header Last-Modified
	o.LastModified = response.GetHeader("Last-Modified")

	o.Payload = new(models.GetFleetsFleetIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFleetsFleetIDForbidden creates a GetFleetsFleetIDForbidden with default headers values
func NewGetFleetsFleetIDForbidden() *GetFleetsFleetIDForbidden {
	return &GetFleetsFleetIDForbidden{}
}

/*GetFleetsFleetIDForbidden handles this case with default header values.

Forbidden
*/
type GetFleetsFleetIDForbidden struct {
	Payload *models.GetFleetsFleetIDForbiddenBody
}

func (o *GetFleetsFleetIDForbidden) Error() string {
	return fmt.Sprintf("[GET /fleets/{fleet_id}/][%d] getFleetsFleetIdForbidden  %+v", 403, o.Payload)
}

func (o *GetFleetsFleetIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetFleetsFleetIDForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFleetsFleetIDNotFound creates a GetFleetsFleetIDNotFound with default headers values
func NewGetFleetsFleetIDNotFound() *GetFleetsFleetIDNotFound {
	return &GetFleetsFleetIDNotFound{}
}

/*GetFleetsFleetIDNotFound handles this case with default header values.

The fleet does not exist or you don't have access to it
*/
type GetFleetsFleetIDNotFound struct {
	Payload *models.GetFleetsFleetIDNotFoundBody
}

func (o *GetFleetsFleetIDNotFound) Error() string {
	return fmt.Sprintf("[GET /fleets/{fleet_id}/][%d] getFleetsFleetIdNotFound  %+v", 404, o.Payload)
}

func (o *GetFleetsFleetIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetFleetsFleetIDNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFleetsFleetIDInternalServerError creates a GetFleetsFleetIDInternalServerError with default headers values
func NewGetFleetsFleetIDInternalServerError() *GetFleetsFleetIDInternalServerError {
	return &GetFleetsFleetIDInternalServerError{}
}

/*GetFleetsFleetIDInternalServerError handles this case with default header values.

Internal server error
*/
type GetFleetsFleetIDInternalServerError struct {
	Payload *models.GetFleetsFleetIDInternalServerErrorBody
}

func (o *GetFleetsFleetIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /fleets/{fleet_id}/][%d] getFleetsFleetIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetFleetsFleetIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetFleetsFleetIDInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}