# UpdateCouponRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**IsActive** | Pointer to **NullableBool** |  | [optional] 
**RedeemBy** | Pointer to **NullableTime** |  | [optional] 
**MaxRedemptions** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewUpdateCouponRequest

`func NewUpdateCouponRequest() *UpdateCouponRequest`

NewUpdateCouponRequest instantiates a new UpdateCouponRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCouponRequestWithDefaults

`func NewUpdateCouponRequestWithDefaults() *UpdateCouponRequest`

NewUpdateCouponRequestWithDefaults instantiates a new UpdateCouponRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateCouponRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateCouponRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateCouponRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateCouponRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateCouponRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateCouponRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIsActive

`func (o *UpdateCouponRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *UpdateCouponRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *UpdateCouponRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *UpdateCouponRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *UpdateCouponRequest) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *UpdateCouponRequest) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetRedeemBy

`func (o *UpdateCouponRequest) GetRedeemBy() time.Time`

GetRedeemBy returns the RedeemBy field if non-nil, zero value otherwise.

### GetRedeemByOk

`func (o *UpdateCouponRequest) GetRedeemByOk() (*time.Time, bool)`

GetRedeemByOk returns a tuple with the RedeemBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemBy

`func (o *UpdateCouponRequest) SetRedeemBy(v time.Time)`

SetRedeemBy sets RedeemBy field to given value.

### HasRedeemBy

`func (o *UpdateCouponRequest) HasRedeemBy() bool`

HasRedeemBy returns a boolean if a field has been set.

### SetRedeemByNil

`func (o *UpdateCouponRequest) SetRedeemByNil(b bool)`

 SetRedeemByNil sets the value for RedeemBy to be an explicit nil

### UnsetRedeemBy
`func (o *UpdateCouponRequest) UnsetRedeemBy()`

UnsetRedeemBy ensures that no value is present for RedeemBy, not even an explicit nil
### GetMaxRedemptions

`func (o *UpdateCouponRequest) GetMaxRedemptions() int32`

GetMaxRedemptions returns the MaxRedemptions field if non-nil, zero value otherwise.

### GetMaxRedemptionsOk

`func (o *UpdateCouponRequest) GetMaxRedemptionsOk() (*int32, bool)`

GetMaxRedemptionsOk returns a tuple with the MaxRedemptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRedemptions

`func (o *UpdateCouponRequest) SetMaxRedemptions(v int32)`

SetMaxRedemptions sets MaxRedemptions field to given value.

### HasMaxRedemptions

`func (o *UpdateCouponRequest) HasMaxRedemptions() bool`

HasMaxRedemptions returns a boolean if a field has been set.

### SetMaxRedemptionsNil

`func (o *UpdateCouponRequest) SetMaxRedemptionsNil(b bool)`

 SetMaxRedemptionsNil sets the value for MaxRedemptions to be an explicit nil

### UnsetMaxRedemptions
`func (o *UpdateCouponRequest) UnsetMaxRedemptions()`

UnsetMaxRedemptions ensures that no value is present for MaxRedemptions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


