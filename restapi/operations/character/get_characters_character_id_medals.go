package character

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetCharactersCharacterIDMedalsHandlerFunc turns a function with the right signature into a get characters character id medals handler
type GetCharactersCharacterIDMedalsHandlerFunc func(GetCharactersCharacterIDMedalsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetCharactersCharacterIDMedalsHandlerFunc) Handle(params GetCharactersCharacterIDMedalsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetCharactersCharacterIDMedalsHandler interface for that can handle valid get characters character id medals params
type GetCharactersCharacterIDMedalsHandler interface {
	Handle(GetCharactersCharacterIDMedalsParams, interface{}) middleware.Responder
}

// NewGetCharactersCharacterIDMedals creates a new http.Handler for the get characters character id medals operation
func NewGetCharactersCharacterIDMedals(ctx *middleware.Context, handler GetCharactersCharacterIDMedalsHandler) *GetCharactersCharacterIDMedals {
	return &GetCharactersCharacterIDMedals{Context: ctx, Handler: handler}
}

/*GetCharactersCharacterIDMedals swagger:route GET /characters/{character_id}/medals/ Character getCharactersCharacterIdMedals

Get medals

Return a list of medals the character has

---

Alternate route: `/v1/characters/{character_id}/medals/`

Alternate route: `/legacy/characters/{character_id}/medals/`

Alternate route: `/dev/characters/{character_id}/medals/`


---

This route is cached for up to 3600 seconds

*/
type GetCharactersCharacterIDMedals struct {
	Context *middleware.Context
	Handler GetCharactersCharacterIDMedalsHandler
}

func (o *GetCharactersCharacterIDMedals) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetCharactersCharacterIDMedalsParams()

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
