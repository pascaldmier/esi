package wars

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetWarsHandlerFunc turns a function with the right signature into a get wars handler
type GetWarsHandlerFunc func(GetWarsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetWarsHandlerFunc) Handle(params GetWarsParams) middleware.Responder {
	return fn(params)
}

// GetWarsHandler interface for that can handle valid get wars params
type GetWarsHandler interface {
	Handle(GetWarsParams) middleware.Responder
}

// NewGetWars creates a new http.Handler for the get wars operation
func NewGetWars(ctx *middleware.Context, handler GetWarsHandler) *GetWars {
	return &GetWars{Context: ctx, Handler: handler}
}

/*GetWars swagger:route GET /wars/ Wars getWars

List wars

Return a list of wars

---

Alternate route: `/v1/wars/`

Alternate route: `/legacy/wars/`

Alternate route: `/dev/wars/`


---

This route is cached for up to 3600 seconds

*/
type GetWars struct {
	Context *middleware.Context
	Handler GetWarsHandler
}

func (o *GetWars) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetWarsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
