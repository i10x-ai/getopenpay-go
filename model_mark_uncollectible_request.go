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

// checks if the MarkUncollectibleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MarkUncollectibleRequest{}

// MarkUncollectibleRequest struct for MarkUncollectibleRequest
type MarkUncollectibleRequest struct {
	Comment NullableString `json:"comment,omitempty"`
}

// NewMarkUncollectibleRequest instantiates a new MarkUncollectibleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarkUncollectibleRequest() *MarkUncollectibleRequest {
	this := MarkUncollectibleRequest{}
	return &this
}

// NewMarkUncollectibleRequestWithDefaults instantiates a new MarkUncollectibleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarkUncollectibleRequestWithDefaults() *MarkUncollectibleRequest {
	this := MarkUncollectibleRequest{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MarkUncollectibleRequest) GetComment() string {
	if o == nil || IsNil(o.Comment.Get()) {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MarkUncollectibleRequest) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *MarkUncollectibleRequest) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *MarkUncollectibleRequest) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *MarkUncollectibleRequest) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *MarkUncollectibleRequest) UnsetComment() {
	o.Comment.Unset()
}

func (o MarkUncollectibleRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarkUncollectibleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Comment.IsSet() {
		toSerialize["comment"] = o.Comment.Get()
	}
	return toSerialize, nil
}

type NullableMarkUncollectibleRequest struct {
	value *MarkUncollectibleRequest
	isSet bool
}

func (v NullableMarkUncollectibleRequest) Get() *MarkUncollectibleRequest {
	return v.value
}

func (v *NullableMarkUncollectibleRequest) Set(val *MarkUncollectibleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMarkUncollectibleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMarkUncollectibleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarkUncollectibleRequest(val *MarkUncollectibleRequest) *NullableMarkUncollectibleRequest {
	return &NullableMarkUncollectibleRequest{value: val, isSet: true}
}

func (v NullableMarkUncollectibleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarkUncollectibleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


