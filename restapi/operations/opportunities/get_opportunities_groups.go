package opportunities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetOpportunitiesGroupsHandlerFunc turns a function with the right signature into a get opportunities groups handler
type GetOpportunitiesGroupsHandlerFunc func(GetOpportunitiesGroupsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetOpportunitiesGroupsHandlerFunc) Handle(params GetOpportunitiesGroupsParams) middleware.Responder {
	return fn(params)
}

// GetOpportunitiesGroupsHandler interface for that can handle valid get opportunities groups params
type GetOpportunitiesGroupsHandler interface {
	Handle(GetOpportunitiesGroupsParams) middleware.Responder
}

// NewGetOpportunitiesGroups creates a new http.Handler for the get opportunities groups operation
func NewGetOpportunitiesGroups(ctx *middleware.Context, handler GetOpportunitiesGroupsHandler) *GetOpportunitiesGroups {
	return &GetOpportunitiesGroups{Context: ctx, Handler: handler}
}

/*GetOpportunitiesGroups swagger:route GET /opportunities/groups/ Opportunities getOpportunitiesGroups

Get opportunities groups

Return a list of opportunities groups

---

Alternate route: `/v1/opportunities/groups/`

Alternate route: `/legacy/opportunities/groups/`

Alternate route: `/dev/opportunities/groups/`


---

This route expires daily at 11:05

*/
type GetOpportunitiesGroups struct {
	Context *middleware.Context
	Handler GetOpportunitiesGroupsHandler
}

func (o *GetOpportunitiesGroups) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetOpportunitiesGroupsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
