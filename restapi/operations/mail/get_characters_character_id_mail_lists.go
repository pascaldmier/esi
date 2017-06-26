package mail

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetCharactersCharacterIDMailListsHandlerFunc turns a function with the right signature into a get characters character id mail lists handler
type GetCharactersCharacterIDMailListsHandlerFunc func(GetCharactersCharacterIDMailListsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetCharactersCharacterIDMailListsHandlerFunc) Handle(params GetCharactersCharacterIDMailListsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetCharactersCharacterIDMailListsHandler interface for that can handle valid get characters character id mail lists params
type GetCharactersCharacterIDMailListsHandler interface {
	Handle(GetCharactersCharacterIDMailListsParams, interface{}) middleware.Responder
}

// NewGetCharactersCharacterIDMailLists creates a new http.Handler for the get characters character id mail lists operation
func NewGetCharactersCharacterIDMailLists(ctx *middleware.Context, handler GetCharactersCharacterIDMailListsHandler) *GetCharactersCharacterIDMailLists {
	return &GetCharactersCharacterIDMailLists{Context: ctx, Handler: handler}
}

/*GetCharactersCharacterIDMailLists swagger:route GET /characters/{character_id}/mail/lists/ Mail getCharactersCharacterIdMailLists

Return mailing list subscriptions

Return all mailing lists that the character is subscribed to


---

Alternate route: `/v1/characters/{character_id}/mail/lists/`

Alternate route: `/legacy/characters/{character_id}/mail/lists/`

Alternate route: `/dev/characters/{character_id}/mail/lists/`


---

This route is cached for up to 120 seconds

*/
type GetCharactersCharacterIDMailLists struct {
	Context *middleware.Context
	Handler GetCharactersCharacterIDMailListsHandler
}

func (o *GetCharactersCharacterIDMailLists) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetCharactersCharacterIDMailListsParams()

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