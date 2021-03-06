package incursions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetIncursionsHandlerFunc turns a function with the right signature into a get incursions handler
type GetIncursionsHandlerFunc func(GetIncursionsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetIncursionsHandlerFunc) Handle(params GetIncursionsParams) middleware.Responder {
	return fn(params)
}

// GetIncursionsHandler interface for that can handle valid get incursions params
type GetIncursionsHandler interface {
	Handle(GetIncursionsParams) middleware.Responder
}

// NewGetIncursions creates a new http.Handler for the get incursions operation
func NewGetIncursions(ctx *middleware.Context, handler GetIncursionsHandler) *GetIncursions {
	return &GetIncursions{Context: ctx, Handler: handler}
}

/*GetIncursions swagger:route GET /incursions/ Incursions getIncursions

List incursions

Return a list of current incursions

---

Alternate route: `/v1/incursions/`

Alternate route: `/legacy/incursions/`

Alternate route: `/dev/incursions/`


---

This route is cached for up to 300 seconds

*/
type GetIncursions struct {
	Context *middleware.Context
	Handler GetIncursionsHandler
}

func (o *GetIncursions) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetIncursionsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
