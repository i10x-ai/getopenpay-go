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

// checks if the StringConfigField type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StringConfigField{}

// StringConfigField struct for StringConfigField
type StringConfigField struct {
	DefaultValue NullableString `json:"default_value,omitempty"`
	Description NullableString `json:"description,omitempty"`
	// The key of the field in the config.
	Key string `json:"key"`
	MaxLength NullableInt32 `json:"max_length,omitempty"`
	MinLength NullableInt32 `json:"min_length,omitempty"`
	// The name of the field.
	Name string `json:"name"`
	Type *string `json:"type,omitempty"`
	Value NullableString `json:"value"`
}

type _StringConfigField StringConfigField

// NewStringConfigField instantiates a new StringConfigField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStringConfigField(key string, name string, value NullableString) *StringConfigField {
	this := StringConfigField{}
	this.Key = key
	this.Name = name
	var type_ string = "string"
	this.Type = &type_
	this.Value = value
	return &this
}

// NewStringConfigFieldWithDefaults instantiates a new StringConfigField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStringConfigFieldWithDefaults() *StringConfigField {
	this := StringConfigField{}
	var type_ string = "string"
	this.Type = &type_
	return &this
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StringConfigField) GetDefaultValue() string {
	if o == nil || IsNil(o.DefaultValue.Get()) {
		var ret string
		return ret
	}
	return *o.DefaultValue.Get()
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StringConfigField) GetDefaultValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultValue.Get(), o.DefaultValue.IsSet()
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *StringConfigField) HasDefaultValue() bool {
	if o != nil && o.DefaultValue.IsSet() {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given NullableString and assigns it to the DefaultValue field.
func (o *StringConfigField) SetDefaultValue(v string) {
	o.DefaultValue.Set(&v)
}
// SetDefaultValueNil sets the value for DefaultValue to be an explicit nil
func (o *StringConfigField) SetDefaultValueNil() {
	o.DefaultValue.Set(nil)
}

// UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
func (o *StringConfigField) UnsetDefaultValue() {
	o.DefaultValue.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StringConfigField) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StringConfigField) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *StringConfigField) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *StringConfigField) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *StringConfigField) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *StringConfigField) UnsetDescription() {
	o.Description.Unset()
}

// GetKey returns the Key field value
func (o *StringConfigField) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *StringConfigField) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *StringConfigField) SetKey(v string) {
	o.Key = v
}

// GetMaxLength returns the MaxLength field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StringConfigField) GetMaxLength() int32 {
	if o == nil || IsNil(o.MaxLength.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxLength.Get()
}

// GetMaxLengthOk returns a tuple with the MaxLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StringConfigField) GetMaxLengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxLength.Get(), o.MaxLength.IsSet()
}

// HasMaxLength returns a boolean if a field has been set.
func (o *StringConfigField) HasMaxLength() bool {
	if o != nil && o.MaxLength.IsSet() {
		return true
	}

	return false
}

// SetMaxLength gets a reference to the given NullableInt32 and assigns it to the MaxLength field.
func (o *StringConfigField) SetMaxLength(v int32) {
	o.MaxLength.Set(&v)
}
// SetMaxLengthNil sets the value for MaxLength to be an explicit nil
func (o *StringConfigField) SetMaxLengthNil() {
	o.MaxLength.Set(nil)
}

// UnsetMaxLength ensures that no value is present for MaxLength, not even an explicit nil
func (o *StringConfigField) UnsetMaxLength() {
	o.MaxLength.Unset()
}

// GetMinLength returns the MinLength field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StringConfigField) GetMinLength() int32 {
	if o == nil || IsNil(o.MinLength.Get()) {
		var ret int32
		return ret
	}
	return *o.MinLength.Get()
}

// GetMinLengthOk returns a tuple with the MinLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StringConfigField) GetMinLengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MinLength.Get(), o.MinLength.IsSet()
}

// HasMinLength returns a boolean if a field has been set.
func (o *StringConfigField) HasMinLength() bool {
	if o != nil && o.MinLength.IsSet() {
		return true
	}

	return false
}

// SetMinLength gets a reference to the given NullableInt32 and assigns it to the MinLength field.
func (o *StringConfigField) SetMinLength(v int32) {
	o.MinLength.Set(&v)
}
// SetMinLengthNil sets the value for MinLength to be an explicit nil
func (o *StringConfigField) SetMinLengthNil() {
	o.MinLength.Set(nil)
}

// UnsetMinLength ensures that no value is present for MinLength, not even an explicit nil
func (o *StringConfigField) UnsetMinLength() {
	o.MinLength.Unset()
}

// GetName returns the Name field value
func (o *StringConfigField) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *StringConfigField) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *StringConfigField) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *StringConfigField) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringConfigField) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *StringConfigField) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *StringConfigField) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StringConfigField) GetValue() string {
	if o == nil || o.Value.Get() == nil {
		var ret string
		return ret
	}

	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StringConfigField) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// SetValue sets field value
func (o *StringConfigField) SetValue(v string) {
	o.Value.Set(&v)
}

func (o StringConfigField) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StringConfigField) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DefaultValue.IsSet() {
		toSerialize["default_value"] = o.DefaultValue.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["key"] = o.Key
	if o.MaxLength.IsSet() {
		toSerialize["max_length"] = o.MaxLength.Get()
	}
	if o.MinLength.IsSet() {
		toSerialize["min_length"] = o.MinLength.Get()
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	toSerialize["value"] = o.Value.Get()
	return toSerialize, nil
}

func (o *StringConfigField) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"key",
		"name",
		"value",
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

	varStringConfigField := _StringConfigField{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStringConfigField)

	if err != nil {
		return err
	}

	*o = StringConfigField(varStringConfigField)

	return err
}

type NullableStringConfigField struct {
	value *StringConfigField
	isSet bool
}

func (v NullableStringConfigField) Get() *StringConfigField {
	return v.value
}

func (v *NullableStringConfigField) Set(val *StringConfigField) {
	v.value = val
	v.isSet = true
}

func (v NullableStringConfigField) IsSet() bool {
	return v.isSet
}

func (v *NullableStringConfigField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStringConfigField(val *StringConfigField) *NullableStringConfigField {
	return &NullableStringConfigField{value: val, isSet: true}
}

func (v NullableStringConfigField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStringConfigField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


