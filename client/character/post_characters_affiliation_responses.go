package character

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/pascaldmier/esi/models"
)

// PostCharactersAffiliationReader is a Reader for the PostCharactersAffiliation structure.
type PostCharactersAffiliationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCharactersAffiliationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostCharactersAffiliationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 422:
		result := NewPostCharactersAffiliationUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPostCharactersAffiliationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostCharactersAffiliationOK creates a PostCharactersAffiliationOK with default headers values
func NewPostCharactersAffiliationOK() *PostCharactersAffiliationOK {
	return &PostCharactersAffiliationOK{}
}

/*PostCharactersAffiliationOK handles this case with default header values.

Character corporation, alliance and faction IDs
*/
type PostCharactersAffiliationOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload models.PostCharactersAffiliationOKBody
}

func (o *PostCharactersAffiliationOK) Error() string {
	return fmt.Sprintf("[POST /characters/affiliation/][%d] postCharactersAffiliationOK  %+v", 200, o.Payload)
}

func (o *PostCharactersAffiliationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewPostCharactersAffiliationUnprocessableEntity creates a PostCharactersAffiliationUnprocessableEntity with default headers values
func NewPostCharactersAffiliationUnprocessableEntity() *PostCharactersAffiliationUnprocessableEntity {
	return &PostCharactersAffiliationUnprocessableEntity{}
}

/*PostCharactersAffiliationUnprocessableEntity handles this case with default header values.

No valid character IDs found
*/
type PostCharactersAffiliationUnprocessableEntity struct {
	Payload *models.PostCharactersAffiliationUnprocessableEntityBody
}

func (o *PostCharactersAffiliationUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /characters/affiliation/][%d] postCharactersAffiliationUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PostCharactersAffiliationUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostCharactersAffiliationUnprocessableEntityBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCharactersAffiliationInternalServerError creates a PostCharactersAffiliationInternalServerError with default headers values
func NewPostCharactersAffiliationInternalServerError() *PostCharactersAffiliationInternalServerError {
	return &PostCharactersAffiliationInternalServerError{}
}

/*PostCharactersAffiliationInternalServerError handles this case with default header values.

Internal server error
*/
type PostCharactersAffiliationInternalServerError struct {
	Payload *models.PostCharactersAffiliationInternalServerErrorBody
}

func (o *PostCharactersAffiliationInternalServerError) Error() string {
	return fmt.Sprintf("[POST /characters/affiliation/][%d] postCharactersAffiliationInternalServerError  %+v", 500, o.Payload)
}

func (o *PostCharactersAffiliationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostCharactersAffiliationInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}