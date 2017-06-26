package industry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/pascaldmier/esi/models"
)

// GetIndustrySystemsReader is a Reader for the GetIndustrySystems structure.
type GetIndustrySystemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIndustrySystemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetIndustrySystemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetIndustrySystemsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetIndustrySystemsOK creates a GetIndustrySystemsOK with default headers values
func NewGetIndustrySystemsOK() *GetIndustrySystemsOK {
	return &GetIndustrySystemsOK{}
}

/*GetIndustrySystemsOK handles this case with default header values.

A list of cost indicies
*/
type GetIndustrySystemsOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload models.GetIndustrySystemsOKBody
}

func (o *GetIndustrySystemsOK) Error() string {
	return fmt.Sprintf("[GET /industry/systems/][%d] getIndustrySystemsOK  %+v", 200, o.Payload)
}

func (o *GetIndustrySystemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetIndustrySystemsInternalServerError creates a GetIndustrySystemsInternalServerError with default headers values
func NewGetIndustrySystemsInternalServerError() *GetIndustrySystemsInternalServerError {
	return &GetIndustrySystemsInternalServerError{}
}

/*GetIndustrySystemsInternalServerError handles this case with default header values.

Internal server error
*/
type GetIndustrySystemsInternalServerError struct {
	Payload *models.GetIndustrySystemsInternalServerErrorBody
}

func (o *GetIndustrySystemsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /industry/systems/][%d] getIndustrySystemsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIndustrySystemsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetIndustrySystemsInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
