package universe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetUniverseRacesHandlerFunc turns a function with the right signature into a get universe races handler
type GetUniverseRacesHandlerFunc func(GetUniverseRacesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetUniverseRacesHandlerFunc) Handle(params GetUniverseRacesParams) middleware.Responder {
	return fn(params)
}

// GetUniverseRacesHandler interface for that can handle valid get universe races params
type GetUniverseRacesHandler interface {
	Handle(GetUniverseRacesParams) middleware.Responder
}

// NewGetUniverseRaces creates a new http.Handler for the get universe races operation
func NewGetUniverseRaces(ctx *middleware.Context, handler GetUniverseRacesHandler) *GetUniverseRaces {
	return &GetUniverseRaces{Context: ctx, Handler: handler}
}

/*GetUniverseRaces swagger:route GET /universe/races/ Universe getUniverseRaces

Get character races

Get a list of character races

---

Alternate route: `/v1/universe/races/`

Alternate route: `/legacy/universe/races/`

Alternate route: `/dev/universe/races/`


---

This route expires daily at 11:05

*/
type GetUniverseRaces struct {
	Context *middleware.Context
	Handler GetUniverseRacesHandler
}

func (o *GetUniverseRaces) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetUniverseRacesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
