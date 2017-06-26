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

// PostCharactersCharacterIDMailLabelsReader is a Reader for the PostCharactersCharacterIDMailLabels structure.
type PostCharactersCharacterIDMailLabelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCharactersCharacterIDMailLabelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostCharactersCharacterIDMailLabelsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewPostCharactersCharacterIDMailLabelsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPostCharactersCharacterIDMailLabelsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostCharactersCharacterIDMailLabelsCreated creates a PostCharactersCharacterIDMailLabelsCreated with default headers values
func NewPostCharactersCharacterIDMailLabelsCreated() *PostCharactersCharacterIDMailLabelsCreated {
	return &PostCharactersCharacterIDMailLabelsCreated{}
}

/*PostCharactersCharacterIDMailLabelsCreated handles this case with default header values.

Label created
*/
type PostCharactersCharacterIDMailLabelsCreated struct {
	Payload int64
}

func (o *PostCharactersCharacterIDMailLabelsCreated) Error() string {
	return fmt.Sprintf("[POST /characters/{character_id}/mail/labels/][%d] postCharactersCharacterIdMailLabelsCreated  %+v", 201, o.Payload)
}

func (o *PostCharactersCharacterIDMailLabelsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCharactersCharacterIDMailLabelsForbidden creates a PostCharactersCharacterIDMailLabelsForbidden with default headers values
func NewPostCharactersCharacterIDMailLabelsForbidden() *PostCharactersCharacterIDMailLabelsForbidden {
	return &PostCharactersCharacterIDMailLabelsForbidden{}
}

/*PostCharactersCharacterIDMailLabelsForbidden handles this case with default header values.

Forbidden
*/
type PostCharactersCharacterIDMailLabelsForbidden struct {
	Payload *models.PostCharactersCharacterIDMailLabelsForbiddenBody
}

func (o *PostCharactersCharacterIDMailLabelsForbidden) Error() string {
	return fmt.Sprintf("[POST /characters/{character_id}/mail/labels/][%d] postCharactersCharacterIdMailLabelsForbidden  %+v", 403, o.Payload)
}

func (o *PostCharactersCharacterIDMailLabelsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostCharactersCharacterIDMailLabelsForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCharactersCharacterIDMailLabelsInternalServerError creates a PostCharactersCharacterIDMailLabelsInternalServerError with default headers values
func NewPostCharactersCharacterIDMailLabelsInternalServerError() *PostCharactersCharacterIDMailLabelsInternalServerError {
	return &PostCharactersCharacterIDMailLabelsInternalServerError{}
}

/*PostCharactersCharacterIDMailLabelsInternalServerError handles this case with default header values.

Internal server error
*/
type PostCharactersCharacterIDMailLabelsInternalServerError struct {
	Payload *models.PostCharactersCharacterIDMailLabelsInternalServerErrorBody
}

func (o *PostCharactersCharacterIDMailLabelsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /characters/{character_id}/mail/labels/][%d] postCharactersCharacterIdMailLabelsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostCharactersCharacterIDMailLabelsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostCharactersCharacterIDMailLabelsInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
