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

// checks if the ListResponsePromotionCodeExternal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListResponsePromotionCodeExternal{}

// ListResponsePromotionCodeExternal struct for ListResponsePromotionCodeExternal
type ListResponsePromotionCodeExternal struct {
	Data []PromotionCodeExternal `json:"data"`
	PageNumber int32 `json:"page_number"`
	PageSize int32 `json:"page_size"`
	TotalObjects int32 `json:"total_objects"`
}

type _ListResponsePromotionCodeExternal ListResponsePromotionCodeExternal

// NewListResponsePromotionCodeExternal instantiates a new ListResponsePromotionCodeExternal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListResponsePromotionCodeExternal(data []PromotionCodeExternal, pageNumber int32, pageSize int32, totalObjects int32) *ListResponsePromotionCodeExternal {
	this := ListResponsePromotionCodeExternal{}
	this.Data = data
	this.PageNumber = pageNumber
	this.PageSize = pageSize
	this.TotalObjects = totalObjects
	return &this
}

// NewListResponsePromotionCodeExternalWithDefaults instantiates a new ListResponsePromotionCodeExternal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListResponsePromotionCodeExternalWithDefaults() *ListResponsePromotionCodeExternal {
	this := ListResponsePromotionCodeExternal{}
	return &this
}

// GetData returns the Data field value
func (o *ListResponsePromotionCodeExternal) GetData() []PromotionCodeExternal {
	if o == nil {
		var ret []PromotionCodeExternal
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListResponsePromotionCodeExternal) GetDataOk() ([]PromotionCodeExternal, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ListResponsePromotionCodeExternal) SetData(v []PromotionCodeExternal) {
	o.Data = v
}

// GetPageNumber returns the PageNumber field value
func (o *ListResponsePromotionCodeExternal) GetPageNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PageNumber
}

// GetPageNumberOk returns a tuple with the PageNumber field value
// and a boolean to check if the value has been set.
func (o *ListResponsePromotionCodeExternal) GetPageNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageNumber, true
}

// SetPageNumber sets field value
func (o *ListResponsePromotionCodeExternal) SetPageNumber(v int32) {
	o.PageNumber = v
}

// GetPageSize returns the PageSize field value
func (o *ListResponsePromotionCodeExternal) GetPageSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value
// and a boolean to check if the value has been set.
func (o *ListResponsePromotionCodeExternal) GetPageSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageSize, true
}

// SetPageSize sets field value
func (o *ListResponsePromotionCodeExternal) SetPageSize(v int32) {
	o.PageSize = v
}

// GetTotalObjects returns the TotalObjects field value
func (o *ListResponsePromotionCodeExternal) GetTotalObjects() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalObjects
}

// GetTotalObjectsOk returns a tuple with the TotalObjects field value
// and a boolean to check if the value has been set.
func (o *ListResponsePromotionCodeExternal) GetTotalObjectsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalObjects, true
}

// SetTotalObjects sets field value
func (o *ListResponsePromotionCodeExternal) SetTotalObjects(v int32) {
	o.TotalObjects = v
}

func (o ListResponsePromotionCodeExternal) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListResponsePromotionCodeExternal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["page_number"] = o.PageNumber
	toSerialize["page_size"] = o.PageSize
	toSerialize["total_objects"] = o.TotalObjects
	return toSerialize, nil
}

func (o *ListResponsePromotionCodeExternal) UnmarshalJSON(data []byte) (err error) {
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

	varListResponsePromotionCodeExternal := _ListResponsePromotionCodeExternal{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListResponsePromotionCodeExternal)

	if err != nil {
		return err
	}

	*o = ListResponsePromotionCodeExternal(varListResponsePromotionCodeExternal)

	return err
}

type NullableListResponsePromotionCodeExternal struct {
	value *ListResponsePromotionCodeExternal
	isSet bool
}

func (v NullableListResponsePromotionCodeExternal) Get() *ListResponsePromotionCodeExternal {
	return v.value
}

func (v *NullableListResponsePromotionCodeExternal) Set(val *ListResponsePromotionCodeExternal) {
	v.value = val
	v.isSet = true
}

func (v NullableListResponsePromotionCodeExternal) IsSet() bool {
	return v.isSet
}

func (v *NullableListResponsePromotionCodeExternal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListResponsePromotionCodeExternal(val *ListResponsePromotionCodeExternal) *NullableListResponsePromotionCodeExternal {
	return &NullableListResponsePromotionCodeExternal{value: val, isSet: true}
}

func (v NullableListResponsePromotionCodeExternal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListResponsePromotionCodeExternal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


