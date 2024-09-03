# NumberConfigField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the field. | 
**Key** | **string** | The key of the field in the config. | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] [default to "number"]
**Value** | **NullableFloat32** |  | 
**DefaultValue** | Pointer to **NullableFloat32** |  | [optional] 
**MinValue** | Pointer to **NullableFloat32** |  | [optional] 
**MaxValue** | Pointer to **NullableFloat32** |  | [optional] 

## Methods

### NewNumberConfigField

`func NewNumberConfigField(name string, key string, value NullableFloat32, ) *NumberConfigField`

NewNumberConfigField instantiates a new NumberConfigField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNumberConfigFieldWithDefaults

`func NewNumberConfigFieldWithDefaults() *NumberConfigField`

NewNumberConfigFieldWithDefaults instantiates a new NumberConfigField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NumberConfigField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NumberConfigField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NumberConfigField) SetName(v string)`

SetName sets Name field to given value.


### GetKey

`func (o *NumberConfigField) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *NumberConfigField) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *NumberConfigField) SetKey(v string)`

SetKey sets Key field to given value.


### GetDescription

`func (o *NumberConfigField) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NumberConfigField) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NumberConfigField) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NumberConfigField) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *NumberConfigField) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *NumberConfigField) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetType

`func (o *NumberConfigField) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NumberConfigField) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NumberConfigField) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NumberConfigField) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *NumberConfigField) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *NumberConfigField) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *NumberConfigField) SetValue(v float32)`

SetValue sets Value field to given value.


### SetValueNil

`func (o *NumberConfigField) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *NumberConfigField) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetDefaultValue

`func (o *NumberConfigField) GetDefaultValue() float32`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *NumberConfigField) GetDefaultValueOk() (*float32, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *NumberConfigField) SetDefaultValue(v float32)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *NumberConfigField) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *NumberConfigField) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *NumberConfigField) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetMinValue

`func (o *NumberConfigField) GetMinValue() float32`

GetMinValue returns the MinValue field if non-nil, zero value otherwise.

### GetMinValueOk

`func (o *NumberConfigField) GetMinValueOk() (*float32, bool)`

GetMinValueOk returns a tuple with the MinValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinValue

`func (o *NumberConfigField) SetMinValue(v float32)`

SetMinValue sets MinValue field to given value.

### HasMinValue

`func (o *NumberConfigField) HasMinValue() bool`

HasMinValue returns a boolean if a field has been set.

### SetMinValueNil

`func (o *NumberConfigField) SetMinValueNil(b bool)`

 SetMinValueNil sets the value for MinValue to be an explicit nil

### UnsetMinValue
`func (o *NumberConfigField) UnsetMinValue()`

UnsetMinValue ensures that no value is present for MinValue, not even an explicit nil
### GetMaxValue

`func (o *NumberConfigField) GetMaxValue() float32`

GetMaxValue returns the MaxValue field if non-nil, zero value otherwise.

### GetMaxValueOk

`func (o *NumberConfigField) GetMaxValueOk() (*float32, bool)`

GetMaxValueOk returns a tuple with the MaxValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValue

`func (o *NumberConfigField) SetMaxValue(v float32)`

SetMaxValue sets MaxValue field to given value.

### HasMaxValue

`func (o *NumberConfigField) HasMaxValue() bool`

HasMaxValue returns a boolean if a field has been set.

### SetMaxValueNil

`func (o *NumberConfigField) SetMaxValueNil(b bool)`

 SetMaxValueNil sets the value for MaxValue to be an explicit nil

### UnsetMaxValue
`func (o *NumberConfigField) UnsetMaxValue()`

UnsetMaxValue ensures that no value is present for MaxValue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


