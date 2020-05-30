// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MovieImdb movie imdb
//
// swagger:model movieImdb
type MovieImdb struct {

	// id
	ID string `json:"id,omitempty"`

	// rating
	Rating int64 `json:"rating,omitempty"`

	// votes
	Votes int64 `json:"votes,omitempty"`
}

// Validate validates this movie imdb
func (m *MovieImdb) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MovieImdb) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MovieImdb) UnmarshalBinary(b []byte) error {
	var res MovieImdb
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}