package market

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/pascaldmier/esi/models"
)

// GetMarketsRegionIDHistoryReader is a Reader for the GetMarketsRegionIDHistory structure.
type GetMarketsRegionIDHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMarketsRegionIDHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMarketsRegionIDHistoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 422:
		result := NewGetMarketsRegionIDHistoryUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetMarketsRegionIDHistoryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetMarketsRegionIDHistoryOK creates a GetMarketsRegionIDHistoryOK with default headers values
func NewGetMarketsRegionIDHistoryOK() *GetMarketsRegionIDHistoryOK {
	return &GetMarketsRegionIDHistoryOK{}
}

/*GetMarketsRegionIDHistoryOK handles this case with default header values.

A list of historical market statistics
*/
type GetMarketsRegionIDHistoryOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload models.GetMarketsRegionIDHistoryOKBody
}

func (o *GetMarketsRegionIDHistoryOK) Error() string {
	return fmt.Sprintf("[GET /markets/{region_id}/history/][%d] getMarketsRegionIdHistoryOK  %+v", 200, o.Payload)
}

func (o *GetMarketsRegionIDHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetMarketsRegionIDHistoryUnprocessableEntity creates a GetMarketsRegionIDHistoryUnprocessableEntity with default headers values
func NewGetMarketsRegionIDHistoryUnprocessableEntity() *GetMarketsRegionIDHistoryUnprocessableEntity {
	return &GetMarketsRegionIDHistoryUnprocessableEntity{}
}

/*GetMarketsRegionIDHistoryUnprocessableEntity handles this case with default header values.

Not found
*/
type GetMarketsRegionIDHistoryUnprocessableEntity struct {
	Payload *models.GetMarketsRegionIDHistoryUnprocessableEntityBody
}

func (o *GetMarketsRegionIDHistoryUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /markets/{region_id}/history/][%d] getMarketsRegionIdHistoryUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GetMarketsRegionIDHistoryUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetMarketsRegionIDHistoryUnprocessableEntityBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMarketsRegionIDHistoryInternalServerError creates a GetMarketsRegionIDHistoryInternalServerError with default headers values
func NewGetMarketsRegionIDHistoryInternalServerError() *GetMarketsRegionIDHistoryInternalServerError {
	return &GetMarketsRegionIDHistoryInternalServerError{}
}

/*GetMarketsRegionIDHistoryInternalServerError handles this case with default header values.

Internal server error
*/
type GetMarketsRegionIDHistoryInternalServerError struct {
	Payload *models.GetMarketsRegionIDHistoryInternalServerErrorBody
}

func (o *GetMarketsRegionIDHistoryInternalServerError) Error() string {
	return fmt.Sprintf("[GET /markets/{region_id}/history/][%d] getMarketsRegionIdHistoryInternalServerError  %+v", 500, o.Payload)
}

func (o *GetMarketsRegionIDHistoryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetMarketsRegionIDHistoryInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
