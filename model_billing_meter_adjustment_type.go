/*
OpenPay API

super charge your subscription management.

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package getopenpay

import (
	"encoding/json"
	"fmt"
)

// BillingMeterAdjustmentType the model 'BillingMeterAdjustmentType'
type BillingMeterAdjustmentType string

// List of BillingMeterAdjustmentType
const (
	BILLINGMETERADJUSTMENTTYPE_CANCEL BillingMeterAdjustmentType = "cancel"
)

// All allowed values of BillingMeterAdjustmentType enum
var AllowedBillingMeterAdjustmentTypeEnumValues = []BillingMeterAdjustmentType{
	"cancel",
}

func (v *BillingMeterAdjustmentType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BillingMeterAdjustmentType(value)
	for _, existing := range AllowedBillingMeterAdjustmentTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BillingMeterAdjustmentType", value)
}

// NewBillingMeterAdjustmentTypeFromValue returns a pointer to a valid BillingMeterAdjustmentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBillingMeterAdjustmentTypeFromValue(v string) (*BillingMeterAdjustmentType, error) {
	ev := BillingMeterAdjustmentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BillingMeterAdjustmentType: valid values are %v", v, AllowedBillingMeterAdjustmentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BillingMeterAdjustmentType) IsValid() bool {
	for _, existing := range AllowedBillingMeterAdjustmentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BillingMeterAdjustmentType value
func (v BillingMeterAdjustmentType) Ptr() *BillingMeterAdjustmentType {
	return &v
}

type NullableBillingMeterAdjustmentType struct {
	value *BillingMeterAdjustmentType
	isSet bool
}

func (v NullableBillingMeterAdjustmentType) Get() *BillingMeterAdjustmentType {
	return v.value
}

func (v *NullableBillingMeterAdjustmentType) Set(val *BillingMeterAdjustmentType) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingMeterAdjustmentType) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingMeterAdjustmentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingMeterAdjustmentType(val *BillingMeterAdjustmentType) *NullableBillingMeterAdjustmentType {
	return &NullableBillingMeterAdjustmentType{value: val, isSet: true}
}

func (v NullableBillingMeterAdjustmentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingMeterAdjustmentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

