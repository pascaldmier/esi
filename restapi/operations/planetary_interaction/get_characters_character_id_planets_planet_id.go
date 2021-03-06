package planetary_interaction

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetCharactersCharacterIDPlanetsPlanetIDHandlerFunc turns a function with the right signature into a get characters character id planets planet id handler
type GetCharactersCharacterIDPlanetsPlanetIDHandlerFunc func(GetCharactersCharacterIDPlanetsPlanetIDParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetCharactersCharacterIDPlanetsPlanetIDHandlerFunc) Handle(params GetCharactersCharacterIDPlanetsPlanetIDParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetCharactersCharacterIDPlanetsPlanetIDHandler interface for that can handle valid get characters character id planets planet id params
type GetCharactersCharacterIDPlanetsPlanetIDHandler interface {
	Handle(GetCharactersCharacterIDPlanetsPlanetIDParams, interface{}) middleware.Responder
}

// NewGetCharactersCharacterIDPlanetsPlanetID creates a new http.Handler for the get characters character id planets planet id operation
func NewGetCharactersCharacterIDPlanetsPlanetID(ctx *middleware.Context, handler GetCharactersCharacterIDPlanetsPlanetIDHandler) *GetCharactersCharacterIDPlanetsPlanetID {
	return &GetCharactersCharacterIDPlanetsPlanetID{Context: ctx, Handler: handler}
}

/*GetCharactersCharacterIDPlanetsPlanetID swagger:route GET /characters/{character_id}/planets/{planet_id}/ Planetary Interaction getCharactersCharacterIdPlanetsPlanetId

Get colony layout

Returns full details on the layout of a single planetary colony, including links, pins and routes. Note: Planetary information is only recalculated when the colony is viewed through the client. Information will not update until this criteria is met.

---

Alternate route: `/v2/characters/{character_id}/planets/{planet_id}/`

Alternate route: `/dev/characters/{character_id}/planets/{planet_id}/`


---

This route is cached for up to 600 seconds

*/
type GetCharactersCharacterIDPlanetsPlanetID struct {
	Context *middleware.Context
	Handler GetCharactersCharacterIDPlanetsPlanetIDHandler
}

func (o *GetCharactersCharacterIDPlanetsPlanetID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetCharactersCharacterIDPlanetsPlanetIDParams()

	uprinc, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
