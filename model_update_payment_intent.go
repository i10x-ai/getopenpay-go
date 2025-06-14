/*
OpenPay API

super charge your subscription management.

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package getopenpay

import (
	"encoding/json"
)

// checks if the UpdatePaymentIntent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdatePaymentIntent{}

// UpdatePaymentIntent struct for UpdatePaymentIntent
type UpdatePaymentIntent struct {
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
}

// NewUpdatePaymentIntent instantiates a new UpdatePaymentIntent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatePaymentIntent() *UpdatePaymentIntent {
	this := UpdatePaymentIntent{}
	return &this
}

// NewUpdatePaymentIntentWithDefaults instantiates a new UpdatePaymentIntent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatePaymentIntentWithDefaults() *UpdatePaymentIntent {
	this := UpdatePaymentIntent{}
	return &this
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdatePaymentIntent) GetCustomFields() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdatePaymentIntent) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *UpdatePaymentIntent) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *UpdatePaymentIntent) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o UpdatePaymentIntent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdatePaymentIntent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CustomFields != nil {
		toSerialize["custom_fields"] = o.CustomFields
	}
	return toSerialize, nil
}

type NullableUpdatePaymentIntent struct {
	value *UpdatePaymentIntent
	isSet bool
}

func (v NullableUpdatePaymentIntent) Get() *UpdatePaymentIntent {
	return v.value
}

func (v *NullableUpdatePaymentIntent) Set(val *UpdatePaymentIntent) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatePaymentIntent) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatePaymentIntent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatePaymentIntent(val *UpdatePaymentIntent) *NullableUpdatePaymentIntent {
	return &NullableUpdatePaymentIntent{value: val, isSet: true}
}

func (v NullableUpdatePaymentIntent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdatePaymentIntent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


