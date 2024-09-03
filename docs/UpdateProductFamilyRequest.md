# UpdateProductFamilyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Products** | **[]string** |  | 
**Hierarchy** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUpdateProductFamilyRequest

`func NewUpdateProductFamilyRequest(name NullableString, products []string, ) *UpdateProductFamilyRequest`

NewUpdateProductFamilyRequest instantiates a new UpdateProductFamilyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProductFamilyRequestWithDefaults

`func NewUpdateProductFamilyRequestWithDefaults() *UpdateProductFamilyRequest`

NewUpdateProductFamilyRequestWithDefaults instantiates a new UpdateProductFamilyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateProductFamilyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateProductFamilyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateProductFamilyRequest) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *UpdateProductFamilyRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateProductFamilyRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *UpdateProductFamilyRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateProductFamilyRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateProductFamilyRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateProductFamilyRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateProductFamilyRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateProductFamilyRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetProducts

`func (o *UpdateProductFamilyRequest) GetProducts() []string`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *UpdateProductFamilyRequest) GetProductsOk() (*[]string, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *UpdateProductFamilyRequest) SetProducts(v []string)`

SetProducts sets Products field to given value.


### GetHierarchy

`func (o *UpdateProductFamilyRequest) GetHierarchy() string`

GetHierarchy returns the Hierarchy field if non-nil, zero value otherwise.

### GetHierarchyOk

`func (o *UpdateProductFamilyRequest) GetHierarchyOk() (*string, bool)`

GetHierarchyOk returns a tuple with the Hierarchy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHierarchy

`func (o *UpdateProductFamilyRequest) SetHierarchy(v string)`

SetHierarchy sets Hierarchy field to given value.

### HasHierarchy

`func (o *UpdateProductFamilyRequest) HasHierarchy() bool`

HasHierarchy returns a boolean if a field has been set.

### SetHierarchyNil

`func (o *UpdateProductFamilyRequest) SetHierarchyNil(b bool)`

 SetHierarchyNil sets the value for Hierarchy to be an explicit nil

### UnsetHierarchy
`func (o *UpdateProductFamilyRequest) UnsetHierarchy()`

UnsetHierarchy ensures that no value is present for Hierarchy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


