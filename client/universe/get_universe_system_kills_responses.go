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

// GetUniverseSystemKillsReader is a Reader for the GetUniverseSystemKills structure.
type GetUniverseSystemKillsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUniverseSystemKillsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUniverseSystemKillsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetUniverseSystemKillsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUniverseSystemKillsOK creates a GetUniverseSystemKillsOK with default headers values
func NewGetUniverseSystemKillsOK() *GetUniverseSystemKillsOK {
	return &GetUniverseSystemKillsOK{}
}

/*GetUniverseSystemKillsOK handles this case with default header values.

A list of systems and number of ship, pod and NPC kills
*/
type GetUniverseSystemKillsOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload models.GetUniverseSystemKillsOKBody
}

func (o *GetUniverseSystemKillsOK) Error() string {
	return fmt.Sprintf("[GET /universe/system_kills/][%d] getUniverseSystemKillsOK  %+v", 200, o.Payload)
}

func (o *GetUniverseSystemKillsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetUniverseSystemKillsInternalServerError creates a GetUniverseSystemKillsInternalServerError with default headers values
func NewGetUniverseSystemKillsInternalServerError() *GetUniverseSystemKillsInternalServerError {
	return &GetUniverseSystemKillsInternalServerError{}
}

/*GetUniverseSystemKillsInternalServerError handles this case with default header values.

Internal server error
*/
type GetUniverseSystemKillsInternalServerError struct {
	Payload *models.GetUniverseSystemKillsInternalServerErrorBody
}

func (o *GetUniverseSystemKillsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /universe/system_kills/][%d] getUniverseSystemKillsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUniverseSystemKillsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetUniverseSystemKillsInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
