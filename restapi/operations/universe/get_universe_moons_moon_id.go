package universe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetUniverseMoonsMoonIDHandlerFunc turns a function with the right signature into a get universe moons moon id handler
type GetUniverseMoonsMoonIDHandlerFunc func(GetUniverseMoonsMoonIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetUniverseMoonsMoonIDHandlerFunc) Handle(params GetUniverseMoonsMoonIDParams) middleware.Responder {
	return fn(params)
}

// GetUniverseMoonsMoonIDHandler interface for that can handle valid get universe moons moon id params
type GetUniverseMoonsMoonIDHandler interface {
	Handle(GetUniverseMoonsMoonIDParams) middleware.Responder
}

// NewGetUniverseMoonsMoonID creates a new http.Handler for the get universe moons moon id operation
func NewGetUniverseMoonsMoonID(ctx *middleware.Context, handler GetUniverseMoonsMoonIDHandler) *GetUniverseMoonsMoonID {
	return &GetUniverseMoonsMoonID{Context: ctx, Handler: handler}
}

/*GetUniverseMoonsMoonID swagger:route GET /universe/moons/{moon_id}/ Universe getUniverseMoonsMoonId

Get moon information

Get information on a moon

---

Alternate route: `/v1/universe/moons/{moon_id}/`

Alternate route: `/legacy/universe/moons/{moon_id}/`

Alternate route: `/dev/universe/moons/{moon_id}/`


---

This route expires daily at 11:05

*/
type GetUniverseMoonsMoonID struct {
	Context *middleware.Context
	Handler GetUniverseMoonsMoonIDHandler
}

func (o *GetUniverseMoonsMoonID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetUniverseMoonsMoonIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
