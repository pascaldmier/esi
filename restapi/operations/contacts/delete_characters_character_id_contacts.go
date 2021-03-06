package contacts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteCharactersCharacterIDContactsHandlerFunc turns a function with the right signature into a delete characters character id contacts handler
type DeleteCharactersCharacterIDContactsHandlerFunc func(DeleteCharactersCharacterIDContactsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteCharactersCharacterIDContactsHandlerFunc) Handle(params DeleteCharactersCharacterIDContactsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeleteCharactersCharacterIDContactsHandler interface for that can handle valid delete characters character id contacts params
type DeleteCharactersCharacterIDContactsHandler interface {
	Handle(DeleteCharactersCharacterIDContactsParams, interface{}) middleware.Responder
}

// NewDeleteCharactersCharacterIDContacts creates a new http.Handler for the delete characters character id contacts operation
func NewDeleteCharactersCharacterIDContacts(ctx *middleware.Context, handler DeleteCharactersCharacterIDContactsHandler) *DeleteCharactersCharacterIDContacts {
	return &DeleteCharactersCharacterIDContacts{Context: ctx, Handler: handler}
}

/*DeleteCharactersCharacterIDContacts swagger:route DELETE /characters/{character_id}/contacts/ Contacts deleteCharactersCharacterIdContacts

Delete contacts

Bulk delete contacts

---

Alternate route: `/v1/characters/{character_id}/contacts/`

Alternate route: `/legacy/characters/{character_id}/contacts/`


*/
type DeleteCharactersCharacterIDContacts struct {
	Context *middleware.Context
	Handler DeleteCharactersCharacterIDContactsHandler
}

func (o *DeleteCharactersCharacterIDContacts) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewDeleteCharactersCharacterIDContactsParams()

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
