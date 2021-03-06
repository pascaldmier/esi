package market

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetMarketsPricesHandlerFunc turns a function with the right signature into a get markets prices handler
type GetMarketsPricesHandlerFunc func(GetMarketsPricesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetMarketsPricesHandlerFunc) Handle(params GetMarketsPricesParams) middleware.Responder {
	return fn(params)
}

// GetMarketsPricesHandler interface for that can handle valid get markets prices params
type GetMarketsPricesHandler interface {
	Handle(GetMarketsPricesParams) middleware.Responder
}

// NewGetMarketsPrices creates a new http.Handler for the get markets prices operation
func NewGetMarketsPrices(ctx *middleware.Context, handler GetMarketsPricesHandler) *GetMarketsPrices {
	return &GetMarketsPrices{Context: ctx, Handler: handler}
}

/*GetMarketsPrices swagger:route GET /markets/prices/ Market getMarketsPrices

List market prices

Return a list of prices

---

Alternate route: `/v1/markets/prices/`

Alternate route: `/legacy/markets/prices/`

Alternate route: `/dev/markets/prices/`


---

This route is cached for up to 3600 seconds

*/
type GetMarketsPrices struct {
	Context *middleware.Context
	Handler GetMarketsPricesHandler
}

func (o *GetMarketsPrices) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetMarketsPricesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
