package universe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetUniverseGraphicsGraphicIDHandlerFunc turns a function with the right signature into a get universe graphics graphic id handler
type GetUniverseGraphicsGraphicIDHandlerFunc func(GetUniverseGraphicsGraphicIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetUniverseGraphicsGraphicIDHandlerFunc) Handle(params GetUniverseGraphicsGraphicIDParams) middleware.Responder {
	return fn(params)
}

// GetUniverseGraphicsGraphicIDHandler interface for that can handle valid get universe graphics graphic id params
type GetUniverseGraphicsGraphicIDHandler interface {
	Handle(GetUniverseGraphicsGraphicIDParams) middleware.Responder
}

// NewGetUniverseGraphicsGraphicID creates a new http.Handler for the get universe graphics graphic id operation
func NewGetUniverseGraphicsGraphicID(ctx *middleware.Context, handler GetUniverseGraphicsGraphicIDHandler) *GetUniverseGraphicsGraphicID {
	return &GetUniverseGraphicsGraphicID{Context: ctx, Handler: handler}
}

/*GetUniverseGraphicsGraphicID swagger:route GET /universe/graphics/{graphic_id}/ Universe getUniverseGraphicsGraphicId

Get graphic information

Get information on a graphic

---

Alternate route: `/v1/universe/graphics/{graphic_id}/`

Alternate route: `/legacy/universe/graphics/{graphic_id}/`

Alternate route: `/dev/universe/graphics/{graphic_id}/`


---

This route expires daily at 11:05

*/
type GetUniverseGraphicsGraphicID struct {
	Context *middleware.Context
	Handler GetUniverseGraphicsGraphicIDHandler
}

func (o *GetUniverseGraphicsGraphicID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetUniverseGraphicsGraphicIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
