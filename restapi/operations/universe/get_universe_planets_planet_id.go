package universe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetUniversePlanetsPlanetIDHandlerFunc turns a function with the right signature into a get universe planets planet id handler
type GetUniversePlanetsPlanetIDHandlerFunc func(GetUniversePlanetsPlanetIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetUniversePlanetsPlanetIDHandlerFunc) Handle(params GetUniversePlanetsPlanetIDParams) middleware.Responder {
	return fn(params)
}

// GetUniversePlanetsPlanetIDHandler interface for that can handle valid get universe planets planet id params
type GetUniversePlanetsPlanetIDHandler interface {
	Handle(GetUniversePlanetsPlanetIDParams) middleware.Responder
}

// NewGetUniversePlanetsPlanetID creates a new http.Handler for the get universe planets planet id operation
func NewGetUniversePlanetsPlanetID(ctx *middleware.Context, handler GetUniversePlanetsPlanetIDHandler) *GetUniversePlanetsPlanetID {
	return &GetUniversePlanetsPlanetID{Context: ctx, Handler: handler}
}

/*GetUniversePlanetsPlanetID swagger:route GET /universe/planets/{planet_id}/ Universe getUniversePlanetsPlanetId

Get planet information

Get information on a planet

---

Alternate route: `/v1/universe/planets/{planet_id}/`

Alternate route: `/legacy/universe/planets/{planet_id}/`

Alternate route: `/dev/universe/planets/{planet_id}/`


---

This route expires daily at 11:05

*/
type GetUniversePlanetsPlanetID struct {
	Context *middleware.Context
	Handler GetUniversePlanetsPlanetIDHandler
}

func (o *GetUniversePlanetsPlanetID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetUniversePlanetsPlanetIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
