// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostMedicinesHandlerFunc turns a function with the right signature into a post medicines handler
type PostMedicinesHandlerFunc func(PostMedicinesParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn PostMedicinesHandlerFunc) Handle(params PostMedicinesParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// PostMedicinesHandler interface for that can handle valid post medicines params
type PostMedicinesHandler interface {
	Handle(PostMedicinesParams, interface{}) middleware.Responder
}

// NewPostMedicines creates a new http.Handler for the post medicines operation
func NewPostMedicines(ctx *middleware.Context, handler PostMedicinesHandler) *PostMedicines {
	return &PostMedicines{Context: ctx, Handler: handler}
}

/*
	PostMedicines swagger:route POST /medicines postMedicines

Add medicine
*/
type PostMedicines struct {
	Context *middleware.Context
	Handler PostMedicinesHandler
}

func (o *PostMedicines) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostMedicinesParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
