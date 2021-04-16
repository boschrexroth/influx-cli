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
	"fmt"
)

// ErrorCode code is the machine-readable error code.
type ErrorCode string

// List of ErrorCode
const (
	ERRORCODE_INTERNAL_ERROR       ErrorCode = "internal error"
	ERRORCODE_NOT_FOUND            ErrorCode = "not found"
	ERRORCODE_CONFLICT             ErrorCode = "conflict"
	ERRORCODE_INVALID              ErrorCode = "invalid"
	ERRORCODE_UNPROCESSABLE_ENTITY ErrorCode = "unprocessable entity"
	ERRORCODE_EMPTY_VALUE          ErrorCode = "empty value"
	ERRORCODE_UNAVAILABLE          ErrorCode = "unavailable"
	ERRORCODE_FORBIDDEN            ErrorCode = "forbidden"
	ERRORCODE_TOO_MANY_REQUESTS    ErrorCode = "too many requests"
	ERRORCODE_UNAUTHORIZED         ErrorCode = "unauthorized"
	ERRORCODE_METHOD_NOT_ALLOWED   ErrorCode = "method not allowed"
	ERRORCODE_REQUEST_TOO_LARGE    ErrorCode = "request too large"
)

func (v *ErrorCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ErrorCode(value)
	for _, existing := range []ErrorCode{"internal error", "not found", "conflict", "invalid", "unprocessable entity", "empty value", "unavailable", "forbidden", "too many requests", "unauthorized", "method not allowed", "request too large"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ErrorCode", value)
}

// Ptr returns reference to ErrorCode value
func (v ErrorCode) Ptr() *ErrorCode {
	return &v
}

type NullableErrorCode struct {
	value *ErrorCode
	isSet bool
}

func (v NullableErrorCode) Get() *ErrorCode {
	return v.value
}

func (v *NullableErrorCode) Set(val *ErrorCode) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorCode) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorCode(val *ErrorCode) *NullableErrorCode {
	return &NullableErrorCode{value: val, isSet: true}
}

func (v NullableErrorCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
