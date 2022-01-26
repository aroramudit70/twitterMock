// Code generated by go-swagger; DO NOT EDIT.

package operation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewGetFeedParams creates a new GetFeedParams object
//
// There are no default values defined in the spec.
func NewGetFeedParams() GetFeedParams {

	return GetFeedParams{}
}

// GetFeedParams contains all the bound params for the get feed operation
// typically these are obtained from a http.Request
//
// swagger:parameters getFeed
type GetFeedParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: header
	*/
	UserName string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetFeedParams() beforehand.
func (o *GetFeedParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := o.bindUserName(r.Header[http.CanonicalHeaderKey("userName")], true, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindUserName binds and validates parameter UserName from header.
func (o *GetFeedParams) bindUserName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("userName", "header", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("userName", "header", raw); err != nil {
		return err
	}
	o.UserName = raw

	return nil
}