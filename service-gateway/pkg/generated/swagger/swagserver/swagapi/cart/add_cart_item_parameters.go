// Code generated by go-swagger; DO NOT EDIT.

package cart

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
)

// NewAddCartItemParams creates a new AddCartItemParams object
// no default values defined in spec.
func NewAddCartItemParams() AddCartItemParams {

	return AddCartItemParams{}
}

// AddCartItemParams contains all the bound params for the add cart item operation
// typically these are obtained from a http.Request
//
// swagger:parameters addCartItem
type AddCartItemParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewAddCartItemParams() beforehand.
func (o *AddCartItemParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}