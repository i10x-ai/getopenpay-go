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

// checks if the NonPciIntegration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NonPciIntegration{}

// NonPciIntegration struct for NonPciIntegration
type NonPciIntegration struct {
	// The account ID
	AccountId string `json:"account_id"`
	// The configuration fields
	Configuration []NonPciIntegrationConfigurationInner `json:"configuration,omitempty"`
	// The integration ID
	IntegrationId string `json:"integration_id"`
	// The integration type
	IntegrationType NonPciIntegrationEnum `json:"integration_type"`
}

type _NonPciIntegration NonPciIntegration

// NewNonPciIntegration instantiates a new NonPciIntegration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNonPciIntegration(accountId string, integrationId string, integrationType NonPciIntegrationEnum) *NonPciIntegration {
	this := NonPciIntegration{}
	this.AccountId = accountId
	this.IntegrationId = integrationId
	this.IntegrationType = integrationType
	return &this
}

// NewNonPciIntegrationWithDefaults instantiates a new NonPciIntegration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNonPciIntegrationWithDefaults() *NonPciIntegration {
	this := NonPciIntegration{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *NonPciIntegration) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *NonPciIntegration) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *NonPciIntegration) SetAccountId(v string) {
	o.AccountId = v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *NonPciIntegration) GetConfiguration() []NonPciIntegrationConfigurationInner {
	if o == nil || IsNil(o.Configuration) {
		var ret []NonPciIntegrationConfigurationInner
		return ret
	}
	return o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonPciIntegration) GetConfigurationOk() ([]NonPciIntegrationConfigurationInner, bool) {
	if o == nil || IsNil(o.Configuration) {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *NonPciIntegration) HasConfiguration() bool {
	if o != nil && !IsNil(o.Configuration) {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given []NonPciIntegrationConfigurationInner and assigns it to the Configuration field.
func (o *NonPciIntegration) SetConfiguration(v []NonPciIntegrationConfigurationInner) {
	o.Configuration = v
}

// GetIntegrationId returns the IntegrationId field value
func (o *NonPciIntegration) GetIntegrationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IntegrationId
}

// GetIntegrationIdOk returns a tuple with the IntegrationId field value
// and a boolean to check if the value has been set.
func (o *NonPciIntegration) GetIntegrationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IntegrationId, true
}

// SetIntegrationId sets field value
func (o *NonPciIntegration) SetIntegrationId(v string) {
	o.IntegrationId = v
}

// GetIntegrationType returns the IntegrationType field value
func (o *NonPciIntegration) GetIntegrationType() NonPciIntegrationEnum {
	if o == nil {
		var ret NonPciIntegrationEnum
		return ret
	}

	return o.IntegrationType
}

// GetIntegrationTypeOk returns a tuple with the IntegrationType field value
// and a boolean to check if the value has been set.
func (o *NonPciIntegration) GetIntegrationTypeOk() (*NonPciIntegrationEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IntegrationType, true
}

// SetIntegrationType sets field value
func (o *NonPciIntegration) SetIntegrationType(v NonPciIntegrationEnum) {
	o.IntegrationType = v
}

func (o NonPciIntegration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NonPciIntegration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account_id"] = o.AccountId
	if !IsNil(o.Configuration) {
		toSerialize["configuration"] = o.Configuration
	}
	toSerialize["integration_id"] = o.IntegrationId
	toSerialize["integration_type"] = o.IntegrationType
	return toSerialize, nil
}

func (o *NonPciIntegration) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"account_id",
		"integration_id",
		"integration_type",
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

	varNonPciIntegration := _NonPciIntegration{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNonPciIntegration)

	if err != nil {
		return err
	}

	*o = NonPciIntegration(varNonPciIntegration)

	return err
}

type NullableNonPciIntegration struct {
	value *NonPciIntegration
	isSet bool
}

func (v NullableNonPciIntegration) Get() *NonPciIntegration {
	return v.value
}

func (v *NullableNonPciIntegration) Set(val *NonPciIntegration) {
	v.value = val
	v.isSet = true
}

func (v NullableNonPciIntegration) IsSet() bool {
	return v.isSet
}

func (v *NullableNonPciIntegration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNonPciIntegration(val *NonPciIntegration) *NullableNonPciIntegration {
	return &NullableNonPciIntegration{value: val, isSet: true}
}

func (v NullableNonPciIntegration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNonPciIntegration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


