# ProductFamilyExternal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | DateTime at which the object was created, in &#39;ISO 8601&#39; format. | 
**Description** | **NullableString** |  | 
**Hierarchy** | Pointer to **string** | A JSON object representing the hierarchy within the family. | [optional] [default to "{}"]
**Id** | **string** | Unique identifier of the product family. | 
**IsDeleted** | Pointer to **bool** | If true, indicates that this object has been deleted | [optional] [default to false]
**Name** | **string** | Name of the product family. | 
**Object** | Pointer to [**ObjectName**](ObjectName.md) |  | [optional] 
**Products** | **[]string** | List of products in the family | 
**UpdatedAt** | **time.Time** | DateTime at which the object was updated, in &#39;ISO 8601&#39; format. | 

## Methods

### NewProductFamilyExternal

`func NewProductFamilyExternal(createdAt time.Time, description NullableString, id string, name string, products []string, updatedAt time.Time, ) *ProductFamilyExternal`

NewProductFamilyExternal instantiates a new ProductFamilyExternal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductFamilyExternalWithDefaults

`func NewProductFamilyExternalWithDefaults() *ProductFamilyExternal`

NewProductFamilyExternalWithDefaults instantiates a new ProductFamilyExternal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ProductFamilyExternal) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProductFamilyExternal) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProductFamilyExternal) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDescription

`func (o *ProductFamilyExternal) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProductFamilyExternal) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProductFamilyExternal) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *ProductFamilyExternal) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ProductFamilyExternal) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetHierarchy

`func (o *ProductFamilyExternal) GetHierarchy() string`

GetHierarchy returns the Hierarchy field if non-nil, zero value otherwise.

### GetHierarchyOk

`func (o *ProductFamilyExternal) GetHierarchyOk() (*string, bool)`

GetHierarchyOk returns a tuple with the Hierarchy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHierarchy

`func (o *ProductFamilyExternal) SetHierarchy(v string)`

SetHierarchy sets Hierarchy field to given value.

### HasHierarchy

`func (o *ProductFamilyExternal) HasHierarchy() bool`

HasHierarchy returns a boolean if a field has been set.

### GetId

`func (o *ProductFamilyExternal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductFamilyExternal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductFamilyExternal) SetId(v string)`

SetId sets Id field to given value.


### GetIsDeleted

`func (o *ProductFamilyExternal) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *ProductFamilyExternal) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *ProductFamilyExternal) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *ProductFamilyExternal) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetName

`func (o *ProductFamilyExternal) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductFamilyExternal) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductFamilyExternal) SetName(v string)`

SetName sets Name field to given value.


### GetObject

`func (o *ProductFamilyExternal) GetObject() ObjectName`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ProductFamilyExternal) GetObjectOk() (*ObjectName, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ProductFamilyExternal) SetObject(v ObjectName)`

SetObject sets Object field to given value.

### HasObject

`func (o *ProductFamilyExternal) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetProducts

`func (o *ProductFamilyExternal) GetProducts() []string`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *ProductFamilyExternal) GetProductsOk() (*[]string, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *ProductFamilyExternal) SetProducts(v []string)`

SetProducts sets Products field to given value.


### GetUpdatedAt

`func (o *ProductFamilyExternal) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ProductFamilyExternal) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ProductFamilyExternal) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


