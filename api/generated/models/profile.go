// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Profile profile
//
// swagger:model Profile
type Profile struct {

	// email Id
	// Example: abx@xyz.com
	EmailID string `json:"emailId,omitempty"`

	// following list
	// Read Only: true
	FollowingList []string `json:"followingList"`

	// name
	// Required: true
	Name *string `json:"name"`

	// password
	// Required: true
	Password *string `json:"password"`

	// Unique Name to indentify a user. Also termed as handle
	// Required: true
	UserName *string `json:"userName"`
}

// Validate validates this profile
func (m *Profile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Profile) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Profile) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	return nil
}

func (m *Profile) validateUserName(formats strfmt.Registry) error {

	if err := validate.Required("userName", "body", m.UserName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this profile based on the context it is used
func (m *Profile) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFollowingList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Profile) contextValidateFollowingList(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "followingList", "body", []string(m.FollowingList)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Profile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Profile) UnmarshalBinary(b []byte) error {
	var res Profile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}