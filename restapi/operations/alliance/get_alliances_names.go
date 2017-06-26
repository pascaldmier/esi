package alliance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetAlliancesNamesHandlerFunc turns a function with the right signature into a get alliances names handler
type GetAlliancesNamesHandlerFunc func(GetAlliancesNamesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAlliancesNamesHandlerFunc) Handle(params GetAlliancesNamesParams) middleware.Responder {
	return fn(params)
}

// GetAlliancesNamesHandler interface for that can handle valid get alliances names params
type GetAlliancesNamesHandler interface {
	Handle(GetAlliancesNamesParams) middleware.Responder
}

// NewGetAlliancesNames creates a new http.Handler for the get alliances names operation
func NewGetAlliancesNames(ctx *middleware.Context, handler GetAlliancesNamesHandler) *GetAlliancesNames {
	return &GetAlliancesNames{Context: ctx, Handler: handler}
}

/*GetAlliancesNames swagger:route GET /alliances/names/ Alliance getAlliancesNames

Get alliance names

Resolve a set of alliance IDs to alliance names

---

Alternate route: `/v1/alliances/names/`

Alternate route: `/legacy/alliances/names/`

Alternate route: `/dev/alliances/names/`


---

This route is cached for up to 3600 seconds

*/
type GetAlliancesNames struct {
	Context *middleware.Context
	Handler GetAlliancesNamesHandler
}

func (o *GetAlliancesNames) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetAlliancesNamesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}