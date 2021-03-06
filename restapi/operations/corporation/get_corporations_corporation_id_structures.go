package corporation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetCorporationsCorporationIDStructuresHandlerFunc turns a function with the right signature into a get corporations corporation id structures handler
type GetCorporationsCorporationIDStructuresHandlerFunc func(GetCorporationsCorporationIDStructuresParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetCorporationsCorporationIDStructuresHandlerFunc) Handle(params GetCorporationsCorporationIDStructuresParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetCorporationsCorporationIDStructuresHandler interface for that can handle valid get corporations corporation id structures params
type GetCorporationsCorporationIDStructuresHandler interface {
	Handle(GetCorporationsCorporationIDStructuresParams, interface{}) middleware.Responder
}

// NewGetCorporationsCorporationIDStructures creates a new http.Handler for the get corporations corporation id structures operation
func NewGetCorporationsCorporationIDStructures(ctx *middleware.Context, handler GetCorporationsCorporationIDStructuresHandler) *GetCorporationsCorporationIDStructures {
	return &GetCorporationsCorporationIDStructures{Context: ctx, Handler: handler}
}

/*GetCorporationsCorporationIDStructures swagger:route GET /corporations/{corporation_id}/structures/ Corporation getCorporationsCorporationIdStructures

Get corporation structures

Get a list of corporation structures

---

Alternate route: `/v1/corporations/{corporation_id}/structures/`

Alternate route: `/legacy/corporations/{corporation_id}/structures/`

Alternate route: `/dev/corporations/{corporation_id}/structures/`


---

This route is cached for up to 3600 seconds

*/
type GetCorporationsCorporationIDStructures struct {
	Context *middleware.Context
	Handler GetCorporationsCorporationIDStructuresHandler
}

func (o *GetCorporationsCorporationIDStructures) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetCorporationsCorporationIDStructuresParams()

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
