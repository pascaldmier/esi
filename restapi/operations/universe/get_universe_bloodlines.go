package universe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetUniverseBloodlinesHandlerFunc turns a function with the right signature into a get universe bloodlines handler
type GetUniverseBloodlinesHandlerFunc func(GetUniverseBloodlinesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetUniverseBloodlinesHandlerFunc) Handle(params GetUniverseBloodlinesParams) middleware.Responder {
	return fn(params)
}

// GetUniverseBloodlinesHandler interface for that can handle valid get universe bloodlines params
type GetUniverseBloodlinesHandler interface {
	Handle(GetUniverseBloodlinesParams) middleware.Responder
}

// NewGetUniverseBloodlines creates a new http.Handler for the get universe bloodlines operation
func NewGetUniverseBloodlines(ctx *middleware.Context, handler GetUniverseBloodlinesHandler) *GetUniverseBloodlines {
	return &GetUniverseBloodlines{Context: ctx, Handler: handler}
}

/*GetUniverseBloodlines swagger:route GET /universe/bloodlines/ Universe getUniverseBloodlines

Get bloodlines

Get a list of bloodlines

---

Alternate route: `/v1/universe/bloodlines/`

Alternate route: `/legacy/universe/bloodlines/`

Alternate route: `/dev/universe/bloodlines/`


---

This route expires daily at 11:05

*/
type GetUniverseBloodlines struct {
	Context *middleware.Context
	Handler GetUniverseBloodlinesHandler
}

func (o *GetUniverseBloodlines) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetUniverseBloodlinesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}