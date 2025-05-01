# FallbackConfigurationOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FallbackProcessorPreferences** | Pointer to [**NullableCheckoutProcessorsPreferences**](CheckoutProcessorsPreferences.md) |  | [optional] 
**PriceGroups** | Pointer to [**[]SelectedPriceQuantity**](SelectedPriceQuantity.md) | List of price groups that will be used to fall back to if the selected product price quantity fails to be created. | [optional] [default to []]

## Methods

### NewFallbackConfigurationOutput

`func NewFallbackConfigurationOutput() *FallbackConfigurationOutput`

NewFallbackConfigurationOutput instantiates a new FallbackConfigurationOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFallbackConfigurationOutputWithDefaults

`func NewFallbackConfigurationOutputWithDefaults() *FallbackConfigurationOutput`

NewFallbackConfigurationOutputWithDefaults instantiates a new FallbackConfigurationOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFallbackProcessorPreferences

`func (o *FallbackConfigurationOutput) GetFallbackProcessorPreferences() CheckoutProcessorsPreferences`

GetFallbackProcessorPreferences returns the FallbackProcessorPreferences field if non-nil, zero value otherwise.

### GetFallbackProcessorPreferencesOk

`func (o *FallbackConfigurationOutput) GetFallbackProcessorPreferencesOk() (*CheckoutProcessorsPreferences, bool)`

GetFallbackProcessorPreferencesOk returns a tuple with the FallbackProcessorPreferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackProcessorPreferences

`func (o *FallbackConfigurationOutput) SetFallbackProcessorPreferences(v CheckoutProcessorsPreferences)`

SetFallbackProcessorPreferences sets FallbackProcessorPreferences field to given value.

### HasFallbackProcessorPreferences

`func (o *FallbackConfigurationOutput) HasFallbackProcessorPreferences() bool`

HasFallbackProcessorPreferences returns a boolean if a field has been set.

### SetFallbackProcessorPreferencesNil

`func (o *FallbackConfigurationOutput) SetFallbackProcessorPreferencesNil(b bool)`

 SetFallbackProcessorPreferencesNil sets the value for FallbackProcessorPreferences to be an explicit nil

### UnsetFallbackProcessorPreferences
`func (o *FallbackConfigurationOutput) UnsetFallbackProcessorPreferences()`

UnsetFallbackProcessorPreferences ensures that no value is present for FallbackProcessorPreferences, not even an explicit nil
### GetPriceGroups

`func (o *FallbackConfigurationOutput) GetPriceGroups() []SelectedPriceQuantity`

GetPriceGroups returns the PriceGroups field if non-nil, zero value otherwise.

### GetPriceGroupsOk

`func (o *FallbackConfigurationOutput) GetPriceGroupsOk() (*[]SelectedPriceQuantity, bool)`

GetPriceGroupsOk returns a tuple with the PriceGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceGroups

`func (o *FallbackConfigurationOutput) SetPriceGroups(v []SelectedPriceQuantity)`

SetPriceGroups sets PriceGroups field to given value.

### HasPriceGroups

`func (o *FallbackConfigurationOutput) HasPriceGroups() bool`

HasPriceGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


