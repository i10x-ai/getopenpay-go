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

// checks if the PriceTierParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PriceTierParams{}

// PriceTierParams struct for PriceTierParams
type PriceTierParams struct {
	FlatAmountAtom NullableInt32 `json:"flat_amount_atom,omitempty"`
	// The price for the unit in the smallest denomination.It is in atomic units (in USD this is cents).
	UnitAmountAtom int32 `json:"unit_amount_atom"`
	// The starting unit for the price tier.
	UnitsFrom int32 `json:"units_from"`
	UnitsUpto NullableInt32 `json:"units_upto,omitempty"`
}

type _PriceTierParams PriceTierParams

// NewPriceTierParams instantiates a new PriceTierParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceTierParams(unitAmountAtom int32, unitsFrom int32) *PriceTierParams {
	this := PriceTierParams{}
	this.UnitAmountAtom = unitAmountAtom
	this.UnitsFrom = unitsFrom
	return &this
}

// NewPriceTierParamsWithDefaults instantiates a new PriceTierParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceTierParamsWithDefaults() *PriceTierParams {
	this := PriceTierParams{}
	return &this
}

// GetFlatAmountAtom returns the FlatAmountAtom field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PriceTierParams) GetFlatAmountAtom() int32 {
	if o == nil || IsNil(o.FlatAmountAtom.Get()) {
		var ret int32
		return ret
	}
	return *o.FlatAmountAtom.Get()
}

// GetFlatAmountAtomOk returns a tuple with the FlatAmountAtom field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PriceTierParams) GetFlatAmountAtomOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.FlatAmountAtom.Get(), o.FlatAmountAtom.IsSet()
}

// HasFlatAmountAtom returns a boolean if a field has been set.
func (o *PriceTierParams) HasFlatAmountAtom() bool {
	if o != nil && o.FlatAmountAtom.IsSet() {
		return true
	}

	return false
}

// SetFlatAmountAtom gets a reference to the given NullableInt32 and assigns it to the FlatAmountAtom field.
func (o *PriceTierParams) SetFlatAmountAtom(v int32) {
	o.FlatAmountAtom.Set(&v)
}
// SetFlatAmountAtomNil sets the value for FlatAmountAtom to be an explicit nil
func (o *PriceTierParams) SetFlatAmountAtomNil() {
	o.FlatAmountAtom.Set(nil)
}

// UnsetFlatAmountAtom ensures that no value is present for FlatAmountAtom, not even an explicit nil
func (o *PriceTierParams) UnsetFlatAmountAtom() {
	o.FlatAmountAtom.Unset()
}

// GetUnitAmountAtom returns the UnitAmountAtom field value
func (o *PriceTierParams) GetUnitAmountAtom() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UnitAmountAtom
}

// GetUnitAmountAtomOk returns a tuple with the UnitAmountAtom field value
// and a boolean to check if the value has been set.
func (o *PriceTierParams) GetUnitAmountAtomOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnitAmountAtom, true
}

// SetUnitAmountAtom sets field value
func (o *PriceTierParams) SetUnitAmountAtom(v int32) {
	o.UnitAmountAtom = v
}

// GetUnitsFrom returns the UnitsFrom field value
func (o *PriceTierParams) GetUnitsFrom() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UnitsFrom
}

// GetUnitsFromOk returns a tuple with the UnitsFrom field value
// and a boolean to check if the value has been set.
func (o *PriceTierParams) GetUnitsFromOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnitsFrom, true
}

// SetUnitsFrom sets field value
func (o *PriceTierParams) SetUnitsFrom(v int32) {
	o.UnitsFrom = v
}

// GetUnitsUpto returns the UnitsUpto field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PriceTierParams) GetUnitsUpto() int32 {
	if o == nil || IsNil(o.UnitsUpto.Get()) {
		var ret int32
		return ret
	}
	return *o.UnitsUpto.Get()
}

// GetUnitsUptoOk returns a tuple with the UnitsUpto field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PriceTierParams) GetUnitsUptoOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.UnitsUpto.Get(), o.UnitsUpto.IsSet()
}

// HasUnitsUpto returns a boolean if a field has been set.
func (o *PriceTierParams) HasUnitsUpto() bool {
	if o != nil && o.UnitsUpto.IsSet() {
		return true
	}

	return false
}

// SetUnitsUpto gets a reference to the given NullableInt32 and assigns it to the UnitsUpto field.
func (o *PriceTierParams) SetUnitsUpto(v int32) {
	o.UnitsUpto.Set(&v)
}
// SetUnitsUptoNil sets the value for UnitsUpto to be an explicit nil
func (o *PriceTierParams) SetUnitsUptoNil() {
	o.UnitsUpto.Set(nil)
}

// UnsetUnitsUpto ensures that no value is present for UnitsUpto, not even an explicit nil
func (o *PriceTierParams) UnsetUnitsUpto() {
	o.UnitsUpto.Unset()
}

func (o PriceTierParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PriceTierParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.FlatAmountAtom.IsSet() {
		toSerialize["flat_amount_atom"] = o.FlatAmountAtom.Get()
	}
	toSerialize["unit_amount_atom"] = o.UnitAmountAtom
	toSerialize["units_from"] = o.UnitsFrom
	if o.UnitsUpto.IsSet() {
		toSerialize["units_upto"] = o.UnitsUpto.Get()
	}
	return toSerialize, nil
}

func (o *PriceTierParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"unit_amount_atom",
		"units_from",
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

	varPriceTierParams := _PriceTierParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPriceTierParams)

	if err != nil {
		return err
	}

	*o = PriceTierParams(varPriceTierParams)

	return err
}

type NullablePriceTierParams struct {
	value *PriceTierParams
	isSet bool
}

func (v NullablePriceTierParams) Get() *PriceTierParams {
	return v.value
}

func (v *NullablePriceTierParams) Set(val *PriceTierParams) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceTierParams) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceTierParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceTierParams(val *PriceTierParams) *NullablePriceTierParams {
	return &NullablePriceTierParams{value: val, isSet: true}
}

func (v NullablePriceTierParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceTierParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


