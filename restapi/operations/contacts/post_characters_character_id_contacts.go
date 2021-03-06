package contacts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostCharactersCharacterIDContactsHandlerFunc turns a function with the right signature into a post characters character id contacts handler
type PostCharactersCharacterIDContactsHandlerFunc func(PostCharactersCharacterIDContactsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn PostCharactersCharacterIDContactsHandlerFunc) Handle(params PostCharactersCharacterIDContactsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// PostCharactersCharacterIDContactsHandler interface for that can handle valid post characters character id contacts params
type PostCharactersCharacterIDContactsHandler interface {
	Handle(PostCharactersCharacterIDContactsParams, interface{}) middleware.Responder
}

// NewPostCharactersCharacterIDContacts creates a new http.Handler for the post characters character id contacts operation
func NewPostCharactersCharacterIDContacts(ctx *middleware.Context, handler PostCharactersCharacterIDContactsHandler) *PostCharactersCharacterIDContacts {
	return &PostCharactersCharacterIDContacts{Context: ctx, Handler: handler}
}

/*PostCharactersCharacterIDContacts swagger:route POST /characters/{character_id}/contacts/ Contacts postCharactersCharacterIdContacts

Add contacts

Bulk add contacts with same settings

---

Alternate route: `/v1/characters/{character_id}/contacts/`

Alternate route: `/legacy/characters/{character_id}/contacts/`

Alternate route: `/dev/characters/{character_id}/contacts/`


*/
type PostCharactersCharacterIDContacts struct {
	Context *middleware.Context
	Handler PostCharactersCharacterIDContactsHandler
}

func (o *PostCharactersCharacterIDContacts) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewPostCharactersCharacterIDContactsParams()

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
