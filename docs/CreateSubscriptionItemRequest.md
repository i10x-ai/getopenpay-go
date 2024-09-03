# CreateSubscriptionItemRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionId** | **string** | The identifier of the subscription to modify | 
**AddAtPeriodEnd** | Pointer to **bool** | If the flag is set to True, item will be added when renewing the subscription at next billing cycle. | [optional] [default to false]
**PriceId** | **string** | The ID of the price. | 
**Quantity** | Pointer to **int32** | The quantity you’d like to apply to the subscription item you’re creating. | [optional] [default to 1]
**ProrationBehavior** | Pointer to [**ProrationEnum**](ProrationEnum.md) |  | [optional] [default to PRORATIONENUM_ALWAYS_INVOICE]

## Methods

### NewCreateSubscriptionItemRequest

`func NewCreateSubscriptionItemRequest(subscriptionId string, priceId string, ) *CreateSubscriptionItemRequest`

NewCreateSubscriptionItemRequest instantiates a new CreateSubscriptionItemRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSubscriptionItemRequestWithDefaults

`func NewCreateSubscriptionItemRequestWithDefaults() *CreateSubscriptionItemRequest`

NewCreateSubscriptionItemRequestWithDefaults instantiates a new CreateSubscriptionItemRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionId

`func (o *CreateSubscriptionItemRequest) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *CreateSubscriptionItemRequest) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *CreateSubscriptionItemRequest) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetAddAtPeriodEnd

`func (o *CreateSubscriptionItemRequest) GetAddAtPeriodEnd() bool`

GetAddAtPeriodEnd returns the AddAtPeriodEnd field if non-nil, zero value otherwise.

### GetAddAtPeriodEndOk

`func (o *CreateSubscriptionItemRequest) GetAddAtPeriodEndOk() (*bool, bool)`

GetAddAtPeriodEndOk returns a tuple with the AddAtPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddAtPeriodEnd

`func (o *CreateSubscriptionItemRequest) SetAddAtPeriodEnd(v bool)`

SetAddAtPeriodEnd sets AddAtPeriodEnd field to given value.

### HasAddAtPeriodEnd

`func (o *CreateSubscriptionItemRequest) HasAddAtPeriodEnd() bool`

HasAddAtPeriodEnd returns a boolean if a field has been set.

### GetPriceId

`func (o *CreateSubscriptionItemRequest) GetPriceId() string`

GetPriceId returns the PriceId field if non-nil, zero value otherwise.

### GetPriceIdOk

`func (o *CreateSubscriptionItemRequest) GetPriceIdOk() (*string, bool)`

GetPriceIdOk returns a tuple with the PriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceId

`func (o *CreateSubscriptionItemRequest) SetPriceId(v string)`

SetPriceId sets PriceId field to given value.


### GetQuantity

`func (o *CreateSubscriptionItemRequest) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CreateSubscriptionItemRequest) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CreateSubscriptionItemRequest) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *CreateSubscriptionItemRequest) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetProrationBehavior

`func (o *CreateSubscriptionItemRequest) GetProrationBehavior() ProrationEnum`

GetProrationBehavior returns the ProrationBehavior field if non-nil, zero value otherwise.

### GetProrationBehaviorOk

`func (o *CreateSubscriptionItemRequest) GetProrationBehaviorOk() (*ProrationEnum, bool)`

GetProrationBehaviorOk returns a tuple with the ProrationBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProrationBehavior

`func (o *CreateSubscriptionItemRequest) SetProrationBehavior(v ProrationEnum)`

SetProrationBehavior sets ProrationBehavior field to given value.

### HasProrationBehavior

`func (o *CreateSubscriptionItemRequest) HasProrationBehavior() bool`

HasProrationBehavior returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


