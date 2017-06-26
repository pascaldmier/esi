package corporation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetCorporationsCorporationIDIconsHandlerFunc turns a function with the right signature into a get corporations corporation id icons handler
type GetCorporationsCorporationIDIconsHandlerFunc func(GetCorporationsCorporationIDIconsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetCorporationsCorporationIDIconsHandlerFunc) Handle(params GetCorporationsCorporationIDIconsParams) middleware.Responder {
	return fn(params)
}

// GetCorporationsCorporationIDIconsHandler interface for that can handle valid get corporations corporation id icons params
type GetCorporationsCorporationIDIconsHandler interface {
	Handle(GetCorporationsCorporationIDIconsParams) middleware.Responder
}

// NewGetCorporationsCorporationIDIcons creates a new http.Handler for the get corporations corporation id icons operation
func NewGetCorporationsCorporationIDIcons(ctx *middleware.Context, handler GetCorporationsCorporationIDIconsHandler) *GetCorporationsCorporationIDIcons {
	return &GetCorporationsCorporationIDIcons{Context: ctx, Handler: handler}
}

/*GetCorporationsCorporationIDIcons swagger:route GET /corporations/{corporation_id}/icons/ Corporation getCorporationsCorporationIdIcons

Get corporation icon

Get the icon urls for a corporation

---

Alternate route: `/v1/corporations/{corporation_id}/icons/`

Alternate route: `/legacy/corporations/{corporation_id}/icons/`

Alternate route: `/dev/corporations/{corporation_id}/icons/`


---

This route is cached for up to 3600 seconds

*/
type GetCorporationsCorporationIDIcons struct {
	Context *middleware.Context
	Handler GetCorporationsCorporationIDIconsHandler
}

func (o *GetCorporationsCorporationIDIcons) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetCorporationsCorporationIDIconsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}