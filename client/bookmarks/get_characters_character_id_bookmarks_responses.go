package bookmarks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/pascaldmier/esi/models"
)

// GetCharactersCharacterIDBookmarksReader is a Reader for the GetCharactersCharacterIDBookmarks structure.
type GetCharactersCharacterIDBookmarksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDBookmarksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCharactersCharacterIDBookmarksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetCharactersCharacterIDBookmarksForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCharactersCharacterIDBookmarksInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCharactersCharacterIDBookmarksOK creates a GetCharactersCharacterIDBookmarksOK with default headers values
func NewGetCharactersCharacterIDBookmarksOK() *GetCharactersCharacterIDBookmarksOK {
	return &GetCharactersCharacterIDBookmarksOK{}
}

/*GetCharactersCharacterIDBookmarksOK handles this case with default header values.

A list of bookmarks
*/
type GetCharactersCharacterIDBookmarksOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload models.GetCharactersCharacterIDBookmarksOKBody
}

func (o *GetCharactersCharacterIDBookmarksOK) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/bookmarks/][%d] getCharactersCharacterIdBookmarksOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDBookmarksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDBookmarksForbidden creates a GetCharactersCharacterIDBookmarksForbidden with default headers values
func NewGetCharactersCharacterIDBookmarksForbidden() *GetCharactersCharacterIDBookmarksForbidden {
	return &GetCharactersCharacterIDBookmarksForbidden{}
}

/*GetCharactersCharacterIDBookmarksForbidden handles this case with default header values.

Forbidden
*/
type GetCharactersCharacterIDBookmarksForbidden struct {
	Payload *models.GetCharactersCharacterIDBookmarksForbiddenBody
}

func (o *GetCharactersCharacterIDBookmarksForbidden) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/bookmarks/][%d] getCharactersCharacterIdBookmarksForbidden  %+v", 403, o.Payload)
}

func (o *GetCharactersCharacterIDBookmarksForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCharactersCharacterIDBookmarksForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDBookmarksInternalServerError creates a GetCharactersCharacterIDBookmarksInternalServerError with default headers values
func NewGetCharactersCharacterIDBookmarksInternalServerError() *GetCharactersCharacterIDBookmarksInternalServerError {
	return &GetCharactersCharacterIDBookmarksInternalServerError{}
}

/*GetCharactersCharacterIDBookmarksInternalServerError handles this case with default header values.

Internal server error
*/
type GetCharactersCharacterIDBookmarksInternalServerError struct {
	Payload *models.GetCharactersCharacterIDBookmarksInternalServerErrorBody
}

func (o *GetCharactersCharacterIDBookmarksInternalServerError) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/bookmarks/][%d] getCharactersCharacterIdBookmarksInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDBookmarksInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCharactersCharacterIDBookmarksInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}