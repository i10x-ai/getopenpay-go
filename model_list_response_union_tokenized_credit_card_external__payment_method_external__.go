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

// checks if the ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternal{}

// ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternal struct for ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternal
type ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternal struct {
	Data []ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner `json:"data"`
	PageNumber int32 `json:"page_number"`
	PageSize int32 `json:"page_size"`
	TotalObjects int32 `json:"total_objects"`
}

type _ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternal ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternal

// NewListResponseUnionTokenizedCreditCardExternalPaymentMethodExternal instantiates a new ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListResponseUnionTokenizedCreditCardExternalPaymentMethodExternal(data []ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner, pageNumber int32, pageSize int32, totalObjects int32) *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternal {
	this := ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternal{}
	this.Data = data
	this.PageNumber = pageNumber
	this.PageSize = pageSize
	this.TotalObjects = totalObjects
	return &this
}

// NewListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalWithDefaults instantiates a new ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalWithDefaults() *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternal {
	this := ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternal{}
	return &this
}

// GetData returns the Data field value
func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternal) GetData() []ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner {
	if o == nil {
		var ret []ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternal) GetDataOk() ([]ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternal) SetData(v []ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternalDataInner) {
	o.Data = v
}

// GetPageNumber returns the PageNumber field value
func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternal) GetPageNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PageNumber
}

// GetPageNumberOk returns a tuple with the PageNumber field value
// and a boolean to check if the value has been set.
func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternal) GetPageNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageNumber, true
}

// SetPageNumber sets field value
func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternal) SetPageNumber(v int32) {
	o.PageNumber = v
}

// GetPageSize returns the PageSize field value
func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternal) GetPageSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value
// and a boolean to check if the value has been set.
func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternal) GetPageSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageSize, true
}

// SetPageSize sets field value
func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternal) SetPageSize(v int32) {
	o.PageSize = v
}

// GetTotalObjects returns the TotalObjects field value
func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternal) GetTotalObjects() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalObjects
}

// GetTotalObjectsOk returns a tuple with the TotalObjects field value
// and a boolean to check if the value has been set.
func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternal) GetTotalObjectsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalObjects, true
}

// SetTotalObjects sets field value
func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternal) SetTotalObjects(v int32) {
	o.TotalObjects = v
}

func (o ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternal) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["page_number"] = o.PageNumber
	toSerialize["page_size"] = o.PageSize
	toSerialize["total_objects"] = o.TotalObjects
	return toSerialize, nil
}

func (o *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternal) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"page_number",
		"page_size",
		"total_objects",
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

	varListResponseUnionTokenizedCreditCardExternalPaymentMethodExternal := _ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternal{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListResponseUnionTokenizedCreditCardExternalPaymentMethodExternal)

	if err != nil {
		return err
	}

	*o = ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternal(varListResponseUnionTokenizedCreditCardExternalPaymentMethodExternal)

	return err
}

type NullableListResponseUnionTokenizedCreditCardExternalPaymentMethodExternal struct {
	value *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternal
	isSet bool
}

func (v NullableListResponseUnionTokenizedCreditCardExternalPaymentMethodExternal) Get() *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternal {
	return v.value
}

func (v *NullableListResponseUnionTokenizedCreditCardExternalPaymentMethodExternal) Set(val *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternal) {
	v.value = val
	v.isSet = true
}

func (v NullableListResponseUnionTokenizedCreditCardExternalPaymentMethodExternal) IsSet() bool {
	return v.isSet
}

func (v *NullableListResponseUnionTokenizedCreditCardExternalPaymentMethodExternal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListResponseUnionTokenizedCreditCardExternalPaymentMethodExternal(val *ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternal) *NullableListResponseUnionTokenizedCreditCardExternalPaymentMethodExternal {
	return &NullableListResponseUnionTokenizedCreditCardExternalPaymentMethodExternal{value: val, isSet: true}
}

func (v NullableListResponseUnionTokenizedCreditCardExternalPaymentMethodExternal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListResponseUnionTokenizedCreditCardExternalPaymentMethodExternal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


