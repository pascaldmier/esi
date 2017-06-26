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

// GetCharactersCharacterIDChatChannelsReader is a Reader for the GetCharactersCharacterIDChatChannels structure.
type GetCharactersCharacterIDChatChannelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDChatChannelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCharactersCharacterIDChatChannelsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetCharactersCharacterIDChatChannelsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCharactersCharacterIDChatChannelsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCharactersCharacterIDChatChannelsOK creates a GetCharactersCharacterIDChatChannelsOK with default headers values
func NewGetCharactersCharacterIDChatChannelsOK() *GetCharactersCharacterIDChatChannelsOK {
	return &GetCharactersCharacterIDChatChannelsOK{}
}

/*GetCharactersCharacterIDChatChannelsOK handles this case with default header values.

A list of chat channels
*/
type GetCharactersCharacterIDChatChannelsOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload models.GetCharactersCharacterIDChatChannelsOKBody
}

func (o *GetCharactersCharacterIDChatChannelsOK) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/chat_channels/][%d] getCharactersCharacterIdChatChannelsOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDChatChannelsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDChatChannelsForbidden creates a GetCharactersCharacterIDChatChannelsForbidden with default headers values
func NewGetCharactersCharacterIDChatChannelsForbidden() *GetCharactersCharacterIDChatChannelsForbidden {
	return &GetCharactersCharacterIDChatChannelsForbidden{}
}

/*GetCharactersCharacterIDChatChannelsForbidden handles this case with default header values.

Forbidden
*/
type GetCharactersCharacterIDChatChannelsForbidden struct {
	Payload *models.GetCharactersCharacterIDChatChannelsForbiddenBody
}

func (o *GetCharactersCharacterIDChatChannelsForbidden) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/chat_channels/][%d] getCharactersCharacterIdChatChannelsForbidden  %+v", 403, o.Payload)
}

func (o *GetCharactersCharacterIDChatChannelsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCharactersCharacterIDChatChannelsForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDChatChannelsInternalServerError creates a GetCharactersCharacterIDChatChannelsInternalServerError with default headers values
func NewGetCharactersCharacterIDChatChannelsInternalServerError() *GetCharactersCharacterIDChatChannelsInternalServerError {
	return &GetCharactersCharacterIDChatChannelsInternalServerError{}
}

/*GetCharactersCharacterIDChatChannelsInternalServerError handles this case with default header values.

Internal server error
*/
type GetCharactersCharacterIDChatChannelsInternalServerError struct {
	Payload *models.GetCharactersCharacterIDChatChannelsInternalServerErrorBody
}

func (o *GetCharactersCharacterIDChatChannelsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/chat_channels/][%d] getCharactersCharacterIdChatChannelsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDChatChannelsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCharactersCharacterIDChatChannelsInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
