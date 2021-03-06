package universe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetUniverseConstellationsHandlerFunc turns a function with the right signature into a get universe constellations handler
type GetUniverseConstellationsHandlerFunc func(GetUniverseConstellationsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetUniverseConstellationsHandlerFunc) Handle(params GetUniverseConstellationsParams) middleware.Responder {
	return fn(params)
}

// GetUniverseConstellationsHandler interface for that can handle valid get universe constellations params
type GetUniverseConstellationsHandler interface {
	Handle(GetUniverseConstellationsParams) middleware.Responder
}

// NewGetUniverseConstellations creates a new http.Handler for the get universe constellations operation
func NewGetUniverseConstellations(ctx *middleware.Context, handler GetUniverseConstellationsHandler) *GetUniverseConstellations {
	return &GetUniverseConstellations{Context: ctx, Handler: handler}
}

/*GetUniverseConstellations swagger:route GET /universe/constellations/ Universe getUniverseConstellations

Get constellations

Get a list of constellations

---

Alternate route: `/v1/universe/constellations/`

Alternate route: `/legacy/universe/constellations/`

Alternate route: `/dev/universe/constellations/`


---

This route expires daily at 11:05

*/
type GetUniverseConstellations struct {
	Context *middleware.Context
	Handler GetUniverseConstellationsHandler
}

func (o *GetUniverseConstellations) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetUniverseConstellationsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
