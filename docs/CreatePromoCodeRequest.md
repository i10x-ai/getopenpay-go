# CreatePromoCodeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CouponId** | **string** | The unique id of coupon for this promotion code. | 
**Code** | **string** | The customer-facing code. This code must be unique across all active promotion codes for a specific customer. Case insensitive. | 
**CustomerIds** | Pointer to **[]string** |  | [optional] 
**ExpiresAt** | Pointer to **NullableTime** |  | [optional] 
**MaxRedemptions** | Pointer to **NullableInt32** |  | [optional] 
**MaxRedemptionsPerCustomer** | Pointer to **NullableInt32** |  | [optional] 
**Restrictions** | Pointer to [**NullablePromoRestrictions**](PromoRestrictions.md) |  | [optional] 

## Methods

### NewCreatePromoCodeRequest

`func NewCreatePromoCodeRequest(couponId string, code string, ) *CreatePromoCodeRequest`

NewCreatePromoCodeRequest instantiates a new CreatePromoCodeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePromoCodeRequestWithDefaults

`func NewCreatePromoCodeRequestWithDefaults() *CreatePromoCodeRequest`

NewCreatePromoCodeRequestWithDefaults instantiates a new CreatePromoCodeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCouponId

`func (o *CreatePromoCodeRequest) GetCouponId() string`

GetCouponId returns the CouponId field if non-nil, zero value otherwise.

### GetCouponIdOk

`func (o *CreatePromoCodeRequest) GetCouponIdOk() (*string, bool)`

GetCouponIdOk returns a tuple with the CouponId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponId

`func (o *CreatePromoCodeRequest) SetCouponId(v string)`

SetCouponId sets CouponId field to given value.


### GetCode

`func (o *CreatePromoCodeRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CreatePromoCodeRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CreatePromoCodeRequest) SetCode(v string)`

SetCode sets Code field to given value.


### GetCustomerIds

`func (o *CreatePromoCodeRequest) GetCustomerIds() []string`

GetCustomerIds returns the CustomerIds field if non-nil, zero value otherwise.

### GetCustomerIdsOk

`func (o *CreatePromoCodeRequest) GetCustomerIdsOk() (*[]string, bool)`

GetCustomerIdsOk returns a tuple with the CustomerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerIds

`func (o *CreatePromoCodeRequest) SetCustomerIds(v []string)`

SetCustomerIds sets CustomerIds field to given value.

### HasCustomerIds

`func (o *CreatePromoCodeRequest) HasCustomerIds() bool`

HasCustomerIds returns a boolean if a field has been set.

### GetExpiresAt

`func (o *CreatePromoCodeRequest) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CreatePromoCodeRequest) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CreatePromoCodeRequest) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *CreatePromoCodeRequest) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *CreatePromoCodeRequest) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *CreatePromoCodeRequest) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetMaxRedemptions

`func (o *CreatePromoCodeRequest) GetMaxRedemptions() int32`

GetMaxRedemptions returns the MaxRedemptions field if non-nil, zero value otherwise.

### GetMaxRedemptionsOk

`func (o *CreatePromoCodeRequest) GetMaxRedemptionsOk() (*int32, bool)`

GetMaxRedemptionsOk returns a tuple with the MaxRedemptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRedemptions

`func (o *CreatePromoCodeRequest) SetMaxRedemptions(v int32)`

SetMaxRedemptions sets MaxRedemptions field to given value.

### HasMaxRedemptions

`func (o *CreatePromoCodeRequest) HasMaxRedemptions() bool`

HasMaxRedemptions returns a boolean if a field has been set.

### SetMaxRedemptionsNil

`func (o *CreatePromoCodeRequest) SetMaxRedemptionsNil(b bool)`

 SetMaxRedemptionsNil sets the value for MaxRedemptions to be an explicit nil

### UnsetMaxRedemptions
`func (o *CreatePromoCodeRequest) UnsetMaxRedemptions()`

UnsetMaxRedemptions ensures that no value is present for MaxRedemptions, not even an explicit nil
### GetMaxRedemptionsPerCustomer

`func (o *CreatePromoCodeRequest) GetMaxRedemptionsPerCustomer() int32`

GetMaxRedemptionsPerCustomer returns the MaxRedemptionsPerCustomer field if non-nil, zero value otherwise.

### GetMaxRedemptionsPerCustomerOk

`func (o *CreatePromoCodeRequest) GetMaxRedemptionsPerCustomerOk() (*int32, bool)`

GetMaxRedemptionsPerCustomerOk returns a tuple with the MaxRedemptionsPerCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRedemptionsPerCustomer

`func (o *CreatePromoCodeRequest) SetMaxRedemptionsPerCustomer(v int32)`

SetMaxRedemptionsPerCustomer sets MaxRedemptionsPerCustomer field to given value.

### HasMaxRedemptionsPerCustomer

`func (o *CreatePromoCodeRequest) HasMaxRedemptionsPerCustomer() bool`

HasMaxRedemptionsPerCustomer returns a boolean if a field has been set.

### SetMaxRedemptionsPerCustomerNil

`func (o *CreatePromoCodeRequest) SetMaxRedemptionsPerCustomerNil(b bool)`

 SetMaxRedemptionsPerCustomerNil sets the value for MaxRedemptionsPerCustomer to be an explicit nil

### UnsetMaxRedemptionsPerCustomer
`func (o *CreatePromoCodeRequest) UnsetMaxRedemptionsPerCustomer()`

UnsetMaxRedemptionsPerCustomer ensures that no value is present for MaxRedemptionsPerCustomer, not even an explicit nil
### GetRestrictions

`func (o *CreatePromoCodeRequest) GetRestrictions() PromoRestrictions`

GetRestrictions returns the Restrictions field if non-nil, zero value otherwise.

### GetRestrictionsOk

`func (o *CreatePromoCodeRequest) GetRestrictionsOk() (*PromoRestrictions, bool)`

GetRestrictionsOk returns a tuple with the Restrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictions

`func (o *CreatePromoCodeRequest) SetRestrictions(v PromoRestrictions)`

SetRestrictions sets Restrictions field to given value.

### HasRestrictions

`func (o *CreatePromoCodeRequest) HasRestrictions() bool`

HasRestrictions returns a boolean if a field has been set.

### SetRestrictionsNil

`func (o *CreatePromoCodeRequest) SetRestrictionsNil(b bool)`

 SetRestrictionsNil sets the value for Restrictions to be an explicit nil

### UnsetRestrictions
`func (o *CreatePromoCodeRequest) UnsetRestrictions()`

UnsetRestrictions ensures that no value is present for Restrictions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


