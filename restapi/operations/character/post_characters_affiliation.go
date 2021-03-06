package character

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostCharactersAffiliationHandlerFunc turns a function with the right signature into a post characters affiliation handler
type PostCharactersAffiliationHandlerFunc func(PostCharactersAffiliationParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostCharactersAffiliationHandlerFunc) Handle(params PostCharactersAffiliationParams) middleware.Responder {
	return fn(params)
}

// PostCharactersAffiliationHandler interface for that can handle valid post characters affiliation params
type PostCharactersAffiliationHandler interface {
	Handle(PostCharactersAffiliationParams) middleware.Responder
}

// NewPostCharactersAffiliation creates a new http.Handler for the post characters affiliation operation
func NewPostCharactersAffiliation(ctx *middleware.Context, handler PostCharactersAffiliationHandler) *PostCharactersAffiliation {
	return &PostCharactersAffiliation{Context: ctx, Handler: handler}
}

/*PostCharactersAffiliation swagger:route POST /characters/affiliation/ Character postCharactersAffiliation

Character affiliation

Bulk lookup of character IDs to corporation, alliance and faction

---

Alternate route: `/v1/characters/affiliation/`

Alternate route: `/legacy/characters/affiliation/`

Alternate route: `/dev/characters/affiliation/`


---

This route is cached for up to 3600 seconds

*/
type PostCharactersAffiliation struct {
	Context *middleware.Context
	Handler PostCharactersAffiliationHandler
}

func (o *PostCharactersAffiliation) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewPostCharactersAffiliationParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
