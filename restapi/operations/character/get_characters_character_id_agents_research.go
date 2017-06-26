package character

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetCharactersCharacterIDAgentsResearchHandlerFunc turns a function with the right signature into a get characters character id agents research handler
type GetCharactersCharacterIDAgentsResearchHandlerFunc func(GetCharactersCharacterIDAgentsResearchParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetCharactersCharacterIDAgentsResearchHandlerFunc) Handle(params GetCharactersCharacterIDAgentsResearchParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetCharactersCharacterIDAgentsResearchHandler interface for that can handle valid get characters character id agents research params
type GetCharactersCharacterIDAgentsResearchHandler interface {
	Handle(GetCharactersCharacterIDAgentsResearchParams, interface{}) middleware.Responder
}

// NewGetCharactersCharacterIDAgentsResearch creates a new http.Handler for the get characters character id agents research operation
func NewGetCharactersCharacterIDAgentsResearch(ctx *middleware.Context, handler GetCharactersCharacterIDAgentsResearchHandler) *GetCharactersCharacterIDAgentsResearch {
	return &GetCharactersCharacterIDAgentsResearch{Context: ctx, Handler: handler}
}

/*GetCharactersCharacterIDAgentsResearch swagger:route GET /characters/{character_id}/agents_research/ Character getCharactersCharacterIdAgentsResearch

Get agents research

Return a list of agents research information for a character. The formula for finding the current research points with an agent is: currentPoints = remainderPoints + pointsPerDay * days(currentTime - researchStartDate)

---

Alternate route: `/v1/characters/{character_id}/agents_research/`

Alternate route: `/legacy/characters/{character_id}/agents_research/`

Alternate route: `/dev/characters/{character_id}/agents_research/`


---

This route is cached for up to 3600 seconds

*/
type GetCharactersCharacterIDAgentsResearch struct {
	Context *middleware.Context
	Handler GetCharactersCharacterIDAgentsResearchHandler
}

func (o *GetCharactersCharacterIDAgentsResearch) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetCharactersCharacterIDAgentsResearchParams()

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
