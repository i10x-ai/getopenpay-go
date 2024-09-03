# UpdateSubscriptionItemRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PriceId** | Pointer to **NullableString** |  | [optional] 
**Quantity** | Pointer to **NullableInt32** |  | [optional] 
**ProrationBehavior** | Pointer to [**ProrationEnum**](ProrationEnum.md) |  | [optional] [default to PRORATIONENUM_ALWAYS_INVOICE]
**DropAtEnd** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewUpdateSubscriptionItemRequest

`func NewUpdateSubscriptionItemRequest() *UpdateSubscriptionItemRequest`

NewUpdateSubscriptionItemRequest instantiates a new UpdateSubscriptionItemRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSubscriptionItemRequestWithDefaults

`func NewUpdateSubscriptionItemRequestWithDefaults() *UpdateSubscriptionItemRequest`

NewUpdateSubscriptionItemRequestWithDefaults instantiates a new UpdateSubscriptionItemRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPriceId

`func (o *UpdateSubscriptionItemRequest) GetPriceId() string`

GetPriceId returns the PriceId field if non-nil, zero value otherwise.

### GetPriceIdOk

`func (o *UpdateSubscriptionItemRequest) GetPriceIdOk() (*string, bool)`

GetPriceIdOk returns a tuple with the PriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceId

`func (o *UpdateSubscriptionItemRequest) SetPriceId(v string)`

SetPriceId sets PriceId field to given value.

### HasPriceId

`func (o *UpdateSubscriptionItemRequest) HasPriceId() bool`

HasPriceId returns a boolean if a field has been set.

### SetPriceIdNil

`func (o *UpdateSubscriptionItemRequest) SetPriceIdNil(b bool)`

 SetPriceIdNil sets the value for PriceId to be an explicit nil

### UnsetPriceId
`func (o *UpdateSubscriptionItemRequest) UnsetPriceId()`

UnsetPriceId ensures that no value is present for PriceId, not even an explicit nil
### GetQuantity

`func (o *UpdateSubscriptionItemRequest) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *UpdateSubscriptionItemRequest) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *UpdateSubscriptionItemRequest) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *UpdateSubscriptionItemRequest) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### SetQuantityNil

`func (o *UpdateSubscriptionItemRequest) SetQuantityNil(b bool)`

 SetQuantityNil sets the value for Quantity to be an explicit nil

### UnsetQuantity
`func (o *UpdateSubscriptionItemRequest) UnsetQuantity()`

UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil
### GetProrationBehavior

`func (o *UpdateSubscriptionItemRequest) GetProrationBehavior() ProrationEnum`

GetProrationBehavior returns the ProrationBehavior field if non-nil, zero value otherwise.

### GetProrationBehaviorOk

`func (o *UpdateSubscriptionItemRequest) GetProrationBehaviorOk() (*ProrationEnum, bool)`

GetProrationBehaviorOk returns a tuple with the ProrationBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProrationBehavior

`func (o *UpdateSubscriptionItemRequest) SetProrationBehavior(v ProrationEnum)`

SetProrationBehavior sets ProrationBehavior field to given value.

### HasProrationBehavior

`func (o *UpdateSubscriptionItemRequest) HasProrationBehavior() bool`

HasProrationBehavior returns a boolean if a field has been set.

### GetDropAtEnd

`func (o *UpdateSubscriptionItemRequest) GetDropAtEnd() bool`

GetDropAtEnd returns the DropAtEnd field if non-nil, zero value otherwise.

### GetDropAtEndOk

`func (o *UpdateSubscriptionItemRequest) GetDropAtEndOk() (*bool, bool)`

GetDropAtEndOk returns a tuple with the DropAtEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropAtEnd

`func (o *UpdateSubscriptionItemRequest) SetDropAtEnd(v bool)`

SetDropAtEnd sets DropAtEnd field to given value.

### HasDropAtEnd

`func (o *UpdateSubscriptionItemRequest) HasDropAtEnd() bool`

HasDropAtEnd returns a boolean if a field has been set.

### SetDropAtEndNil

`func (o *UpdateSubscriptionItemRequest) SetDropAtEndNil(b bool)`

 SetDropAtEndNil sets the value for DropAtEnd to be an explicit nil

### UnsetDropAtEnd
`func (o *UpdateSubscriptionItemRequest) UnsetDropAtEnd()`

UnsetDropAtEnd ensures that no value is present for DropAtEnd, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


