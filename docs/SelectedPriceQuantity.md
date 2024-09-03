# SelectedPriceQuantity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PriceId** | **string** | Unique Identifier of the price. | 
**Quantity** | **int32** | Quantity of the product selected for the subscription.This field is ignored for metered prices | 

## Methods

### NewSelectedPriceQuantity

`func NewSelectedPriceQuantity(priceId string, quantity int32, ) *SelectedPriceQuantity`

NewSelectedPriceQuantity instantiates a new SelectedPriceQuantity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectedPriceQuantityWithDefaults

`func NewSelectedPriceQuantityWithDefaults() *SelectedPriceQuantity`

NewSelectedPriceQuantityWithDefaults instantiates a new SelectedPriceQuantity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPriceId

`func (o *SelectedPriceQuantity) GetPriceId() string`

GetPriceId returns the PriceId field if non-nil, zero value otherwise.

### GetPriceIdOk

`func (o *SelectedPriceQuantity) GetPriceIdOk() (*string, bool)`

GetPriceIdOk returns a tuple with the PriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceId

`func (o *SelectedPriceQuantity) SetPriceId(v string)`

SetPriceId sets PriceId field to given value.


### GetQuantity

`func (o *SelectedPriceQuantity) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *SelectedPriceQuantity) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *SelectedPriceQuantity) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


