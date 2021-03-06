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

// PostFleetsFleetIDMembersReader is a Reader for the PostFleetsFleetIDMembers structure.
type PostFleetsFleetIDMembersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostFleetsFleetIDMembersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewPostFleetsFleetIDMembersNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewPostFleetsFleetIDMembersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPostFleetsFleetIDMembersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewPostFleetsFleetIDMembersUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPostFleetsFleetIDMembersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostFleetsFleetIDMembersNoContent creates a PostFleetsFleetIDMembersNoContent with default headers values
func NewPostFleetsFleetIDMembersNoContent() *PostFleetsFleetIDMembersNoContent {
	return &PostFleetsFleetIDMembersNoContent{}
}

/*PostFleetsFleetIDMembersNoContent handles this case with default header values.

Fleet invitation sent
*/
type PostFleetsFleetIDMembersNoContent struct {
}

func (o *PostFleetsFleetIDMembersNoContent) Error() string {
	return fmt.Sprintf("[POST /fleets/{fleet_id}/members/][%d] postFleetsFleetIdMembersNoContent ", 204)
}

func (o *PostFleetsFleetIDMembersNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostFleetsFleetIDMembersForbidden creates a PostFleetsFleetIDMembersForbidden with default headers values
func NewPostFleetsFleetIDMembersForbidden() *PostFleetsFleetIDMembersForbidden {
	return &PostFleetsFleetIDMembersForbidden{}
}

/*PostFleetsFleetIDMembersForbidden handles this case with default header values.

Forbidden
*/
type PostFleetsFleetIDMembersForbidden struct {
	Payload *models.PostFleetsFleetIDMembersForbiddenBody
}

func (o *PostFleetsFleetIDMembersForbidden) Error() string {
	return fmt.Sprintf("[POST /fleets/{fleet_id}/members/][%d] postFleetsFleetIdMembersForbidden  %+v", 403, o.Payload)
}

func (o *PostFleetsFleetIDMembersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostFleetsFleetIDMembersForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFleetsFleetIDMembersNotFound creates a PostFleetsFleetIDMembersNotFound with default headers values
func NewPostFleetsFleetIDMembersNotFound() *PostFleetsFleetIDMembersNotFound {
	return &PostFleetsFleetIDMembersNotFound{}
}

/*PostFleetsFleetIDMembersNotFound handles this case with default header values.

The fleet does not exist or you don't have access to it
*/
type PostFleetsFleetIDMembersNotFound struct {
	Payload *models.PostFleetsFleetIDMembersNotFoundBody
}

func (o *PostFleetsFleetIDMembersNotFound) Error() string {
	return fmt.Sprintf("[POST /fleets/{fleet_id}/members/][%d] postFleetsFleetIdMembersNotFound  %+v", 404, o.Payload)
}

func (o *PostFleetsFleetIDMembersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostFleetsFleetIDMembersNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFleetsFleetIDMembersUnprocessableEntity creates a PostFleetsFleetIDMembersUnprocessableEntity with default headers values
func NewPostFleetsFleetIDMembersUnprocessableEntity() *PostFleetsFleetIDMembersUnprocessableEntity {
	return &PostFleetsFleetIDMembersUnprocessableEntity{}
}

/*PostFleetsFleetIDMembersUnprocessableEntity handles this case with default header values.

Errors in invitation
*/
type PostFleetsFleetIDMembersUnprocessableEntity struct {
	Payload *models.PostFleetsFleetIDMembersUnprocessableEntityBody
}

func (o *PostFleetsFleetIDMembersUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /fleets/{fleet_id}/members/][%d] postFleetsFleetIdMembersUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PostFleetsFleetIDMembersUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostFleetsFleetIDMembersUnprocessableEntityBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFleetsFleetIDMembersInternalServerError creates a PostFleetsFleetIDMembersInternalServerError with default headers values
func NewPostFleetsFleetIDMembersInternalServerError() *PostFleetsFleetIDMembersInternalServerError {
	return &PostFleetsFleetIDMembersInternalServerError{}
}

/*PostFleetsFleetIDMembersInternalServerError handles this case with default header values.

Internal server error
*/
type PostFleetsFleetIDMembersInternalServerError struct {
	Payload *models.PostFleetsFleetIDMembersInternalServerErrorBody
}

func (o *PostFleetsFleetIDMembersInternalServerError) Error() string {
	return fmt.Sprintf("[POST /fleets/{fleet_id}/members/][%d] postFleetsFleetIdMembersInternalServerError  %+v", 500, o.Payload)
}

func (o *PostFleetsFleetIDMembersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostFleetsFleetIDMembersInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
