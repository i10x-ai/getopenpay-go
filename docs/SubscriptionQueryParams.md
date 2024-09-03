# SubscriptionQueryParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageNumber** | Pointer to **int32** | Page number | [optional] [default to 1]
**PageSize** | Pointer to **int32** | Page size | [optional] [default to 100]
**SortKey** | Pointer to **string** | Key name based on which data is sorted. | [optional] [default to "created_at"]
**SortDescending** | Pointer to **bool** | Sort direction. | [optional] [default to false]
**CreatedAt** | Pointer to [**NullableDateTimeFilter**](DateTimeFilter.md) |  | [optional] 
**UpdatedAt** | Pointer to [**NullableDateTimeFilter**](DateTimeFilter.md) |  | [optional] 
**Expand** | Pointer to **[]string** |  | [optional] 
**CustomerId** | Pointer to **NullableString** |  | [optional] 
**ProductId** | Pointer to **NullableString** |  | [optional] 
**PriceId** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to [**NullableSubscriptionStatusEnum**](SubscriptionStatusEnum.md) |  | [optional] 
**CurrentPeriodStart** | Pointer to [**NullableDateTimeFilter**](DateTimeFilter.md) |  | [optional] 
**CurrentPeriodEnd** | Pointer to [**NullableDateTimeFilter**](DateTimeFilter.md) |  | [optional] 
**CouponId** | Pointer to **NullableString** |  | [optional] 
**Search** | Pointer to [**NullableSearchFilter**](SearchFilter.md) |  | [optional] 

## Methods

### NewSubscriptionQueryParams

`func NewSubscriptionQueryParams() *SubscriptionQueryParams`

NewSubscriptionQueryParams instantiates a new SubscriptionQueryParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionQueryParamsWithDefaults

`func NewSubscriptionQueryParamsWithDefaults() *SubscriptionQueryParams`

NewSubscriptionQueryParamsWithDefaults instantiates a new SubscriptionQueryParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageNumber

`func (o *SubscriptionQueryParams) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *SubscriptionQueryParams) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *SubscriptionQueryParams) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *SubscriptionQueryParams) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetPageSize

`func (o *SubscriptionQueryParams) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *SubscriptionQueryParams) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *SubscriptionQueryParams) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *SubscriptionQueryParams) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetSortKey

`func (o *SubscriptionQueryParams) GetSortKey() string`

GetSortKey returns the SortKey field if non-nil, zero value otherwise.

### GetSortKeyOk

`func (o *SubscriptionQueryParams) GetSortKeyOk() (*string, bool)`

GetSortKeyOk returns a tuple with the SortKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortKey

`func (o *SubscriptionQueryParams) SetSortKey(v string)`

SetSortKey sets SortKey field to given value.

### HasSortKey

`func (o *SubscriptionQueryParams) HasSortKey() bool`

HasSortKey returns a boolean if a field has been set.

### GetSortDescending

`func (o *SubscriptionQueryParams) GetSortDescending() bool`

GetSortDescending returns the SortDescending field if non-nil, zero value otherwise.

### GetSortDescendingOk

`func (o *SubscriptionQueryParams) GetSortDescendingOk() (*bool, bool)`

GetSortDescendingOk returns a tuple with the SortDescending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortDescending

`func (o *SubscriptionQueryParams) SetSortDescending(v bool)`

SetSortDescending sets SortDescending field to given value.

### HasSortDescending

`func (o *SubscriptionQueryParams) HasSortDescending() bool`

HasSortDescending returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SubscriptionQueryParams) GetCreatedAt() DateTimeFilter`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SubscriptionQueryParams) GetCreatedAtOk() (*DateTimeFilter, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SubscriptionQueryParams) SetCreatedAt(v DateTimeFilter)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SubscriptionQueryParams) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *SubscriptionQueryParams) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *SubscriptionQueryParams) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *SubscriptionQueryParams) GetUpdatedAt() DateTimeFilter`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SubscriptionQueryParams) GetUpdatedAtOk() (*DateTimeFilter, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SubscriptionQueryParams) SetUpdatedAt(v DateTimeFilter)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SubscriptionQueryParams) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *SubscriptionQueryParams) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *SubscriptionQueryParams) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetExpand

`func (o *SubscriptionQueryParams) GetExpand() []string`

GetExpand returns the Expand field if non-nil, zero value otherwise.

### GetExpandOk

`func (o *SubscriptionQueryParams) GetExpandOk() (*[]string, bool)`

GetExpandOk returns a tuple with the Expand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpand

`func (o *SubscriptionQueryParams) SetExpand(v []string)`

SetExpand sets Expand field to given value.

### HasExpand

`func (o *SubscriptionQueryParams) HasExpand() bool`

HasExpand returns a boolean if a field has been set.

### GetCustomerId

`func (o *SubscriptionQueryParams) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *SubscriptionQueryParams) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *SubscriptionQueryParams) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *SubscriptionQueryParams) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### SetCustomerIdNil

`func (o *SubscriptionQueryParams) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *SubscriptionQueryParams) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetProductId

`func (o *SubscriptionQueryParams) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *SubscriptionQueryParams) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *SubscriptionQueryParams) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *SubscriptionQueryParams) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### SetProductIdNil

`func (o *SubscriptionQueryParams) SetProductIdNil(b bool)`

 SetProductIdNil sets the value for ProductId to be an explicit nil

### UnsetProductId
`func (o *SubscriptionQueryParams) UnsetProductId()`

UnsetProductId ensures that no value is present for ProductId, not even an explicit nil
### GetPriceId

`func (o *SubscriptionQueryParams) GetPriceId() string`

GetPriceId returns the PriceId field if non-nil, zero value otherwise.

### GetPriceIdOk

`func (o *SubscriptionQueryParams) GetPriceIdOk() (*string, bool)`

GetPriceIdOk returns a tuple with the PriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceId

`func (o *SubscriptionQueryParams) SetPriceId(v string)`

SetPriceId sets PriceId field to given value.

### HasPriceId

`func (o *SubscriptionQueryParams) HasPriceId() bool`

HasPriceId returns a boolean if a field has been set.

### SetPriceIdNil

`func (o *SubscriptionQueryParams) SetPriceIdNil(b bool)`

 SetPriceIdNil sets the value for PriceId to be an explicit nil

### UnsetPriceId
`func (o *SubscriptionQueryParams) UnsetPriceId()`

UnsetPriceId ensures that no value is present for PriceId, not even an explicit nil
### GetStatus

`func (o *SubscriptionQueryParams) GetStatus() SubscriptionStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubscriptionQueryParams) GetStatusOk() (*SubscriptionStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubscriptionQueryParams) SetStatus(v SubscriptionStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SubscriptionQueryParams) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *SubscriptionQueryParams) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *SubscriptionQueryParams) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetCurrentPeriodStart

`func (o *SubscriptionQueryParams) GetCurrentPeriodStart() DateTimeFilter`

GetCurrentPeriodStart returns the CurrentPeriodStart field if non-nil, zero value otherwise.

### GetCurrentPeriodStartOk

`func (o *SubscriptionQueryParams) GetCurrentPeriodStartOk() (*DateTimeFilter, bool)`

GetCurrentPeriodStartOk returns a tuple with the CurrentPeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriodStart

`func (o *SubscriptionQueryParams) SetCurrentPeriodStart(v DateTimeFilter)`

SetCurrentPeriodStart sets CurrentPeriodStart field to given value.

### HasCurrentPeriodStart

`func (o *SubscriptionQueryParams) HasCurrentPeriodStart() bool`

HasCurrentPeriodStart returns a boolean if a field has been set.

### SetCurrentPeriodStartNil

`func (o *SubscriptionQueryParams) SetCurrentPeriodStartNil(b bool)`

 SetCurrentPeriodStartNil sets the value for CurrentPeriodStart to be an explicit nil

### UnsetCurrentPeriodStart
`func (o *SubscriptionQueryParams) UnsetCurrentPeriodStart()`

UnsetCurrentPeriodStart ensures that no value is present for CurrentPeriodStart, not even an explicit nil
### GetCurrentPeriodEnd

`func (o *SubscriptionQueryParams) GetCurrentPeriodEnd() DateTimeFilter`

GetCurrentPeriodEnd returns the CurrentPeriodEnd field if non-nil, zero value otherwise.

### GetCurrentPeriodEndOk

`func (o *SubscriptionQueryParams) GetCurrentPeriodEndOk() (*DateTimeFilter, bool)`

GetCurrentPeriodEndOk returns a tuple with the CurrentPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriodEnd

`func (o *SubscriptionQueryParams) SetCurrentPeriodEnd(v DateTimeFilter)`

SetCurrentPeriodEnd sets CurrentPeriodEnd field to given value.

### HasCurrentPeriodEnd

`func (o *SubscriptionQueryParams) HasCurrentPeriodEnd() bool`

HasCurrentPeriodEnd returns a boolean if a field has been set.

### SetCurrentPeriodEndNil

`func (o *SubscriptionQueryParams) SetCurrentPeriodEndNil(b bool)`

 SetCurrentPeriodEndNil sets the value for CurrentPeriodEnd to be an explicit nil

### UnsetCurrentPeriodEnd
`func (o *SubscriptionQueryParams) UnsetCurrentPeriodEnd()`

UnsetCurrentPeriodEnd ensures that no value is present for CurrentPeriodEnd, not even an explicit nil
### GetCouponId

`func (o *SubscriptionQueryParams) GetCouponId() string`

GetCouponId returns the CouponId field if non-nil, zero value otherwise.

### GetCouponIdOk

`func (o *SubscriptionQueryParams) GetCouponIdOk() (*string, bool)`

GetCouponIdOk returns a tuple with the CouponId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponId

`func (o *SubscriptionQueryParams) SetCouponId(v string)`

SetCouponId sets CouponId field to given value.

### HasCouponId

`func (o *SubscriptionQueryParams) HasCouponId() bool`

HasCouponId returns a boolean if a field has been set.

### SetCouponIdNil

`func (o *SubscriptionQueryParams) SetCouponIdNil(b bool)`

 SetCouponIdNil sets the value for CouponId to be an explicit nil

### UnsetCouponId
`func (o *SubscriptionQueryParams) UnsetCouponId()`

UnsetCouponId ensures that no value is present for CouponId, not even an explicit nil
### GetSearch

`func (o *SubscriptionQueryParams) GetSearch() SearchFilter`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *SubscriptionQueryParams) GetSearchOk() (*SearchFilter, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *SubscriptionQueryParams) SetSearch(v SearchFilter)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *SubscriptionQueryParams) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

### SetSearchNil

`func (o *SubscriptionQueryParams) SetSearchNil(b bool)`

 SetSearchNil sets the value for Search to be an explicit nil

### UnsetSearch
`func (o *SubscriptionQueryParams) UnsetSearch()`

UnsetSearch ensures that no value is present for Search, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


