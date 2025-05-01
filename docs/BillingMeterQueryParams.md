# BillingMeterQueryParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to [**NullableDateTimeFilter**](DateTimeFilter.md) |  | [optional] 
**Expand** | Pointer to **[]string** | Specifies which fields in the response should be expanded. | [optional] 
**IsActive** | **NullableBool** |  | 
**PageNumber** | Pointer to **int32** | Page number | [optional] [default to 1]
**PageSize** | Pointer to **int32** | Page size | [optional] [default to 100]
**SortDescending** | Pointer to **bool** | Sort direction. | [optional] [default to false]
**SortKey** | Pointer to **string** | Key name based on which data is sorted. | [optional] [default to "created_at"]
**UpdatedAt** | Pointer to [**NullableDateTimeFilter**](DateTimeFilter.md) |  | [optional] 

## Methods

### NewBillingMeterQueryParams

`func NewBillingMeterQueryParams(isActive NullableBool, ) *BillingMeterQueryParams`

NewBillingMeterQueryParams instantiates a new BillingMeterQueryParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingMeterQueryParamsWithDefaults

`func NewBillingMeterQueryParamsWithDefaults() *BillingMeterQueryParams`

NewBillingMeterQueryParamsWithDefaults instantiates a new BillingMeterQueryParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *BillingMeterQueryParams) GetCreatedAt() DateTimeFilter`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BillingMeterQueryParams) GetCreatedAtOk() (*DateTimeFilter, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BillingMeterQueryParams) SetCreatedAt(v DateTimeFilter)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BillingMeterQueryParams) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *BillingMeterQueryParams) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *BillingMeterQueryParams) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetExpand

`func (o *BillingMeterQueryParams) GetExpand() []string`

GetExpand returns the Expand field if non-nil, zero value otherwise.

### GetExpandOk

`func (o *BillingMeterQueryParams) GetExpandOk() (*[]string, bool)`

GetExpandOk returns a tuple with the Expand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpand

`func (o *BillingMeterQueryParams) SetExpand(v []string)`

SetExpand sets Expand field to given value.

### HasExpand

`func (o *BillingMeterQueryParams) HasExpand() bool`

HasExpand returns a boolean if a field has been set.

### GetIsActive

`func (o *BillingMeterQueryParams) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *BillingMeterQueryParams) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *BillingMeterQueryParams) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### SetIsActiveNil

`func (o *BillingMeterQueryParams) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *BillingMeterQueryParams) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetPageNumber

`func (o *BillingMeterQueryParams) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *BillingMeterQueryParams) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *BillingMeterQueryParams) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *BillingMeterQueryParams) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetPageSize

`func (o *BillingMeterQueryParams) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *BillingMeterQueryParams) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *BillingMeterQueryParams) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *BillingMeterQueryParams) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetSortDescending

`func (o *BillingMeterQueryParams) GetSortDescending() bool`

GetSortDescending returns the SortDescending field if non-nil, zero value otherwise.

### GetSortDescendingOk

`func (o *BillingMeterQueryParams) GetSortDescendingOk() (*bool, bool)`

GetSortDescendingOk returns a tuple with the SortDescending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortDescending

`func (o *BillingMeterQueryParams) SetSortDescending(v bool)`

SetSortDescending sets SortDescending field to given value.

### HasSortDescending

`func (o *BillingMeterQueryParams) HasSortDescending() bool`

HasSortDescending returns a boolean if a field has been set.

### GetSortKey

`func (o *BillingMeterQueryParams) GetSortKey() string`

GetSortKey returns the SortKey field if non-nil, zero value otherwise.

### GetSortKeyOk

`func (o *BillingMeterQueryParams) GetSortKeyOk() (*string, bool)`

GetSortKeyOk returns a tuple with the SortKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortKey

`func (o *BillingMeterQueryParams) SetSortKey(v string)`

SetSortKey sets SortKey field to given value.

### HasSortKey

`func (o *BillingMeterQueryParams) HasSortKey() bool`

HasSortKey returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *BillingMeterQueryParams) GetUpdatedAt() DateTimeFilter`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BillingMeterQueryParams) GetUpdatedAtOk() (*DateTimeFilter, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BillingMeterQueryParams) SetUpdatedAt(v DateTimeFilter)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BillingMeterQueryParams) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *BillingMeterQueryParams) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *BillingMeterQueryParams) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


