package character

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostCharactersCharacterIDCspaHandlerFunc turns a function with the right signature into a post characters character id cspa handler
type PostCharactersCharacterIDCspaHandlerFunc func(PostCharactersCharacterIDCspaParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn PostCharactersCharacterIDCspaHandlerFunc) Handle(params PostCharactersCharacterIDCspaParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// PostCharactersCharacterIDCspaHandler interface for that can handle valid post characters character id cspa params
type PostCharactersCharacterIDCspaHandler interface {
	Handle(PostCharactersCharacterIDCspaParams, interface{}) middleware.Responder
}

// NewPostCharactersCharacterIDCspa creates a new http.Handler for the post characters character id cspa operation
func NewPostCharactersCharacterIDCspa(ctx *middleware.Context, handler PostCharactersCharacterIDCspaHandler) *PostCharactersCharacterIDCspa {
	return &PostCharactersCharacterIDCspa{Context: ctx, Handler: handler}
}

/*PostCharactersCharacterIDCspa swagger:route POST /characters/{character_id}/cspa/ Character postCharactersCharacterIdCspa

Calculate a CSPA charge cost

Takes a source character ID in the url and a set of target character ID's in the body, returns a CSPA charge cost

---

Alternate route: `/v3/characters/{character_id}/cspa/`

Alternate route: `/legacy/characters/{character_id}/cspa/`

Alternate route: `/dev/characters/{character_id}/cspa/`


*/
type PostCharactersCharacterIDCspa struct {
	Context *middleware.Context
	Handler PostCharactersCharacterIDCspaHandler
}

func (o *PostCharactersCharacterIDCspa) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewPostCharactersCharacterIDCspaParams()

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
