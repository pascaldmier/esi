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

// GetUniverseCategoriesReader is a Reader for the GetUniverseCategories structure.
type GetUniverseCategoriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUniverseCategoriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUniverseCategoriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetUniverseCategoriesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUniverseCategoriesOK creates a GetUniverseCategoriesOK with default headers values
func NewGetUniverseCategoriesOK() *GetUniverseCategoriesOK {
	return &GetUniverseCategoriesOK{}
}

/*GetUniverseCategoriesOK handles this case with default header values.

A list of item category ids
*/
type GetUniverseCategoriesOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload []int32
}

func (o *GetUniverseCategoriesOK) Error() string {
	return fmt.Sprintf("[GET /universe/categories/][%d] getUniverseCategoriesOK  %+v", 200, o.Payload)
}

func (o *GetUniverseCategoriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetUniverseCategoriesInternalServerError creates a GetUniverseCategoriesInternalServerError with default headers values
func NewGetUniverseCategoriesInternalServerError() *GetUniverseCategoriesInternalServerError {
	return &GetUniverseCategoriesInternalServerError{}
}

/*GetUniverseCategoriesInternalServerError handles this case with default header values.

Internal server error
*/
type GetUniverseCategoriesInternalServerError struct {
	Payload *models.GetUniverseCategoriesInternalServerErrorBody
}

func (o *GetUniverseCategoriesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /universe/categories/][%d] getUniverseCategoriesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUniverseCategoriesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetUniverseCategoriesInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
