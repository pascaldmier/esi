package market

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetMarketsStructuresStructureIDHandlerFunc turns a function with the right signature into a get markets structures structure id handler
type GetMarketsStructuresStructureIDHandlerFunc func(GetMarketsStructuresStructureIDParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetMarketsStructuresStructureIDHandlerFunc) Handle(params GetMarketsStructuresStructureIDParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetMarketsStructuresStructureIDHandler interface for that can handle valid get markets structures structure id params
type GetMarketsStructuresStructureIDHandler interface {
	Handle(GetMarketsStructuresStructureIDParams, interface{}) middleware.Responder
}

// NewGetMarketsStructuresStructureID creates a new http.Handler for the get markets structures structure id operation
func NewGetMarketsStructuresStructureID(ctx *middleware.Context, handler GetMarketsStructuresStructureIDHandler) *GetMarketsStructuresStructureID {
	return &GetMarketsStructuresStructureID{Context: ctx, Handler: handler}
}

/*GetMarketsStructuresStructureID swagger:route GET /markets/structures/{structure_id}/ Market getMarketsStructuresStructureId

List orders in a structure

Return all orders in a structure

---

Alternate route: `/v1/markets/structures/{structure_id}/`

Alternate route: `/legacy/markets/structures/{structure_id}/`

Alternate route: `/dev/markets/structures/{structure_id}/`


---

This route is cached for up to 300 seconds

*/
type GetMarketsStructuresStructureID struct {
	Context *middleware.Context
	Handler GetMarketsStructuresStructureIDHandler
}

func (o *GetMarketsStructuresStructureID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetMarketsStructuresStructureIDParams()

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