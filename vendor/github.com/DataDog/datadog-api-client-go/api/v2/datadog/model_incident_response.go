/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// IncidentResponse Response with an incident.
type IncidentResponse struct {
	Data IncidentResponseData `json:"data"`
	// Included related resources that the user requested.
	Included *[]IncidentResponseIncludedItem `json:"included,omitempty"`
}

// NewIncidentResponse instantiates a new IncidentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncidentResponse(data IncidentResponseData) *IncidentResponse {
	this := IncidentResponse{}
	this.Data = data
	return &this
}

// NewIncidentResponseWithDefaults instantiates a new IncidentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncidentResponseWithDefaults() *IncidentResponse {
	this := IncidentResponse{}
	return &this
}

// GetData returns the Data field value
func (o *IncidentResponse) GetData() IncidentResponseData {
	if o == nil {
		var ret IncidentResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *IncidentResponse) GetDataOk() (*IncidentResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *IncidentResponse) SetData(v IncidentResponseData) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *IncidentResponse) GetIncluded() []IncidentResponseIncludedItem {
	if o == nil || o.Included == nil {
		var ret []IncidentResponseIncludedItem
		return ret
	}
	return *o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentResponse) GetIncludedOk() (*[]IncidentResponseIncludedItem, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *IncidentResponse) HasIncluded() bool {
	if o != nil && o.Included != nil {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []IncidentResponseIncludedItem and assigns it to the Included field.
func (o *IncidentResponse) SetIncluded(v []IncidentResponseIncludedItem) {
	o.Included = &v
}

func (o IncidentResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if o.Included != nil {
		toSerialize["included"] = o.Included
	}
	return json.Marshal(toSerialize)
}

type NullableIncidentResponse struct {
	value *IncidentResponse
	isSet bool
}

func (v NullableIncidentResponse) Get() *IncidentResponse {
	return v.value
}

func (v *NullableIncidentResponse) Set(val *IncidentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIncidentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIncidentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncidentResponse(val *IncidentResponse) *NullableIncidentResponse {
	return &NullableIncidentResponse{value: val, isSet: true}
}

func (v NullableIncidentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncidentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
