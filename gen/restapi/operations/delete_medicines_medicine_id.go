// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteMedicinesMedicineIDHandlerFunc turns a function with the right signature into a delete medicines medicine ID handler
type DeleteMedicinesMedicineIDHandlerFunc func(DeleteMedicinesMedicineIDParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteMedicinesMedicineIDHandlerFunc) Handle(params DeleteMedicinesMedicineIDParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeleteMedicinesMedicineIDHandler interface for that can handle valid delete medicines medicine ID params
type DeleteMedicinesMedicineIDHandler interface {
	Handle(DeleteMedicinesMedicineIDParams, interface{}) middleware.Responder
}

// NewDeleteMedicinesMedicineID creates a new http.Handler for the delete medicines medicine ID operation
func NewDeleteMedicinesMedicineID(ctx *middleware.Context, handler DeleteMedicinesMedicineIDHandler) *DeleteMedicinesMedicineID {
	return &DeleteMedicinesMedicineID{Context: ctx, Handler: handler}
}

/*
	DeleteMedicinesMedicineID swagger:route DELETE /medicines/{medicine_id} deleteMedicinesMedicineId

Delete medicine
*/
type DeleteMedicinesMedicineID struct {
	Context *middleware.Context
	Handler DeleteMedicinesMedicineIDHandler
}

func (o *DeleteMedicinesMedicineID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteMedicinesMedicineIDParams()
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
