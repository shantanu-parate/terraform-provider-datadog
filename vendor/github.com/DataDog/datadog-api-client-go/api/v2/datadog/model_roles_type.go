/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"fmt"
)

// RolesType Roles type.
type RolesType string

// List of RolesType
const (
	ROLESTYPE_ROLES RolesType = "roles"
)

func (v *RolesType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RolesType(value)
	for _, existing := range []RolesType{"roles"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RolesType", *v)
}

// Ptr returns reference to RolesType value
func (v RolesType) Ptr() *RolesType {
	return &v
}

type NullableRolesType struct {
	value *RolesType
	isSet bool
}

func (v NullableRolesType) Get() *RolesType {
	return v.value
}

func (v *NullableRolesType) Set(val *RolesType) {
	v.value = val
	v.isSet = true
}

func (v NullableRolesType) IsSet() bool {
	return v.isSet
}

func (v *NullableRolesType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRolesType(val *RolesType) *NullableRolesType {
	return &NullableRolesType{value: val, isSet: true}
}

func (v NullableRolesType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRolesType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}