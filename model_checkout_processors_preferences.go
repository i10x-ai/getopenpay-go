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

// checks if the CheckoutProcessorsPreferences type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CheckoutProcessorsPreferences{}

// CheckoutProcessorsPreferences Allows customization of which payment processors are available during each checkout session.
type CheckoutProcessorsPreferences struct {
	// If not empty, the enabled providers will be filtered by this list.
	FilterProviders []PaymentProviderType `json:"filter_providers,omitempty"`
	// If not empty, only the processors with these ids will be accepted.
	IdsWhitelist []string `json:"ids_whitelist,omitempty"`
	// If not empty, only the processors with these names will be accepted. Valid values are: adyen, airwallex, authorize_net, checkout_com, braintree, stripe, foobar, pockyt, cybersource, loop, paypal, nmi.
	NamesWhitelist []string `json:"names_whitelist,omitempty"`
}

// NewCheckoutProcessorsPreferences instantiates a new CheckoutProcessorsPreferences object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckoutProcessorsPreferences() *CheckoutProcessorsPreferences {
	this := CheckoutProcessorsPreferences{}
	return &this
}

// NewCheckoutProcessorsPreferencesWithDefaults instantiates a new CheckoutProcessorsPreferences object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckoutProcessorsPreferencesWithDefaults() *CheckoutProcessorsPreferences {
	this := CheckoutProcessorsPreferences{}
	return &this
}

// GetFilterProviders returns the FilterProviders field value if set, zero value otherwise.
func (o *CheckoutProcessorsPreferences) GetFilterProviders() []PaymentProviderType {
	if o == nil || IsNil(o.FilterProviders) {
		var ret []PaymentProviderType
		return ret
	}
	return o.FilterProviders
}

// GetFilterProvidersOk returns a tuple with the FilterProviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutProcessorsPreferences) GetFilterProvidersOk() ([]PaymentProviderType, bool) {
	if o == nil || IsNil(o.FilterProviders) {
		return nil, false
	}
	return o.FilterProviders, true
}

// HasFilterProviders returns a boolean if a field has been set.
func (o *CheckoutProcessorsPreferences) HasFilterProviders() bool {
	if o != nil && !IsNil(o.FilterProviders) {
		return true
	}

	return false
}

// SetFilterProviders gets a reference to the given []PaymentProviderType and assigns it to the FilterProviders field.
func (o *CheckoutProcessorsPreferences) SetFilterProviders(v []PaymentProviderType) {
	o.FilterProviders = v
}

// GetIdsWhitelist returns the IdsWhitelist field value if set, zero value otherwise.
func (o *CheckoutProcessorsPreferences) GetIdsWhitelist() []string {
	if o == nil || IsNil(o.IdsWhitelist) {
		var ret []string
		return ret
	}
	return o.IdsWhitelist
}

// GetIdsWhitelistOk returns a tuple with the IdsWhitelist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutProcessorsPreferences) GetIdsWhitelistOk() ([]string, bool) {
	if o == nil || IsNil(o.IdsWhitelist) {
		return nil, false
	}
	return o.IdsWhitelist, true
}

// HasIdsWhitelist returns a boolean if a field has been set.
func (o *CheckoutProcessorsPreferences) HasIdsWhitelist() bool {
	if o != nil && !IsNil(o.IdsWhitelist) {
		return true
	}

	return false
}

// SetIdsWhitelist gets a reference to the given []string and assigns it to the IdsWhitelist field.
func (o *CheckoutProcessorsPreferences) SetIdsWhitelist(v []string) {
	o.IdsWhitelist = v
}

// GetNamesWhitelist returns the NamesWhitelist field value if set, zero value otherwise.
func (o *CheckoutProcessorsPreferences) GetNamesWhitelist() []string {
	if o == nil || IsNil(o.NamesWhitelist) {
		var ret []string
		return ret
	}
	return o.NamesWhitelist
}

// GetNamesWhitelistOk returns a tuple with the NamesWhitelist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutProcessorsPreferences) GetNamesWhitelistOk() ([]string, bool) {
	if o == nil || IsNil(o.NamesWhitelist) {
		return nil, false
	}
	return o.NamesWhitelist, true
}

// HasNamesWhitelist returns a boolean if a field has been set.
func (o *CheckoutProcessorsPreferences) HasNamesWhitelist() bool {
	if o != nil && !IsNil(o.NamesWhitelist) {
		return true
	}

	return false
}

// SetNamesWhitelist gets a reference to the given []string and assigns it to the NamesWhitelist field.
func (o *CheckoutProcessorsPreferences) SetNamesWhitelist(v []string) {
	o.NamesWhitelist = v
}

func (o CheckoutProcessorsPreferences) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckoutProcessorsPreferences) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FilterProviders) {
		toSerialize["filter_providers"] = o.FilterProviders
	}
	if !IsNil(o.IdsWhitelist) {
		toSerialize["ids_whitelist"] = o.IdsWhitelist
	}
	if !IsNil(o.NamesWhitelist) {
		toSerialize["names_whitelist"] = o.NamesWhitelist
	}
	return toSerialize, nil
}

type NullableCheckoutProcessorsPreferences struct {
	value *CheckoutProcessorsPreferences
	isSet bool
}

func (v NullableCheckoutProcessorsPreferences) Get() *CheckoutProcessorsPreferences {
	return v.value
}

func (v *NullableCheckoutProcessorsPreferences) Set(val *CheckoutProcessorsPreferences) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutProcessorsPreferences) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutProcessorsPreferences) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutProcessorsPreferences(val *CheckoutProcessorsPreferences) *NullableCheckoutProcessorsPreferences {
	return &NullableCheckoutProcessorsPreferences{value: val, isSet: true}
}

func (v NullableCheckoutProcessorsPreferences) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutProcessorsPreferences) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


