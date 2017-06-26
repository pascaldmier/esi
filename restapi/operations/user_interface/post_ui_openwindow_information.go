package user_interface

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostUIOpenwindowInformationHandlerFunc turns a function with the right signature into a post ui openwindow information handler
type PostUIOpenwindowInformationHandlerFunc func(PostUIOpenwindowInformationParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn PostUIOpenwindowInformationHandlerFunc) Handle(params PostUIOpenwindowInformationParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// PostUIOpenwindowInformationHandler interface for that can handle valid post ui openwindow information params
type PostUIOpenwindowInformationHandler interface {
	Handle(PostUIOpenwindowInformationParams, interface{}) middleware.Responder
}

// NewPostUIOpenwindowInformation creates a new http.Handler for the post ui openwindow information operation
func NewPostUIOpenwindowInformation(ctx *middleware.Context, handler PostUIOpenwindowInformationHandler) *PostUIOpenwindowInformation {
	return &PostUIOpenwindowInformation{Context: ctx, Handler: handler}
}

/*PostUIOpenwindowInformation swagger:route POST /ui/openwindow/information/ User Interface postUiOpenwindowInformation

Open Information Window

Open the information window for a character, corporation or alliance inside the client

---

Alternate route: `/v1/ui/openwindow/information/`

Alternate route: `/legacy/ui/openwindow/information/`

Alternate route: `/dev/ui/openwindow/information/`


*/
type PostUIOpenwindowInformation struct {
	Context *middleware.Context
	Handler PostUIOpenwindowInformationHandler
}

func (o *PostUIOpenwindowInformation) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewPostUIOpenwindowInformationParams()

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
