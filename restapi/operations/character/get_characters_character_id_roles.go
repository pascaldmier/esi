package character

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetCharactersCharacterIDRolesHandlerFunc turns a function with the right signature into a get characters character id roles handler
type GetCharactersCharacterIDRolesHandlerFunc func(GetCharactersCharacterIDRolesParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetCharactersCharacterIDRolesHandlerFunc) Handle(params GetCharactersCharacterIDRolesParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetCharactersCharacterIDRolesHandler interface for that can handle valid get characters character id roles params
type GetCharactersCharacterIDRolesHandler interface {
	Handle(GetCharactersCharacterIDRolesParams, interface{}) middleware.Responder
}

// NewGetCharactersCharacterIDRoles creates a new http.Handler for the get characters character id roles operation
func NewGetCharactersCharacterIDRoles(ctx *middleware.Context, handler GetCharactersCharacterIDRolesHandler) *GetCharactersCharacterIDRoles {
	return &GetCharactersCharacterIDRoles{Context: ctx, Handler: handler}
}

/*GetCharactersCharacterIDRoles swagger:route GET /characters/{character_id}/roles/ Character getCharactersCharacterIdRoles

Get character corporation roles

Returns a character's corporation roles

---

Alternate route: `/v1/characters/{character_id}/roles/`

Alternate route: `/legacy/characters/{character_id}/roles/`

Alternate route: `/dev/characters/{character_id}/roles/`


---

This route is cached for up to 3600 seconds

*/
type GetCharactersCharacterIDRoles struct {
	Context *middleware.Context
	Handler GetCharactersCharacterIDRolesHandler
}

func (o *GetCharactersCharacterIDRoles) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetCharactersCharacterIDRolesParams()

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
