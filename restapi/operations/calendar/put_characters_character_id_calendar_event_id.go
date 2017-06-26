package calendar

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PutCharactersCharacterIDCalendarEventIDHandlerFunc turns a function with the right signature into a put characters character id calendar event id handler
type PutCharactersCharacterIDCalendarEventIDHandlerFunc func(PutCharactersCharacterIDCalendarEventIDParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn PutCharactersCharacterIDCalendarEventIDHandlerFunc) Handle(params PutCharactersCharacterIDCalendarEventIDParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// PutCharactersCharacterIDCalendarEventIDHandler interface for that can handle valid put characters character id calendar event id params
type PutCharactersCharacterIDCalendarEventIDHandler interface {
	Handle(PutCharactersCharacterIDCalendarEventIDParams, interface{}) middleware.Responder
}

// NewPutCharactersCharacterIDCalendarEventID creates a new http.Handler for the put characters character id calendar event id operation
func NewPutCharactersCharacterIDCalendarEventID(ctx *middleware.Context, handler PutCharactersCharacterIDCalendarEventIDHandler) *PutCharactersCharacterIDCalendarEventID {
	return &PutCharactersCharacterIDCalendarEventID{Context: ctx, Handler: handler}
}

/*PutCharactersCharacterIDCalendarEventID swagger:route PUT /characters/{character_id}/calendar/{event_id}/ Calendar putCharactersCharacterIdCalendarEventId

Respond to an event

Set your response status to an event

---

Alternate route: `/v3/characters/{character_id}/calendar/{event_id}/`

Alternate route: `/dev/characters/{character_id}/calendar/{event_id}/`


*/
type PutCharactersCharacterIDCalendarEventID struct {
	Context *middleware.Context
	Handler PutCharactersCharacterIDCalendarEventIDHandler
}

func (o *PutCharactersCharacterIDCalendarEventID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewPutCharactersCharacterIDCalendarEventIDParams()

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
