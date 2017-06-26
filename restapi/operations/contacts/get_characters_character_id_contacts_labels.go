package contacts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetCharactersCharacterIDContactsLabelsHandlerFunc turns a function with the right signature into a get characters character id contacts labels handler
type GetCharactersCharacterIDContactsLabelsHandlerFunc func(GetCharactersCharacterIDContactsLabelsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetCharactersCharacterIDContactsLabelsHandlerFunc) Handle(params GetCharactersCharacterIDContactsLabelsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetCharactersCharacterIDContactsLabelsHandler interface for that can handle valid get characters character id contacts labels params
type GetCharactersCharacterIDContactsLabelsHandler interface {
	Handle(GetCharactersCharacterIDContactsLabelsParams, interface{}) middleware.Responder
}

// NewGetCharactersCharacterIDContactsLabels creates a new http.Handler for the get characters character id contacts labels operation
func NewGetCharactersCharacterIDContactsLabels(ctx *middleware.Context, handler GetCharactersCharacterIDContactsLabelsHandler) *GetCharactersCharacterIDContactsLabels {
	return &GetCharactersCharacterIDContactsLabels{Context: ctx, Handler: handler}
}

/*GetCharactersCharacterIDContactsLabels swagger:route GET /characters/{character_id}/contacts/labels/ Contacts getCharactersCharacterIdContactsLabels

Get contact labels

Return custom labels for contacts the character defined

---

Alternate route: `/v1/characters/{character_id}/contacts/labels/`

Alternate route: `/legacy/characters/{character_id}/contacts/labels/`

Alternate route: `/dev/characters/{character_id}/contacts/labels/`


---

This route is cached for up to 300 seconds

*/
type GetCharactersCharacterIDContactsLabels struct {
	Context *middleware.Context
	Handler GetCharactersCharacterIDContactsLabelsHandler
}

func (o *GetCharactersCharacterIDContactsLabels) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetCharactersCharacterIDContactsLabelsParams()

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
