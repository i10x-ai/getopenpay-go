# CreateProductRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSku** | Pointer to **NullableString** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Description** | **string** | Product&#39;s description. | 
**Features** | Pointer to **[]string** | List of product features. | [optional] 
**IsActive** | Pointer to **bool** | Whether or not this product is accepting new subscriptions. | [optional] [default to true]
**Name** | **string** | Name of product. | 
**UnitLabel** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreateProductRequest

`func NewCreateProductRequest(description string, name string, ) *CreateProductRequest`

NewCreateProductRequest instantiates a new CreateProductRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProductRequestWithDefaults

`func NewCreateProductRequestWithDefaults() *CreateProductRequest`

NewCreateProductRequestWithDefaults instantiates a new CreateProductRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSku

`func (o *CreateProductRequest) GetAccountSku() string`

GetAccountSku returns the AccountSku field if non-nil, zero value otherwise.

### GetAccountSkuOk

`func (o *CreateProductRequest) GetAccountSkuOk() (*string, bool)`

GetAccountSkuOk returns a tuple with the AccountSku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSku

`func (o *CreateProductRequest) SetAccountSku(v string)`

SetAccountSku sets AccountSku field to given value.

### HasAccountSku

`func (o *CreateProductRequest) HasAccountSku() bool`

HasAccountSku returns a boolean if a field has been set.

### SetAccountSkuNil

`func (o *CreateProductRequest) SetAccountSkuNil(b bool)`

 SetAccountSkuNil sets the value for AccountSku to be an explicit nil

### UnsetAccountSku
`func (o *CreateProductRequest) UnsetAccountSku()`

UnsetAccountSku ensures that no value is present for AccountSku, not even an explicit nil
### GetCustomFields

`func (o *CreateProductRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CreateProductRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CreateProductRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CreateProductRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *CreateProductRequest) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *CreateProductRequest) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil
### GetDescription

`func (o *CreateProductRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateProductRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateProductRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetFeatures

`func (o *CreateProductRequest) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *CreateProductRequest) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *CreateProductRequest) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *CreateProductRequest) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetIsActive

`func (o *CreateProductRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *CreateProductRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *CreateProductRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *CreateProductRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetName

`func (o *CreateProductRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateProductRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateProductRequest) SetName(v string)`

SetName sets Name field to given value.


### GetUnitLabel

`func (o *CreateProductRequest) GetUnitLabel() string`

GetUnitLabel returns the UnitLabel field if non-nil, zero value otherwise.

### GetUnitLabelOk

`func (o *CreateProductRequest) GetUnitLabelOk() (*string, bool)`

GetUnitLabelOk returns a tuple with the UnitLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitLabel

`func (o *CreateProductRequest) SetUnitLabel(v string)`

SetUnitLabel sets UnitLabel field to given value.

### HasUnitLabel

`func (o *CreateProductRequest) HasUnitLabel() bool`

HasUnitLabel returns a boolean if a field has been set.

### SetUnitLabelNil

`func (o *CreateProductRequest) SetUnitLabelNil(b bool)`

 SetUnitLabelNil sets the value for UnitLabel to be an explicit nil

### UnsetUnitLabel
`func (o *CreateProductRequest) UnsetUnitLabel()`

UnsetUnitLabel ensures that no value is present for UnitLabel, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


