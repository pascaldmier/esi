package character

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetCharactersCharacterIDPortraitHandlerFunc turns a function with the right signature into a get characters character id portrait handler
type GetCharactersCharacterIDPortraitHandlerFunc func(GetCharactersCharacterIDPortraitParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetCharactersCharacterIDPortraitHandlerFunc) Handle(params GetCharactersCharacterIDPortraitParams) middleware.Responder {
	return fn(params)
}

// GetCharactersCharacterIDPortraitHandler interface for that can handle valid get characters character id portrait params
type GetCharactersCharacterIDPortraitHandler interface {
	Handle(GetCharactersCharacterIDPortraitParams) middleware.Responder
}

// NewGetCharactersCharacterIDPortrait creates a new http.Handler for the get characters character id portrait operation
func NewGetCharactersCharacterIDPortrait(ctx *middleware.Context, handler GetCharactersCharacterIDPortraitHandler) *GetCharactersCharacterIDPortrait {
	return &GetCharactersCharacterIDPortrait{Context: ctx, Handler: handler}
}

/*GetCharactersCharacterIDPortrait swagger:route GET /characters/{character_id}/portrait/ Character getCharactersCharacterIdPortrait

Get character portraits

Get portrait urls for a character

---

Alternate route: `/v2/characters/{character_id}/portrait/`

Alternate route: `/dev/characters/{character_id}/portrait/`


---

This route is cached for up to 3600 seconds

*/
type GetCharactersCharacterIDPortrait struct {
	Context *middleware.Context
	Handler GetCharactersCharacterIDPortraitHandler
}

func (o *GetCharactersCharacterIDPortrait) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetCharactersCharacterIDPortraitParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
