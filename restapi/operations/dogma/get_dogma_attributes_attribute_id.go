package dogma

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetDogmaAttributesAttributeIDHandlerFunc turns a function with the right signature into a get dogma attributes attribute id handler
type GetDogmaAttributesAttributeIDHandlerFunc func(GetDogmaAttributesAttributeIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetDogmaAttributesAttributeIDHandlerFunc) Handle(params GetDogmaAttributesAttributeIDParams) middleware.Responder {
	return fn(params)
}

// GetDogmaAttributesAttributeIDHandler interface for that can handle valid get dogma attributes attribute id params
type GetDogmaAttributesAttributeIDHandler interface {
	Handle(GetDogmaAttributesAttributeIDParams) middleware.Responder
}

// NewGetDogmaAttributesAttributeID creates a new http.Handler for the get dogma attributes attribute id operation
func NewGetDogmaAttributesAttributeID(ctx *middleware.Context, handler GetDogmaAttributesAttributeIDHandler) *GetDogmaAttributesAttributeID {
	return &GetDogmaAttributesAttributeID{Context: ctx, Handler: handler}
}

/*GetDogmaAttributesAttributeID swagger:route GET /dogma/attributes/{attribute_id}/ Dogma getDogmaAttributesAttributeId

Get attribute information

Get information on a dogma attribute

---

Alternate route: `/v1/dogma/attributes/{attribute_id}/`

Alternate route: `/legacy/dogma/attributes/{attribute_id}/`

Alternate route: `/dev/dogma/attributes/{attribute_id}/`


---

This route expires daily at 11:05

*/
type GetDogmaAttributesAttributeID struct {
	Context *middleware.Context
	Handler GetDogmaAttributesAttributeIDHandler
}

func (o *GetDogmaAttributesAttributeID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetDogmaAttributesAttributeIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
