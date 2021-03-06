package alliance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetAlliancesHandlerFunc turns a function with the right signature into a get alliances handler
type GetAlliancesHandlerFunc func(GetAlliancesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAlliancesHandlerFunc) Handle(params GetAlliancesParams) middleware.Responder {
	return fn(params)
}

// GetAlliancesHandler interface for that can handle valid get alliances params
type GetAlliancesHandler interface {
	Handle(GetAlliancesParams) middleware.Responder
}

// NewGetAlliances creates a new http.Handler for the get alliances operation
func NewGetAlliances(ctx *middleware.Context, handler GetAlliancesHandler) *GetAlliances {
	return &GetAlliances{Context: ctx, Handler: handler}
}

/*GetAlliances swagger:route GET /alliances/ Alliance getAlliances

List all alliances

List all active player alliances

---

Alternate route: `/v1/alliances/`

Alternate route: `/legacy/alliances/`

Alternate route: `/dev/alliances/`


---

This route is cached for up to 3600 seconds

*/
type GetAlliances struct {
	Context *middleware.Context
	Handler GetAlliancesHandler
}

func (o *GetAlliances) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetAlliancesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
