# DiscountExternal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Coupon** | [**CouponExternal**](CouponExternal.md) |  | 
**CouponId** | **string** |  | 
**CreatedAt** | **time.Time** | DateTime at which the object was created, in &#39;ISO 8601&#39; format. | 
**CustomerId** | **NullableString** |  | 
**EndDate** | **NullableTime** |  | 
**Id** | **string** |  | 
**InvoiceId** | **NullableString** |  | 
**InvoiceItemId** | **NullableString** |  | 
**IsDeleted** | Pointer to **bool** | If true, indicates that this object has been deleted | [optional] [default to false]
**Object** | Pointer to [**ObjectName**](ObjectName.md) |  | [optional] 
**PromotionCode** | [**NullablePromotionCodeExternal**](PromotionCodeExternal.md) |  | 
**PromotionCodeId** | **NullableString** |  | 
**StartDate** | **time.Time** |  | 
**SubscriptionId** | **NullableString** |  | 
**UpdatedAt** | **time.Time** | DateTime at which the object was updated, in &#39;ISO 8601&#39; format. | 

## Methods

### NewDiscountExternal

`func NewDiscountExternal(coupon CouponExternal, couponId string, createdAt time.Time, customerId NullableString, endDate NullableTime, id string, invoiceId NullableString, invoiceItemId NullableString, promotionCode NullablePromotionCodeExternal, promotionCodeId NullableString, startDate time.Time, subscriptionId NullableString, updatedAt time.Time, ) *DiscountExternal`

NewDiscountExternal instantiates a new DiscountExternal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscountExternalWithDefaults

`func NewDiscountExternalWithDefaults() *DiscountExternal`

NewDiscountExternalWithDefaults instantiates a new DiscountExternal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoupon

`func (o *DiscountExternal) GetCoupon() CouponExternal`

GetCoupon returns the Coupon field if non-nil, zero value otherwise.

### GetCouponOk

`func (o *DiscountExternal) GetCouponOk() (*CouponExternal, bool)`

GetCouponOk returns a tuple with the Coupon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoupon

`func (o *DiscountExternal) SetCoupon(v CouponExternal)`

SetCoupon sets Coupon field to given value.


### GetCouponId

`func (o *DiscountExternal) GetCouponId() string`

GetCouponId returns the CouponId field if non-nil, zero value otherwise.

### GetCouponIdOk

`func (o *DiscountExternal) GetCouponIdOk() (*string, bool)`

GetCouponIdOk returns a tuple with the CouponId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponId

`func (o *DiscountExternal) SetCouponId(v string)`

SetCouponId sets CouponId field to given value.


### GetCreatedAt

`func (o *DiscountExternal) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DiscountExternal) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DiscountExternal) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCustomerId

`func (o *DiscountExternal) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *DiscountExternal) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *DiscountExternal) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### SetCustomerIdNil

`func (o *DiscountExternal) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *DiscountExternal) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetEndDate

`func (o *DiscountExternal) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *DiscountExternal) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *DiscountExternal) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.


### SetEndDateNil

`func (o *DiscountExternal) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *DiscountExternal) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetId

`func (o *DiscountExternal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DiscountExternal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DiscountExternal) SetId(v string)`

SetId sets Id field to given value.


### GetInvoiceId

`func (o *DiscountExternal) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *DiscountExternal) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *DiscountExternal) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### SetInvoiceIdNil

`func (o *DiscountExternal) SetInvoiceIdNil(b bool)`

 SetInvoiceIdNil sets the value for InvoiceId to be an explicit nil

### UnsetInvoiceId
`func (o *DiscountExternal) UnsetInvoiceId()`

UnsetInvoiceId ensures that no value is present for InvoiceId, not even an explicit nil
### GetInvoiceItemId

`func (o *DiscountExternal) GetInvoiceItemId() string`

GetInvoiceItemId returns the InvoiceItemId field if non-nil, zero value otherwise.

### GetInvoiceItemIdOk

`func (o *DiscountExternal) GetInvoiceItemIdOk() (*string, bool)`

GetInvoiceItemIdOk returns a tuple with the InvoiceItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceItemId

`func (o *DiscountExternal) SetInvoiceItemId(v string)`

SetInvoiceItemId sets InvoiceItemId field to given value.


### SetInvoiceItemIdNil

`func (o *DiscountExternal) SetInvoiceItemIdNil(b bool)`

 SetInvoiceItemIdNil sets the value for InvoiceItemId to be an explicit nil

### UnsetInvoiceItemId
`func (o *DiscountExternal) UnsetInvoiceItemId()`

UnsetInvoiceItemId ensures that no value is present for InvoiceItemId, not even an explicit nil
### GetIsDeleted

`func (o *DiscountExternal) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *DiscountExternal) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *DiscountExternal) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *DiscountExternal) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetObject

`func (o *DiscountExternal) GetObject() ObjectName`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *DiscountExternal) GetObjectOk() (*ObjectName, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *DiscountExternal) SetObject(v ObjectName)`

SetObject sets Object field to given value.

### HasObject

`func (o *DiscountExternal) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetPromotionCode

`func (o *DiscountExternal) GetPromotionCode() PromotionCodeExternal`

GetPromotionCode returns the PromotionCode field if non-nil, zero value otherwise.

### GetPromotionCodeOk

`func (o *DiscountExternal) GetPromotionCodeOk() (*PromotionCodeExternal, bool)`

GetPromotionCodeOk returns a tuple with the PromotionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotionCode

`func (o *DiscountExternal) SetPromotionCode(v PromotionCodeExternal)`

SetPromotionCode sets PromotionCode field to given value.


### SetPromotionCodeNil

`func (o *DiscountExternal) SetPromotionCodeNil(b bool)`

 SetPromotionCodeNil sets the value for PromotionCode to be an explicit nil

### UnsetPromotionCode
`func (o *DiscountExternal) UnsetPromotionCode()`

UnsetPromotionCode ensures that no value is present for PromotionCode, not even an explicit nil
### GetPromotionCodeId

`func (o *DiscountExternal) GetPromotionCodeId() string`

GetPromotionCodeId returns the PromotionCodeId field if non-nil, zero value otherwise.

### GetPromotionCodeIdOk

`func (o *DiscountExternal) GetPromotionCodeIdOk() (*string, bool)`

GetPromotionCodeIdOk returns a tuple with the PromotionCodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotionCodeId

`func (o *DiscountExternal) SetPromotionCodeId(v string)`

SetPromotionCodeId sets PromotionCodeId field to given value.


### SetPromotionCodeIdNil

`func (o *DiscountExternal) SetPromotionCodeIdNil(b bool)`

 SetPromotionCodeIdNil sets the value for PromotionCodeId to be an explicit nil

### UnsetPromotionCodeId
`func (o *DiscountExternal) UnsetPromotionCodeId()`

UnsetPromotionCodeId ensures that no value is present for PromotionCodeId, not even an explicit nil
### GetStartDate

`func (o *DiscountExternal) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *DiscountExternal) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *DiscountExternal) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.


### GetSubscriptionId

`func (o *DiscountExternal) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *DiscountExternal) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *DiscountExternal) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### SetSubscriptionIdNil

`func (o *DiscountExternal) SetSubscriptionIdNil(b bool)`

 SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil

### UnsetSubscriptionId
`func (o *DiscountExternal) UnsetSubscriptionId()`

UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
### GetUpdatedAt

`func (o *DiscountExternal) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DiscountExternal) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DiscountExternal) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


