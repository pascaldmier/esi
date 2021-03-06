package fleets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetFleetsFleetIDMembersHandlerFunc turns a function with the right signature into a get fleets fleet id members handler
type GetFleetsFleetIDMembersHandlerFunc func(GetFleetsFleetIDMembersParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetFleetsFleetIDMembersHandlerFunc) Handle(params GetFleetsFleetIDMembersParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetFleetsFleetIDMembersHandler interface for that can handle valid get fleets fleet id members params
type GetFleetsFleetIDMembersHandler interface {
	Handle(GetFleetsFleetIDMembersParams, interface{}) middleware.Responder
}

// NewGetFleetsFleetIDMembers creates a new http.Handler for the get fleets fleet id members operation
func NewGetFleetsFleetIDMembers(ctx *middleware.Context, handler GetFleetsFleetIDMembersHandler) *GetFleetsFleetIDMembers {
	return &GetFleetsFleetIDMembers{Context: ctx, Handler: handler}
}

/*GetFleetsFleetIDMembers swagger:route GET /fleets/{fleet_id}/members/ Fleets getFleetsFleetIdMembers

Get fleet members

Return information about fleet members

---

Alternate route: `/v1/fleets/{fleet_id}/members/`

Alternate route: `/legacy/fleets/{fleet_id}/members/`

Alternate route: `/dev/fleets/{fleet_id}/members/`


---

This route is cached for up to 5 seconds

*/
type GetFleetsFleetIDMembers struct {
	Context *middleware.Context
	Handler GetFleetsFleetIDMembersHandler
}

func (o *GetFleetsFleetIDMembers) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetFleetsFleetIDMembersParams()

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
