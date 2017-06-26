package industry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetCharactersCharacterIDIndustryJobsHandlerFunc turns a function with the right signature into a get characters character id industry jobs handler
type GetCharactersCharacterIDIndustryJobsHandlerFunc func(GetCharactersCharacterIDIndustryJobsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetCharactersCharacterIDIndustryJobsHandlerFunc) Handle(params GetCharactersCharacterIDIndustryJobsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetCharactersCharacterIDIndustryJobsHandler interface for that can handle valid get characters character id industry jobs params
type GetCharactersCharacterIDIndustryJobsHandler interface {
	Handle(GetCharactersCharacterIDIndustryJobsParams, interface{}) middleware.Responder
}

// NewGetCharactersCharacterIDIndustryJobs creates a new http.Handler for the get characters character id industry jobs operation
func NewGetCharactersCharacterIDIndustryJobs(ctx *middleware.Context, handler GetCharactersCharacterIDIndustryJobsHandler) *GetCharactersCharacterIDIndustryJobs {
	return &GetCharactersCharacterIDIndustryJobs{Context: ctx, Handler: handler}
}

/*GetCharactersCharacterIDIndustryJobs swagger:route GET /characters/{character_id}/industry/jobs/ Industry getCharactersCharacterIdIndustryJobs

List character industry jobs

List industry jobs placed by a character

---

Alternate route: `/v1/characters/{character_id}/industry/jobs/`

Alternate route: `/legacy/characters/{character_id}/industry/jobs/`

Alternate route: `/dev/characters/{character_id}/industry/jobs/`


---

This route is cached for up to 300 seconds

*/
type GetCharactersCharacterIDIndustryJobs struct {
	Context *middleware.Context
	Handler GetCharactersCharacterIDIndustryJobsHandler
}

func (o *GetCharactersCharacterIDIndustryJobs) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetCharactersCharacterIDIndustryJobsParams()

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
