# CouponQueryParams

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
**IsActive** | Pointer to **NullableBool** |  | [optional] 
**Search** | Pointer to [**NullableSearchFilter**](SearchFilter.md) |  | [optional] 
**Currency** | Pointer to [**NullableCurrencyEnum**](CurrencyEnum.md) |  | [optional] 
**Duration** | Pointer to [**NullableCouponDuration**](CouponDuration.md) |  | [optional] 
**MaxRedemptions** | Pointer to [**NullableIntRangeFilter**](IntRangeFilter.md) |  | [optional] 
**RedeemBy** | Pointer to [**NullableDateTimeFilter**](DateTimeFilter.md) |  | [optional] 

## Methods

### NewCouponQueryParams

`func NewCouponQueryParams() *CouponQueryParams`

NewCouponQueryParams instantiates a new CouponQueryParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCouponQueryParamsWithDefaults

`func NewCouponQueryParamsWithDefaults() *CouponQueryParams`

NewCouponQueryParamsWithDefaults instantiates a new CouponQueryParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageNumber

`func (o *CouponQueryParams) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *CouponQueryParams) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *CouponQueryParams) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *CouponQueryParams) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetPageSize

`func (o *CouponQueryParams) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *CouponQueryParams) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *CouponQueryParams) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *CouponQueryParams) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetSortKey

`func (o *CouponQueryParams) GetSortKey() string`

GetSortKey returns the SortKey field if non-nil, zero value otherwise.

### GetSortKeyOk

`func (o *CouponQueryParams) GetSortKeyOk() (*string, bool)`

GetSortKeyOk returns a tuple with the SortKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortKey

`func (o *CouponQueryParams) SetSortKey(v string)`

SetSortKey sets SortKey field to given value.

### HasSortKey

`func (o *CouponQueryParams) HasSortKey() bool`

HasSortKey returns a boolean if a field has been set.

### GetSortDescending

`func (o *CouponQueryParams) GetSortDescending() bool`

GetSortDescending returns the SortDescending field if non-nil, zero value otherwise.

### GetSortDescendingOk

`func (o *CouponQueryParams) GetSortDescendingOk() (*bool, bool)`

GetSortDescendingOk returns a tuple with the SortDescending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortDescending

`func (o *CouponQueryParams) SetSortDescending(v bool)`

SetSortDescending sets SortDescending field to given value.

### HasSortDescending

`func (o *CouponQueryParams) HasSortDescending() bool`

HasSortDescending returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CouponQueryParams) GetCreatedAt() DateTimeFilter`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CouponQueryParams) GetCreatedAtOk() (*DateTimeFilter, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CouponQueryParams) SetCreatedAt(v DateTimeFilter)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CouponQueryParams) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *CouponQueryParams) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *CouponQueryParams) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *CouponQueryParams) GetUpdatedAt() DateTimeFilter`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CouponQueryParams) GetUpdatedAtOk() (*DateTimeFilter, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CouponQueryParams) SetUpdatedAt(v DateTimeFilter)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CouponQueryParams) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *CouponQueryParams) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *CouponQueryParams) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetExpand

`func (o *CouponQueryParams) GetExpand() []string`

GetExpand returns the Expand field if non-nil, zero value otherwise.

### GetExpandOk

`func (o *CouponQueryParams) GetExpandOk() (*[]string, bool)`

GetExpandOk returns a tuple with the Expand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpand

`func (o *CouponQueryParams) SetExpand(v []string)`

SetExpand sets Expand field to given value.

### HasExpand

`func (o *CouponQueryParams) HasExpand() bool`

HasExpand returns a boolean if a field has been set.

### GetIsActive

`func (o *CouponQueryParams) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *CouponQueryParams) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *CouponQueryParams) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *CouponQueryParams) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *CouponQueryParams) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *CouponQueryParams) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetSearch

`func (o *CouponQueryParams) GetSearch() SearchFilter`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *CouponQueryParams) GetSearchOk() (*SearchFilter, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *CouponQueryParams) SetSearch(v SearchFilter)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *CouponQueryParams) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

### SetSearchNil

`func (o *CouponQueryParams) SetSearchNil(b bool)`

 SetSearchNil sets the value for Search to be an explicit nil

### UnsetSearch
`func (o *CouponQueryParams) UnsetSearch()`

UnsetSearch ensures that no value is present for Search, not even an explicit nil
### GetCurrency

`func (o *CouponQueryParams) GetCurrency() CurrencyEnum`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CouponQueryParams) GetCurrencyOk() (*CurrencyEnum, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CouponQueryParams) SetCurrency(v CurrencyEnum)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CouponQueryParams) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *CouponQueryParams) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *CouponQueryParams) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetDuration

`func (o *CouponQueryParams) GetDuration() CouponDuration`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *CouponQueryParams) GetDurationOk() (*CouponDuration, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *CouponQueryParams) SetDuration(v CouponDuration)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *CouponQueryParams) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *CouponQueryParams) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *CouponQueryParams) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetMaxRedemptions

`func (o *CouponQueryParams) GetMaxRedemptions() IntRangeFilter`

GetMaxRedemptions returns the MaxRedemptions field if non-nil, zero value otherwise.

### GetMaxRedemptionsOk

`func (o *CouponQueryParams) GetMaxRedemptionsOk() (*IntRangeFilter, bool)`

GetMaxRedemptionsOk returns a tuple with the MaxRedemptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRedemptions

`func (o *CouponQueryParams) SetMaxRedemptions(v IntRangeFilter)`

SetMaxRedemptions sets MaxRedemptions field to given value.

### HasMaxRedemptions

`func (o *CouponQueryParams) HasMaxRedemptions() bool`

HasMaxRedemptions returns a boolean if a field has been set.

### SetMaxRedemptionsNil

`func (o *CouponQueryParams) SetMaxRedemptionsNil(b bool)`

 SetMaxRedemptionsNil sets the value for MaxRedemptions to be an explicit nil

### UnsetMaxRedemptions
`func (o *CouponQueryParams) UnsetMaxRedemptions()`

UnsetMaxRedemptions ensures that no value is present for MaxRedemptions, not even an explicit nil
### GetRedeemBy

`func (o *CouponQueryParams) GetRedeemBy() DateTimeFilter`

GetRedeemBy returns the RedeemBy field if non-nil, zero value otherwise.

### GetRedeemByOk

`func (o *CouponQueryParams) GetRedeemByOk() (*DateTimeFilter, bool)`

GetRedeemByOk returns a tuple with the RedeemBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemBy

`func (o *CouponQueryParams) SetRedeemBy(v DateTimeFilter)`

SetRedeemBy sets RedeemBy field to given value.

### HasRedeemBy

`func (o *CouponQueryParams) HasRedeemBy() bool`

HasRedeemBy returns a boolean if a field has been set.

### SetRedeemByNil

`func (o *CouponQueryParams) SetRedeemByNil(b bool)`

 SetRedeemByNil sets the value for RedeemBy to be an explicit nil

### UnsetRedeemBy
`func (o *CouponQueryParams) UnsetRedeemBy()`

UnsetRedeemBy ensures that no value is present for RedeemBy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


