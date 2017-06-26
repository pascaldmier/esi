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

// GetUniverseGroupsGroupIDReader is a Reader for the GetUniverseGroupsGroupID structure.
type GetUniverseGroupsGroupIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUniverseGroupsGroupIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUniverseGroupsGroupIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetUniverseGroupsGroupIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetUniverseGroupsGroupIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUniverseGroupsGroupIDOK creates a GetUniverseGroupsGroupIDOK with default headers values
func NewGetUniverseGroupsGroupIDOK() *GetUniverseGroupsGroupIDOK {
	return &GetUniverseGroupsGroupIDOK{}
}

/*GetUniverseGroupsGroupIDOK handles this case with default header values.

Information about an item group
*/
type GetUniverseGroupsGroupIDOK struct {
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

	Payload *models.GetUniverseGroupsGroupIDOKBody
}

func (o *GetUniverseGroupsGroupIDOK) Error() string {
	return fmt.Sprintf("[GET /universe/groups/{group_id}/][%d] getUniverseGroupsGroupIdOK  %+v", 200, o.Payload)
}

func (o *GetUniverseGroupsGroupIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Cache-Control
	o.CacheControl = response.GetHeader("Cache-Control")

	// response header Content-Language
	o.ContentLanguage = response.GetHeader("Content-Language")

	// response header Expires
	o.Expires = response.GetHeader("Expires")

	// response header Last-Modified
	o.LastModified = response.GetHeader("Last-Modified")

	o.Payload = new(models.GetUniverseGroupsGroupIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseGroupsGroupIDNotFound creates a GetUniverseGroupsGroupIDNotFound with default headers values
func NewGetUniverseGroupsGroupIDNotFound() *GetUniverseGroupsGroupIDNotFound {
	return &GetUniverseGroupsGroupIDNotFound{}
}

/*GetUniverseGroupsGroupIDNotFound handles this case with default header values.

Group not found
*/
type GetUniverseGroupsGroupIDNotFound struct {
	Payload *models.GetUniverseGroupsGroupIDNotFoundBody
}

func (o *GetUniverseGroupsGroupIDNotFound) Error() string {
	return fmt.Sprintf("[GET /universe/groups/{group_id}/][%d] getUniverseGroupsGroupIdNotFound  %+v", 404, o.Payload)
}

func (o *GetUniverseGroupsGroupIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetUniverseGroupsGroupIDNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseGroupsGroupIDInternalServerError creates a GetUniverseGroupsGroupIDInternalServerError with default headers values
func NewGetUniverseGroupsGroupIDInternalServerError() *GetUniverseGroupsGroupIDInternalServerError {
	return &GetUniverseGroupsGroupIDInternalServerError{}
}

/*GetUniverseGroupsGroupIDInternalServerError handles this case with default header values.

Internal server error
*/
type GetUniverseGroupsGroupIDInternalServerError struct {
	Payload *models.GetUniverseGroupsGroupIDInternalServerErrorBody
}

func (o *GetUniverseGroupsGroupIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /universe/groups/{group_id}/][%d] getUniverseGroupsGroupIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUniverseGroupsGroupIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetUniverseGroupsGroupIDInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
