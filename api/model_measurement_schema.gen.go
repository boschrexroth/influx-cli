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
	"time"
)

// MeasurementSchema Definition of a measurement schema.
type MeasurementSchema struct {
	Id string `json:"id" yaml:"id"`
	// The ID of the organization.
	OrgID *string `json:"orgID,omitempty" yaml:"orgID,omitempty"`
	// The ID of the bucket that the measurement schema is associated with.
	BucketID *string `json:"bucketID,omitempty" yaml:"bucketID,omitempty"`
	Name     string  `json:"name" yaml:"name"`
	// Ordered collection of column definitions.
	Columns   []MeasurementSchemaColumn `json:"columns" yaml:"columns"`
	CreatedAt time.Time                 `json:"createdAt" yaml:"createdAt"`
	UpdatedAt time.Time                 `json:"updatedAt" yaml:"updatedAt"`
}

// NewMeasurementSchema instantiates a new MeasurementSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMeasurementSchema(id string, name string, columns []MeasurementSchemaColumn, createdAt time.Time, updatedAt time.Time) *MeasurementSchema {
	this := MeasurementSchema{}
	this.Id = id
	this.Name = name
	this.Columns = columns
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewMeasurementSchemaWithDefaults instantiates a new MeasurementSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMeasurementSchemaWithDefaults() *MeasurementSchema {
	this := MeasurementSchema{}
	return &this
}

// GetId returns the Id field value
func (o *MeasurementSchema) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MeasurementSchema) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MeasurementSchema) SetId(v string) {
	o.Id = v
}

// GetOrgID returns the OrgID field value if set, zero value otherwise.
func (o *MeasurementSchema) GetOrgID() string {
	if o == nil || o.OrgID == nil {
		var ret string
		return ret
	}
	return *o.OrgID
}

// GetOrgIDOk returns a tuple with the OrgID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeasurementSchema) GetOrgIDOk() (*string, bool) {
	if o == nil || o.OrgID == nil {
		return nil, false
	}
	return o.OrgID, true
}

// HasOrgID returns a boolean if a field has been set.
func (o *MeasurementSchema) HasOrgID() bool {
	if o != nil && o.OrgID != nil {
		return true
	}

	return false
}

// SetOrgID gets a reference to the given string and assigns it to the OrgID field.
func (o *MeasurementSchema) SetOrgID(v string) {
	o.OrgID = &v
}

// GetBucketID returns the BucketID field value if set, zero value otherwise.
func (o *MeasurementSchema) GetBucketID() string {
	if o == nil || o.BucketID == nil {
		var ret string
		return ret
	}
	return *o.BucketID
}

// GetBucketIDOk returns a tuple with the BucketID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeasurementSchema) GetBucketIDOk() (*string, bool) {
	if o == nil || o.BucketID == nil {
		return nil, false
	}
	return o.BucketID, true
}

// HasBucketID returns a boolean if a field has been set.
func (o *MeasurementSchema) HasBucketID() bool {
	if o != nil && o.BucketID != nil {
		return true
	}

	return false
}

// SetBucketID gets a reference to the given string and assigns it to the BucketID field.
func (o *MeasurementSchema) SetBucketID(v string) {
	o.BucketID = &v
}

// GetName returns the Name field value
func (o *MeasurementSchema) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MeasurementSchema) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MeasurementSchema) SetName(v string) {
	o.Name = v
}

// GetColumns returns the Columns field value
func (o *MeasurementSchema) GetColumns() []MeasurementSchemaColumn {
	if o == nil {
		var ret []MeasurementSchemaColumn
		return ret
	}

	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value
// and a boolean to check if the value has been set.
func (o *MeasurementSchema) GetColumnsOk() (*[]MeasurementSchemaColumn, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Columns, true
}

// SetColumns sets field value
func (o *MeasurementSchema) SetColumns(v []MeasurementSchemaColumn) {
	o.Columns = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *MeasurementSchema) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *MeasurementSchema) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *MeasurementSchema) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *MeasurementSchema) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *MeasurementSchema) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *MeasurementSchema) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o MeasurementSchema) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.OrgID != nil {
		toSerialize["orgID"] = o.OrgID
	}
	if o.BucketID != nil {
		toSerialize["bucketID"] = o.BucketID
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["columns"] = o.Columns
	}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if true {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableMeasurementSchema struct {
	value *MeasurementSchema
	isSet bool
}

func (v NullableMeasurementSchema) Get() *MeasurementSchema {
	return v.value
}

func (v *NullableMeasurementSchema) Set(val *MeasurementSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableMeasurementSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableMeasurementSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMeasurementSchema(val *MeasurementSchema) *NullableMeasurementSchema {
	return &NullableMeasurementSchema{value: val, isSet: true}
}

func (v NullableMeasurementSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMeasurementSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
