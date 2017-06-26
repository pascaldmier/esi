package mail

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostCharactersCharacterIDMailLabelsHandlerFunc turns a function with the right signature into a post characters character id mail labels handler
type PostCharactersCharacterIDMailLabelsHandlerFunc func(PostCharactersCharacterIDMailLabelsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn PostCharactersCharacterIDMailLabelsHandlerFunc) Handle(params PostCharactersCharacterIDMailLabelsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// PostCharactersCharacterIDMailLabelsHandler interface for that can handle valid post characters character id mail labels params
type PostCharactersCharacterIDMailLabelsHandler interface {
	Handle(PostCharactersCharacterIDMailLabelsParams, interface{}) middleware.Responder
}

// NewPostCharactersCharacterIDMailLabels creates a new http.Handler for the post characters character id mail labels operation
func NewPostCharactersCharacterIDMailLabels(ctx *middleware.Context, handler PostCharactersCharacterIDMailLabelsHandler) *PostCharactersCharacterIDMailLabels {
	return &PostCharactersCharacterIDMailLabels{Context: ctx, Handler: handler}
}

/*PostCharactersCharacterIDMailLabels swagger:route POST /characters/{character_id}/mail/labels/ Mail postCharactersCharacterIdMailLabels

Create a mail label

Create a mail label

---

Alternate route: `/v2/characters/{character_id}/mail/labels/`

Alternate route: `/legacy/characters/{character_id}/mail/labels/`

Alternate route: `/dev/characters/{character_id}/mail/labels/`


*/
type PostCharactersCharacterIDMailLabels struct {
	Context *middleware.Context
	Handler PostCharactersCharacterIDMailLabelsHandler
}

func (o *PostCharactersCharacterIDMailLabels) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewPostCharactersCharacterIDMailLabelsParams()

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