# CreateCheckoutLineItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**PriceId** | **string** | Unique ID of the price corresponding to this line item | 
**Quantity** | **int32** | The quantity of the line item being purchased. | 

## Methods

### NewCreateCheckoutLineItem

`func NewCreateCheckoutLineItem(priceId string, quantity int32, ) *CreateCheckoutLineItem`

NewCreateCheckoutLineItem instantiates a new CreateCheckoutLineItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCheckoutLineItemWithDefaults

`func NewCreateCheckoutLineItemWithDefaults() *CreateCheckoutLineItem`

NewCreateCheckoutLineItemWithDefaults instantiates a new CreateCheckoutLineItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomFields

`func (o *CreateCheckoutLineItem) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CreateCheckoutLineItem) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CreateCheckoutLineItem) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CreateCheckoutLineItem) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *CreateCheckoutLineItem) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *CreateCheckoutLineItem) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil
### GetDescription

`func (o *CreateCheckoutLineItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateCheckoutLineItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateCheckoutLineItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateCheckoutLineItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateCheckoutLineItem) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateCheckoutLineItem) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPriceId

`func (o *CreateCheckoutLineItem) GetPriceId() string`

GetPriceId returns the PriceId field if non-nil, zero value otherwise.

### GetPriceIdOk

`func (o *CreateCheckoutLineItem) GetPriceIdOk() (*string, bool)`

GetPriceIdOk returns a tuple with the PriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceId

`func (o *CreateCheckoutLineItem) SetPriceId(v string)`

SetPriceId sets PriceId field to given value.


### GetQuantity

`func (o *CreateCheckoutLineItem) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CreateCheckoutLineItem) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CreateCheckoutLineItem) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


