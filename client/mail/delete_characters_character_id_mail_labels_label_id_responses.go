package mail

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/pascaldmier/esi/models"
)

// DeleteCharactersCharacterIDMailLabelsLabelIDReader is a Reader for the DeleteCharactersCharacterIDMailLabelsLabelID structure.
type DeleteCharactersCharacterIDMailLabelsLabelIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCharactersCharacterIDMailLabelsLabelIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteCharactersCharacterIDMailLabelsLabelIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewDeleteCharactersCharacterIDMailLabelsLabelIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewDeleteCharactersCharacterIDMailLabelsLabelIDUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteCharactersCharacterIDMailLabelsLabelIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteCharactersCharacterIDMailLabelsLabelIDNoContent creates a DeleteCharactersCharacterIDMailLabelsLabelIDNoContent with default headers values
func NewDeleteCharactersCharacterIDMailLabelsLabelIDNoContent() *DeleteCharactersCharacterIDMailLabelsLabelIDNoContent {
	return &DeleteCharactersCharacterIDMailLabelsLabelIDNoContent{}
}

/*DeleteCharactersCharacterIDMailLabelsLabelIDNoContent handles this case with default header values.

Label deleted
*/
type DeleteCharactersCharacterIDMailLabelsLabelIDNoContent struct {
}

func (o *DeleteCharactersCharacterIDMailLabelsLabelIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /characters/{character_id}/mail/labels/{label_id}/][%d] deleteCharactersCharacterIdMailLabelsLabelIdNoContent ", 204)
}

func (o *DeleteCharactersCharacterIDMailLabelsLabelIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCharactersCharacterIDMailLabelsLabelIDForbidden creates a DeleteCharactersCharacterIDMailLabelsLabelIDForbidden with default headers values
func NewDeleteCharactersCharacterIDMailLabelsLabelIDForbidden() *DeleteCharactersCharacterIDMailLabelsLabelIDForbidden {
	return &DeleteCharactersCharacterIDMailLabelsLabelIDForbidden{}
}

/*DeleteCharactersCharacterIDMailLabelsLabelIDForbidden handles this case with default header values.

Forbidden
*/
type DeleteCharactersCharacterIDMailLabelsLabelIDForbidden struct {
	Payload *models.DeleteCharactersCharacterIDMailLabelsLabelIDForbiddenBody
}

func (o *DeleteCharactersCharacterIDMailLabelsLabelIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /characters/{character_id}/mail/labels/{label_id}/][%d] deleteCharactersCharacterIdMailLabelsLabelIdForbidden  %+v", 403, o.Payload)
}

func (o *DeleteCharactersCharacterIDMailLabelsLabelIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteCharactersCharacterIDMailLabelsLabelIDForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCharactersCharacterIDMailLabelsLabelIDUnprocessableEntity creates a DeleteCharactersCharacterIDMailLabelsLabelIDUnprocessableEntity with default headers values
func NewDeleteCharactersCharacterIDMailLabelsLabelIDUnprocessableEntity() *DeleteCharactersCharacterIDMailLabelsLabelIDUnprocessableEntity {
	return &DeleteCharactersCharacterIDMailLabelsLabelIDUnprocessableEntity{}
}

/*DeleteCharactersCharacterIDMailLabelsLabelIDUnprocessableEntity handles this case with default header values.

Default labels cannot be deleted
*/
type DeleteCharactersCharacterIDMailLabelsLabelIDUnprocessableEntity struct {
	Payload *models.DeleteCharactersCharacterIDMailLabelsLabelIDUnprocessableEntityBody
}

func (o *DeleteCharactersCharacterIDMailLabelsLabelIDUnprocessableEntity) Error() string {
	return fmt.Sprintf("[DELETE /characters/{character_id}/mail/labels/{label_id}/][%d] deleteCharactersCharacterIdMailLabelsLabelIdUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *DeleteCharactersCharacterIDMailLabelsLabelIDUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteCharactersCharacterIDMailLabelsLabelIDUnprocessableEntityBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCharactersCharacterIDMailLabelsLabelIDInternalServerError creates a DeleteCharactersCharacterIDMailLabelsLabelIDInternalServerError with default headers values
func NewDeleteCharactersCharacterIDMailLabelsLabelIDInternalServerError() *DeleteCharactersCharacterIDMailLabelsLabelIDInternalServerError {
	return &DeleteCharactersCharacterIDMailLabelsLabelIDInternalServerError{}
}

/*DeleteCharactersCharacterIDMailLabelsLabelIDInternalServerError handles this case with default header values.

Internal server error
*/
type DeleteCharactersCharacterIDMailLabelsLabelIDInternalServerError struct {
	Payload *models.DeleteCharactersCharacterIDMailLabelsLabelIDInternalServerErrorBody
}

func (o *DeleteCharactersCharacterIDMailLabelsLabelIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /characters/{character_id}/mail/labels/{label_id}/][%d] deleteCharactersCharacterIdMailLabelsLabelIdInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteCharactersCharacterIDMailLabelsLabelIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteCharactersCharacterIDMailLabelsLabelIDInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
