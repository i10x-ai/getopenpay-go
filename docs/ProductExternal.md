# ProductExternal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the product. | 
**Object** | Pointer to [**ObjectName**](ObjectName.md) |  | [optional] [default to OBJECTNAME_PRODUCT]
**CreatedAt** | **time.Time** | DateTime at which the object was created, in &#39;ISO 8601&#39; format. | 
**UpdatedAt** | **time.Time** | DateTime at which the object was updated, in &#39;ISO 8601&#39; format. | 
**IsDeleted** | Pointer to **bool** | If true, indicates that this object has been deleted | [optional] [default to false]
**Name** | **string** | Name of product. | 
**Description** | **NullableString** |  | 
**Features** | **[]string** |  | 
**DefaultPrice** | **NullableString** |  | 
**IsActive** | **bool** | Whether the product is currently available for purchase. | 
**Prices** | [**[]PriceExternal**](PriceExternal.md) |  | 
**AccountSku** | **NullableString** |  | 
**UnitLabel** | **NullableString** |  | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewProductExternal

`func NewProductExternal(id string, createdAt time.Time, updatedAt time.Time, name string, description NullableString, features []string, defaultPrice NullableString, isActive bool, prices []PriceExternal, accountSku NullableString, unitLabel NullableString, ) *ProductExternal`

NewProductExternal instantiates a new ProductExternal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductExternalWithDefaults

`func NewProductExternalWithDefaults() *ProductExternal`

NewProductExternalWithDefaults instantiates a new ProductExternal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProductExternal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductExternal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductExternal) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *ProductExternal) GetObject() ObjectName`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ProductExternal) GetObjectOk() (*ObjectName, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ProductExternal) SetObject(v ObjectName)`

SetObject sets Object field to given value.

### HasObject

`func (o *ProductExternal) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ProductExternal) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProductExternal) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProductExternal) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ProductExternal) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ProductExternal) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ProductExternal) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetIsDeleted

`func (o *ProductExternal) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *ProductExternal) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *ProductExternal) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *ProductExternal) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetName

`func (o *ProductExternal) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductExternal) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductExternal) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ProductExternal) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProductExternal) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProductExternal) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *ProductExternal) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ProductExternal) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetFeatures

`func (o *ProductExternal) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *ProductExternal) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *ProductExternal) SetFeatures(v []string)`

SetFeatures sets Features field to given value.


### GetDefaultPrice

`func (o *ProductExternal) GetDefaultPrice() string`

GetDefaultPrice returns the DefaultPrice field if non-nil, zero value otherwise.

### GetDefaultPriceOk

`func (o *ProductExternal) GetDefaultPriceOk() (*string, bool)`

GetDefaultPriceOk returns a tuple with the DefaultPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPrice

`func (o *ProductExternal) SetDefaultPrice(v string)`

SetDefaultPrice sets DefaultPrice field to given value.


### SetDefaultPriceNil

`func (o *ProductExternal) SetDefaultPriceNil(b bool)`

 SetDefaultPriceNil sets the value for DefaultPrice to be an explicit nil

### UnsetDefaultPrice
`func (o *ProductExternal) UnsetDefaultPrice()`

UnsetDefaultPrice ensures that no value is present for DefaultPrice, not even an explicit nil
### GetIsActive

`func (o *ProductExternal) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ProductExternal) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ProductExternal) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetPrices

`func (o *ProductExternal) GetPrices() []PriceExternal`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *ProductExternal) GetPricesOk() (*[]PriceExternal, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *ProductExternal) SetPrices(v []PriceExternal)`

SetPrices sets Prices field to given value.


### GetAccountSku

`func (o *ProductExternal) GetAccountSku() string`

GetAccountSku returns the AccountSku field if non-nil, zero value otherwise.

### GetAccountSkuOk

`func (o *ProductExternal) GetAccountSkuOk() (*string, bool)`

GetAccountSkuOk returns a tuple with the AccountSku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSku

`func (o *ProductExternal) SetAccountSku(v string)`

SetAccountSku sets AccountSku field to given value.


### SetAccountSkuNil

`func (o *ProductExternal) SetAccountSkuNil(b bool)`

 SetAccountSkuNil sets the value for AccountSku to be an explicit nil

### UnsetAccountSku
`func (o *ProductExternal) UnsetAccountSku()`

UnsetAccountSku ensures that no value is present for AccountSku, not even an explicit nil
### GetUnitLabel

`func (o *ProductExternal) GetUnitLabel() string`

GetUnitLabel returns the UnitLabel field if non-nil, zero value otherwise.

### GetUnitLabelOk

`func (o *ProductExternal) GetUnitLabelOk() (*string, bool)`

GetUnitLabelOk returns a tuple with the UnitLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitLabel

`func (o *ProductExternal) SetUnitLabel(v string)`

SetUnitLabel sets UnitLabel field to given value.


### SetUnitLabelNil

`func (o *ProductExternal) SetUnitLabelNil(b bool)`

 SetUnitLabelNil sets the value for UnitLabel to be an explicit nil

### UnsetUnitLabel
`func (o *ProductExternal) UnsetUnitLabel()`

UnsetUnitLabel ensures that no value is present for UnitLabel, not even an explicit nil
### GetMetadata

`func (o *ProductExternal) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ProductExternal) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ProductExternal) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ProductExternal) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *ProductExternal) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *ProductExternal) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetCustomFields

`func (o *ProductExternal) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ProductExternal) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ProductExternal) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ProductExternal) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *ProductExternal) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *ProductExternal) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


