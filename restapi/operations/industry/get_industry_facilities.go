package industry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetIndustryFacilitiesHandlerFunc turns a function with the right signature into a get industry facilities handler
type GetIndustryFacilitiesHandlerFunc func(GetIndustryFacilitiesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetIndustryFacilitiesHandlerFunc) Handle(params GetIndustryFacilitiesParams) middleware.Responder {
	return fn(params)
}

// GetIndustryFacilitiesHandler interface for that can handle valid get industry facilities params
type GetIndustryFacilitiesHandler interface {
	Handle(GetIndustryFacilitiesParams) middleware.Responder
}

// NewGetIndustryFacilities creates a new http.Handler for the get industry facilities operation
func NewGetIndustryFacilities(ctx *middleware.Context, handler GetIndustryFacilitiesHandler) *GetIndustryFacilities {
	return &GetIndustryFacilities{Context: ctx, Handler: handler}
}

/*GetIndustryFacilities swagger:route GET /industry/facilities/ Industry getIndustryFacilities

List industry facilities

Return a list of industry facilities

---

Alternate route: `/v1/industry/facilities/`

Alternate route: `/legacy/industry/facilities/`

Alternate route: `/dev/industry/facilities/`


---

This route is cached for up to 3600 seconds

*/
type GetIndustryFacilities struct {
	Context *middleware.Context
	Handler GetIndustryFacilitiesHandler
}

func (o *GetIndustryFacilities) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetIndustryFacilitiesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
