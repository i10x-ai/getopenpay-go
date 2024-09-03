# PromotionCodeExternal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Object** | Pointer to [**ObjectName**](ObjectName.md) |  | [optional] [default to OBJECTNAME_PROMOTION_CODE]
**CreatedAt** | **time.Time** | DateTime at which the object was created, in &#39;ISO 8601&#39; format. | 
**UpdatedAt** | **time.Time** | DateTime at which the object was updated, in &#39;ISO 8601&#39; format. | 
**IsDeleted** | Pointer to **bool** | If true, indicates that this object has been deleted | [optional] [default to false]
**Code** | **string** |  | 
**IsActive** | **bool** |  | 
**CouponId** | **string** |  | 
**Coupon** | [**CouponExternal**](CouponExternal.md) |  | 
**CustomerIds** | Pointer to **[]string** |  | [optional] 
**ExpiresAt** | Pointer to **NullableTime** |  | [optional] 
**MaxRedemptions** | Pointer to **NullableInt32** |  | [optional] 
**TimesRedeemed** | Pointer to **NullableInt32** |  | [optional] 
**MaxRedemptionsPerCustomer** | Pointer to **NullableInt32** |  | [optional] 
**Restrictions** | [**NullablePromoRestrictions**](PromoRestrictions.md) |  | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPromotionCodeExternal

`func NewPromotionCodeExternal(id string, createdAt time.Time, updatedAt time.Time, code string, isActive bool, couponId string, coupon CouponExternal, restrictions NullablePromoRestrictions, ) *PromotionCodeExternal`

NewPromotionCodeExternal instantiates a new PromotionCodeExternal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromotionCodeExternalWithDefaults

`func NewPromotionCodeExternalWithDefaults() *PromotionCodeExternal`

NewPromotionCodeExternalWithDefaults instantiates a new PromotionCodeExternal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PromotionCodeExternal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PromotionCodeExternal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PromotionCodeExternal) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *PromotionCodeExternal) GetObject() ObjectName`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *PromotionCodeExternal) GetObjectOk() (*ObjectName, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *PromotionCodeExternal) SetObject(v ObjectName)`

SetObject sets Object field to given value.

### HasObject

`func (o *PromotionCodeExternal) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PromotionCodeExternal) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PromotionCodeExternal) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PromotionCodeExternal) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *PromotionCodeExternal) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PromotionCodeExternal) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PromotionCodeExternal) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetIsDeleted

`func (o *PromotionCodeExternal) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *PromotionCodeExternal) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *PromotionCodeExternal) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *PromotionCodeExternal) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetCode

`func (o *PromotionCodeExternal) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PromotionCodeExternal) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PromotionCodeExternal) SetCode(v string)`

SetCode sets Code field to given value.


### GetIsActive

`func (o *PromotionCodeExternal) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *PromotionCodeExternal) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *PromotionCodeExternal) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetCouponId

`func (o *PromotionCodeExternal) GetCouponId() string`

GetCouponId returns the CouponId field if non-nil, zero value otherwise.

### GetCouponIdOk

`func (o *PromotionCodeExternal) GetCouponIdOk() (*string, bool)`

GetCouponIdOk returns a tuple with the CouponId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponId

`func (o *PromotionCodeExternal) SetCouponId(v string)`

SetCouponId sets CouponId field to given value.


### GetCoupon

`func (o *PromotionCodeExternal) GetCoupon() CouponExternal`

GetCoupon returns the Coupon field if non-nil, zero value otherwise.

### GetCouponOk

`func (o *PromotionCodeExternal) GetCouponOk() (*CouponExternal, bool)`

GetCouponOk returns a tuple with the Coupon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoupon

`func (o *PromotionCodeExternal) SetCoupon(v CouponExternal)`

SetCoupon sets Coupon field to given value.


### GetCustomerIds

`func (o *PromotionCodeExternal) GetCustomerIds() []string`

GetCustomerIds returns the CustomerIds field if non-nil, zero value otherwise.

### GetCustomerIdsOk

`func (o *PromotionCodeExternal) GetCustomerIdsOk() (*[]string, bool)`

GetCustomerIdsOk returns a tuple with the CustomerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerIds

`func (o *PromotionCodeExternal) SetCustomerIds(v []string)`

SetCustomerIds sets CustomerIds field to given value.

### HasCustomerIds

`func (o *PromotionCodeExternal) HasCustomerIds() bool`

HasCustomerIds returns a boolean if a field has been set.

### GetExpiresAt

`func (o *PromotionCodeExternal) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *PromotionCodeExternal) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *PromotionCodeExternal) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *PromotionCodeExternal) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *PromotionCodeExternal) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *PromotionCodeExternal) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetMaxRedemptions

`func (o *PromotionCodeExternal) GetMaxRedemptions() int32`

GetMaxRedemptions returns the MaxRedemptions field if non-nil, zero value otherwise.

### GetMaxRedemptionsOk

`func (o *PromotionCodeExternal) GetMaxRedemptionsOk() (*int32, bool)`

GetMaxRedemptionsOk returns a tuple with the MaxRedemptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRedemptions

`func (o *PromotionCodeExternal) SetMaxRedemptions(v int32)`

SetMaxRedemptions sets MaxRedemptions field to given value.

### HasMaxRedemptions

`func (o *PromotionCodeExternal) HasMaxRedemptions() bool`

HasMaxRedemptions returns a boolean if a field has been set.

### SetMaxRedemptionsNil

`func (o *PromotionCodeExternal) SetMaxRedemptionsNil(b bool)`

 SetMaxRedemptionsNil sets the value for MaxRedemptions to be an explicit nil

### UnsetMaxRedemptions
`func (o *PromotionCodeExternal) UnsetMaxRedemptions()`

UnsetMaxRedemptions ensures that no value is present for MaxRedemptions, not even an explicit nil
### GetTimesRedeemed

`func (o *PromotionCodeExternal) GetTimesRedeemed() int32`

GetTimesRedeemed returns the TimesRedeemed field if non-nil, zero value otherwise.

### GetTimesRedeemedOk

`func (o *PromotionCodeExternal) GetTimesRedeemedOk() (*int32, bool)`

GetTimesRedeemedOk returns a tuple with the TimesRedeemed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimesRedeemed

`func (o *PromotionCodeExternal) SetTimesRedeemed(v int32)`

SetTimesRedeemed sets TimesRedeemed field to given value.

### HasTimesRedeemed

`func (o *PromotionCodeExternal) HasTimesRedeemed() bool`

HasTimesRedeemed returns a boolean if a field has been set.

### SetTimesRedeemedNil

`func (o *PromotionCodeExternal) SetTimesRedeemedNil(b bool)`

 SetTimesRedeemedNil sets the value for TimesRedeemed to be an explicit nil

### UnsetTimesRedeemed
`func (o *PromotionCodeExternal) UnsetTimesRedeemed()`

UnsetTimesRedeemed ensures that no value is present for TimesRedeemed, not even an explicit nil
### GetMaxRedemptionsPerCustomer

`func (o *PromotionCodeExternal) GetMaxRedemptionsPerCustomer() int32`

GetMaxRedemptionsPerCustomer returns the MaxRedemptionsPerCustomer field if non-nil, zero value otherwise.

### GetMaxRedemptionsPerCustomerOk

`func (o *PromotionCodeExternal) GetMaxRedemptionsPerCustomerOk() (*int32, bool)`

GetMaxRedemptionsPerCustomerOk returns a tuple with the MaxRedemptionsPerCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRedemptionsPerCustomer

`func (o *PromotionCodeExternal) SetMaxRedemptionsPerCustomer(v int32)`

SetMaxRedemptionsPerCustomer sets MaxRedemptionsPerCustomer field to given value.

### HasMaxRedemptionsPerCustomer

`func (o *PromotionCodeExternal) HasMaxRedemptionsPerCustomer() bool`

HasMaxRedemptionsPerCustomer returns a boolean if a field has been set.

### SetMaxRedemptionsPerCustomerNil

`func (o *PromotionCodeExternal) SetMaxRedemptionsPerCustomerNil(b bool)`

 SetMaxRedemptionsPerCustomerNil sets the value for MaxRedemptionsPerCustomer to be an explicit nil

### UnsetMaxRedemptionsPerCustomer
`func (o *PromotionCodeExternal) UnsetMaxRedemptionsPerCustomer()`

UnsetMaxRedemptionsPerCustomer ensures that no value is present for MaxRedemptionsPerCustomer, not even an explicit nil
### GetRestrictions

`func (o *PromotionCodeExternal) GetRestrictions() PromoRestrictions`

GetRestrictions returns the Restrictions field if non-nil, zero value otherwise.

### GetRestrictionsOk

`func (o *PromotionCodeExternal) GetRestrictionsOk() (*PromoRestrictions, bool)`

GetRestrictionsOk returns a tuple with the Restrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictions

`func (o *PromotionCodeExternal) SetRestrictions(v PromoRestrictions)`

SetRestrictions sets Restrictions field to given value.


### SetRestrictionsNil

`func (o *PromotionCodeExternal) SetRestrictionsNil(b bool)`

 SetRestrictionsNil sets the value for Restrictions to be an explicit nil

### UnsetRestrictions
`func (o *PromotionCodeExternal) UnsetRestrictions()`

UnsetRestrictions ensures that no value is present for Restrictions, not even an explicit nil
### GetMetadata

`func (o *PromotionCodeExternal) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PromotionCodeExternal) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PromotionCodeExternal) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PromotionCodeExternal) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PromotionCodeExternal) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PromotionCodeExternal) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


