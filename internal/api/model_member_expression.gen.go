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

// MemberExpression Represents accessing a property of an object
type MemberExpression struct {
	// Type of AST node
	Type     *string      `json:"type,omitempty"`
	Object   *Expression  `json:"object,omitempty"`
	Property *PropertyKey `json:"property,omitempty"`
}

// NewMemberExpression instantiates a new MemberExpression object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMemberExpression() *MemberExpression {
	this := MemberExpression{}
	return &this
}

// NewMemberExpressionWithDefaults instantiates a new MemberExpression object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemberExpressionWithDefaults() *MemberExpression {
	this := MemberExpression{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MemberExpression) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberExpression) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MemberExpression) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *MemberExpression) SetType(v string) {
	o.Type = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *MemberExpression) GetObject() Expression {
	if o == nil || o.Object == nil {
		var ret Expression
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberExpression) GetObjectOk() (*Expression, bool) {
	if o == nil || o.Object == nil {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *MemberExpression) HasObject() bool {
	if o != nil && o.Object != nil {
		return true
	}

	return false
}

// SetObject gets a reference to the given Expression and assigns it to the Object field.
func (o *MemberExpression) SetObject(v Expression) {
	o.Object = &v
}

// GetProperty returns the Property field value if set, zero value otherwise.
func (o *MemberExpression) GetProperty() PropertyKey {
	if o == nil || o.Property == nil {
		var ret PropertyKey
		return ret
	}
	return *o.Property
}

// GetPropertyOk returns a tuple with the Property field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberExpression) GetPropertyOk() (*PropertyKey, bool) {
	if o == nil || o.Property == nil {
		return nil, false
	}
	return o.Property, true
}

// HasProperty returns a boolean if a field has been set.
func (o *MemberExpression) HasProperty() bool {
	if o != nil && o.Property != nil {
		return true
	}

	return false
}

// SetProperty gets a reference to the given PropertyKey and assigns it to the Property field.
func (o *MemberExpression) SetProperty(v PropertyKey) {
	o.Property = &v
}

func (o MemberExpression) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Object != nil {
		toSerialize["object"] = o.Object
	}
	if o.Property != nil {
		toSerialize["property"] = o.Property
	}
	return json.Marshal(toSerialize)
}

type NullableMemberExpression struct {
	value *MemberExpression
	isSet bool
}

func (v NullableMemberExpression) Get() *MemberExpression {
	return v.value
}

func (v *NullableMemberExpression) Set(val *MemberExpression) {
	v.value = val
	v.isSet = true
}

func (v NullableMemberExpression) IsSet() bool {
	return v.isSet
}

func (v *NullableMemberExpression) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemberExpression(val *MemberExpression) *NullableMemberExpression {
	return &NullableMemberExpression{value: val, isSet: true}
}

func (v NullableMemberExpression) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemberExpression) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
