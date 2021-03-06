package character

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetCharactersCharacterIDChatChannelsHandlerFunc turns a function with the right signature into a get characters character id chat channels handler
type GetCharactersCharacterIDChatChannelsHandlerFunc func(GetCharactersCharacterIDChatChannelsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetCharactersCharacterIDChatChannelsHandlerFunc) Handle(params GetCharactersCharacterIDChatChannelsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetCharactersCharacterIDChatChannelsHandler interface for that can handle valid get characters character id chat channels params
type GetCharactersCharacterIDChatChannelsHandler interface {
	Handle(GetCharactersCharacterIDChatChannelsParams, interface{}) middleware.Responder
}

// NewGetCharactersCharacterIDChatChannels creates a new http.Handler for the get characters character id chat channels operation
func NewGetCharactersCharacterIDChatChannels(ctx *middleware.Context, handler GetCharactersCharacterIDChatChannelsHandler) *GetCharactersCharacterIDChatChannels {
	return &GetCharactersCharacterIDChatChannels{Context: ctx, Handler: handler}
}

/*GetCharactersCharacterIDChatChannels swagger:route GET /characters/{character_id}/chat_channels/ Character getCharactersCharacterIdChatChannels

Get chat channels

Return chat channels that a character is the owner or an operator of

---

Alternate route: `/v1/characters/{character_id}/chat_channels/`

Alternate route: `/legacy/characters/{character_id}/chat_channels/`

Alternate route: `/dev/characters/{character_id}/chat_channels/`


---

This route is cached for up to 300 seconds

*/
type GetCharactersCharacterIDChatChannels struct {
	Context *middleware.Context
	Handler GetCharactersCharacterIDChatChannelsHandler
}

func (o *GetCharactersCharacterIDChatChannels) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetCharactersCharacterIDChatChannelsParams()

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
