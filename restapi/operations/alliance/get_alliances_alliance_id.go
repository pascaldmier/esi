package alliance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetAlliancesAllianceIDHandlerFunc turns a function with the right signature into a get alliances alliance id handler
type GetAlliancesAllianceIDHandlerFunc func(GetAlliancesAllianceIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAlliancesAllianceIDHandlerFunc) Handle(params GetAlliancesAllianceIDParams) middleware.Responder {
	return fn(params)
}

// GetAlliancesAllianceIDHandler interface for that can handle valid get alliances alliance id params
type GetAlliancesAllianceIDHandler interface {
	Handle(GetAlliancesAllianceIDParams) middleware.Responder
}

// NewGetAlliancesAllianceID creates a new http.Handler for the get alliances alliance id operation
func NewGetAlliancesAllianceID(ctx *middleware.Context, handler GetAlliancesAllianceIDHandler) *GetAlliancesAllianceID {
	return &GetAlliancesAllianceID{Context: ctx, Handler: handler}
}

/*GetAlliancesAllianceID swagger:route GET /alliances/{alliance_id}/ Alliance getAlliancesAllianceId

Get alliance information

Public information about an alliance

---

Alternate route: `/v2/alliances/{alliance_id}/`


---

This route is cached for up to 3600 seconds

*/
type GetAlliancesAllianceID struct {
	Context *middleware.Context
	Handler GetAlliancesAllianceIDHandler
}

func (o *GetAlliancesAllianceID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetAlliancesAllianceIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
