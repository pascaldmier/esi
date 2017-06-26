package fleets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteFleetsFleetIDSquadsSquadIDHandlerFunc turns a function with the right signature into a delete fleets fleet id squads squad id handler
type DeleteFleetsFleetIDSquadsSquadIDHandlerFunc func(DeleteFleetsFleetIDSquadsSquadIDParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteFleetsFleetIDSquadsSquadIDHandlerFunc) Handle(params DeleteFleetsFleetIDSquadsSquadIDParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeleteFleetsFleetIDSquadsSquadIDHandler interface for that can handle valid delete fleets fleet id squads squad id params
type DeleteFleetsFleetIDSquadsSquadIDHandler interface {
	Handle(DeleteFleetsFleetIDSquadsSquadIDParams, interface{}) middleware.Responder
}

// NewDeleteFleetsFleetIDSquadsSquadID creates a new http.Handler for the delete fleets fleet id squads squad id operation
func NewDeleteFleetsFleetIDSquadsSquadID(ctx *middleware.Context, handler DeleteFleetsFleetIDSquadsSquadIDHandler) *DeleteFleetsFleetIDSquadsSquadID {
	return &DeleteFleetsFleetIDSquadsSquadID{Context: ctx, Handler: handler}
}

/*DeleteFleetsFleetIDSquadsSquadID swagger:route DELETE /fleets/{fleet_id}/squads/{squad_id}/ Fleets deleteFleetsFleetIdSquadsSquadId

Delete fleet squad

Delete a fleet squad, only empty squads can be deleted

---

Alternate route: `/v1/fleets/{fleet_id}/squads/{squad_id}/`

Alternate route: `/legacy/fleets/{fleet_id}/squads/{squad_id}/`

Alternate route: `/dev/fleets/{fleet_id}/squads/{squad_id}/`


*/
type DeleteFleetsFleetIDSquadsSquadID struct {
	Context *middleware.Context
	Handler DeleteFleetsFleetIDSquadsSquadIDHandler
}

func (o *DeleteFleetsFleetIDSquadsSquadID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewDeleteFleetsFleetIDSquadsSquadIDParams()

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
