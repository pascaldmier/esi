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

// PutFleetsFleetIDReader is a Reader for the PutFleetsFleetID structure.
type PutFleetsFleetIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutFleetsFleetIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewPutFleetsFleetIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPutFleetsFleetIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPutFleetsFleetIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPutFleetsFleetIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPutFleetsFleetIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutFleetsFleetIDNoContent creates a PutFleetsFleetIDNoContent with default headers values
func NewPutFleetsFleetIDNoContent() *PutFleetsFleetIDNoContent {
	return &PutFleetsFleetIDNoContent{}
}

/*PutFleetsFleetIDNoContent handles this case with default header values.

Fleet updated
*/
type PutFleetsFleetIDNoContent struct {
}

func (o *PutFleetsFleetIDNoContent) Error() string {
	return fmt.Sprintf("[PUT /fleets/{fleet_id}/][%d] putFleetsFleetIdNoContent ", 204)
}

func (o *PutFleetsFleetIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutFleetsFleetIDBadRequest creates a PutFleetsFleetIDBadRequest with default headers values
func NewPutFleetsFleetIDBadRequest() *PutFleetsFleetIDBadRequest {
	return &PutFleetsFleetIDBadRequest{}
}

/*PutFleetsFleetIDBadRequest handles this case with default header values.

Invalid request body
*/
type PutFleetsFleetIDBadRequest struct {
	Payload *models.PutFleetsFleetIDBadRequestBody
}

func (o *PutFleetsFleetIDBadRequest) Error() string {
	return fmt.Sprintf("[PUT /fleets/{fleet_id}/][%d] putFleetsFleetIdBadRequest  %+v", 400, o.Payload)
}

func (o *PutFleetsFleetIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PutFleetsFleetIDBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFleetsFleetIDForbidden creates a PutFleetsFleetIDForbidden with default headers values
func NewPutFleetsFleetIDForbidden() *PutFleetsFleetIDForbidden {
	return &PutFleetsFleetIDForbidden{}
}

/*PutFleetsFleetIDForbidden handles this case with default header values.

Forbidden
*/
type PutFleetsFleetIDForbidden struct {
	Payload *models.PutFleetsFleetIDForbiddenBody
}

func (o *PutFleetsFleetIDForbidden) Error() string {
	return fmt.Sprintf("[PUT /fleets/{fleet_id}/][%d] putFleetsFleetIdForbidden  %+v", 403, o.Payload)
}

func (o *PutFleetsFleetIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PutFleetsFleetIDForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFleetsFleetIDNotFound creates a PutFleetsFleetIDNotFound with default headers values
func NewPutFleetsFleetIDNotFound() *PutFleetsFleetIDNotFound {
	return &PutFleetsFleetIDNotFound{}
}

/*PutFleetsFleetIDNotFound handles this case with default header values.

The fleet does not exist or you don't have access to it
*/
type PutFleetsFleetIDNotFound struct {
	Payload *models.PutFleetsFleetIDNotFoundBody
}

func (o *PutFleetsFleetIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /fleets/{fleet_id}/][%d] putFleetsFleetIdNotFound  %+v", 404, o.Payload)
}

func (o *PutFleetsFleetIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PutFleetsFleetIDNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFleetsFleetIDInternalServerError creates a PutFleetsFleetIDInternalServerError with default headers values
func NewPutFleetsFleetIDInternalServerError() *PutFleetsFleetIDInternalServerError {
	return &PutFleetsFleetIDInternalServerError{}
}

/*PutFleetsFleetIDInternalServerError handles this case with default header values.

Internal server error
*/
type PutFleetsFleetIDInternalServerError struct {
	Payload *models.PutFleetsFleetIDInternalServerErrorBody
}

func (o *PutFleetsFleetIDInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /fleets/{fleet_id}/][%d] putFleetsFleetIdInternalServerError  %+v", 500, o.Payload)
}

func (o *PutFleetsFleetIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PutFleetsFleetIDInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
