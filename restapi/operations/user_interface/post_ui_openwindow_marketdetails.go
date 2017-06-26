package user_interface

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostUIOpenwindowMarketdetailsHandlerFunc turns a function with the right signature into a post ui openwindow marketdetails handler
type PostUIOpenwindowMarketdetailsHandlerFunc func(PostUIOpenwindowMarketdetailsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn PostUIOpenwindowMarketdetailsHandlerFunc) Handle(params PostUIOpenwindowMarketdetailsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// PostUIOpenwindowMarketdetailsHandler interface for that can handle valid post ui openwindow marketdetails params
type PostUIOpenwindowMarketdetailsHandler interface {
	Handle(PostUIOpenwindowMarketdetailsParams, interface{}) middleware.Responder
}

// NewPostUIOpenwindowMarketdetails creates a new http.Handler for the post ui openwindow marketdetails operation
func NewPostUIOpenwindowMarketdetails(ctx *middleware.Context, handler PostUIOpenwindowMarketdetailsHandler) *PostUIOpenwindowMarketdetails {
	return &PostUIOpenwindowMarketdetails{Context: ctx, Handler: handler}
}

/*PostUIOpenwindowMarketdetails swagger:route POST /ui/openwindow/marketdetails/ User Interface postUiOpenwindowMarketdetails

Open Market Details

Open the market details window for a specific typeID inside the client

---

Alternate route: `/v1/ui/openwindow/marketdetails/`

Alternate route: `/legacy/ui/openwindow/marketdetails/`

Alternate route: `/dev/ui/openwindow/marketdetails/`


*/
type PostUIOpenwindowMarketdetails struct {
	Context *middleware.Context
	Handler PostUIOpenwindowMarketdetailsHandler
}

func (o *PostUIOpenwindowMarketdetails) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewPostUIOpenwindowMarketdetailsParams()

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