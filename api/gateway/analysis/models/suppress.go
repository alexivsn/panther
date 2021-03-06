// Code generated by go-swagger; DO NOT EDIT.

// Panther is a scalable, powerful, cloud-native SIEM written in Golang/React.
// Copyright (C) 2020 Panther Labs Inc
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Suppress suppress
//
// swagger:model Suppress
type Suppress struct {

	// policy ids
	// Required: true
	// Min Items: 1
	PolicyIds []ID `json:"policyIds"`

	// resource patterns
	// Required: true
	ResourcePatterns Suppressions `json:"resourcePatterns"`
}

// Validate validates this suppress
func (m *Suppress) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePolicyIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourcePatterns(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Suppress) validatePolicyIds(formats strfmt.Registry) error {

	if err := validate.Required("policyIds", "body", m.PolicyIds); err != nil {
		return err
	}

	iPolicyIdsSize := int64(len(m.PolicyIds))

	if err := validate.MinItems("policyIds", "body", iPolicyIdsSize, 1); err != nil {
		return err
	}

	for i := 0; i < len(m.PolicyIds); i++ {

		if err := m.PolicyIds[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("policyIds" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *Suppress) validateResourcePatterns(formats strfmt.Registry) error {

	if err := validate.Required("resourcePatterns", "body", m.ResourcePatterns); err != nil {
		return err
	}

	if err := m.ResourcePatterns.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("resourcePatterns")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Suppress) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Suppress) UnmarshalBinary(b []byte) error {
	var res Suppress
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
