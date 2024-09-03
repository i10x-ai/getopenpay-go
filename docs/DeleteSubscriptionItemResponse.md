# DeleteSubscriptionItemResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Message describing result of API call. | [optional] [default to "SubscriptionItem deleted successfully."]
**SubscriptionItemId** | **string** | Unique identifier of the subscription_item. | 
**DeletedAt** | **NullableTime** |  | 
**DropAtEnd** | Pointer to **bool** | Whether or not this item will be dropped from subscription before next renewal | [optional] [default to false]

## Methods

### NewDeleteSubscriptionItemResponse

`func NewDeleteSubscriptionItemResponse(subscriptionItemId string, deletedAt NullableTime, ) *DeleteSubscriptionItemResponse`

NewDeleteSubscriptionItemResponse instantiates a new DeleteSubscriptionItemResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteSubscriptionItemResponseWithDefaults

`func NewDeleteSubscriptionItemResponseWithDefaults() *DeleteSubscriptionItemResponse`

NewDeleteSubscriptionItemResponseWithDefaults instantiates a new DeleteSubscriptionItemResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *DeleteSubscriptionItemResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DeleteSubscriptionItemResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DeleteSubscriptionItemResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DeleteSubscriptionItemResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSubscriptionItemId

`func (o *DeleteSubscriptionItemResponse) GetSubscriptionItemId() string`

GetSubscriptionItemId returns the SubscriptionItemId field if non-nil, zero value otherwise.

### GetSubscriptionItemIdOk

`func (o *DeleteSubscriptionItemResponse) GetSubscriptionItemIdOk() (*string, bool)`

GetSubscriptionItemIdOk returns a tuple with the SubscriptionItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionItemId

`func (o *DeleteSubscriptionItemResponse) SetSubscriptionItemId(v string)`

SetSubscriptionItemId sets SubscriptionItemId field to given value.


### GetDeletedAt

`func (o *DeleteSubscriptionItemResponse) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *DeleteSubscriptionItemResponse) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *DeleteSubscriptionItemResponse) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.


### SetDeletedAtNil

`func (o *DeleteSubscriptionItemResponse) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *DeleteSubscriptionItemResponse) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetDropAtEnd

`func (o *DeleteSubscriptionItemResponse) GetDropAtEnd() bool`

GetDropAtEnd returns the DropAtEnd field if non-nil, zero value otherwise.

### GetDropAtEndOk

`func (o *DeleteSubscriptionItemResponse) GetDropAtEndOk() (*bool, bool)`

GetDropAtEndOk returns a tuple with the DropAtEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropAtEnd

`func (o *DeleteSubscriptionItemResponse) SetDropAtEnd(v bool)`

SetDropAtEnd sets DropAtEnd field to given value.

### HasDropAtEnd

`func (o *DeleteSubscriptionItemResponse) HasDropAtEnd() bool`

HasDropAtEnd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


