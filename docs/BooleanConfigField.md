# BooleanConfigField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultValue** | Pointer to **NullableBool** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Key** | **string** | The key of the field in the config. | 
**Name** | **string** | The name of the field. | 
**Type** | Pointer to **string** |  | [optional] [default to "boolean"]
**Value** | **NullableBool** |  | 

## Methods

### NewBooleanConfigField

`func NewBooleanConfigField(key string, name string, value NullableBool, ) *BooleanConfigField`

NewBooleanConfigField instantiates a new BooleanConfigField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBooleanConfigFieldWithDefaults

`func NewBooleanConfigFieldWithDefaults() *BooleanConfigField`

NewBooleanConfigFieldWithDefaults instantiates a new BooleanConfigField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultValue

`func (o *BooleanConfigField) GetDefaultValue() bool`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *BooleanConfigField) GetDefaultValueOk() (*bool, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *BooleanConfigField) SetDefaultValue(v bool)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *BooleanConfigField) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *BooleanConfigField) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *BooleanConfigField) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetDescription

`func (o *BooleanConfigField) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BooleanConfigField) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BooleanConfigField) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BooleanConfigField) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BooleanConfigField) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BooleanConfigField) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetKey

`func (o *BooleanConfigField) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *BooleanConfigField) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *BooleanConfigField) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *BooleanConfigField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BooleanConfigField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BooleanConfigField) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *BooleanConfigField) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BooleanConfigField) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BooleanConfigField) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BooleanConfigField) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *BooleanConfigField) GetValue() bool`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BooleanConfigField) GetValueOk() (*bool, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BooleanConfigField) SetValue(v bool)`

SetValue sets Value field to given value.


### SetValueNil

`func (o *BooleanConfigField) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *BooleanConfigField) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


