# EnumConfigField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultValue** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Key** | **string** | The key of the field in the config. | 
**Name** | **string** | The name of the field. | 
**Options** | **[]string** | The options for the enum. | 
**Type** | Pointer to **string** |  | [optional] [default to "enum"]
**Value** | **NullableString** |  | 

## Methods

### NewEnumConfigField

`func NewEnumConfigField(key string, name string, options []string, value NullableString, ) *EnumConfigField`

NewEnumConfigField instantiates a new EnumConfigField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnumConfigFieldWithDefaults

`func NewEnumConfigFieldWithDefaults() *EnumConfigField`

NewEnumConfigFieldWithDefaults instantiates a new EnumConfigField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultValue

`func (o *EnumConfigField) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *EnumConfigField) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *EnumConfigField) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *EnumConfigField) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *EnumConfigField) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *EnumConfigField) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetDescription

`func (o *EnumConfigField) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnumConfigField) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnumConfigField) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EnumConfigField) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *EnumConfigField) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *EnumConfigField) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetKey

`func (o *EnumConfigField) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *EnumConfigField) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *EnumConfigField) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *EnumConfigField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnumConfigField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnumConfigField) SetName(v string)`

SetName sets Name field to given value.


### GetOptions

`func (o *EnumConfigField) GetOptions() []string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *EnumConfigField) GetOptionsOk() (*[]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *EnumConfigField) SetOptions(v []string)`

SetOptions sets Options field to given value.


### GetType

`func (o *EnumConfigField) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EnumConfigField) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EnumConfigField) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EnumConfigField) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *EnumConfigField) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EnumConfigField) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EnumConfigField) SetValue(v string)`

SetValue sets Value field to given value.


### SetValueNil

`func (o *EnumConfigField) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *EnumConfigField) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


