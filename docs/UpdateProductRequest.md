# UpdateProductRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**AccountSku** | Pointer to **NullableString** |  | [optional] 
**Features** | Pointer to **[]string** |  | [optional] 
**UnitLabel** | Pointer to **NullableString** |  | [optional] 
**IsActive** | Pointer to **NullableBool** |  | [optional] 
**DefaultPriceId** | Pointer to **NullableString** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewUpdateProductRequest

`func NewUpdateProductRequest() *UpdateProductRequest`

NewUpdateProductRequest instantiates a new UpdateProductRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProductRequestWithDefaults

`func NewUpdateProductRequestWithDefaults() *UpdateProductRequest`

NewUpdateProductRequestWithDefaults instantiates a new UpdateProductRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateProductRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateProductRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateProductRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateProductRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateProductRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateProductRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *UpdateProductRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateProductRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateProductRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateProductRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateProductRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateProductRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetAccountSku

`func (o *UpdateProductRequest) GetAccountSku() string`

GetAccountSku returns the AccountSku field if non-nil, zero value otherwise.

### GetAccountSkuOk

`func (o *UpdateProductRequest) GetAccountSkuOk() (*string, bool)`

GetAccountSkuOk returns a tuple with the AccountSku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSku

`func (o *UpdateProductRequest) SetAccountSku(v string)`

SetAccountSku sets AccountSku field to given value.

### HasAccountSku

`func (o *UpdateProductRequest) HasAccountSku() bool`

HasAccountSku returns a boolean if a field has been set.

### SetAccountSkuNil

`func (o *UpdateProductRequest) SetAccountSkuNil(b bool)`

 SetAccountSkuNil sets the value for AccountSku to be an explicit nil

### UnsetAccountSku
`func (o *UpdateProductRequest) UnsetAccountSku()`

UnsetAccountSku ensures that no value is present for AccountSku, not even an explicit nil
### GetFeatures

`func (o *UpdateProductRequest) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *UpdateProductRequest) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *UpdateProductRequest) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *UpdateProductRequest) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetUnitLabel

`func (o *UpdateProductRequest) GetUnitLabel() string`

GetUnitLabel returns the UnitLabel field if non-nil, zero value otherwise.

### GetUnitLabelOk

`func (o *UpdateProductRequest) GetUnitLabelOk() (*string, bool)`

GetUnitLabelOk returns a tuple with the UnitLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitLabel

`func (o *UpdateProductRequest) SetUnitLabel(v string)`

SetUnitLabel sets UnitLabel field to given value.

### HasUnitLabel

`func (o *UpdateProductRequest) HasUnitLabel() bool`

HasUnitLabel returns a boolean if a field has been set.

### SetUnitLabelNil

`func (o *UpdateProductRequest) SetUnitLabelNil(b bool)`

 SetUnitLabelNil sets the value for UnitLabel to be an explicit nil

### UnsetUnitLabel
`func (o *UpdateProductRequest) UnsetUnitLabel()`

UnsetUnitLabel ensures that no value is present for UnitLabel, not even an explicit nil
### GetIsActive

`func (o *UpdateProductRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *UpdateProductRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *UpdateProductRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *UpdateProductRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *UpdateProductRequest) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *UpdateProductRequest) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetDefaultPriceId

`func (o *UpdateProductRequest) GetDefaultPriceId() string`

GetDefaultPriceId returns the DefaultPriceId field if non-nil, zero value otherwise.

### GetDefaultPriceIdOk

`func (o *UpdateProductRequest) GetDefaultPriceIdOk() (*string, bool)`

GetDefaultPriceIdOk returns a tuple with the DefaultPriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPriceId

`func (o *UpdateProductRequest) SetDefaultPriceId(v string)`

SetDefaultPriceId sets DefaultPriceId field to given value.

### HasDefaultPriceId

`func (o *UpdateProductRequest) HasDefaultPriceId() bool`

HasDefaultPriceId returns a boolean if a field has been set.

### SetDefaultPriceIdNil

`func (o *UpdateProductRequest) SetDefaultPriceIdNil(b bool)`

 SetDefaultPriceIdNil sets the value for DefaultPriceId to be an explicit nil

### UnsetDefaultPriceId
`func (o *UpdateProductRequest) UnsetDefaultPriceId()`

UnsetDefaultPriceId ensures that no value is present for DefaultPriceId, not even an explicit nil
### GetCustomFields

`func (o *UpdateProductRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *UpdateProductRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *UpdateProductRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *UpdateProductRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *UpdateProductRequest) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *UpdateProductRequest) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


