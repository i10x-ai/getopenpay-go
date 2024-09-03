# StringConfigField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the field. | 
**Key** | **string** | The key of the field in the config. | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] [default to "string"]
**Value** | **NullableString** |  | 
**DefaultValue** | Pointer to **NullableString** |  | [optional] 
**MinLength** | Pointer to **NullableInt32** |  | [optional] 
**MaxLength** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewStringConfigField

`func NewStringConfigField(name string, key string, value NullableString, ) *StringConfigField`

NewStringConfigField instantiates a new StringConfigField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStringConfigFieldWithDefaults

`func NewStringConfigFieldWithDefaults() *StringConfigField`

NewStringConfigFieldWithDefaults instantiates a new StringConfigField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StringConfigField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StringConfigField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StringConfigField) SetName(v string)`

SetName sets Name field to given value.


### GetKey

`func (o *StringConfigField) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *StringConfigField) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *StringConfigField) SetKey(v string)`

SetKey sets Key field to given value.


### GetDescription

`func (o *StringConfigField) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StringConfigField) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StringConfigField) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StringConfigField) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *StringConfigField) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *StringConfigField) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetType

`func (o *StringConfigField) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StringConfigField) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StringConfigField) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StringConfigField) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *StringConfigField) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *StringConfigField) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *StringConfigField) SetValue(v string)`

SetValue sets Value field to given value.


### SetValueNil

`func (o *StringConfigField) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *StringConfigField) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetDefaultValue

`func (o *StringConfigField) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *StringConfigField) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *StringConfigField) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *StringConfigField) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *StringConfigField) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *StringConfigField) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetMinLength

`func (o *StringConfigField) GetMinLength() int32`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *StringConfigField) GetMinLengthOk() (*int32, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *StringConfigField) SetMinLength(v int32)`

SetMinLength sets MinLength field to given value.

### HasMinLength

`func (o *StringConfigField) HasMinLength() bool`

HasMinLength returns a boolean if a field has been set.

### SetMinLengthNil

`func (o *StringConfigField) SetMinLengthNil(b bool)`

 SetMinLengthNil sets the value for MinLength to be an explicit nil

### UnsetMinLength
`func (o *StringConfigField) UnsetMinLength()`

UnsetMinLength ensures that no value is present for MinLength, not even an explicit nil
### GetMaxLength

`func (o *StringConfigField) GetMaxLength() int32`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *StringConfigField) GetMaxLengthOk() (*int32, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *StringConfigField) SetMaxLength(v int32)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *StringConfigField) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.

### SetMaxLengthNil

`func (o *StringConfigField) SetMaxLengthNil(b bool)`

 SetMaxLengthNil sets the value for MaxLength to be an explicit nil

### UnsetMaxLength
`func (o *StringConfigField) UnsetMaxLength()`

UnsetMaxLength ensures that no value is present for MaxLength, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


