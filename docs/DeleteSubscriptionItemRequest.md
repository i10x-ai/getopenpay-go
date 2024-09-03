# DeleteSubscriptionItemRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DropAtEnd** | Pointer to **bool** | Whether or not this item will be dropped from subscription before next renewal | [optional] [default to true]
**ProrationBehavior** | Pointer to [**ProrationEnum**](ProrationEnum.md) |  | [optional] [default to PRORATIONENUM_ALWAYS_INVOICE]

## Methods

### NewDeleteSubscriptionItemRequest

`func NewDeleteSubscriptionItemRequest() *DeleteSubscriptionItemRequest`

NewDeleteSubscriptionItemRequest instantiates a new DeleteSubscriptionItemRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteSubscriptionItemRequestWithDefaults

`func NewDeleteSubscriptionItemRequestWithDefaults() *DeleteSubscriptionItemRequest`

NewDeleteSubscriptionItemRequestWithDefaults instantiates a new DeleteSubscriptionItemRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDropAtEnd

`func (o *DeleteSubscriptionItemRequest) GetDropAtEnd() bool`

GetDropAtEnd returns the DropAtEnd field if non-nil, zero value otherwise.

### GetDropAtEndOk

`func (o *DeleteSubscriptionItemRequest) GetDropAtEndOk() (*bool, bool)`

GetDropAtEndOk returns a tuple with the DropAtEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropAtEnd

`func (o *DeleteSubscriptionItemRequest) SetDropAtEnd(v bool)`

SetDropAtEnd sets DropAtEnd field to given value.

### HasDropAtEnd

`func (o *DeleteSubscriptionItemRequest) HasDropAtEnd() bool`

HasDropAtEnd returns a boolean if a field has been set.

### GetProrationBehavior

`func (o *DeleteSubscriptionItemRequest) GetProrationBehavior() ProrationEnum`

GetProrationBehavior returns the ProrationBehavior field if non-nil, zero value otherwise.

### GetProrationBehaviorOk

`func (o *DeleteSubscriptionItemRequest) GetProrationBehaviorOk() (*ProrationEnum, bool)`

GetProrationBehaviorOk returns a tuple with the ProrationBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProrationBehavior

`func (o *DeleteSubscriptionItemRequest) SetProrationBehavior(v ProrationEnum)`

SetProrationBehavior sets ProrationBehavior field to given value.

### HasProrationBehavior

`func (o *DeleteSubscriptionItemRequest) HasProrationBehavior() bool`

HasProrationBehavior returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


