# CreateProductFamilyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** |  | [optional] 
**Hierarchy** | Pointer to **string** | A JSON object representing the hierarchy within the family. | [optional] [default to "{}"]
**Name** | **string** | The name of the product family. | 
**Products** | **[]string** | List of unique id&#39;s for the products in this family. | 

## Methods

### NewCreateProductFamilyRequest

`func NewCreateProductFamilyRequest(name string, products []string, ) *CreateProductFamilyRequest`

NewCreateProductFamilyRequest instantiates a new CreateProductFamilyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProductFamilyRequestWithDefaults

`func NewCreateProductFamilyRequestWithDefaults() *CreateProductFamilyRequest`

NewCreateProductFamilyRequestWithDefaults instantiates a new CreateProductFamilyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateProductFamilyRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateProductFamilyRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateProductFamilyRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateProductFamilyRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateProductFamilyRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateProductFamilyRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetHierarchy

`func (o *CreateProductFamilyRequest) GetHierarchy() string`

GetHierarchy returns the Hierarchy field if non-nil, zero value otherwise.

### GetHierarchyOk

`func (o *CreateProductFamilyRequest) GetHierarchyOk() (*string, bool)`

GetHierarchyOk returns a tuple with the Hierarchy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHierarchy

`func (o *CreateProductFamilyRequest) SetHierarchy(v string)`

SetHierarchy sets Hierarchy field to given value.

### HasHierarchy

`func (o *CreateProductFamilyRequest) HasHierarchy() bool`

HasHierarchy returns a boolean if a field has been set.

### GetName

`func (o *CreateProductFamilyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateProductFamilyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateProductFamilyRequest) SetName(v string)`

SetName sets Name field to given value.


### GetProducts

`func (o *CreateProductFamilyRequest) GetProducts() []string`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *CreateProductFamilyRequest) GetProductsOk() (*[]string, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *CreateProductFamilyRequest) SetProducts(v []string)`

SetProducts sets Products field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


