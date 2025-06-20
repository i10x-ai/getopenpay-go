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


// ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner struct for ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner
type ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner struct {
	PaymentMethodExternal *PaymentMethodExternal
	TokenizedCreditCardExternal *TokenizedCreditCardExternal
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into PaymentMethodExternal
	err = json.Unmarshal(data, &dst.PaymentMethodExternal);
	if err == nil {
		jsonPaymentMethodExternal, _ := json.Marshal(dst.PaymentMethodExternal)
		if string(jsonPaymentMethodExternal) == "{}" { // empty struct
			dst.PaymentMethodExternal = nil
		} else {
			return nil // data stored in dst.PaymentMethodExternal, return on the first match
		}
	} else {
		dst.PaymentMethodExternal = nil
	}

	// try to unmarshal JSON data into TokenizedCreditCardExternal
	err = json.Unmarshal(data, &dst.TokenizedCreditCardExternal);
	if err == nil {
		jsonTokenizedCreditCardExternal, _ := json.Marshal(dst.TokenizedCreditCardExternal)
		if string(jsonTokenizedCreditCardExternal) == "{}" { // empty struct
			dst.TokenizedCreditCardExternal = nil
		} else {
			return nil // data stored in dst.TokenizedCreditCardExternal, return on the first match
		}
	} else {
		dst.TokenizedCreditCardExternal = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) MarshalJSON() ([]byte, error) {
	if src.PaymentMethodExternal != nil {
		return json.Marshal(&src.PaymentMethodExternal)
	}

	if src.TokenizedCreditCardExternal != nil {
		return json.Marshal(&src.TokenizedCreditCardExternal)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner struct {
	value *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner
	isSet bool
}

func (v NullableListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) Get() *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner {
	return v.value
}

func (v *NullableListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) Set(val *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner(val *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) *NullableListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner {
	return &NullableListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner{value: val, isSet: true}
}

func (v NullableListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


