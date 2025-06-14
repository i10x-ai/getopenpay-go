/*
OpenPay API

super charge your subscription management.

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package getopenpay

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the InvoiceItemDiscountAmountsPublic type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoiceItemDiscountAmountsPublic{}

// InvoiceItemDiscountAmountsPublic struct for InvoiceItemDiscountAmountsPublic
type InvoiceItemDiscountAmountsPublic struct {
	// The amount_atom of the discount.
	AmountAtom int32 `json:"amount_atom"`
	CouponDescription NullableString `json:"coupon_description,omitempty"`
	// Name of the coupon that was applied to get this discount.
	CouponName string `json:"coupon_name"`
}

type _InvoiceItemDiscountAmountsPublic InvoiceItemDiscountAmountsPublic

// NewInvoiceItemDiscountAmountsPublic instantiates a new InvoiceItemDiscountAmountsPublic object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceItemDiscountAmountsPublic(amountAtom int32, couponName string) *InvoiceItemDiscountAmountsPublic {
	this := InvoiceItemDiscountAmountsPublic{}
	this.AmountAtom = amountAtom
	this.CouponName = couponName
	return &this
}

// NewInvoiceItemDiscountAmountsPublicWithDefaults instantiates a new InvoiceItemDiscountAmountsPublic object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceItemDiscountAmountsPublicWithDefaults() *InvoiceItemDiscountAmountsPublic {
	this := InvoiceItemDiscountAmountsPublic{}
	return &this
}

// GetAmountAtom returns the AmountAtom field value
func (o *InvoiceItemDiscountAmountsPublic) GetAmountAtom() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AmountAtom
}

// GetAmountAtomOk returns a tuple with the AmountAtom field value
// and a boolean to check if the value has been set.
func (o *InvoiceItemDiscountAmountsPublic) GetAmountAtomOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountAtom, true
}

// SetAmountAtom sets field value
func (o *InvoiceItemDiscountAmountsPublic) SetAmountAtom(v int32) {
	o.AmountAtom = v
}

// GetCouponDescription returns the CouponDescription field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceItemDiscountAmountsPublic) GetCouponDescription() string {
	if o == nil || IsNil(o.CouponDescription.Get()) {
		var ret string
		return ret
	}
	return *o.CouponDescription.Get()
}

// GetCouponDescriptionOk returns a tuple with the CouponDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceItemDiscountAmountsPublic) GetCouponDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CouponDescription.Get(), o.CouponDescription.IsSet()
}

// HasCouponDescription returns a boolean if a field has been set.
func (o *InvoiceItemDiscountAmountsPublic) HasCouponDescription() bool {
	if o != nil && o.CouponDescription.IsSet() {
		return true
	}

	return false
}

// SetCouponDescription gets a reference to the given NullableString and assigns it to the CouponDescription field.
func (o *InvoiceItemDiscountAmountsPublic) SetCouponDescription(v string) {
	o.CouponDescription.Set(&v)
}
// SetCouponDescriptionNil sets the value for CouponDescription to be an explicit nil
func (o *InvoiceItemDiscountAmountsPublic) SetCouponDescriptionNil() {
	o.CouponDescription.Set(nil)
}

// UnsetCouponDescription ensures that no value is present for CouponDescription, not even an explicit nil
func (o *InvoiceItemDiscountAmountsPublic) UnsetCouponDescription() {
	o.CouponDescription.Unset()
}

// GetCouponName returns the CouponName field value
func (o *InvoiceItemDiscountAmountsPublic) GetCouponName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CouponName
}

// GetCouponNameOk returns a tuple with the CouponName field value
// and a boolean to check if the value has been set.
func (o *InvoiceItemDiscountAmountsPublic) GetCouponNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CouponName, true
}

// SetCouponName sets field value
func (o *InvoiceItemDiscountAmountsPublic) SetCouponName(v string) {
	o.CouponName = v
}

func (o InvoiceItemDiscountAmountsPublic) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceItemDiscountAmountsPublic) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount_atom"] = o.AmountAtom
	if o.CouponDescription.IsSet() {
		toSerialize["coupon_description"] = o.CouponDescription.Get()
	}
	toSerialize["coupon_name"] = o.CouponName
	return toSerialize, nil
}

func (o *InvoiceItemDiscountAmountsPublic) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amount_atom",
		"coupon_name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varInvoiceItemDiscountAmountsPublic := _InvoiceItemDiscountAmountsPublic{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInvoiceItemDiscountAmountsPublic)

	if err != nil {
		return err
	}

	*o = InvoiceItemDiscountAmountsPublic(varInvoiceItemDiscountAmountsPublic)

	return err
}

type NullableInvoiceItemDiscountAmountsPublic struct {
	value *InvoiceItemDiscountAmountsPublic
	isSet bool
}

func (v NullableInvoiceItemDiscountAmountsPublic) Get() *InvoiceItemDiscountAmountsPublic {
	return v.value
}

func (v *NullableInvoiceItemDiscountAmountsPublic) Set(val *InvoiceItemDiscountAmountsPublic) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceItemDiscountAmountsPublic) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceItemDiscountAmountsPublic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceItemDiscountAmountsPublic(val *InvoiceItemDiscountAmountsPublic) *NullableInvoiceItemDiscountAmountsPublic {
	return &NullableInvoiceItemDiscountAmountsPublic{value: val, isSet: true}
}

func (v NullableInvoiceItemDiscountAmountsPublic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceItemDiscountAmountsPublic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


