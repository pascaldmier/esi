package bookmarks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetCharactersCharacterIDBookmarksFoldersHandlerFunc turns a function with the right signature into a get characters character id bookmarks folders handler
type GetCharactersCharacterIDBookmarksFoldersHandlerFunc func(GetCharactersCharacterIDBookmarksFoldersParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetCharactersCharacterIDBookmarksFoldersHandlerFunc) Handle(params GetCharactersCharacterIDBookmarksFoldersParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetCharactersCharacterIDBookmarksFoldersHandler interface for that can handle valid get characters character id bookmarks folders params
type GetCharactersCharacterIDBookmarksFoldersHandler interface {
	Handle(GetCharactersCharacterIDBookmarksFoldersParams, interface{}) middleware.Responder
}

// NewGetCharactersCharacterIDBookmarksFolders creates a new http.Handler for the get characters character id bookmarks folders operation
func NewGetCharactersCharacterIDBookmarksFolders(ctx *middleware.Context, handler GetCharactersCharacterIDBookmarksFoldersHandler) *GetCharactersCharacterIDBookmarksFolders {
	return &GetCharactersCharacterIDBookmarksFolders{Context: ctx, Handler: handler}
}

/*GetCharactersCharacterIDBookmarksFolders swagger:route GET /characters/{character_id}/bookmarks/folders/ Bookmarks getCharactersCharacterIdBookmarksFolders

List bookmark folders

List your character's personal bookmark folders

---

Alternate route: `/v1/characters/{character_id}/bookmarks/folders/`

Alternate route: `/legacy/characters/{character_id}/bookmarks/folders/`

Alternate route: `/dev/characters/{character_id}/bookmarks/folders/`


---

This route is cached for up to 3600 seconds

*/
type GetCharactersCharacterIDBookmarksFolders struct {
	Context *middleware.Context
	Handler GetCharactersCharacterIDBookmarksFoldersHandler
}

func (o *GetCharactersCharacterIDBookmarksFolders) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetCharactersCharacterIDBookmarksFoldersParams()

	uprinc, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
