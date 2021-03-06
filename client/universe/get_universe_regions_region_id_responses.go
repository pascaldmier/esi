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

// GetUniverseRegionsRegionIDReader is a Reader for the GetUniverseRegionsRegionID structure.
type GetUniverseRegionsRegionIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUniverseRegionsRegionIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUniverseRegionsRegionIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetUniverseRegionsRegionIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetUniverseRegionsRegionIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUniverseRegionsRegionIDOK creates a GetUniverseRegionsRegionIDOK with default headers values
func NewGetUniverseRegionsRegionIDOK() *GetUniverseRegionsRegionIDOK {
	return &GetUniverseRegionsRegionIDOK{}
}

/*GetUniverseRegionsRegionIDOK handles this case with default header values.

Information about a region
*/
type GetUniverseRegionsRegionIDOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*The language used in the response
	 */
	ContentLanguage string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload *models.GetUniverseRegionsRegionIDOKBody
}

func (o *GetUniverseRegionsRegionIDOK) Error() string {
	return fmt.Sprintf("[GET /universe/regions/{region_id}/][%d] getUniverseRegionsRegionIdOK  %+v", 200, o.Payload)
}

func (o *GetUniverseRegionsRegionIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Cache-Control
	o.CacheControl = response.GetHeader("Cache-Control")

	// response header Content-Language
	o.ContentLanguage = response.GetHeader("Content-Language")

	// response header Expires
	o.Expires = response.GetHeader("Expires")

	// response header Last-Modified
	o.LastModified = response.GetHeader("Last-Modified")

	o.Payload = new(models.GetUniverseRegionsRegionIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseRegionsRegionIDNotFound creates a GetUniverseRegionsRegionIDNotFound with default headers values
func NewGetUniverseRegionsRegionIDNotFound() *GetUniverseRegionsRegionIDNotFound {
	return &GetUniverseRegionsRegionIDNotFound{}
}

/*GetUniverseRegionsRegionIDNotFound handles this case with default header values.

Region not found
*/
type GetUniverseRegionsRegionIDNotFound struct {
	Payload *models.GetUniverseRegionsRegionIDNotFoundBody
}

func (o *GetUniverseRegionsRegionIDNotFound) Error() string {
	return fmt.Sprintf("[GET /universe/regions/{region_id}/][%d] getUniverseRegionsRegionIdNotFound  %+v", 404, o.Payload)
}

func (o *GetUniverseRegionsRegionIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetUniverseRegionsRegionIDNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseRegionsRegionIDInternalServerError creates a GetUniverseRegionsRegionIDInternalServerError with default headers values
func NewGetUniverseRegionsRegionIDInternalServerError() *GetUniverseRegionsRegionIDInternalServerError {
	return &GetUniverseRegionsRegionIDInternalServerError{}
}

/*GetUniverseRegionsRegionIDInternalServerError handles this case with default header values.

Internal server error
*/
type GetUniverseRegionsRegionIDInternalServerError struct {
	Payload *models.GetUniverseRegionsRegionIDInternalServerErrorBody
}

func (o *GetUniverseRegionsRegionIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /universe/regions/{region_id}/][%d] getUniverseRegionsRegionIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUniverseRegionsRegionIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetUniverseRegionsRegionIDInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
