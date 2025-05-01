# NonPciIntegrationConfigurationInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultValue** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Key** | **string** | The key of the field in the config. | 
**MaxLength** | Pointer to **int32** |  | [optional] 
**MinLength** | Pointer to **int32** |  | [optional] 
**Name** | **string** | The name of the field. | 
**Type** | Pointer to **string** |  | [optional] [default to "string"]
**Value** | **string** |  | 
**MaxValue** | Pointer to **float32** |  | [optional] 
**MinValue** | Pointer to **float32** |  | [optional] 
**Options** | **[]string** | The options for the enum. | 

## Methods

### NewNonPciIntegrationConfigurationInner

`func NewNonPciIntegrationConfigurationInner(key string, name string, value string, options []string, ) *NonPciIntegrationConfigurationInner`

NewNonPciIntegrationConfigurationInner instantiates a new NonPciIntegrationConfigurationInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNonPciIntegrationConfigurationInnerWithDefaults

`func NewNonPciIntegrationConfigurationInnerWithDefaults() *NonPciIntegrationConfigurationInner`

NewNonPciIntegrationConfigurationInnerWithDefaults instantiates a new NonPciIntegrationConfigurationInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultValue

`func (o *NonPciIntegrationConfigurationInner) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *NonPciIntegrationConfigurationInner) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *NonPciIntegrationConfigurationInner) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *NonPciIntegrationConfigurationInner) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetDescription

`func (o *NonPciIntegrationConfigurationInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NonPciIntegrationConfigurationInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NonPciIntegrationConfigurationInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NonPciIntegrationConfigurationInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetKey

`func (o *NonPciIntegrationConfigurationInner) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *NonPciIntegrationConfigurationInner) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *NonPciIntegrationConfigurationInner) SetKey(v string)`

SetKey sets Key field to given value.


### GetMaxLength

`func (o *NonPciIntegrationConfigurationInner) GetMaxLength() int32`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *NonPciIntegrationConfigurationInner) GetMaxLengthOk() (*int32, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *NonPciIntegrationConfigurationInner) SetMaxLength(v int32)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *NonPciIntegrationConfigurationInner) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.

### GetMinLength

`func (o *NonPciIntegrationConfigurationInner) GetMinLength() int32`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *NonPciIntegrationConfigurationInner) GetMinLengthOk() (*int32, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *NonPciIntegrationConfigurationInner) SetMinLength(v int32)`

SetMinLength sets MinLength field to given value.

### HasMinLength

`func (o *NonPciIntegrationConfigurationInner) HasMinLength() bool`

HasMinLength returns a boolean if a field has been set.

### GetName

`func (o *NonPciIntegrationConfigurationInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NonPciIntegrationConfigurationInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NonPciIntegrationConfigurationInner) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *NonPciIntegrationConfigurationInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NonPciIntegrationConfigurationInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NonPciIntegrationConfigurationInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NonPciIntegrationConfigurationInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *NonPciIntegrationConfigurationInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *NonPciIntegrationConfigurationInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *NonPciIntegrationConfigurationInner) SetValue(v string)`

SetValue sets Value field to given value.


### GetMaxValue

`func (o *NonPciIntegrationConfigurationInner) GetMaxValue() float32`

GetMaxValue returns the MaxValue field if non-nil, zero value otherwise.

### GetMaxValueOk

`func (o *NonPciIntegrationConfigurationInner) GetMaxValueOk() (*float32, bool)`

GetMaxValueOk returns a tuple with the MaxValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValue

`func (o *NonPciIntegrationConfigurationInner) SetMaxValue(v float32)`

SetMaxValue sets MaxValue field to given value.

### HasMaxValue

`func (o *NonPciIntegrationConfigurationInner) HasMaxValue() bool`

HasMaxValue returns a boolean if a field has been set.

### GetMinValue

`func (o *NonPciIntegrationConfigurationInner) GetMinValue() float32`

GetMinValue returns the MinValue field if non-nil, zero value otherwise.

### GetMinValueOk

`func (o *NonPciIntegrationConfigurationInner) GetMinValueOk() (*float32, bool)`

GetMinValueOk returns a tuple with the MinValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinValue

`func (o *NonPciIntegrationConfigurationInner) SetMinValue(v float32)`

SetMinValue sets MinValue field to given value.

### HasMinValue

`func (o *NonPciIntegrationConfigurationInner) HasMinValue() bool`

HasMinValue returns a boolean if a field has been set.

### GetOptions

`func (o *NonPciIntegrationConfigurationInner) GetOptions() []string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *NonPciIntegrationConfigurationInner) GetOptionsOk() (*[]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *NonPciIntegrationConfigurationInner) SetOptions(v []string)`

SetOptions sets Options field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


