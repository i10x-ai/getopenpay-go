# CreateCheckoutLineItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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


