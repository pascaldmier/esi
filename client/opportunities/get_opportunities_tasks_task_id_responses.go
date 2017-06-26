package opportunities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/pascaldmier/esi/models"
)

// GetOpportunitiesTasksTaskIDReader is a Reader for the GetOpportunitiesTasksTaskID structure.
type GetOpportunitiesTasksTaskIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOpportunitiesTasksTaskIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetOpportunitiesTasksTaskIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetOpportunitiesTasksTaskIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetOpportunitiesTasksTaskIDOK creates a GetOpportunitiesTasksTaskIDOK with default headers values
func NewGetOpportunitiesTasksTaskIDOK() *GetOpportunitiesTasksTaskIDOK {
	return &GetOpportunitiesTasksTaskIDOK{}
}

/*GetOpportunitiesTasksTaskIDOK handles this case with default header values.

Details of an opportunities task
*/
type GetOpportunitiesTasksTaskIDOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload *models.GetOpportunitiesTasksTaskIDOKBody
}

func (o *GetOpportunitiesTasksTaskIDOK) Error() string {
	return fmt.Sprintf("[GET /opportunities/tasks/{task_id}/][%d] getOpportunitiesTasksTaskIdOK  %+v", 200, o.Payload)
}

func (o *GetOpportunitiesTasksTaskIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Cache-Control
	o.CacheControl = response.GetHeader("Cache-Control")

	// response header Expires
	o.Expires = response.GetHeader("Expires")

	// response header Last-Modified
	o.LastModified = response.GetHeader("Last-Modified")

	o.Payload = new(models.GetOpportunitiesTasksTaskIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOpportunitiesTasksTaskIDInternalServerError creates a GetOpportunitiesTasksTaskIDInternalServerError with default headers values
func NewGetOpportunitiesTasksTaskIDInternalServerError() *GetOpportunitiesTasksTaskIDInternalServerError {
	return &GetOpportunitiesTasksTaskIDInternalServerError{}
}

/*GetOpportunitiesTasksTaskIDInternalServerError handles this case with default header values.

Internal server error
*/
type GetOpportunitiesTasksTaskIDInternalServerError struct {
	Payload *models.GetOpportunitiesTasksTaskIDInternalServerErrorBody
}

func (o *GetOpportunitiesTasksTaskIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /opportunities/tasks/{task_id}/][%d] getOpportunitiesTasksTaskIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOpportunitiesTasksTaskIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetOpportunitiesTasksTaskIDInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
