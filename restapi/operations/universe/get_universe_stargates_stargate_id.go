package universe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetUniverseStargatesStargateIDHandlerFunc turns a function with the right signature into a get universe stargates stargate id handler
type GetUniverseStargatesStargateIDHandlerFunc func(GetUniverseStargatesStargateIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetUniverseStargatesStargateIDHandlerFunc) Handle(params GetUniverseStargatesStargateIDParams) middleware.Responder {
	return fn(params)
}

// GetUniverseStargatesStargateIDHandler interface for that can handle valid get universe stargates stargate id params
type GetUniverseStargatesStargateIDHandler interface {
	Handle(GetUniverseStargatesStargateIDParams) middleware.Responder
}

// NewGetUniverseStargatesStargateID creates a new http.Handler for the get universe stargates stargate id operation
func NewGetUniverseStargatesStargateID(ctx *middleware.Context, handler GetUniverseStargatesStargateIDHandler) *GetUniverseStargatesStargateID {
	return &GetUniverseStargatesStargateID{Context: ctx, Handler: handler}
}

/*GetUniverseStargatesStargateID swagger:route GET /universe/stargates/{stargate_id}/ Universe getUniverseStargatesStargateId

Get stargate information

Get information on a stargate

---

Alternate route: `/v1/universe/stargates/{stargate_id}/`

Alternate route: `/legacy/universe/stargates/{stargate_id}/`

Alternate route: `/dev/universe/stargates/{stargate_id}/`


---

This route expires daily at 11:05

*/
type GetUniverseStargatesStargateID struct {
	Context *middleware.Context
	Handler GetUniverseStargatesStargateIDHandler
}

func (o *GetUniverseStargatesStargateID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetUniverseStargatesStargateIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
