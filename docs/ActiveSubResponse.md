# ActiveSubResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasActiveSubscriptions** | **bool** |  | 
**Subscriptions** | [**[]SubscriptionExternal**](SubscriptionExternal.md) |  | 

## Methods

### NewActiveSubResponse

`func NewActiveSubResponse(hasActiveSubscriptions bool, subscriptions []SubscriptionExternal, ) *ActiveSubResponse`

NewActiveSubResponse instantiates a new ActiveSubResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActiveSubResponseWithDefaults

`func NewActiveSubResponseWithDefaults() *ActiveSubResponse`

NewActiveSubResponseWithDefaults instantiates a new ActiveSubResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasActiveSubscriptions

`func (o *ActiveSubResponse) GetHasActiveSubscriptions() bool`

GetHasActiveSubscriptions returns the HasActiveSubscriptions field if non-nil, zero value otherwise.

### GetHasActiveSubscriptionsOk

`func (o *ActiveSubResponse) GetHasActiveSubscriptionsOk() (*bool, bool)`

GetHasActiveSubscriptionsOk returns a tuple with the HasActiveSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasActiveSubscriptions

`func (o *ActiveSubResponse) SetHasActiveSubscriptions(v bool)`

SetHasActiveSubscriptions sets HasActiveSubscriptions field to given value.


### GetSubscriptions

`func (o *ActiveSubResponse) GetSubscriptions() []SubscriptionExternal`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *ActiveSubResponse) GetSubscriptionsOk() (*[]SubscriptionExternal, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *ActiveSubResponse) SetSubscriptions(v []SubscriptionExternal)`

SetSubscriptions sets Subscriptions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


