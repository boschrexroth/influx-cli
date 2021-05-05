/*
 * Subset of Influx API covered by Influx CLI
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PipeLiteral Represents a specialized literal value, indicating the left hand value of a pipe expression
type PipeLiteral struct {
	// Type of AST node
	Type *string `json:"type,omitempty"`
}

// NewPipeLiteral instantiates a new PipeLiteral object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipeLiteral() *PipeLiteral {
	this := PipeLiteral{}
	return &this
}

// NewPipeLiteralWithDefaults instantiates a new PipeLiteral object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipeLiteralWithDefaults() *PipeLiteral {
	this := PipeLiteral{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PipeLiteral) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipeLiteral) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PipeLiteral) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PipeLiteral) SetType(v string) {
	o.Type = &v
}

func (o PipeLiteral) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullablePipeLiteral struct {
	value *PipeLiteral
	isSet bool
}

func (v NullablePipeLiteral) Get() *PipeLiteral {
	return v.value
}

func (v *NullablePipeLiteral) Set(val *PipeLiteral) {
	v.value = val
	v.isSet = true
}

func (v NullablePipeLiteral) IsSet() bool {
	return v.isSet
}

func (v *NullablePipeLiteral) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipeLiteral(val *PipeLiteral) *NullablePipeLiteral {
	return &NullablePipeLiteral{value: val, isSet: true}
}

func (v NullablePipeLiteral) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipeLiteral) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
