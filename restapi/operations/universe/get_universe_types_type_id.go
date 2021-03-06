package universe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetUniverseTypesTypeIDHandlerFunc turns a function with the right signature into a get universe types type id handler
type GetUniverseTypesTypeIDHandlerFunc func(GetUniverseTypesTypeIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetUniverseTypesTypeIDHandlerFunc) Handle(params GetUniverseTypesTypeIDParams) middleware.Responder {
	return fn(params)
}

// GetUniverseTypesTypeIDHandler interface for that can handle valid get universe types type id params
type GetUniverseTypesTypeIDHandler interface {
	Handle(GetUniverseTypesTypeIDParams) middleware.Responder
}

// NewGetUniverseTypesTypeID creates a new http.Handler for the get universe types type id operation
func NewGetUniverseTypesTypeID(ctx *middleware.Context, handler GetUniverseTypesTypeIDHandler) *GetUniverseTypesTypeID {
	return &GetUniverseTypesTypeID{Context: ctx, Handler: handler}
}

/*GetUniverseTypesTypeID swagger:route GET /universe/types/{type_id}/ Universe getUniverseTypesTypeId

Get type information

Get information on a type

---

Alternate route: `/v2/universe/types/{type_id}/`


---

This route expires daily at 11:05

*/
type GetUniverseTypesTypeID struct {
	Context *middleware.Context
	Handler GetUniverseTypesTypeIDHandler
}

func (o *GetUniverseTypesTypeID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetUniverseTypesTypeIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
