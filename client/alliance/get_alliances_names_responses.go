package alliance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/pascaldmier/esi/models"
)

// GetAlliancesNamesReader is a Reader for the GetAlliancesNames structure.
type GetAlliancesNamesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAlliancesNamesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAlliancesNamesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetAlliancesNamesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAlliancesNamesOK creates a GetAlliancesNamesOK with default headers values
func NewGetAlliancesNamesOK() *GetAlliancesNamesOK {
	return &GetAlliancesNamesOK{}
}

/*GetAlliancesNamesOK handles this case with default header values.

List of id/name associations
*/
type GetAlliancesNamesOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload models.GetAlliancesNamesOKBody
}

func (o *GetAlliancesNamesOK) Error() string {
	return fmt.Sprintf("[GET /alliances/names/][%d] getAlliancesNamesOK  %+v", 200, o.Payload)
}

func (o *GetAlliancesNamesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAlliancesNamesInternalServerError creates a GetAlliancesNamesInternalServerError with default headers values
func NewGetAlliancesNamesInternalServerError() *GetAlliancesNamesInternalServerError {
	return &GetAlliancesNamesInternalServerError{}
}

/*GetAlliancesNamesInternalServerError handles this case with default header values.

Internal server error
*/
type GetAlliancesNamesInternalServerError struct {
	Payload *models.GetAlliancesNamesInternalServerErrorBody
}

func (o *GetAlliancesNamesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /alliances/names/][%d] getAlliancesNamesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAlliancesNamesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetAlliancesNamesInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
