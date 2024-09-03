# UpdatePromoCodeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsActive** | **bool** | Whether the promotion code is currently active. A promotion code can only be reactivated when the coupon is still valid and the promotion code is otherwise redeemable. | 

## Methods

### NewUpdatePromoCodeRequest

`func NewUpdatePromoCodeRequest(isActive bool, ) *UpdatePromoCodeRequest`

NewUpdatePromoCodeRequest instantiates a new UpdatePromoCodeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePromoCodeRequestWithDefaults

`func NewUpdatePromoCodeRequestWithDefaults() *UpdatePromoCodeRequest`

NewUpdatePromoCodeRequestWithDefaults instantiates a new UpdatePromoCodeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsActive

`func (o *UpdatePromoCodeRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *UpdatePromoCodeRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *UpdatePromoCodeRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


