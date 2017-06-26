package universe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetUniverseStructuresHandlerFunc turns a function with the right signature into a get universe structures handler
type GetUniverseStructuresHandlerFunc func(GetUniverseStructuresParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetUniverseStructuresHandlerFunc) Handle(params GetUniverseStructuresParams) middleware.Responder {
	return fn(params)
}

// GetUniverseStructuresHandler interface for that can handle valid get universe structures params
type GetUniverseStructuresHandler interface {
	Handle(GetUniverseStructuresParams) middleware.Responder
}

// NewGetUniverseStructures creates a new http.Handler for the get universe structures operation
func NewGetUniverseStructures(ctx *middleware.Context, handler GetUniverseStructuresHandler) *GetUniverseStructures {
	return &GetUniverseStructures{Context: ctx, Handler: handler}
}

/*GetUniverseStructures swagger:route GET /universe/structures/ Universe getUniverseStructures

List all public structures

List all public structures

---

Alternate route: `/v1/universe/structures/`

Alternate route: `/legacy/universe/structures/`

Alternate route: `/dev/universe/structures/`


---

This route is cached for up to 3600 seconds

*/
type GetUniverseStructures struct {
	Context *middleware.Context
	Handler GetUniverseStructuresHandler
}

func (o *GetUniverseStructures) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetUniverseStructuresParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}