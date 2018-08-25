// Code generated by go-swagger; DO NOT EDIT.

package swagmodel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Cart cart
// swagger:model cart
type Cart struct {

	// cart ID
	// Required: true
	CartID *int64 `json:"cartID"`

	// item count
	// Required: true
	ItemCount *string `json:"itemCount"`

	// total price
	// Required: true
	TotalPrice *string `json:"totalPrice"`

	// updated at
	// Required: true
	UpdatedAt *string `json:"updatedAt"`
}

// Validate validates this cart
func (m *Cart) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCartID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItemCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalPrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Cart) validateCartID(formats strfmt.Registry) error {

	if err := validate.Required("cartID", "body", m.CartID); err != nil {
		return err
	}

	return nil
}

func (m *Cart) validateItemCount(formats strfmt.Registry) error {

	if err := validate.Required("itemCount", "body", m.ItemCount); err != nil {
		return err
	}

	return nil
}

func (m *Cart) validateTotalPrice(formats strfmt.Registry) error {

	if err := validate.Required("totalPrice", "body", m.TotalPrice); err != nil {
		return err
	}

	return nil
}

func (m *Cart) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updatedAt", "body", m.UpdatedAt); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Cart) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Cart) UnmarshalBinary(b []byte) error {
	var res Cart
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}