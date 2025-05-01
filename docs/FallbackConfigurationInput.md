# FallbackConfigurationInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FallbackProcessorPreferences** | Pointer to [**NullableCheckoutProcessorsPreferences**](CheckoutProcessorsPreferences.md) |  | [optional] 
**PriceGroups** | Pointer to [**[]SelectedPriceQuantity**](SelectedPriceQuantity.md) | List of price groups that will be used to fall back to if the selected product price quantity fails to be created. | [optional] [default to []]

## Methods

### NewFallbackConfigurationInput

`func NewFallbackConfigurationInput() *FallbackConfigurationInput`

NewFallbackConfigurationInput instantiates a new FallbackConfigurationInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFallbackConfigurationInputWithDefaults

`func NewFallbackConfigurationInputWithDefaults() *FallbackConfigurationInput`

NewFallbackConfigurationInputWithDefaults instantiates a new FallbackConfigurationInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFallbackProcessorPreferences

`func (o *FallbackConfigurationInput) GetFallbackProcessorPreferences() CheckoutProcessorsPreferences`

GetFallbackProcessorPreferences returns the FallbackProcessorPreferences field if non-nil, zero value otherwise.

### GetFallbackProcessorPreferencesOk

`func (o *FallbackConfigurationInput) GetFallbackProcessorPreferencesOk() (*CheckoutProcessorsPreferences, bool)`

GetFallbackProcessorPreferencesOk returns a tuple with the FallbackProcessorPreferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackProcessorPreferences

`func (o *FallbackConfigurationInput) SetFallbackProcessorPreferences(v CheckoutProcessorsPreferences)`

SetFallbackProcessorPreferences sets FallbackProcessorPreferences field to given value.

### HasFallbackProcessorPreferences

`func (o *FallbackConfigurationInput) HasFallbackProcessorPreferences() bool`

HasFallbackProcessorPreferences returns a boolean if a field has been set.

### SetFallbackProcessorPreferencesNil

`func (o *FallbackConfigurationInput) SetFallbackProcessorPreferencesNil(b bool)`

 SetFallbackProcessorPreferencesNil sets the value for FallbackProcessorPreferences to be an explicit nil

### UnsetFallbackProcessorPreferences
`func (o *FallbackConfigurationInput) UnsetFallbackProcessorPreferences()`

UnsetFallbackProcessorPreferences ensures that no value is present for FallbackProcessorPreferences, not even an explicit nil
### GetPriceGroups

`func (o *FallbackConfigurationInput) GetPriceGroups() []SelectedPriceQuantity`

GetPriceGroups returns the PriceGroups field if non-nil, zero value otherwise.

### GetPriceGroupsOk

`func (o *FallbackConfigurationInput) GetPriceGroupsOk() (*[]SelectedPriceQuantity, bool)`

GetPriceGroupsOk returns a tuple with the PriceGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceGroups

`func (o *FallbackConfigurationInput) SetPriceGroups(v []SelectedPriceQuantity)`

SetPriceGroups sets PriceGroups field to given value.

### HasPriceGroups

`func (o *FallbackConfigurationInput) HasPriceGroups() bool`

HasPriceGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


