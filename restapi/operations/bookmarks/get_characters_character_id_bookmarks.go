package bookmarks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetCharactersCharacterIDBookmarksHandlerFunc turns a function with the right signature into a get characters character id bookmarks handler
type GetCharactersCharacterIDBookmarksHandlerFunc func(GetCharactersCharacterIDBookmarksParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetCharactersCharacterIDBookmarksHandlerFunc) Handle(params GetCharactersCharacterIDBookmarksParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetCharactersCharacterIDBookmarksHandler interface for that can handle valid get characters character id bookmarks params
type GetCharactersCharacterIDBookmarksHandler interface {
	Handle(GetCharactersCharacterIDBookmarksParams, interface{}) middleware.Responder
}

// NewGetCharactersCharacterIDBookmarks creates a new http.Handler for the get characters character id bookmarks operation
func NewGetCharactersCharacterIDBookmarks(ctx *middleware.Context, handler GetCharactersCharacterIDBookmarksHandler) *GetCharactersCharacterIDBookmarks {
	return &GetCharactersCharacterIDBookmarks{Context: ctx, Handler: handler}
}

/*GetCharactersCharacterIDBookmarks swagger:route GET /characters/{character_id}/bookmarks/ Bookmarks getCharactersCharacterIdBookmarks

List bookmarks

List your character's personal bookmarks

---

Alternate route: `/v1/characters/{character_id}/bookmarks/`

Alternate route: `/legacy/characters/{character_id}/bookmarks/`

Alternate route: `/dev/characters/{character_id}/bookmarks/`


---

This route is cached for up to 3600 seconds

*/
type GetCharactersCharacterIDBookmarks struct {
	Context *middleware.Context
	Handler GetCharactersCharacterIDBookmarksHandler
}

func (o *GetCharactersCharacterIDBookmarks) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetCharactersCharacterIDBookmarksParams()

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