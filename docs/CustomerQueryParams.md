# CustomerQueryParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CouponId** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to [**NullableDateTimeFilter**](DateTimeFilter.md) |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**Emails** | Pointer to **[]string** | List of customer email address. | [optional] [default to []]
**Expand** | Pointer to **[]string** | Specifies which fields in the response should be expanded. | [optional] 
**PageNumber** | Pointer to **int32** | Page number | [optional] [default to 1]
**PageSize** | Pointer to **int32** | Page size | [optional] [default to 100]
**Search** | Pointer to [**NullableSearchFilter**](SearchFilter.md) |  | [optional] 
**SortDescending** | Pointer to **bool** | Sort direction. | [optional] [default to false]
**SortKey** | Pointer to **string** | Key name based on which data is sorted. | [optional] [default to "created_at"]
**UpdatedAt** | Pointer to [**NullableDateTimeFilter**](DateTimeFilter.md) |  | [optional] 

## Methods

### NewCustomerQueryParams

`func NewCustomerQueryParams() *CustomerQueryParams`

NewCustomerQueryParams instantiates a new CustomerQueryParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerQueryParamsWithDefaults

`func NewCustomerQueryParamsWithDefaults() *CustomerQueryParams`

NewCustomerQueryParamsWithDefaults instantiates a new CustomerQueryParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCouponId

`func (o *CustomerQueryParams) GetCouponId() string`

GetCouponId returns the CouponId field if non-nil, zero value otherwise.

### GetCouponIdOk

`func (o *CustomerQueryParams) GetCouponIdOk() (*string, bool)`

GetCouponIdOk returns a tuple with the CouponId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponId

`func (o *CustomerQueryParams) SetCouponId(v string)`

SetCouponId sets CouponId field to given value.

### HasCouponId

`func (o *CustomerQueryParams) HasCouponId() bool`

HasCouponId returns a boolean if a field has been set.

### SetCouponIdNil

`func (o *CustomerQueryParams) SetCouponIdNil(b bool)`

 SetCouponIdNil sets the value for CouponId to be an explicit nil

### UnsetCouponId
`func (o *CustomerQueryParams) UnsetCouponId()`

UnsetCouponId ensures that no value is present for CouponId, not even an explicit nil
### GetCreatedAt

`func (o *CustomerQueryParams) GetCreatedAt() DateTimeFilter`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CustomerQueryParams) GetCreatedAtOk() (*DateTimeFilter, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CustomerQueryParams) SetCreatedAt(v DateTimeFilter)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CustomerQueryParams) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *CustomerQueryParams) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *CustomerQueryParams) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetEmail

`func (o *CustomerQueryParams) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CustomerQueryParams) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CustomerQueryParams) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CustomerQueryParams) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *CustomerQueryParams) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *CustomerQueryParams) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetEmails

`func (o *CustomerQueryParams) GetEmails() []string`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *CustomerQueryParams) GetEmailsOk() (*[]string, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *CustomerQueryParams) SetEmails(v []string)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *CustomerQueryParams) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### GetExpand

`func (o *CustomerQueryParams) GetExpand() []string`

GetExpand returns the Expand field if non-nil, zero value otherwise.

### GetExpandOk

`func (o *CustomerQueryParams) GetExpandOk() (*[]string, bool)`

GetExpandOk returns a tuple with the Expand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpand

`func (o *CustomerQueryParams) SetExpand(v []string)`

SetExpand sets Expand field to given value.

### HasExpand

`func (o *CustomerQueryParams) HasExpand() bool`

HasExpand returns a boolean if a field has been set.

### GetPageNumber

`func (o *CustomerQueryParams) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *CustomerQueryParams) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *CustomerQueryParams) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *CustomerQueryParams) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetPageSize

`func (o *CustomerQueryParams) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *CustomerQueryParams) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *CustomerQueryParams) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *CustomerQueryParams) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetSearch

`func (o *CustomerQueryParams) GetSearch() SearchFilter`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *CustomerQueryParams) GetSearchOk() (*SearchFilter, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *CustomerQueryParams) SetSearch(v SearchFilter)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *CustomerQueryParams) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

### SetSearchNil

`func (o *CustomerQueryParams) SetSearchNil(b bool)`

 SetSearchNil sets the value for Search to be an explicit nil

### UnsetSearch
`func (o *CustomerQueryParams) UnsetSearch()`

UnsetSearch ensures that no value is present for Search, not even an explicit nil
### GetSortDescending

`func (o *CustomerQueryParams) GetSortDescending() bool`

GetSortDescending returns the SortDescending field if non-nil, zero value otherwise.

### GetSortDescendingOk

`func (o *CustomerQueryParams) GetSortDescendingOk() (*bool, bool)`

GetSortDescendingOk returns a tuple with the SortDescending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortDescending

`func (o *CustomerQueryParams) SetSortDescending(v bool)`

SetSortDescending sets SortDescending field to given value.

### HasSortDescending

`func (o *CustomerQueryParams) HasSortDescending() bool`

HasSortDescending returns a boolean if a field has been set.

### GetSortKey

`func (o *CustomerQueryParams) GetSortKey() string`

GetSortKey returns the SortKey field if non-nil, zero value otherwise.

### GetSortKeyOk

`func (o *CustomerQueryParams) GetSortKeyOk() (*string, bool)`

GetSortKeyOk returns a tuple with the SortKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortKey

`func (o *CustomerQueryParams) SetSortKey(v string)`

SetSortKey sets SortKey field to given value.

### HasSortKey

`func (o *CustomerQueryParams) HasSortKey() bool`

HasSortKey returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CustomerQueryParams) GetUpdatedAt() DateTimeFilter`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CustomerQueryParams) GetUpdatedAtOk() (*DateTimeFilter, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CustomerQueryParams) SetUpdatedAt(v DateTimeFilter)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CustomerQueryParams) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *CustomerQueryParams) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *CustomerQueryParams) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


