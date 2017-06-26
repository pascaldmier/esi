package sovereignty

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/pascaldmier/esi/models"
)

// GetSovereigntyStructuresReader is a Reader for the GetSovereigntyStructures structure.
type GetSovereigntyStructuresReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSovereigntyStructuresReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSovereigntyStructuresOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetSovereigntyStructuresInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSovereigntyStructuresOK creates a GetSovereigntyStructuresOK with default headers values
func NewGetSovereigntyStructuresOK() *GetSovereigntyStructuresOK {
	return &GetSovereigntyStructuresOK{}
}

/*GetSovereigntyStructuresOK handles this case with default header values.

A list of sovereignty structures
*/
type GetSovereigntyStructuresOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload models.GetSovereigntyStructuresOKBody
}

func (o *GetSovereigntyStructuresOK) Error() string {
	return fmt.Sprintf("[GET /sovereignty/structures/][%d] getSovereigntyStructuresOK  %+v", 200, o.Payload)
}

func (o *GetSovereigntyStructuresOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSovereigntyStructuresInternalServerError creates a GetSovereigntyStructuresInternalServerError with default headers values
func NewGetSovereigntyStructuresInternalServerError() *GetSovereigntyStructuresInternalServerError {
	return &GetSovereigntyStructuresInternalServerError{}
}

/*GetSovereigntyStructuresInternalServerError handles this case with default header values.

Internal server error
*/
type GetSovereigntyStructuresInternalServerError struct {
	Payload *models.GetSovereigntyStructuresInternalServerErrorBody
}

func (o *GetSovereigntyStructuresInternalServerError) Error() string {
	return fmt.Sprintf("[GET /sovereignty/structures/][%d] getSovereigntyStructuresInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSovereigntyStructuresInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetSovereigntyStructuresInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}